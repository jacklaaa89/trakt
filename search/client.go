package search

import (
	"net/http"

	"github.com/jacklaaa89/trakt"
)

type Client struct{ b *trakt.BaseClient }

// wrappedSearchQuery this is only required because there seems to be
// a weird bug with the "query" package in which it only runs the custom
// encoder on sub-fields and not the top level interface originally supplied to
// the "query.Values" func.
// see: "github.com/google/go-querystring/query".reflectValues to see what happens.
type wrappedSearchQuery struct{ *trakt.SearchQueryParams }

func TextQuery(params *trakt.SearchQueryParams) *trakt.SearchResultIterator {
	return getC().TextQuery(params)
}

// TextQuery performs a search using a textual based query.
// the query can be restricted to certain search fields by providing the Fields parameter.
func (c *Client) TextQuery(params *trakt.SearchQueryParams) *trakt.SearchResultIterator {
	path := trakt.FormatURLPath("/search/%s", params.Type)
	return &trakt.SearchResultIterator{Iterator: c.b.NewIterator(http.MethodGet, path, wrappedSearchQuery{params})}
}

func IDLookup(id trakt.SearchID, params *trakt.IDLookupParams) *trakt.SearchResultIterator {
	return getC().IDLookup(id, params)
}

func (c *Client) IDLookup(id trakt.SearchID, params *trakt.IDLookupParams) *trakt.SearchResultIterator {
	path := trakt.FormatURLPath(trakt.IDPath(id), id)
	return &trakt.SearchResultIterator{Iterator: c.b.NewIterator(http.MethodGet, path, params)}
}

func getC() *Client { return &Client{trakt.NewClient(trakt.GetBackend())} }
