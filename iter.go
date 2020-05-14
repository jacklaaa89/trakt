package trakt

import (
	"reflect"
)

// NoLimit defines that we have no limit on the amount
// of pages to paginate through.
const NoLimit int64 = -1

func newEmptyFrame(rcv interface{}) iterationFrame { return &frame{r: rcv} }

type frame struct {
	listMeta
	r interface{}
}

func (f *frame) rcv() interface{} { return f.r }

// iterationFrame represents a window or slice of
// an entire pagination result.
type iterationFrame interface {
	// rcv returns the receiver which the response data is unmarshalled into.
	rcv() interface{}
	// meta the current metadata for the position we are in the complete set.
	meta() listMeta
}

// BasicIterator represents an iterator which does not deal with paging.
// this can be used where the iterator does not require a page limit be set.
type BasicIterator interface {
	// Current retrieves the current object in view.
	Current() interface{}
	// Err retrieves any error that occurred performing
	// an operation to get a "frame". A frame represents a single
	// paginated set of results (the results for a single page)
	Err() error
	// Next moves the cursor to the next point in the results
	// performing an operation to retrieve the next result frame
	// from the API if required.
	Next() bool
	// getPage internal function which allows us to retrieve the next page
	// of results.
	getPage() iterationFrame
}

// Iterator a generic representation of an iterator.
type Iterator interface {
	BasicIterator
	// PageLimit sets an absolute limit on how many pages to iterate through.
	// if the result set has less pages than the limit, obviously it will
	// finish before that.
	PageLimit(page int64)
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

// PageLimit is not required on a single iterator as it only deals
// with a single page. So the limit is already enforced to 1.
func (s *singleIter) PageLimit(_ int64) {}

func (s *singleIter) getPage() iterationFrame {
	var window iterationFrame
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
	meta  listMeta
	limit int64
}

// PageLimit allows us to limit the amount of pages to paginate.
func (it *iter) PageLimit(page int64) { it.limit = page }

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
	np := lm.CurrentPage + 1
	if len(it.values) == 0 && (lm.TotalPages > lm.CurrentPage) && !it.hasReachedLimit(np) {
		// update the page number.
		it.listParams.setPagination(np, lm.Limit)
		it.getPage()
	}

	return it.singleIter.Next()
}

// hasReachedLimit determines if we have reached the defined page limit.
func (it *iter) hasReachedLimit(nextPage int64) bool {
	if it.limit == NoLimit {
		return false
	}

	return it.limit < nextPage
}

func (it *iter) getPage() iterationFrame {
	window := it.singleIter.getPage()
	if window != nil {
		it.meta = window.meta()
	}

	return window
}

// Query is the function used to get a page listing.
type Query func(ListParamsContainer) (iterationFrame, error)

// newIterator returns a new iter for a given query and its options.
func newIterator(container ListParamsContainer, query Query) Iterator {
	iter := &iter{
		singleIter: singleIter{
			listParams: container,
			query:      query,
		},
		limit: NoLimit, // by default paginate through all results available.
	}

	iter.getPage()

	return iter
}

// newSimulatedIterator a simulated iterator is an iterator which
// only works its way though the single frame (result set) it will never
// attempt to move to another page. This is used so we can have all lists of
// results (paginated or not) will utilise an iterator so the library is consistent.
func newSimulatedIterator(container ListParamsContainer, query Query) Iterator {
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
