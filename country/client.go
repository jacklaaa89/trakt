package country

import (
	"net/http"

	"github.com/jacklaaa89/trakt"
)

type Client struct{ b *trakt.BaseClient }

func List(params *trakt.ListByTypeParams) *trakt.CountryIterator {
	return getC().List(params)
}

func (c *Client) List(params *trakt.ListByTypeParams) *trakt.CountryIterator {
	path := trakt.FormatURLPath("/countries/%s", params.Type)
	return &trakt.CountryIterator{BasicIterator: c.b.NewSimulatedIterator(http.MethodGet, path, params)}
}

func getC() *Client { return &Client{trakt.NewClient(trakt.GetBackend())} }
