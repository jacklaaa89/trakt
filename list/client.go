package list

import (
	"net/http"

	"github.com/jacklaaa89/trakt"
)

type Client struct{ b *trakt.BaseClient }

func Trending(params *trakt.BasicListParams) *trakt.RecentListIterator {
	return getC().Trending(params)
}

func (c *Client) Trending(params *trakt.BasicListParams) *trakt.RecentListIterator {
	return c.generateListIterator("trending", params)
}

func Popular(params *trakt.BasicListParams) *trakt.RecentListIterator {
	return getC().Popular(params)
}

func (c *Client) Popular(params *trakt.BasicListParams) *trakt.RecentListIterator {
	return c.generateListIterator("popular", params)
}

func (c *Client) generateListIterator(action string, params *trakt.BasicListParams) *trakt.RecentListIterator {
	path := trakt.FormatURLPath("/lists/%s", action)
	return &trakt.RecentListIterator{Iterator: c.b.NewIterator(http.MethodGet, path, params)}
}

func getC() *Client { return &Client{trakt.NewClient(trakt.GetBackend())} }
