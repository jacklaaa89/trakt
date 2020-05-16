package authorization

import (
	"context"
	"net/http"
	"time"

	"github.com/jacklaaa89/trakt"
)

type Client struct{ b trakt.BaseClient }

type newDeviceCodeParams struct {
	*trakt.BasicParams
	ClientID string `json:"client_id" url:"-"`
}

func NewCode(params *trakt.BasicParams) (*trakt.DeviceCode, error) {
	return getC().NewCode(params)
}

func (c *Client) NewCode(params *trakt.BasicParams) (*trakt.DeviceCode, error) {
	d := &trakt.DeviceCode{}
	p := &newDeviceCodeParams{params, c.b.Key()}
	err := c.b.Call(http.MethodPost, "/oauth/device/code", p, &d)
	return d, err
}

type wrappedPollCodeParams struct {
	*trakt.PollCodeParams
	ClientID string `json:"client_id" url:"-"`
}

// Code implements ErrorHandler interface.
// Device polling status codes mean different things.
func (wrappedPollCodeParams) Code(statusCode int) trakt.ErrorCode {
	switch statusCode {
	case http.StatusBadRequest:
		return trakt.ErrorCodePendingDeviceCode
	case http.StatusNotFound:
		return trakt.ErrorCodeInvalidDeviceCode
	case http.StatusConflict:
		return trakt.ErrorCodeDeviceCodeUsed
	case http.StatusGone:
		return trakt.ErrorCodeDeviceCodeExpired
	case http.StatusTeapot: // fair play
		return trakt.ErrorCodeDeviceCodeDenied
	}

	return trakt.DefaultErrorHandler.Code(statusCode)
}

// Poll attempts to poll on the supplied interval to determine if the user has entered
// the device code.
func Poll(params *trakt.PollCodeParams) (*trakt.Token, error) {
	return getC().Poll(params)
}

// Poll attempts to poll on the supplied interval to determine if the user has entered
// the device code.
// this function is blocking until either:
// - the supplied context is marked as done
// - the supplied expires in duration exceeds its deadline
// - any errors which deem that the token cannot be processed are returned from the API.
func (c *Client) Poll(params *trakt.PollCodeParams) (*trakt.Token, error) {
	r := <-c.PollAsync(params)
	return r.Token, r.Err
}

// PollAsync attempts to poll on the supplied interval to determine if the user has entered
// the device code.
func PollAsync(params *trakt.PollCodeParams) <-chan *trakt.PollResult {
	return getC().PollAsync(params)
}

// PollAsync attempts to poll on the supplied interval to determine if the user has entered
// the device code.
func (c *Client) PollAsync(params *trakt.PollCodeParams) <-chan *trakt.PollResult {
	cCtx := params.Context
	if cCtx == nil {
		cCtx = context.Background()
	}

	ch := make(chan *trakt.PollResult, 1)

	go func() {
		ctx, cnl := context.WithTimeout(cCtx, params.ExpiresIn)
		defer cnl()

		// perform initial request.
		t, err := c.poll(withContext(ctx, params))
		if err == nil {
			ch <- &trakt.PollResult{Token: t}
			return
		}

		traktError, ok := err.(*trakt.Error)
		if !ok || !canContinuePolling(traktError.Code) {
			ch <- &trakt.PollResult{Err: err}
			return
		}

		tk := time.NewTicker(params.Interval)
		defer func() {
			tk.Stop()
			close(ch)
		}()

		for {
			select {
			case <-tk.C:
				// perform a new poll.
				t, err = c.poll(withContext(ctx, params))
				if err == nil {
					ch <- &trakt.PollResult{Token: t}
					return
				}
				traktError, ok := err.(*trakt.Error)
				if !ok || !canContinuePolling(traktError.Code) {
					ch <- &trakt.PollResult{Err: err}
					return
				}
			case <-ctx.Done():
				ch <- &trakt.PollResult{Err: ctx.Err()}
				return
			}
		}
	}()

	return ch
}

func (c *Client) poll(params *trakt.PollCodeParams) (*trakt.Token, error) {
	t := &trakt.Token{}
	p := &wrappedPollCodeParams{params, c.b.Key()}
	err := c.b.Call(http.MethodPost, "/oauth/device/token", p, t)
	return t, err
}

// withContext makes a shallow copy of the supplied params and returns it with the
// supplied context attached.
func withContext(ctx context.Context, w *trakt.PollCodeParams) *trakt.PollCodeParams {
	cp := *w
	cp.Context = ctx
	return &cp
}

// canContinuePolling determines if the status code returned means that we can
// continue polling.
func canContinuePolling(e trakt.ErrorCode) bool { return e == trakt.ErrorCodePendingDeviceCode }

func getC() *Client { return &Client{trakt.NewClient()} }
