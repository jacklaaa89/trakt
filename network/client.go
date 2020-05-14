package network

import (
	"net/http"

	"github.com/jackaaa89/trakt"
)

type Client struct{ b *trakt.BaseClient }

func List(params *trakt.BasicParams) *trakt.NetworkIterator {
	return getC().List(params)
}

func (c *Client) List(params *trakt.BasicParams) *trakt.NetworkIterator {
	l := make([]*trakt.Network, 0)
	return &trakt.NetworkIterator{BasicIterator: c.b.NewSimulatedIterator(http.MethodGet, "/networks", params, &l)}
}

func getC() *Client { return &Client{trakt.NewClient(trakt.GetBackend())} }
