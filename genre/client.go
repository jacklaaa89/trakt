package genre

import (
	"net/http"

	"github.com/jackaaa89/trakt"
)

type Client struct{ b *trakt.BaseClient }

func List(params *trakt.ListByTypeParams) *trakt.GenreIterator {
	return getC().List(params)
}

func (c *Client) List(params *trakt.ListByTypeParams) *trakt.GenreIterator {
	gl := make([]*trakt.Genre, 0)
	path := trakt.FormatURLPath("/genres/%s", params.Type)
	return &trakt.GenreIterator{BasicIterator: c.b.NewSimulatedIterator(http.MethodGet, path, params, &gl)}
}

func getC() *Client { return &Client{trakt.NewClient(trakt.GetBackend())} }
