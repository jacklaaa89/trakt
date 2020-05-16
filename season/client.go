package season

import (
	"net/http"

	"github.com/jacklaaa89/trakt"
)

type Client struct{ b trakt.BaseClient }

func Episodes(id trakt.SearchID, season int64, params *trakt.EpisodeListParams) *trakt.EpisodeWithTranslationsIterator {
	return getC().Episodes(id, season, params)
}

func (c *Client) Episodes(id trakt.SearchID, season int64, params *trakt.EpisodeListParams) *trakt.EpisodeWithTranslationsIterator {
	path := trakt.FormatURLPath("/shows/%s/seasons/%s", id, season)
	return &trakt.EpisodeWithTranslationsIterator{
		BasicIterator: c.b.NewSimulatedIterator(http.MethodGet, path, params),
	}
}

func Comments(id trakt.SearchID, season int64, params *trakt.CommentListParams) *trakt.CommentIterator {
	return getC().Comments(id, season, params)
}

func (c *Client) Comments(id trakt.SearchID, season int64, params *trakt.CommentListParams) *trakt.CommentIterator {
	path := trakt.FormatURLPath("shows/%s/seasons/%s/comments/%s", id, season, params.Sort)
	return &trakt.CommentIterator{Iterator: c.b.NewIterator(http.MethodGet, path, params)}
}

func Lists(id trakt.SearchID, season int64, params *trakt.GetListParams) *trakt.ListIterator {
	return getC().Lists(id, season, params)
}

func (c *Client) Lists(id trakt.SearchID, season int64, params *trakt.GetListParams) *trakt.ListIterator {
	path := trakt.FormatURLPath(
		"/shows/%s/seasons/%s/lists/%s/%s", id, season, params.ListType, params.SortType,
	)
	return &trakt.ListIterator{Iterator: c.b.NewIterator(http.MethodGet, path, params)}
}

func People(id trakt.SearchID, season int64, params *trakt.ExtendedParams) (*trakt.CastAndCrew, error) {
	return getC().People(id, season, params)
}

func (c *Client) People(id trakt.SearchID, season int64, params *trakt.ExtendedParams) (*trakt.CastAndCrew, error) {
	path := trakt.FormatURLPath("/shows/%s/seasons/%s/people", id, season)
	cc := &trakt.CastAndCrew{}
	err := c.b.Call(http.MethodGet, path, params, cc)
	return cc, err
}

func Ratings(id trakt.SearchID, season int64, params *trakt.BasicParams) (*trakt.RatingDistribution, error) {
	return getC().Ratings(id, season, params)
}

func (c *Client) Ratings(id trakt.SearchID, season int64, params *trakt.BasicParams) (*trakt.RatingDistribution, error) {
	path := trakt.FormatURLPath("/shows/%s/seasons/%s/ratings", id, season)
	r := &trakt.RatingDistribution{}
	err := c.b.Call(http.MethodGet, path, params, r)
	return r, err
}

func Statistics(id trakt.SearchID, season int64, params *trakt.BasicParams) (*trakt.Statistics, error) {
	return getC().Statistics(id, season, params)
}

func (c *Client) Statistics(id trakt.SearchID, season int64, params *trakt.BasicParams) (*trakt.Statistics, error) {
	path := trakt.FormatURLPath("/shows/%s/seasons/%s/stats", id, season)
	stats := &trakt.Statistics{}
	err := c.b.Call(http.MethodGet, path, params, stats)
	return stats, err
}

func WatchingNow(id trakt.SearchID, season int64, params *trakt.BasicListParams) *trakt.UserIterator {
	return getC().WatchingNow(id, season, params)
}

func (c *Client) WatchingNow(id trakt.SearchID, season int64, params *trakt.BasicListParams) *trakt.UserIterator {
	path := trakt.FormatURLPath("/shows/%s/seasons/%s/watching", id, season)
	return &trakt.UserIterator{Iterator: c.b.NewIterator(http.MethodGet, path, params)}
}

func getC() *Client { return &Client{trakt.NewClient()} }
