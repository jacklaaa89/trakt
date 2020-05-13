// Package stripe provides the binding for Stripe REST APIs.
package trakt

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/google/go-querystring/query"
)

const (
	// APIVersion is the currently supported API version
	APIVersion string = "2"
	OAuthURL   string = "https://trakt.tv"
)

// Key is the Stripe API key used globally in the binding.
var Key string

// Backend is an interface for making calls against a Stripe service.
// This interface exists to enable mocking for during testing if needed.
type Backend interface {
	Call(method, path, key string, params ParamsContainer, v interface{}) error
	CallWithFrame(method, path, key string, params ParamsContainer, v IterationFrame) error
	SetMaxNetworkRetries(maxNetworkRetries int)
}

// BackendConfig is used to configure a new Stripe backend.
type BackendConfig struct {
	// HTTPClient is an HTTP client instance to use when making API requests.
	//
	// If left unset, it'll be set to a default HTTP client for the package.
	HTTPClient *http.Client

	// LeveledLogger is the logger that the backend will use to log errors,
	// warnings, and informational messages.
	//
	// LeveledLoggerInterface is implemented by LeveledLogger, and one can be
	// initialized at the desired level of logging.  LeveledLoggerInterface
	// also provides out-of-the-box compatibility with a Logrus Logger, but may
	// require a thin shim for use with other logging libraries that use less
	// standard conventions like Zap.
	LeveledLogger LeveledLoggerInterface

	// MaxNetworkRetries sets maximum number of times that the library will
	// retry requests that appear to have failed due to an intermittent
	// problem.
	//
	// Defaults to 0.
	MaxNetworkRetries int

	// URL is the base URL to use for API paths.
	//
	// If left empty, it'll be set to the default for the SupportedBackend.
	URL string
}

// BackendImplementation is the internal implementation for making HTTP calls
// to Stripe.
//
// The public use of this struct is deprecated. It will be unexported in a
// future version.
type BackendImplementation struct {
	URL               string
	HTTPClient        *http.Client
	LeveledLogger     LeveledLoggerInterface
	MaxNetworkRetries int

	// networkRetriesSleep indicates whether the backend should use the normal
	// sleep between retries.
	//
	// See also SetNetworkRetriesSleep.
	networkRetriesSleep bool
}

// CallWithFrame called for pagination type functions where we are typically querying for a single frame / segment
// in the entire set. This puts a higher level of abstraction above that of the basic Call function.
func (s *BackendImplementation) CallWithFrame(method, path, key string, params ParamsContainer, v IterationFrame) error {
	if v == nil {
		return errors.New(`frame data is required`)
	}

	rcv := v.rcv()
	vo := reflect.ValueOf(rcv)
	if !vo.IsValid() {
		return errors.New(`frame receiver is invalid`)
	}

	if vo.Kind() != reflect.Ptr {
		return errors.New(`frame receiver has to be a pointer`)
	}

	return s.callRaw(method, path, key, params, rcv, v)
}

func (s *BackendImplementation) Call(method, path, key string, params ParamsContainer, v interface{}) error {
	return s.callRaw(method, path, key, params, v, nil)
}

// CallRaw is the implementation for invoking Stripe APIs internally without a backend.
func (s *BackendImplementation) callRaw(method, path, key string, params ParamsContainer, v, h interface{}) error {
	var body []byte
	if isHTTPWriteMethod(method) {
		var encodeErr error
		// perform a JSON encode of the q as the body.
		// the tags defined will dictate what is set in the body JSON
		body, encodeErr = json.Marshal(params)
		if encodeErr != nil {
			return encodeErr
		}
	} else {
		uv, err := query.Values(params)
		if err != nil {
			return err
		}

		path += `?` + uv.Encode()
	}

	bodyBuffer := bytes.NewBuffer(body)

	req, err := s.newRequest(method, path, key, "application/json", params)
	if err != nil {
		return err
	}

	if err := s.do(req, bodyBuffer, params, v, h); err != nil {
		return err
	}

	return nil
}

// NewRequest is used by Call to generate an http.Request. It handles encoding
// parameters and attaching the appropriate headers.
func (s *BackendImplementation) newRequest(method, path, key, contentType string, params ParamsContainer) (*http.Request, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = s.URL + path

	// Body is set later by `Do`.
	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		s.LeveledLogger.Errorf("Cannot create trakt request: %v", err)
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	req.Header.Add("User-Agent", encodedUserAgent)

	// add trakt specific headers.
	req.Header.Add("trakt-api-version", APIVersion)
	req.Header.Add("trakt-api-key", key)

	// add oauth token if supplied.
	token := params.oauth()
	if token != "" {
		authorization := "Bearer " + token
		req.Header.Add("Authorization", authorization)
	}

	if params != nil {
		req = req.WithContext(params.context())

		for k, v := range params.headers() {
			for _, line := range v {
				// Use Set to override the default value possibly set before
				req.Header.Set(k, line)
			}
		}
	}

	return req, nil
}

// Do is used by Call to execute an API request and parse the response. It uses
// the backend's HTTP client to execute the request and unmarshals the response
// into v. It also handles unmarshaling errors returned by the API.
func (s *BackendImplementation) do(req *http.Request, body *bytes.Buffer, params ParamsContainer, v, h interface{}) error {
	s.LeveledLogger.Infof("Requesting %v %v%v\n", req.Method, req.URL.Host, req.URL.Path)

	var res *http.Response
	var err error
	var requestDuration time.Duration
	var resBody []byte
	for retry := 0; ; {
		start := time.Now()

		if body != nil {
			reader := bytes.NewReader(body.Bytes())

			req.Body = nopReadCloser{reader}
			req.GetBody = func() (io.ReadCloser, error) {
				reader := bytes.NewReader(body.Bytes())
				return nopReadCloser{reader}, nil
			}
		}

		res, err = s.HTTPClient.Do(req)

		requestDuration = time.Since(start)
		s.LeveledLogger.Infof("Request completed in %v (retry: %v)", requestDuration, retry)

		if err == nil {
			resBody, err = ioutil.ReadAll(res.Body)
			res.Body.Close()
		}

		if err != nil {
			s.LeveledLogger.Errorf("Request failed with error: %v", err)
		} else if res.StatusCode >= 400 {
			err = s.responseToError(res, resBody, params)
		}

		// If the response was okay, or an error that shouldn't be retried,
		// we're done, and it's safe to leave the retry loop.
		if !s.shouldRetry(err, req, res, retry) {
			break
		}

		sleepDuration := s.sleepTime(retry)
		retry++

		s.LeveledLogger.Warnf("Initiating retry %v for request %v %v%v after sleeping %v",
			retry, req.Method, req.URL.Host, req.URL.Path, sleepDuration)

		time.Sleep(sleepDuration)
	}

	if err != nil {
		return err
	}

	s.LeveledLogger.Debugf("Response: %s\n", string(resBody))

	if v != nil {
		return s.unmarshalJSONVerbose(res, resBody, v, h)
	}

	return nil
}

// responseToError converts a stripe response to an Error.
func (s *BackendImplementation) responseToError(res *http.Response, resBody []byte, p ParamsContainer) error {
	var errorHandler = DefaultErrorHandler
	switch e := p.(type) {
	case ErrorHandler:
		errorHandler = e
	}

	return &Error{
		HTTPStatusCode: res.StatusCode,
		RequestID:      res.Header.Get(`X-Request-ID`),
		Body:           string(resBody),
		Resource:       res.Request.URL.Path,
		Code:           errorHandler.Code(res.StatusCode),
	}
}

// SetMaxNetworkRetries sets max number of retries on failed requests
//
// This function is deprecated. Please use GetBackendWithConfig instead.
func (s *BackendImplementation) SetMaxNetworkRetries(maxNetworkRetries int) {
	s.MaxNetworkRetries = maxNetworkRetries
}

// unmarshalJSONVerbose unmarshals JSON, but in case of a failure logs and
// produces a more descriptive error.
func (s *BackendImplementation) unmarshalJSONVerbose(res *http.Response, body []byte, v, h interface{}) error {
	err := json.Unmarshal(body, v)
	if err != nil {
		bodySample := string(body)
		if len(bodySample) > 500 {
			bodySample = bodySample[0:500] + " ..."
		}

		// Make sure a multi-line response ends up all on one line
		bodySample = strings.Replace(bodySample, "\n", "\\n", -1)

		newErr := fmt.Errorf(
			"couldn't deserialize JSON (response status: %v, body sample: '%s'): %v",
			res.StatusCode, bodySample, err,
		)
		s.LeveledLogger.Errorf("%s", newErr.Error())
		return newErr
	}

	// perform some custom unmarshal's if required.
	// we always want to perform JSON unmarshalling to.
	// this is just to grab additional information that is
	// not just provided in the response body.
	switch m := h.(type) {
	case headerUnmarshaller:
		err = m.UnmarshalHeaders(res.Header)
	}

	if err != nil {
		s.LeveledLogger.Errorf("%s", err)
	}
	return err
}

// Checks if an error is a problem that we should retry on. This includes both
// socket errors that may represent an intermittent problem and some special
// HTTP statuses.
func (s *BackendImplementation) shouldRetry(err error, req *http.Request, resp *http.Response, numRetries int) bool {
	if numRetries >= s.MaxNetworkRetries {
		return false
	}

	traktErr, _ := err.(*Error)

	if traktErr == nil && err != nil {
		return true
	}

	// 409 Conflict
	if resp.StatusCode == http.StatusConflict {
		return true
	}

	// 429 Too Many Requests
	//
	// There are a few different problems that can lead to a 429. The most
	// common is rate limiting, on which we *don't* want to retry because
	// that'd likely contribute to more contention problems. However, some 429s
	// are lock timeouts, which is when a request conflicted with another
	// request or an internal process on some particular object. These 429s are
	// safe to retry.
	if resp.StatusCode == http.StatusTooManyRequests {
		return false
	}

	// 500 Internal Server Error
	if resp.StatusCode >= http.StatusInternalServerError {
		return true
	}

	// 503 Service Unavailable
	if resp.StatusCode == http.StatusServiceUnavailable {
		return true
	}

	// 503 Gateway Timeout
	if resp.StatusCode == http.StatusGatewayTimeout {
		return true
	}

	// if we have a cloudflare error (which dont use standard HTTP status codes)
	if isCloudflareError(resp.StatusCode) {
		return true
	}

	return false
}

func isCloudflareError(code int) bool {
	return code == 520 || code == 521 || code == 522
}

// sleepTime calculates sleeping/delay time in milliseconds between failure and a new one request.
func (s *BackendImplementation) sleepTime(numRetries int) time.Duration {
	// We disable sleeping in some cases for tests.
	if !s.networkRetriesSleep {
		return 0 * time.Second
	}

	// Apply exponential backoff with minNetworkRetriesDelay on the
	// number of num_retries so far as inputs.
	delay := minNetworkRetriesDelay + minNetworkRetriesDelay*time.Duration(numRetries*numRetries)

	// Do not allow the number to exceed maxNetworkRetriesDelay.
	if delay > maxNetworkRetriesDelay {
		delay = maxNetworkRetriesDelay
	}

	// Apply some jitter by randomizing the value in the range of 75%-100%.
	jitter := rand.Int63n(int64(delay / 4))
	delay -= time.Duration(jitter)

	// But never sleep less than the base sleep seconds.
	if delay < minNetworkRetriesDelay {
		delay = minNetworkRetriesDelay
	}

	return delay
}

// Backends are the currently supported endpoints.
type Backends struct {
	API Backend
	mu  sync.RWMutex
}

// Bool returns a pointer to the bool value passed in.
func Bool(v bool) *bool {
	return &v
}

// BoolValue returns the value of the bool pointer passed in or
// false if the pointer is nil.
func BoolValue(v *bool) bool {
	if v != nil {
		return *v
	}
	return false
}

// BoolSlice returns a slice of bool pointers given a slice of bools.
func BoolSlice(v []bool) []*bool {
	out := make([]*bool, len(v))
	for i := range v {
		out[i] = &v[i]
	}
	return out
}

// Float64 returns a pointer to the float64 value passed in.
func Float64(v float64) *float64 {
	return &v
}

// Float64Value returns the value of the float64 pointer passed in or
// 0 if the pointer is nil.
func Float64Value(v *float64) float64 {
	if v != nil {
		return *v
	}
	return 0
}

// Float64Slice returns a slice of float64 pointers given a slice of float64s.
func Float64Slice(v []float64) []*float64 {
	out := make([]*float64, len(v))
	for i := range v {
		out[i] = &v[i]
	}
	return out
}

// FormatURLPath takes a format string (of the kind used in the fmt package)
// representing a URL path with a number of parameters that belong in the path
// and returns a formatted string.
//
// we perform a quick type comparison to handle some well-known types
// including a SearchID or a fmt.Stringer or string, it any of the
// parameters are not of these types, then we ignore the value
// as it should only be of of these three types.
// this would cause come very noticeable formatting errors.
func FormatURLPath(format string, params ...interface{}) string {
	// Convert parameters to interface{} and URL-escape them
	untypedParams := make([]interface{}, len(params))
	var i int
	for _, param := range params {
		var formatted string
		switch p := param.(type) {
		case string:
			formatted = p
		case fmt.Stringer:
			formatted = p.String()
		case SearchID:
			formatted = p.id()
		default:
			continue
		}

		untypedParams[i] = interface{}(url.QueryEscape(formatted))
		i++
	}

	return fmt.Sprintf(format, untypedParams[:i]...)
}

// GetBackend returns one of the library's supported backends based off of the
// given argument.
//
// It returns an existing default backend if one's already been created.
func GetBackend() Backend {
	var backend Backend

	backends.mu.RLock()
	backend = backends.API
	backends.mu.RUnlock()

	if backend != nil {
		return backend
	}

	backend = GetBackendWithConfig(
		&BackendConfig{
			HTTPClient:        httpClient,
			LeveledLogger:     DefaultLeveledLogger,
			MaxNetworkRetries: 0,
			URL:               "", // Set by GetBackendWithConfiguation when empty
		},
	)

	SetBackend(backend)

	return backend
}

// GetBackendWithConfig is the same as GetBackend except that it can be given a
// configuration struct that will configure certain aspects of the backend
// that's return.
func GetBackendWithConfig(config *BackendConfig) Backend {
	if config.HTTPClient == nil {
		config.HTTPClient = httpClient
	}

	if config.LeveledLogger == nil {
		config.LeveledLogger = DefaultLeveledLogger
	}

	if config.URL == "" {
		config.URL = apiURL
	}

	config.URL = normalizeURL(config.URL)

	return newBackendImplementation(config)
}

// Int64 returns a pointer to the int64 value passed in.
func Int64(v int64) *int64 {
	return &v
}

// Int64Value returns the value of the int64 pointer passed in or
// 0 if the pointer is nil.
func Int64Value(v *int64) int64 {
	if v != nil {
		return *v
	}
	return 0
}

// Int64Slice returns a slice of int64 pointers given a slice of int64s.
func Int64Slice(v []int64) []*int64 {
	out := make([]*int64, len(v))
	for i := range v {
		out[i] = &v[i]
	}
	return out
}

// SetBackend sets the backend used in the binding.
func SetBackend(b Backend) {
	backends.mu.Lock()
	defer backends.mu.Unlock()
	backends.API = b
}

// SetHTTPClient overrides the default HTTP client.
// This is useful if you're running in a Google AppEngine environment
// where the http.DefaultClient is not available.
func SetHTTPClient(client *http.Client) {
	httpClient = client
}

// String returns a pointer to the string value passed in.
func String(v string) *string {
	return &v
}

// StringValue returns the value of the string pointer passed in or
// "" if the pointer is nil.
func StringValue(v *string) string {
	if v != nil {
		return *v
	}
	return ""
}

// StringSlice returns a slice of string pointers given a slice of strings.
func StringSlice(v []string) []*string {
	out := make([]*string, len(v))
	for i := range v {
		out[i] = &v[i]
	}
	return out
}

const apiURL = "https://api.trakt.tv"

// clientVersion is the binding version
const clientVersion = "0.1.0"

// defaultHTTPTimeout is the default timeout on the http.Client used by the library.
// This is chosen to be consistent with the other Stripe language libraries and
// to coordinate with other timeouts configured in the Stripe infrastructure.
const defaultHTTPTimeout = 80 * time.Second

// maxNetworkRetriesDelay and minNetworkRetriesDelay defines sleep time in milliseconds between
// tries to send HTTP request again after network failure.
const maxNetworkRetriesDelay = 5000 * time.Millisecond
const minNetworkRetriesDelay = 500 * time.Millisecond

// nopReadCloser's sole purpose is to give us a way to turn an `io.Reader` into
// an `io.ReadCloser` by adding a no-op implementation of the `Closer`
// interface. We need this because `http.Request`'s `Body` takes an
// `io.ReadCloser` instead of a `io.Reader`.
type nopReadCloser struct {
	io.Reader
}

func (nopReadCloser) Close() error { return nil }

// headerUnmarshaller interface which is used to inform the unmarshaller
// that we need to unmarshal this type using the response headers.
type headerUnmarshaller interface {
	UnmarshalHeaders(h http.Header) error
}

var backends Backends
var encodedUserAgent string

// The default HTTP client used for communication with any of Stripe's
// backends.
//
// Can be overridden with the function `SetHTTPClient` or by setting the
// `HTTPClient` value when using `BackendConfig`.
var httpClient = &http.Client{
	Timeout: defaultHTTPTimeout,
}

//
// Private functions
//

func init() {
	initUserAgent()
}

func initUserAgent() {
	encodedUserAgent = "Trakt/v1 GoBindings/" + clientVersion
}

// newBackendImplementation returns a new Backend based off a given type and
// fully initialized BackendConfig struct.
//
// The vast majority of the time you should be calling GetBackendWithConfig
// instead of this function.
func newBackendImplementation(config *BackendConfig) Backend {
	return &BackendImplementation{
		HTTPClient:          config.HTTPClient,
		LeveledLogger:       config.LeveledLogger,
		MaxNetworkRetries:   config.MaxNetworkRetries,
		URL:                 config.URL,
		networkRetriesSleep: true,
	}
}

func isHTTPWriteMethod(method string) bool {
	return method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch || method == http.MethodDelete
}

func normalizeURL(url string) string {
	// All paths include a leading slash, so to keep logs pretty, trim a
	// trailing slash on the URL.
	url = strings.TrimSuffix(url, "/")

	// For a long time we had the `/v1` suffix as part of a configured URL
	// rather than in the per-package URLs throughout the library. Continue
	// to support this for the time being by stripping one that's been
	// passed for better backwards compatibility.
	url = strings.TrimSuffix(url, "/v1")

	return url
}
