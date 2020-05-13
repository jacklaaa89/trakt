package scrobble

import (
	"net/http"

	"github.com/jackaaa89/trakt"
)

type Client struct {
	B   trakt.Backend
	Key string
}

func Start(params *trakt.ScrobbleParams) (*trakt.Scrobble, error) {
	return getC().Start(params)
}

func (c *Client) Start(params *trakt.ScrobbleParams) (*trakt.Scrobble, error) {
	s := &trakt.Scrobble{}
	err := c.B.Call(http.MethodPost, "/scrobble/start", c.Key, params, s)
	return s, err
}

func Pause(params *trakt.ScrobbleParams) (*trakt.Scrobble, error) {
	return getC().Pause(params)
}

func (c *Client) Pause(params *trakt.ScrobbleParams) (*trakt.Scrobble, error) {
	s := &trakt.Scrobble{}
	err := c.B.Call(http.MethodPost, "/scrobble/pause", c.Key, params, s)
	return s, err
}

func Stop(params *trakt.ScrobbleParams) (*trakt.Scrobble, error) {
	return getC().Stop(params)
}

func (c *Client) Stop(params *trakt.ScrobbleParams) (*trakt.Scrobble, error) {
	s := &trakt.Scrobble{}
	err := c.B.Call(http.MethodPost, "/scrobble/stop", c.Key, params, s)
	return s, err
}

func getC() *Client {
	return &Client{B: trakt.GetBackend(), Key: trakt.Key}
}
