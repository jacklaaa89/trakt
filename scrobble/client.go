package scrobble

import (
	"net/http"

	"github.com/jackaaa89/trakt"
)

type Client struct{ b *trakt.BaseClient }

func Start(params *trakt.ScrobbleParams) (*trakt.ScrobbleEvent, error) {
	return getC().Start(params)
}

func (c *Client) Start(params *trakt.ScrobbleParams) (*trakt.ScrobbleEvent, error) {
	s := &trakt.ScrobbleEvent{}
	err := c.b.Call(http.MethodPost, "/scrobble/start", params, s)
	return s, err
}

func Pause(params *trakt.ScrobbleParams) (*trakt.ScrobbleEvent, error) {
	return getC().Pause(params)
}

func (c *Client) Pause(params *trakt.ScrobbleParams) (*trakt.ScrobbleEvent, error) {
	s := &trakt.ScrobbleEvent{}
	err := c.b.Call(http.MethodPost, "/scrobble/pause", params, s)
	return s, err
}

func Stop(params *trakt.ScrobbleParams) (*trakt.ScrobbleEvent, error) {
	return getC().Stop(params)
}

func (c *Client) Stop(params *trakt.ScrobbleParams) (*trakt.ScrobbleEvent, error) {
	s := &trakt.ScrobbleEvent{}
	err := c.b.Call(http.MethodPost, "/scrobble/stop", params, s)
	return s, err
}

func getC() *Client { return &Client{trakt.NewClient(trakt.GetBackend())} }
