package movie

import (
	"net/http"

	"github.com/jackaaa89/trakt"
)

type Client struct{ b *trakt.BaseClient }

func Trending(params *trakt.FilterListParams) *trakt.TrendingMovieIterator {
	return getC().Trending(params)
}

func (c *Client) Trending(params *trakt.FilterListParams) *trakt.TrendingMovieIterator {
	return &trakt.TrendingMovieIterator{Iterator: c.b.NewIterator(http.MethodGet, "/movies/trending", params)}
}

func Popular(params *trakt.FilterListParams) *trakt.MovieIterator {
	return getC().Popular(params)
}

func (c *Client) Popular(params *trakt.FilterListParams) *trakt.MovieIterator {
	return &trakt.MovieIterator{Iterator: c.b.NewIterator(http.MethodGet, "/movies/popular", params)}
}

func Played(params *trakt.TimePeriodListParams) *trakt.MovieWithStatisticsIterator {
	return getC().Played(params)
}

func (c *Client) Played(params *trakt.TimePeriodListParams) *trakt.MovieWithStatisticsIterator {
	return c.newTimePeriodIterator("played", params)
}

func Watched(params *trakt.TimePeriodListParams) *trakt.MovieWithStatisticsIterator {
	return getC().Watched(params)
}

func (c *Client) Watched(params *trakt.TimePeriodListParams) *trakt.MovieWithStatisticsIterator {
	return c.newTimePeriodIterator("watched", params)
}

func Collected(params *trakt.TimePeriodListParams) *trakt.MovieWithStatisticsIterator {
	return getC().Collected(params)
}

func (c *Client) Collected(params *trakt.TimePeriodListParams) *trakt.MovieWithStatisticsIterator {
	return c.newTimePeriodIterator("collected", params)
}

func Anticipated(params *trakt.FilterListParams) *trakt.AnticipatedMovieIterator {
	return getC().Anticipated(params)
}

func (c *Client) Anticipated(params *trakt.FilterListParams) *trakt.AnticipatedMovieIterator {
	return &trakt.AnticipatedMovieIterator{Iterator: c.b.NewIterator(http.MethodGet, "/movies/anticipated", params)}
}

func BoxOffice(params *trakt.BoxOfficeListParams) *trakt.BoxOfficeMovieIterator {
	return getC().BoxOffice(params)
}

func (c *Client) BoxOffice(params *trakt.BoxOfficeListParams) *trakt.BoxOfficeMovieIterator {
	return &trakt.BoxOfficeMovieIterator{
		BasicIterator: c.b.NewSimulatedIterator(http.MethodGet, "/movies/boxoffice", params),
	}
}

func RecentlyUpdated(params *trakt.RecentlyUpdatedListParams) *trakt.RecentlyUpdatedMovieIterator {
	return getC().RecentlyUpdated(params)
}

func (c *Client) RecentlyUpdated(params *trakt.RecentlyUpdatedListParams) *trakt.RecentlyUpdatedMovieIterator {
	path := trakt.FormatURLPath("/movies/updates/%s", params.StartDate.Format(`2006-01-02`))
	return &trakt.RecentlyUpdatedMovieIterator{Iterator: c.b.NewIterator(http.MethodGet, path, params)}
}

func Get(id trakt.SearchID, params *trakt.ExtendedParams) (*trakt.Movie, error) {
	return getC().Get(id, params)
}

func (c *Client) Get(id trakt.SearchID, params *trakt.ExtendedParams) (*trakt.Movie, error) {
	path := trakt.FormatURLPath("/movies/%s", id)
	mov := &trakt.Movie{}
	err := c.b.Call(http.MethodGet, path, params, mov)
	return mov, err
}

func Aliases(id trakt.SearchID, params *trakt.BasicParams) *trakt.AliasIterator {
	return getC().Aliases(id, params)
}

func (c *Client) Aliases(id trakt.SearchID, params *trakt.BasicParams) *trakt.AliasIterator {
	path := trakt.FormatURLPath("movies/%s/aliases", id)
	return &trakt.AliasIterator{BasicIterator: c.b.NewSimulatedIterator(http.MethodGet, path, params)}
}

func Releases(id trakt.SearchID, params *trakt.ReleaseListParams) *trakt.ReleaseIterator {
	return getC().Releases(id, params)
}

func (c *Client) Releases(id trakt.SearchID, params *trakt.ReleaseListParams) *trakt.ReleaseIterator {
	path := trakt.FormatURLPath("movies/%s/releases/%s", id, params.Country)
	return &trakt.ReleaseIterator{BasicIterator: c.b.NewSimulatedIterator(http.MethodGet, path, params)}
}

func Translations(id trakt.SearchID, params *trakt.TranslationListParams) *trakt.TranslationIterator {
	return getC().Translations(id, params)
}

func (c *Client) Translations(id trakt.SearchID, params *trakt.TranslationListParams) *trakt.TranslationIterator {
	path := trakt.FormatURLPath("movies/%s/translations/%s", id, params.Language)
	return &trakt.TranslationIterator{BasicIterator: c.b.NewSimulatedIterator(http.MethodGet, path, params)}
}

func Comments(id trakt.SearchID, params *trakt.CommentListParams) *trakt.CommentIterator {
	return getC().Comments(id, params)
}

func (c *Client) Comments(id trakt.SearchID, params *trakt.CommentListParams) *trakt.CommentIterator {
	path := trakt.FormatURLPath("movies/%s/comments/%s", id, params.Sort)
	return &trakt.CommentIterator{Iterator: c.b.NewIterator(http.MethodGet, path, params)}
}

func WatchingNow(id trakt.SearchID, params *trakt.BasicListParams) *trakt.UserIterator {
	return getC().WatchingNow(id, params)
}

func (c *Client) WatchingNow(id trakt.SearchID, params *trakt.BasicListParams) *trakt.UserIterator {
	path := trakt.FormatURLPath("movies/%s/watching", id)
	return &trakt.UserIterator{Iterator: c.b.NewIterator(http.MethodGet, path, params)}
}

func Related(id trakt.SearchID, params *trakt.ExtendedListParams) *trakt.MovieIterator {
	return getC().Related(id, params)
}

func (c *Client) Related(id trakt.SearchID, params *trakt.ExtendedListParams) *trakt.MovieIterator {
	path := trakt.FormatURLPath("movies/%s/related", id)
	return &trakt.MovieIterator{Iterator: c.b.NewIterator(http.MethodGet, path, params)}
}

func Ratings(id trakt.SearchID, params *trakt.BasicParams) (*trakt.RatingDistribution, error) {
	return getC().Ratings(id, params)
}

func (c *Client) Ratings(id trakt.SearchID, params *trakt.BasicParams) (*trakt.RatingDistribution, error) {
	path := trakt.FormatURLPath("/movies/%s/ratings", id)
	stats := &trakt.RatingDistribution{}
	err := c.b.Call(http.MethodGet, path, params, stats)
	return stats, err
}

func Statistics(id trakt.SearchID, params *trakt.BasicParams) (*trakt.Statistics, error) {
	return getC().Statistics(id, params)
}

func (c *Client) Statistics(id trakt.SearchID, params *trakt.BasicParams) (*trakt.Statistics, error) {
	path := trakt.FormatURLPath("/movies/%s/stats", id)
	stats := &trakt.Statistics{}
	err := c.b.Call(http.MethodGet, path, params, stats)
	return stats, err
}

func Lists(id trakt.SearchID, params *trakt.GetListParams) *trakt.ListIterator {
	return getC().Lists(id, params)
}

func (c *Client) Lists(id trakt.SearchID, params *trakt.GetListParams) *trakt.ListIterator {
	path := trakt.FormatURLPath("movies/%s/lists/%s/%s", id, params.ListType, params.SortType)
	return &trakt.ListIterator{Iterator: c.b.NewIterator(http.MethodGet, path, params)}
}

func People(id trakt.SearchID, params *trakt.ExtendedParams) (*trakt.CastAndCrew, error) {
	return getC().People(id, params)
}

func (c *Client) People(id trakt.SearchID, params *trakt.ExtendedParams) (*trakt.CastAndCrew, error) {
	path := trakt.FormatURLPath("/movies/%s/people", id)
	cc := &trakt.CastAndCrew{}
	err := c.b.Call(http.MethodGet, path, params, cc)
	return cc, err
}

func (c *Client) newTimePeriodIterator(action string, p *trakt.TimePeriodListParams) *trakt.MovieWithStatisticsIterator {
	var period = trakt.TimePeriodAll
	if p.Period != "" {
		period = p.Period
	}
	path := trakt.FormatURLPath("/movies/%s/%s", action, period)
	return &trakt.MovieWithStatisticsIterator{Iterator: c.b.NewIterator(http.MethodGet, path, p)}
}

func getC() *Client { return &Client{trakt.NewClient(trakt.GetBackend())} }
