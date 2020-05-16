package network

import (
	"net/http"

	"github.com/jacklaaa89/trakt"
)

type Client struct{ b trakt.BaseClient }

func List(params *trakt.BasicParams) *trakt.NetworkIterator {
	return getC().List(params)
}

func (c *Client) List(params *trakt.BasicParams) *trakt.NetworkIterator {
	return &trakt.NetworkIterator{BasicIterator: c.b.NewSimulatedIterator(http.MethodGet, "/networks", params)}
}

func getC() *Client { return &Client{trakt.NewClient()} }
