package network

import (
	"net/http"

	"github.com/jackaaa89/trakt"
)

type Client struct {
	B   trakt.Backend
	Key string
}

func List(params *trakt.BasicParams) *trakt.NetworkIterator {
	return getC().List(params)
}

func (c *Client) List(params *trakt.BasicParams) *trakt.NetworkIterator {
	return &trakt.NetworkIterator{
		Iterator: trakt.NewSimulatedIterator(params, func(p trakt.ListParamsContainer) (trakt.IterationFrame, error) {
			l := make([]*trakt.Network, 0)
			f := trakt.NewEmptyFrame(&l)
			err := c.B.CallWithFrame(http.MethodGet, "/networks", c.Key, p, f)
			return f, err
		}),
	}
}

func getC() *Client {
	return &Client{B: trakt.GetBackend(), Key: trakt.Key}
}
