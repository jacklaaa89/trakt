package trakt

import (
	"encoding/json"
	"strconv"
	"time"
)

type AuthorizationURLParams struct {
	BasicParams
	RedirectURI string `url:"redirect_uri" json:"-"`
	State       string `url:"state" json:"-"`
}

type ExchangeCodeParams struct {
	BasicParams
	RedirectURI  string `url:"-" json:"redirect_uri"`
	Code         string `url:"-" json:"code"`
	ClientSecret string `url:"-" json:"client_secret"`
}

type RefreshTokenParams struct {
	BasicParams
	RedirectURI  string `url:"-" json:"redirect_uri"`
	RefreshToken string `url:"-" json:"refresh_token"`
	ClientSecret string `url:"-" json:"client_secret"`
}

type RevokeTokenParams struct {
	BasicParams
	AccessToken  string `url:"-" json:"token"`
	ClientSecret string `url:"-" json:"client_secret"`
}

type PollCodeParams struct {
	BasicParams
	Code         string        `url:"-" json:"code"`
	ClientSecret string        `url:"-" json:"client_secret"`
	Interval     time.Duration `url:"-" json:"-"`
	ExpiresIn    time.Duration `url:"-" json:"-"`
}

type DeviceCode struct {
	Code            string        `json:"device_code"`
	UserCode        string        `json:"user_code"`
	VerificationURL string        `json:"verification_url"`
	ExpiresIn       time.Duration `json:"-"`
	Interval        time.Duration `json:"-"`
}

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

type PollResult struct {
	Token *Token
	Err   error
}

type Token struct {
	AccessToken  string        `json:"access_token"`
	Type         string        `json:"token_type"`
	Scope        string        `json:"scope"`
	CreatedAt    time.Time     `json:"-"`
	RefreshToken string        `json:"refresh_token"`
	ExpiresIn    time.Duration `json:"-"`
}

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

func parseSeconds(i int64) (time.Duration, error) {
	return time.ParseDuration(strconv.Itoa(int(i)) + `s`)
}
