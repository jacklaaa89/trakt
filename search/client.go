package search

import (
	"net/http"

	"github.com/jackaaa89/trakt"
)

type Client struct{ b *trakt.BaseClient }

func TextQuery(params *trakt.SearchQueryParams) *trakt.SearchResultIterator {
	return getC().TextQuery(params)
}

// TextQuery performs a search using a textual based query.
// the query can be restricted to certain search fields by providing the Fields parameter.
func (c *Client) TextQuery(params *trakt.SearchQueryParams) *trakt.SearchResultIterator {
	rcv := make([]*trakt.SearchResult, 0)
	path := trakt.FormatURLPath("/search/%s", params.Type)
	return &trakt.SearchResultIterator{Iterator: c.b.NewIterator(http.MethodGet, path, params, &rcv)}
}

func IDLookup(id trakt.SearchID, params *trakt.IDLookupParams) *trakt.SearchResultIterator {
	return getC().IDLookup(id, params)
}

func (c *Client) IDLookup(id trakt.SearchID, params *trakt.IDLookupParams) *trakt.SearchResultIterator {
	rcv := make([]*trakt.SearchResult, 0)
	path := trakt.FormatURLPath(trakt.IDPath(id), id)
	return &trakt.SearchResultIterator{Iterator: c.b.NewIterator(http.MethodGet, path, params, &rcv)}
}

func getC() *Client { return &Client{trakt.NewClient(trakt.GetBackend())} }
