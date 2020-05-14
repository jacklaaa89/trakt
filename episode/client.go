package season

import (
	"net/http"

	"github.com/jackaaa89/trakt"
)

type Client struct{ b *trakt.BaseClient }

func Get(id trakt.SearchID, season, episode int64, params *trakt.ExtendedParams) (*trakt.Episode, error) {
	return getC().Get(id, season, episode, params)
}

func (c *Client) Get(id trakt.SearchID, season, episode int64, params *trakt.ExtendedParams) (*trakt.Episode, error) {
	path := trakt.FormatURLPath("/shows/%s/seasons/%s/episodes/%s", id, season, episode)
	ep := &trakt.Episode{}
	err := c.b.Call(http.MethodGet, path, params, ep)
	return ep, err
}

func Translations(id trakt.SearchID, season, episode int64, params *trakt.TranslationListParams) *trakt.TranslationIterator {
	return getC().Translations(id, season, episode, params)
}

func (c *Client) Translations(id trakt.SearchID, season, episode int64, params *trakt.TranslationListParams) *trakt.TranslationIterator {
	rcv := make([]*trakt.Translation, 0)
	path := trakt.FormatURLPath("/shows/%s/seasons/%s/episodes/%s/translations/%s", id, season, episode, params.Language)
	return &trakt.TranslationIterator{BasicIterator: c.b.NewSimulatedIterator(http.MethodGet, path, params, &rcv)}
}

func Comments(id trakt.SearchID, season, episode int64, params *trakt.CommentListParams) *trakt.CommentIterator {
	return getC().Comments(id, season, episode, params)
}

func (c *Client) Comments(id trakt.SearchID, season, episode int64, params *trakt.CommentListParams) *trakt.CommentIterator {
	rcv := make([]*trakt.Comment, 0)
	path := trakt.FormatURLPath(
		"shows/%s/seasons/%s/episodes/%s/comments/%s",
		id, season, episode, params.Sort,
	)
	return &trakt.CommentIterator{Iterator: c.b.NewIterator(http.MethodGet, path, params, &rcv)}
}

func Lists(id trakt.SearchID, season, episode int64, params *trakt.GetListParams) *trakt.ListIterator {
	return getC().Lists(id, season, episode, params)
}

func (c *Client) Lists(id trakt.SearchID, season, episode int64, params *trakt.GetListParams) *trakt.ListIterator {
	rcv := make([]*trakt.List, 0)

	path := trakt.FormatURLPath(
		"/shows/%s/seasons/%s/episodes/%s/lists/%s/%s",
		id, season, episode, params.ListType, params.SortType,
	)
	return &trakt.ListIterator{Iterator: c.b.NewIterator(http.MethodGet, path, params, &rcv)}
}

func People(id trakt.SearchID, season, episode int64, params *trakt.ExtendedParams) (*trakt.CastAndCrew, error) {
	return getC().People(id, season, episode, params)
}

func (c *Client) People(id trakt.SearchID, season, episode int64, params *trakt.ExtendedParams) (*trakt.CastAndCrew, error) {
	path := trakt.FormatURLPath("/shows/%s/seasons/%s/episodes/%s/people", id, season, episode)
	cc := &trakt.CastAndCrew{}
	err := c.b.Call(http.MethodGet, path, params, cc)
	return cc, err
}

func Ratings(id trakt.SearchID, season, episode int64, params *trakt.BasicParams) (*trakt.Rating, error) {
	return getC().Ratings(id, season, episode, params)
}

func (c *Client) Ratings(id trakt.SearchID, season, episode int64, params *trakt.BasicParams) (*trakt.Rating, error) {
	path := trakt.FormatURLPath("/shows/%s/seasons/%s/episodes/%s/ratings", id, season, episode)
	r := &trakt.Rating{}
	err := c.b.Call(http.MethodGet, path, params, r)
	return r, err
}

func Statistics(id trakt.SearchID, season, episode int64, params *trakt.BasicParams) (*trakt.Statistics, error) {
	return getC().Statistics(id, season, episode, params)
}

func (c *Client) Statistics(id trakt.SearchID, season, episode int64, params *trakt.BasicParams) (*trakt.Statistics, error) {
	path := trakt.FormatURLPath("/shows/%s/seasons/%s/episodes/%s/stats", id, season, episode)
	stats := &trakt.Statistics{}
	err := c.b.Call(http.MethodGet, path, params, stats)
	return stats, err
}

func WatchingNow(id trakt.SearchID, season, episode int64, params *trakt.BasicListParams) *trakt.UserIterator {
	return getC().WatchingNow(id, season, episode, params)
}

func (c *Client) WatchingNow(id trakt.SearchID, season, episode int64, params *trakt.BasicListParams) *trakt.UserIterator {
	rcv := make([]*trakt.User, 0)
	path := trakt.FormatURLPath("/shows/%s/seasons/%s/episodes/%s/watching", id, season, episode)
	return &trakt.UserIterator{Iterator: c.b.NewIterator(http.MethodGet, path, params, &rcv)}
}

func getC() *Client { return &Client{trakt.NewClient(trakt.GetBackend())} }
