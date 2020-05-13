package person

import (
	"net/http"

	"github.com/jackaaa89/trakt"
)

type Client struct{ b *trakt.BaseClient }

func Get(id trakt.SearchID, params *trakt.ExtendedParams) (*trakt.Person, error) {
	return getC().Get(id, params)
}

func (c *Client) Get(id trakt.SearchID, params *trakt.ExtendedParams) (*trakt.Person, error) {
	p := &trakt.Person{}
	path := trakt.FormatURLPath("/people/%s", id)
	err := c.b.Call(http.MethodGet, path, params, p)
	return p, err
}

func MovieCredits(id trakt.SearchID, params *trakt.ExtendedParams) (*trakt.Credits, error) {
	return getC().MovieCredits(id, params)
}

func (c *Client) MovieCredits(id trakt.SearchID, params *trakt.ExtendedParams) (*trakt.Credits, error) {
	cr := &trakt.Credits{}
	path := trakt.FormatURLPath("/people/%s/movies", id)
	err := c.b.Call(http.MethodGet, path, params, cr)
	return cr, err
}

func ShowCredits(id trakt.SearchID, params *trakt.ExtendedParams) (*trakt.Credits, error) {
	return getC().ShowCredits(id, params)
}

func (c *Client) ShowCredits(id trakt.SearchID, params *trakt.ExtendedParams) (*trakt.Credits, error) {
	cr := &trakt.Credits{}
	path := trakt.FormatURLPath("/people/%s/shows", id)
	err := c.b.Call(http.MethodGet, path, params, cr)
	return cr, err
}

func Lists(id trakt.SearchID, params *trakt.GetListParams) *trakt.ListIterator {
	return getC().Lists(id, params)
}

func (c *Client) Lists(id trakt.SearchID, params *trakt.GetListParams) *trakt.ListIterator {
	rcv := make([]*trakt.List, 0)
	path := trakt.FormatURLPath("people/%s/lists/%s/%s", id, params.ListType, params.SortType)
	return &trakt.ListIterator{Iterator: c.b.NewIterator(http.MethodGet, path, params, &rcv)}
}

func getC() *Client { return &Client{trakt.NewClient(trakt.GetBackend())} }
