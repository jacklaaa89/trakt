package checkin

import (
	"errors"
	"net/http"

	"github.com/jacklaaa89/trakt"
)

type Client struct{ b trakt.BaseClient }

func Start(params *trakt.StartCheckinParams) (*trakt.Checkin, error) {
	return getC().Start(params)
}

func (c *Client) Start(params *trakt.StartCheckinParams) (*trakt.Checkin, error) {
	if params == nil {
		return nil, errors.New(`params cannot be nil`)
	}

	ci := &trakt.Checkin{}
	p := &wrappedCheckinParams{*params}
	err := c.b.Call(http.MethodPost, "/checkin", p, ci)
	return ci, err
}

func Stop(params *trakt.Params) error {
	return getC().Stop(params)
}

func (c *Client) Stop(params *trakt.Params) error {
	return c.b.Call(http.MethodDelete, "/checkin", params, nil)
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

func getC() *Client { return &Client{trakt.NewClient()} }
