package movie

import (
	"net/http"

	"github.com/jackaaa89/trakt"
)

// Lists
// People

type Client struct {
	B   trakt.Backend
	Key string
}

func Trending(params *trakt.ExtendedListParams) *trakt.TrendingMovieIterator {
	return getC().Trending(params)
}

func (c *Client) Trending(params *trakt.ExtendedListParams) *trakt.TrendingMovieIterator {
	return &trakt.TrendingMovieIterator{
		Iterator: trakt.NewIterator(params, func(p trakt.ListParamsContainer) (trakt.IterationFrame, error) {
			tml := make([]*trakt.TrendingMovie, 0)
			f := trakt.NewEmptyFrame(&tml)
			err := c.B.CallWithFrame(http.MethodGet, "/movies/trending", c.Key, p, f)
			return f, err
		}),
	}
}

func Popular(params *trakt.ExtendedListParams) *trakt.MovieIterator {
	return getC().Popular(params)
}

func (c *Client) Popular(params *trakt.ExtendedListParams) *trakt.MovieIterator {
	return &trakt.MovieIterator{
		Iterator: trakt.NewIterator(params, func(p trakt.ListParamsContainer) (trakt.IterationFrame, error) {
			rcv := make([]*trakt.Movie, 0)
			f := trakt.NewEmptyFrame(&rcv)
			err := c.B.CallWithFrame(http.MethodGet, "/movies/popular", c.Key, p, f)
			return f, err
		}),
	}
}

func Played(params *trakt.TimePeriodListParams) *trakt.MovieWithStatisticsIterator {
	return getC().Played(params)
}

func (c *Client) Played(params *trakt.TimePeriodListParams) *trakt.MovieWithStatisticsIterator {
	return c.generateTimePeriodMovieListIterator("played", params)
}

func Watched(params *trakt.TimePeriodListParams) *trakt.MovieWithStatisticsIterator {
	return getC().Watched(params)
}

func (c *Client) Watched(params *trakt.TimePeriodListParams) *trakt.MovieWithStatisticsIterator {
	return c.generateTimePeriodMovieListIterator("watched", params)
}

func Collected(params *trakt.TimePeriodListParams) *trakt.MovieWithStatisticsIterator {
	return getC().Collected(params)
}

func (c *Client) Collected(params *trakt.TimePeriodListParams) *trakt.MovieWithStatisticsIterator {
	return c.generateTimePeriodMovieListIterator("collected", params)
}

func Anticipated(params *trakt.ExtendedListParams) *trakt.AnticipatedMovieIterator {
	return getC().Anticipated(params)
}

func (c *Client) Anticipated(params *trakt.ExtendedListParams) *trakt.AnticipatedMovieIterator {
	return &trakt.AnticipatedMovieIterator{
		Iterator: trakt.NewIterator(params, func(p trakt.ListParamsContainer) (trakt.IterationFrame, error) {
			rcv := make([]*trakt.AnticipatedMovie, 0)
			f := trakt.NewEmptyFrame(&rcv)
			err := c.B.CallWithFrame(http.MethodGet, "/movies/anticipated", c.Key, p, f)
			return f, err
		}),
	}
}

func BoxOffice(params *trakt.BoxOfficeListParams) *trakt.BoxOfficeMovieIterator {
	return getC().BoxOffice(params)
}

func (c *Client) BoxOffice(params *trakt.BoxOfficeListParams) *trakt.BoxOfficeMovieIterator {
	return &trakt.BoxOfficeMovieIterator{
		Iterator: trakt.NewIterator(params, func(p trakt.ListParamsContainer) (trakt.IterationFrame, error) {
			rcv := make([]*trakt.BoxOfficeMovie, 0)
			f := trakt.NewEmptyFrame(&rcv)
			err := c.B.CallWithFrame(http.MethodGet, "/movies/boxoffice", c.Key, p, f)
			return f, err
		}),
	}
}

func RecentlyUpdated(params *trakt.RecentlyUpdatedListParams) *trakt.RecentlyUpdatedMovieIterator {
	return getC().RecentlyUpdated(params)
}

func (c *Client) RecentlyUpdated(params *trakt.RecentlyUpdatedListParams) *trakt.RecentlyUpdatedMovieIterator {
	return &trakt.RecentlyUpdatedMovieIterator{
		Iterator: trakt.NewIterator(params, func(p trakt.ListParamsContainer) (trakt.IterationFrame, error) {
			rcv := make([]*trakt.RecentlyUpdatedMovie, 0)
			f := trakt.NewEmptyFrame(&rcv)
			err := c.B.CallWithFrame(
				http.MethodGet,
				trakt.FormatURLPath(
					"/movies/updates/%s",
					params.StartDate.Format(`2006-01-02`),
				),
				c.Key, p, f,
			)
			return f, err
		}),
	}
}

func Get(id trakt.SearchID, params *trakt.ExtendedParams) (*trakt.Movie, error) {
	return getC().Get(id, params)
}

func (c *Client) Get(id trakt.SearchID, params *trakt.ExtendedParams) (*trakt.Movie, error) {
	path := trakt.FormatURLPath("/movies/%s", id)
	mov := &trakt.Movie{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, mov)
	return mov, err
}

func Aliases(id trakt.SearchID, params *trakt.BasicParams) *trakt.AliasIterator {
	return getC().Aliases(id, params)
}

func (c *Client) Aliases(id trakt.SearchID, params *trakt.BasicParams) *trakt.AliasIterator {
	return &trakt.AliasIterator{
		Iterator: trakt.NewSimulatedIterator(params, func(p trakt.ListParamsContainer) (trakt.IterationFrame, error) {
			rcv := make([]*trakt.Alias, 0)
			f := trakt.NewEmptyFrame(&rcv)
			path := trakt.FormatURLPath("movies/%s/aliases", id)
			err := c.B.CallWithFrame(http.MethodGet, path, c.Key, p, f)
			return f, err
		}),
	}
}

func Releases(id trakt.SearchID, params *trakt.ReleaseListParams) *trakt.ReleaseIterator {
	return getC().Releases(id, params)
}

func (c *Client) Releases(id trakt.SearchID, params *trakt.ReleaseListParams) *trakt.ReleaseIterator {
	return &trakt.ReleaseIterator{
		Iterator: trakt.NewSimulatedIterator(params, func(p trakt.ListParamsContainer) (trakt.IterationFrame, error) {
			rcv := make([]*trakt.Release, 0)
			f := trakt.NewEmptyFrame(&rcv)
			path := trakt.FormatURLPath("movies/%s/releases/%s", id, params.Country)
			err := c.B.CallWithFrame(http.MethodGet, path, c.Key, p, f)
			return f, err
		}),
	}
}

func Translations(id trakt.SearchID, params *trakt.TranslationListParams) *trakt.TranslationIterator {
	return getC().Translations(id, params)
}

func (c *Client) Translations(id trakt.SearchID, params *trakt.TranslationListParams) *trakt.TranslationIterator {
	return &trakt.TranslationIterator{
		Iterator: trakt.NewSimulatedIterator(params, func(p trakt.ListParamsContainer) (trakt.IterationFrame, error) {
			rcv := make([]*trakt.Translation, 0)
			f := trakt.NewEmptyFrame(&rcv)
			path := trakt.FormatURLPath("movies/%s/translations/%s", id, params.Language)
			err := c.B.CallWithFrame(http.MethodGet, path, c.Key, p, f)
			return f, err
		}),
	}
}

func Comments(id trakt.SearchID, params *trakt.CommentListParams) *trakt.CommentIterator {
	return getC().Comments(id, params)
}

func (c *Client) Comments(id trakt.SearchID, params *trakt.CommentListParams) *trakt.CommentIterator {
	return &trakt.CommentIterator{
		Iterator: trakt.NewIterator(params, func(p trakt.ListParamsContainer) (trakt.IterationFrame, error) {
			rcv := make([]*trakt.Comment, 0)
			f := trakt.NewEmptyFrame(&rcv)
			path := trakt.FormatURLPath("movies/%s/comments/%s", id, params.Sort)
			err := c.B.CallWithFrame(http.MethodGet, path, c.Key, p, f)
			return f, err
		}),
	}
}

func WatchingNow(id trakt.SearchID, params *trakt.BasicListParams) *trakt.UserIterator {
	return getC().WatchingNow(id, params)
}

func (c *Client) WatchingNow(id trakt.SearchID, params *trakt.BasicListParams) *trakt.UserIterator {
	return &trakt.UserIterator{
		Iterator: trakt.NewIterator(params, func(p trakt.ListParamsContainer) (trakt.IterationFrame, error) {
			rcv := make([]*trakt.User, 0)
			f := trakt.NewEmptyFrame(&rcv)
			path := trakt.FormatURLPath("movies/%s/watching", id)
			err := c.B.CallWithFrame(http.MethodGet, path, c.Key, p, f)
			return f, err
		}),
	}
}

func Related(id trakt.SearchID, params *trakt.ExtendedListParams) *trakt.MovieIterator {
	return getC().Related(id, params)
}

func (c *Client) Related(id trakt.SearchID, params *trakt.ExtendedListParams) *trakt.MovieIterator {
	return &trakt.MovieIterator{
		Iterator: trakt.NewIterator(params, func(p trakt.ListParamsContainer) (trakt.IterationFrame, error) {
			rcv := make([]*trakt.Movie, 0)
			f := trakt.NewEmptyFrame(&rcv)
			path := trakt.FormatURLPath("movies/%s/related", id)
			err := c.B.CallWithFrame(http.MethodGet, path, c.Key, p, f)
			return f, err
		}),
	}
}

func Ratings(id trakt.SearchID, params *trakt.BasicParams) (*trakt.Rating, error) {
	return getC().Ratings(id, params)
}

func (c *Client) Ratings(id trakt.SearchID, params *trakt.BasicParams) (*trakt.Rating, error) {
	path := trakt.FormatURLPath("/movies/%s/ratings", id)
	stats := &trakt.Rating{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, stats)
	return stats, err
}

func Statistics(id trakt.SearchID, params *trakt.BasicParams) (*trakt.Statistics, error) {
	return getC().Statistics(id, params)
}

func (c *Client) Statistics(id trakt.SearchID, params *trakt.BasicParams) (*trakt.Statistics, error) {
	path := trakt.FormatURLPath("/movies/%s/stats", id)
	stats := &trakt.Statistics{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, stats)
	return stats, err
}

func (c *Client) generateTimePeriodMovieListIterator(
	action string, params *trakt.TimePeriodListParams,
) *trakt.MovieWithStatisticsIterator {
	return &trakt.MovieWithStatisticsIterator{
		Iterator: trakt.NewIterator(params, func(p trakt.ListParamsContainer) (trakt.IterationFrame, error) {
			rcv := make([]*trakt.MovieWithStatistics, 0)
			f := trakt.NewEmptyFrame(&rcv)
			var period = trakt.TimePeriodAll
			if params.Period != "" {
				period = params.Period
			}

			err := c.B.CallWithFrame(
				http.MethodGet,
				trakt.FormatURLPath("/movies/%s/%s", action, string(period)),
				c.Key, p, f,
			)
			return f, err
		}),
	}
}

func getC() *Client {
	return &Client{B: trakt.GetBackend(), Key: trakt.Key}
}
