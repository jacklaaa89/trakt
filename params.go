package trakt

import (
	"context"
	"net/http"
	"strconv"
)

// listMeta is the structure that contains the common properties
// of List iterators. The Count property is only populated if the
// total_count include option is passed in (see tests for example).
type listMeta struct {
	CurrentPage int64 `json:"-" url:"-"`
	Limit       int64 `json:"-" url:"-"`
	TotalPages  int64 `json:"-" url:"-"`
	TotalCount  int64 `json:"-" url:"-"`
}

func (l listMeta) meta() listMeta { return l }

// UnmarshalHeaders allows us to unmarshal a response from
// the response HTTP headers.
// this is the case for pagination values where they are supplied
// in headers and not the response.
func (l *listMeta) UnmarshalHeaders(h http.Header) error {
	*l = listMeta{
		Limit:       parseInt(h.Get(`X-Pagination-Limit`)),
		CurrentPage: parseInt(h.Get(`X-Pagination-Page`)),
		TotalPages:  parseInt(h.Get(`X-Pagination-Page-Count`)),
		TotalCount:  parseInt(h.Get(`X-Pagination-Item-Count`)),
	}

	return nil
}

// BasicListParams is the structure that contains the common properties
// of any *ListParams structure, it is used when no OAuth token is required.
type BasicListParams struct {
	// Context used for request. It may carry deadlines, cancelation signals,
	// and other request-scoped values across API boundaries and between
	// processes.
	//
	// Note that a cancelled or timed out context does not provide any
	// guarantee whether the operation was or was not completed on Stripe's API
	// servers. For certainty, you must either retry with the same idempotency
	// key or query the state of the API.
	Context context.Context `url:"-" json:"-"`

	// Headers may be used to provide extra header lines on the HTTP request.
	Headers http.Header `url:"-" json:"-"`

	Page  *int64 `url:"page,omitempty" json:"-"`
	Limit *int64 `url:"limit,omitempty" json:"-"`
}

func (p *BasicListParams) context() context.Context {
	if p != nil && p.Context != nil {
		return p.Context
	}

	return context.Background()
}

func (p *BasicListParams) headers() http.Header {
	if p == nil {
		return nil
	}

	return p.Headers
}

func (p *BasicListParams) oauth() string { return `` }

func (p *BasicListParams) setPagination(page, limit int64) {
	p.Page = Int64(page)
	p.Limit = Int64(limit)
}

// ListParams is the structure that contains the common properties
// of any *ListParams structure.
type ListParams struct {
	// Context used for request. It may carry deadlines, cancelation signals,
	// and other request-scoped values across API boundaries and between
	// processes.
	//
	// Note that a cancelled or timed out context does not provide any
	// guarantee whether the operation was or was not completed on Stripe's API
	// servers. For certainty, you must either retry with the same idempotency
	// key or query the state of the API.
	Context context.Context `url:"-" json:"-"`

	// Headers may be used to provide extra header lines on the HTTP request.
	Headers http.Header `url:"-" json:"-"`

	Page  *int64 `url:"page,omitempty" json:"-"`
	Limit *int64 `url:"limit,omitempty" json:"-"`

	// OAuth token to use with the request.
	// this is passed as a header if supplied.
	OAuth string `url:"-" json:"-"`
}

func (p *ListParams) context() context.Context {
	if p != nil && p.Context != nil {
		return p.Context
	}

	return context.Background()
}

func (p *ListParams) headers() http.Header {
	if p == nil {
		return nil
	}

	return p.Headers
}

func (p *ListParams) oauth() string {
	if p == nil {
		return ``
	}

	return p.OAuth
}

func (p *ListParams) setPagination(page, limit int64) {
	p.Page = Int64(page)
	p.Limit = Int64(limit)
}

// ListParamsContainer is a general interface for which all list parameter
// structs should comply. They achieve this by embedding a ListParams struct
// and inheriting its implementation of this interface.
type ListParamsContainer interface {
	ParamsContainer
	setPagination(page, limit int64)
}

// BasicParams parameters which do not require an OAuth token.
type BasicParams struct {
	// Context used for request. It may carry deadlines, cancelation signals,
	// and other request-scoped values across API boundaries and between
	// processes.
	//
	// Note that a cancelled or timed out context does not provide any
	// guarantee whether the operation was or was not completed on Stripe's API
	// servers. For certainty, you must either retry with the same idempotency
	// key or query the state of the API.
	Context context.Context `url:"-" json:"-"`

	// Headers may be used to provide extra header lines on the HTTP request.
	Headers http.Header `url:"-" json:"-"`
}

func (p *BasicParams) setPagination(_, _ int64) {}

func (p *BasicParams) context() context.Context {
	if p != nil && p.Context != nil {
		return p.Context
	}

	return context.Background()
}

func (p *BasicParams) headers() http.Header {
	if p == nil {
		return nil
	}

	return p.Headers
}

func (p *BasicParams) oauth() string { return `` }

// Params is the structure that contains the common properties
// of any *Params structure.
type Params struct {
	// Context used for request. It may carry deadlines, cancelation signals,
	// and other request-scoped values across API boundaries and between
	// processes.
	//
	// Note that a cancelled or timed out context does not provide any
	// guarantee whether the operation was or was not completed on Stripe's API
	// servers. For certainty, you must either retry with the same idempotency
	// key or query the state of the API.
	Context context.Context `url:"-" json:"-"`

	// Headers may be used to provide extra header lines on the HTTP request.
	Headers http.Header `url:"-" json:"-"`

	// OAuth token to use with the request.
	// this is passed as a header if supplied.
	OAuth string `url:"-" json:"-"`
}

func (p *Params) setPagination(_, _ int64) {}

func (p *Params) context() context.Context {
	if p != nil && p.Context != nil {
		return p.Context
	}

	return context.Background()
}

func (p *Params) headers() http.Header {
	if p == nil {
		return nil
	}

	return p.Headers
}

func (p *Params) oauth() string {
	if p == nil {
		return ``
	}

	return p.OAuth
}

// ParamsContainer is a general interface for which all parameter structs
// should comply. They achieve this by embedding a Params struct and inheriting
// its implementation of this interface.
type ParamsContainer interface {
	context() context.Context
	headers() http.Header
	oauth() string
}

// parseInt helper function to parse a uint from a string.
func parseInt(s string) int64 {
	i, _ := strconv.Atoi(s)
	return int64(i)
}
