package trakt

import (
	"reflect"
)

func NewEmptyFrame(rcv interface{}) IterationFrame { return &frame{r: rcv} }

type frame struct {
	ListMeta
	r interface{}
}

func (f *frame) rcv() interface{} { return f.r }

// IterationFrame represents a window or slice of
// an entire pagination result.
type IterationFrame interface {
	// rcv returns the receiver which the response data is unmarshalled into.
	rcv() interface{}
	// meta the current metadata for the position we are in the complete set.
	meta() ListMeta
}

type Iterator interface {
	Current() interface{}
	Err() error
	Next() bool
	getPage() IterationFrame
}

type singleIter struct {
	cur        interface{}
	err        error
	values     []interface{}
	query      Query
	listParams ListParamsContainer
}

func (s *singleIter) Current() interface{} { return s.cur }
func (s *singleIter) Err() error           { return s.err }

func (s *singleIter) getPage() IterationFrame {
	var window IterationFrame
	cp := s.listParams

	// we pass a copy through so that internally we dont update our reference.
	window, s.err = s.query(cp)
	if window != nil {
		s.values = toInterfaceSlice(window.rcv())
	}

	return window
}

func (s *singleIter) Next() bool {
	if s.err != nil {
		return false
	}

	if len(s.values) == 0 {
		return false
	}
	s.cur = s.values[0]
	s.values = s.values[1:]
	return true
}

// iter provides a convenient interface
// for iterating over the elements
// returned from paginated list API calls.
// Successive calls to the Next method
// will step through each item in the list,
// fetching pages of items as needed.
// Iterators are not thread-safe, so they should not be consumed
// across multiple goroutines.
type iter struct {
	singleIter
	meta ListMeta
}

// Next advances the iter to the next item in the list,
// which will then be available
// through the Current method.
// It returns false when the iterator stops
// at the end of the list.
func (it *iter) Next() bool {
	if it.err != nil {
		return false
	}

	lm := it.meta
	if len(it.values) == 0 && (lm.TotalPages > lm.CurrentPage) {
		// update the page number.
		it.listParams.setPagination(lm.CurrentPage+1, lm.Limit)
		it.getPage()
	}

	return it.singleIter.Next()
}

func (it *iter) getPage() IterationFrame {
	window := it.singleIter.getPage()
	if window != nil {
		it.meta = window.meta()
	}

	return window
}

// Query is the function used to get a page listing.
type Query func(ListParamsContainer) (IterationFrame, error)

// NewIterator returns a new iter for a given query and its options.
func NewIterator(container ListParamsContainer, query Query) Iterator {
	iter := &iter{
		singleIter: singleIter{
			listParams: container,
			query:      query,
		},
	}

	iter.getPage()

	return iter
}

// NewSimulatedIterator a simulated iterator is an iterator which
// only works its way though the single frame (result set) it will never
// attempt to move to another page. This is used so we can have all lists of
// results (paginated or not) will utilise an iterator so the library is consistent.
func NewSimulatedIterator(container ListParamsContainer, query Query) Iterator {
	iter := &singleIter{
		query:      query,
		listParams: container,
	}
	iter.getPage()
	return iter
}

func toInterfaceSlice(i interface{}) []interface{} {
	if i == nil {
		return nil
	}

	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil
		}

		return toInterfaceSlice(v.Elem().Interface())
	}

	switch v.Kind() {
	case reflect.Map, reflect.Slice, reflect.Array:
		break
	default:
		return nil // not a valid type.
	}

	var idx int
	s := make([]interface{}, v.Len())
	for idx < v.Len() {
		iv := v.Index(idx)
		if !iv.IsValid() || iv.IsNil() {
			continue
		}

		s[idx] = iv.Interface()
		idx++
	}

	return s[:idx]
}
