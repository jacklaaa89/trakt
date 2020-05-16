package language

import (
	"net/http"

	"github.com/jacklaaa89/trakt"
)

type Client struct{ b *trakt.BaseClient }

func List(params *trakt.ListByTypeParams) *trakt.LanguageIterator {
	return getC().List(params)
}

func (c *Client) List(params *trakt.ListByTypeParams) *trakt.LanguageIterator {
	path := trakt.FormatURLPath("/languages/%s", params.Type)
	return &trakt.LanguageIterator{BasicIterator: c.b.NewSimulatedIterator(http.MethodGet, path, params)}
}

func getC() *Client { return &Client{trakt.NewClient(trakt.GetBackend())} }
