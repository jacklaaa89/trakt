package trakt

import (
	"encoding/json"
	"strconv"
	"time"
)

// AuthorizationURLParams the parameters required to
// generate the authorization URL which is used to authenticate
// users via OAuth 2.
type AuthorizationURLParams struct {
	// BasicParams is the basic parameters which all requests can take.
	BasicParams

	// RedirectURI the URL to redirect to after the user accepts or denies
	// the request to allow access to your app. The access code is send to this
	// URL on a success. This has to match the URL set in the App settings
	// on Trakt.
	RedirectURI string `url:"redirect_uri" json:"-"`
	// State  this should be a random unique string which is send with the request
	// to authenticate. When trakt redirects back to your app, this will be sent back
	// so you can perform a check to see if its the same that you sent originally.
	State string `url:"state" json:"-"`
}

// ExchangeCodeParams parameters required to exchange a authorization code
// for an access token.
type ExchangeCodeParams struct {
	// BasicParams is the basic parameters which all requests can take.
	BasicParams

	// RedirectURI the URL which was set in the original request to authenticate.
	RedirectURI string `url:"-" json:"redirect_uri"`
	// Code the code given to us after a successful authorization.
	Code string `url:"-" json:"code"`
	// ClientSecret the client secret generated by trakt which is unique to our app.
	// this can be found in app settings. DO NOT EXPOSE THIS VALUE.
	ClientSecret string `url:"-" json:"client_secret"`
}

// RefreshTokenParams represents the parameters required in order to
// refresh an access token after it has expired.
type RefreshTokenParams struct {
	// BasicParams is the basic parameters which all requests can take.
	BasicParams

	// RedirectURI the URL which was set in the original request to authenticate.
	RedirectURI string `url:"-" json:"redirect_uri"`
	// RefreshToken the refresh token supplied when we originally retrieved the access token.
	RefreshToken string `url:"-" json:"refresh_token"`
	// ClientSecret the client secret generated by trakt which is unique to our app.
	// this can be found in app settings. DO NOT EXPOSE THIS VALUE.
	ClientSecret string `url:"-" json:"client_secret"`
}

// RevokeTokenParams the parameters required in order to revoke an access token
// once an access token has been revoked, it cannot be used for any API call from there on
// the user will need to be re-authenticated if you require access again.
type RevokeTokenParams struct {
	// BasicParams is the basic parameters which all requests can take.
	BasicParams

	// AccessToken the token to revoke.
	AccessToken string `url:"-" json:"token"`
	// ClientSecret the client secret generated by trakt which is unique to our app.
	// this can be found in app settings. DO NOT EXPOSE THIS VALUE.
	ClientSecret string `url:"-" json:"client_secret"`
}

// PollCodeParams parameters required to poll the status of a device code.
type PollCodeParams struct {
	// BasicParams is the basic parameters which all requests can take.
	BasicParams

	// Code the generated code, retrieved from calling NewCode
	Code string `url:"-" json:"code"`
	// ClientSecret the client secret generated by trakt which is unique to our app.
	// this can be found in app settings. DO NOT EXPOSE THIS VALUE.
	ClientSecret string `url:"-" json:"client_secret"`

	// Interval the interval to poll on. This will be provided
	// in the response from NewCode, the API will throw errors
	// if you poll more frequently.
	Interval time.Duration `url:"-" json:"-"`
	// ExpiresIn the duration in which the code expires.
	// This will be provided in the response from NewCode.
	ExpiresIn time.Duration `url:"-" json:"-"`
}

// DeviceCode a representation of a device code.
type DeviceCode struct {
	// Code this is the code the status of.
	Code string `json:"device_code"`
	// UserCode this is the code that the user needs to enter to
	// confirm authorization.
	UserCode string `json:"user_code"`
	// VerificationURL the URL the user should navigate to to
	// enter their code.
	VerificationURL string `json:"verification_url"`
	// ExpiresIn the length of time expressed as a Duration in which
	// the device code request expires from the time is was requested.
	ExpiresIn time.Duration `json:"-"`
	// Interval the length of time expressed as a Duration in which
	// the system should poll the status of the authorization.
	Interval time.Duration `json:"-"`
}

// UnmarshalJSON implements Unmarshaller interface.
// allows us to convert the interval | expires in into time.Duration.
func (d *DeviceCode) UnmarshalJSON(bytes []byte) error {
	type B DeviceCode
	type A struct {
		B
		IntervalSecs  int64 `json:"interval"`
		ExpiresInSecs int64 `json:"expires_in"`
	}

	var a = new(A)
	err := json.Unmarshal(bytes, a)
	if err != nil {
		return err
	}

	a.B.Interval, err = parseSeconds(a.IntervalSecs)
	if err != nil {
		return err
	}

	a.B.ExpiresIn, err = parseSeconds(a.ExpiresInSecs)
	if err != nil {
		return err
	}

	*d = DeviceCode(a.B)
	return nil
}

// PollResult the result from polling the status of
// a device code.
//
// this will contain the completed token or an error
// if an error occurred which we could not recover from.
type PollResult struct {
	// Token the access token.
	Token *Token
	// Err any error that occurred polling the status
	// of the device code.
	Err error
}

// Token represents an access token.
// A token is usually valid for 3 months and can be
// refreshed using the refresh token.
// the expires in duration determines how long the
// token is valid for.
type Token struct {
	// AccessToken the generated access token key.
	// this is the code which is used with authenticated requests.
	AccessToken string `json:"access_token"`
	// Type the type of token, usually bearer.
	Type string `json:"token_type"`
	// Scope the scopes / level of access this token has to the
	// users account.
	Scope string `json:"scope"`
	// CreatedAt the time when the token was generated.
	CreatedAt time.Time `json:"-"`
	// RefreshToken this can be used to generate a new access token
	// after one expires without having to get the user to authenticate
	// again.
	RefreshToken string `json:"refresh_token"`
	// ExpiresIn the length of time this token is valid for.
	ExpiresIn time.Duration `json:"-"`
}

// UnmarshalJSON implements Unmarshaller interface.
// allows us to convert the created at | expires in into time.Duration.
func (t *Token) UnmarshalJSON(bytes []byte) error {
	type B Token
	type A struct {
		B
		CreatedAtUnix int64 `json:"created_at"`
		ExpiresInSecs int64 `json:"expires_in"`
	}

	var a = new(A)
	err := json.Unmarshal(bytes, a)
	if err != nil {
		return err
	}

	a.B.CreatedAt = time.Unix(a.CreatedAtUnix, 0)
	a.B.ExpiresIn, err = parseSeconds(a.ExpiresInSecs)
	if err != nil {
		return err
	}

	*t = Token(a.B)
	return nil
}

// parseSeconds helper function which converts int i which is assumed
// to be a number of seconds into a time.Duration.
func parseSeconds(i int64) (time.Duration, error) {
	return time.ParseDuration(strconv.Itoa(int(i)) + `s`)
}
