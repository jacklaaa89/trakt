package show

import (
	"net/http"

	"github.com/jacklaaa89/trakt"
)

type Client struct{ b trakt.BaseClient }

func Trending(params *trakt.FilterListParams) *trakt.TrendingShowIterator {
	return getC().Trending(params)
}

func (c *Client) Trending(params *trakt.FilterListParams) *trakt.TrendingShowIterator {
	return &trakt.TrendingShowIterator{Iterator: c.b.NewIterator(http.MethodGet, "/shows/trending", params)}
}

func Popular(params *trakt.FilterListParams) *trakt.ShowIterator {
	return getC().Popular(params)
}

func (c *Client) Popular(params *trakt.FilterListParams) *trakt.ShowIterator {
	return &trakt.ShowIterator{Iterator: c.b.NewIterator(http.MethodGet, "/shows/popular", params)}
}

func Played(params *trakt.TimePeriodListParams) *trakt.ShowWithStatisticsIterator {
	return getC().Played(params)
}

func (c *Client) Played(params *trakt.TimePeriodListParams) *trakt.ShowWithStatisticsIterator {
	return c.newTimePeriodIterator("played", params)
}

func Watched(params *trakt.TimePeriodListParams) *trakt.ShowWithStatisticsIterator {
	return getC().Watched(params)
}

func (c *Client) Watched(params *trakt.TimePeriodListParams) *trakt.ShowWithStatisticsIterator {
	return c.newTimePeriodIterator("watched", params)
}

func Collected(params *trakt.TimePeriodListParams) *trakt.ShowWithStatisticsIterator {
	return getC().Collected(params)
}

func (c *Client) Collected(params *trakt.TimePeriodListParams) *trakt.ShowWithStatisticsIterator {
	return c.newTimePeriodIterator("collected", params)
}

func Anticipated(params *trakt.FilterListParams) *trakt.AnticipatedShowIterator {
	return getC().Anticipated(params)
}

func (c *Client) Anticipated(params *trakt.FilterListParams) *trakt.AnticipatedShowIterator {
	return &trakt.AnticipatedShowIterator{Iterator: c.b.NewIterator(http.MethodGet, "/shows/anticipated", params)}
}

func RecentlyUpdated(params *trakt.RecentlyUpdatedListParams) *trakt.RecentlyUpdatedShowIterator {
	return getC().RecentlyUpdated(params)
}

func (c *Client) RecentlyUpdated(params *trakt.RecentlyUpdatedListParams) *trakt.RecentlyUpdatedShowIterator {
	path := trakt.FormatURLPath("/shows/updates/%s", params.StartDate.Format(`2006-01-02`))
	return &trakt.RecentlyUpdatedShowIterator{Iterator: c.b.NewIterator(http.MethodGet, path, params)}
}

func Get(id trakt.SearchID, params *trakt.ExtendedParams) (*trakt.Show, error) {
	return getC().Get(id, params)
}

func (c *Client) Get(id trakt.SearchID, params *trakt.ExtendedParams) (*trakt.Show, error) {
	path := trakt.FormatURLPath("/shows/%s", id)
	mov := &trakt.Show{}
	err := c.b.Call(http.MethodGet, path, params, mov)
	return mov, err
}

func Aliases(id trakt.SearchID, params *trakt.BasicParams) *trakt.AliasIterator {
	return getC().Aliases(id, params)
}

func (c *Client) Aliases(id trakt.SearchID, params *trakt.BasicParams) *trakt.AliasIterator {
	path := trakt.FormatURLPath("shows/%s/aliases", id)
	return &trakt.AliasIterator{BasicIterator: c.b.NewSimulatedIterator(http.MethodGet, path, params)}
}

func Certifications(id trakt.SearchID, params *trakt.BasicParams) *trakt.CertificationIterator {
	return getC().Certifications(id, params)
}

func (c *Client) Certifications(id trakt.SearchID, params *trakt.BasicParams) *trakt.CertificationIterator {
	path := trakt.FormatURLPath("shows/%s/certifications", id)
	return &trakt.CertificationIterator{BasicIterator: c.b.NewSimulatedIterator(http.MethodGet, path, params)}
}

func Translations(id trakt.SearchID, params *trakt.TranslationListParams) *trakt.TranslationIterator {
	return getC().Translations(id, params)
}

func (c *Client) Translations(id trakt.SearchID, params *trakt.TranslationListParams) *trakt.TranslationIterator {
	path := trakt.FormatURLPath("shows/%s/translations/%s", id, params.Language)
	return &trakt.TranslationIterator{BasicIterator: c.b.NewSimulatedIterator(http.MethodGet, path, params)}
}

func Comments(id trakt.SearchID, params *trakt.CommentListParams) *trakt.CommentIterator {
	return getC().Comments(id, params)
}

func (c *Client) Comments(id trakt.SearchID, params *trakt.CommentListParams) *trakt.CommentIterator {
	path := trakt.FormatURLPath("shows/%s/comments/%s", id, params.Sort)
	return &trakt.CommentIterator{Iterator: c.b.NewIterator(http.MethodGet, path, params)}
}

func Lists(id trakt.SearchID, params *trakt.GetListParams) *trakt.ListIterator {
	return getC().Lists(id, params)
}

func (c *Client) Lists(id trakt.SearchID, params *trakt.GetListParams) *trakt.ListIterator {
	path := trakt.FormatURLPath("shows/%s/lists/%s/%s", id, params.ListType, params.SortType)
	return &trakt.ListIterator{Iterator: c.b.NewIterator(http.MethodGet, path, params)}
}

func CollectionProgress(id trakt.SearchID, params *trakt.ProgressParams) (*trakt.CollectedProgress, error) {
	return getC().CollectionProgress(id, params)
}

func (c *Client) CollectionProgress(id trakt.SearchID, params *trakt.ProgressParams) (*trakt.CollectedProgress, error) {
	path := trakt.FormatURLPath("/shows/%s/progress/collection", id)
	cc := &trakt.CollectedProgress{}
	err := c.b.Call(http.MethodGet, path, params, cc)
	return cc, err
}

func WatchedProgress(id trakt.SearchID, params *trakt.ProgressParams) (*trakt.WatchedProgress, error) {
	return getC().WatchedProgress(id, params)
}

func (c *Client) WatchedProgress(id trakt.SearchID, params *trakt.ProgressParams) (*trakt.WatchedProgress, error) {
	path := trakt.FormatURLPath("/shows/%s/progress/watched", id)
	cc := &trakt.WatchedProgress{}
	err := c.b.Call(http.MethodGet, path, params, cc)
	return cc, err
}

func People(id trakt.SearchID, params *trakt.ExtendedParams) (*trakt.CastAndCrew, error) {
	return getC().People(id, params)
}

func (c *Client) People(id trakt.SearchID, params *trakt.ExtendedParams) (*trakt.CastAndCrew, error) {
	path := trakt.FormatURLPath("/shows/%s/people", id)
	cc := &trakt.CastAndCrew{}
	err := c.b.Call(http.MethodGet, path, params, cc)
	return cc, err
}

func Ratings(id trakt.SearchID, params *trakt.BasicParams) (*trakt.RatingDistribution, error) {
	return getC().Ratings(id, params)
}

func (c *Client) Ratings(id trakt.SearchID, params *trakt.BasicParams) (*trakt.RatingDistribution, error) {
	path := trakt.FormatURLPath("/shows/%s/ratings", id)
	stats := &trakt.RatingDistribution{}
	err := c.b.Call(http.MethodGet, path, params, stats)
	return stats, err
}

func Statistics(id trakt.SearchID, params *trakt.BasicParams) (*trakt.Statistics, error) {
	return getC().Statistics(id, params)
}

func (c *Client) Statistics(id trakt.SearchID, params *trakt.BasicParams) (*trakt.Statistics, error) {
	path := trakt.FormatURLPath("/shows/%s/stats", id)
	stats := &trakt.Statistics{}
	err := c.b.Call(http.MethodGet, path, params, stats)
	return stats, err
}

func WatchingNow(id trakt.SearchID, params *trakt.BasicListParams) *trakt.UserIterator {
	return getC().WatchingNow(id, params)
}

func (c *Client) WatchingNow(id trakt.SearchID, params *trakt.BasicListParams) *trakt.UserIterator {
	path := trakt.FormatURLPath("shows/%s/watching", id)
	return &trakt.UserIterator{Iterator: c.b.NewIterator(http.MethodGet, path, params)}
}

func NextEpisode(id trakt.SearchID, params *trakt.ExtendedParams) (*trakt.Episode, error) {
	return getC().NextEpisode(id, params)
}

func (c *Client) NextEpisode(id trakt.SearchID, params *trakt.ExtendedParams) (*trakt.Episode, error) {
	ep := &trakt.Episode{}
	path := trakt.FormatURLPath("shows/%s/next_episode", id)
	err := c.b.Call(http.MethodGet, path, params, ep)
	return handleNoEpisodeFound(ep, err)
}

func LastEpisode(id trakt.SearchID, params *trakt.ExtendedParams) (*trakt.Episode, error) {
	return getC().LastEpisode(id, params)
}

func (c *Client) LastEpisode(id trakt.SearchID, params *trakt.ExtendedParams) (*trakt.Episode, error) {
	ep := &trakt.Episode{}
	path := trakt.FormatURLPath("shows/%s/last_episode", id)
	err := c.b.Call(http.MethodGet, path, params, ep)
	return handleNoEpisodeFound(ep, err)
}

func Seasons(id trakt.SearchID, params *trakt.BasicListParams) *trakt.SeasonIterator {
	return getC().Seasons(id, params)
}

func (c *Client) Seasons(id trakt.SearchID, params *trakt.BasicListParams) *trakt.SeasonIterator {
	path := trakt.FormatURLPath("shows/%s/seasons", id)
	return &trakt.SeasonIterator{Iterator: c.b.NewIterator(http.MethodGet, path, params)}
}

func (c *Client) newTimePeriodIterator(action string, p *trakt.TimePeriodListParams) *trakt.ShowWithStatisticsIterator {
	var period = trakt.TimePeriodAll
	if p.Period != "" {
		period = p.Period
	}
	path := trakt.FormatURLPath("/shows/%s/%s", action, period)
	return &trakt.ShowWithStatisticsIterator{Iterator: c.b.NewIterator(http.MethodGet, path, p)}
}

// handleNoEpisodeFound helper function to handle if the response was a success but no episode was found
// in any request to trakt the BARE minimum we should receive in all types of request are a trakt UUID
// if we had no error and an episode with no ID then we can assume that the request was a success
// but no episode was found.
func handleNoEpisodeFound(e *trakt.Episode, err error) (*trakt.Episode, error) {
	if err != nil {
		return nil, err
	}

	if e == nil {
		return nil, nil
	}

	if e.Trakt <= 0 {
		return nil, nil
	}

	return e, nil
}

func getC() *Client { return &Client{trakt.NewClient()} }
