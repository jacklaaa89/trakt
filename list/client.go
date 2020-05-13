package list

import (
	"net/http"

	"github.com/jackaaa89/trakt"
)

type Client struct {
	B   trakt.Backend
	Key string
}

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
	return &trakt.RecentListIterator{
		Iterator: trakt.NewIterator(params, func(p trakt.ListParamsContainer) (trakt.IterationFrame, error) {
			ll := make([]*trakt.RecentList, 0)
			f := trakt.NewEmptyFrame(&ll)
			err := c.B.CallWithFrame(http.MethodGet, trakt.FormatURLPath("/lists/%s", action), c.Key, params, f)
			return f, err
		}),
	}
}

func getC() *Client { return &Client{trakt.GetBackend(), trakt.Key} }
