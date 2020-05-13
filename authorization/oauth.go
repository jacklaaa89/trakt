package authorization

import (
	"errors"
	"net/http"

	"github.com/google/go-querystring/query"

	"github.com/jackaaa89/trakt"
)

type grantType string

const (
	authorizationCode grantType = "authorization_code"
	refreshToken      grantType = "refresh_token"
)

type wrappedAuthorizeURLParams struct {
	trakt.AuthorizationURLParams `url:",inline"`

	ClientID     string `url:"client_id"`
	ResponseType string `url:"response_type"`
}

// AuthorizeURL attempts to generates the OAuth URL.
func AuthorizeURL(params *trakt.AuthorizationURLParams) (string, error) {
	return getC().AuthorizeURL(params)
}

// AuthorizeURL attempts to generates the OAuth URL.
func (c *Client) AuthorizeURL(params *trakt.AuthorizationURLParams) (string, error) {
	if params == nil {
		return "", errors.New("params cannot be nil")
	}
	uv, err := query.Values(wrappedAuthorizeURLParams{
		AuthorizationURLParams: *params,
		ClientID:               c.b.Key,
		ResponseType:           "code",
	})

	if err != nil {
		return ``, err
	}

	return trakt.OAuthURL + trakt.FormatURLPath(`/oauth/authorize?`) + uv.Encode(), nil
}

type genericTokenParameters struct {
	ClientID  string    `json:"client_id" url:"-"`
	GrantType grantType `json:"grant_type" url:"-"`
}

type wrappedExchangeCodeParams struct {
	*trakt.ExchangeCodeParams
	genericTokenParameters
}

func ExchangeCode(params *trakt.ExchangeCodeParams) (*trakt.Token, error) {
	return getC().ExchangeCode(params)
}

// ExchangeCode attempts to exchange a code retrieved from the OAuth process
// into an access token, the client id is omitted from the parameters as
// it is derived from trakt.Key and authentication is the
// the only places in the API where the client secret is required.
func (c *Client) ExchangeCode(params *trakt.ExchangeCodeParams) (*trakt.Token, error) {
	t := &trakt.Token{}
	p := &wrappedExchangeCodeParams{params, genericTokenParameters{c.b.Key, authorizationCode}}
	err := c.b.Call(http.MethodPost, `/oauth/token`, p, t)
	return t, err
}

type wrappedRefreshTokenParams struct {
	*trakt.RefreshTokenParams
	genericTokenParameters
}

func RefreshToken(params *trakt.RefreshTokenParams) (*trakt.Token, error) {
	return getC().RefreshToken(params)
}

func (c *Client) RefreshToken(params *trakt.RefreshTokenParams) (*trakt.Token, error) {
	t := &trakt.Token{}
	p := &wrappedRefreshTokenParams{params, genericTokenParameters{c.b.Key, refreshToken}}
	err := c.b.Call(http.MethodPost, `/oauth/token`, p, t)
	return t, err
}

type wrappedRevokeTokenParams struct {
	*trakt.RevokeTokenParams
	ClientID string `json:"client_id" url:"-"`
}

func RevokeToken(params *trakt.RevokeTokenParams) error {
	return getC().RevokeToken(params)
}

func (c *Client) RevokeToken(params *trakt.RevokeTokenParams) error {
	p := &wrappedRevokeTokenParams{params, c.b.Key}
	return c.b.Call(http.MethodPost, `/oauth/revoke`, p, nil)
}
