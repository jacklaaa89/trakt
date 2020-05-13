package checkin

import (
	"net/http"

	"github.com/jackaaa89/trakt"
)

type Client struct {
	B   trakt.Backend
	Key string
}

func Start(params *trakt.StartCheckinParams) (*trakt.Checkin, error) {
	return getC().Start(params)
}

func (c *Client) Start(params *trakt.StartCheckinParams) (*trakt.Checkin, error) {
	ci := &trakt.Checkin{}
	p := &wrappedCheckinParams{*params}
	err := c.B.Call(http.MethodPost, "/checkin", c.Key, p, ci)
	return ci, err
}

func Stop(params *trakt.Params) error {
	return getC().Stop(params)
}

func (c *Client) Stop(params *trakt.Params) error {
	return c.B.Call(http.MethodDelete, "/checkin", c.Key, params, nil)
}

type wrappedCheckinParams struct {
	trakt.StartCheckinParams
}

func (wrappedCheckinParams) Code(statusCode int) trakt.ErrorCode {
	if statusCode == http.StatusConflict {
		return trakt.ErrorCodeCheckinInProgress
	}

	return trakt.DefaultErrorHandler.Code(statusCode)
}

func getC() *Client { return &Client{trakt.GetBackend(), trakt.Key} }
