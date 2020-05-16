package sync

import (
	"errors"
	"net/http"

	"github.com/jackaaa89/trakt"
)

// invalidTypeError error returned when an invalid media type is supplied.
var invalidTypeError = errors.New("invalid type: only movie / show are applicable")

type Client struct{ b *trakt.BaseClient }

func LastActivities(params *trakt.Params) (*trakt.LastActivity, error) {
	return getC().LastActivities(params)
}

func (c *Client) LastActivities(params *trakt.Params) (*trakt.LastActivity, error) {
	l := &trakt.LastActivity{}
	err := c.b.Call(http.MethodGet, "/sync/last_activities", params, l)
	return l, err
}

func Playbacks(params *trakt.ListPlaybackParams) *trakt.PlaybackIterator {
	return getC().Playbacks(params)
}

func (c *Client) Playbacks(params *trakt.ListPlaybackParams) *trakt.PlaybackIterator {
	path := trakt.FormatURLPath("/sync/playback/%s", params.Type)
	return &trakt.PlaybackIterator{BasicIterator: c.b.NewSimulatedIterator(http.MethodGet, path, params)}
}

func RemovePlayback(id int64, params *trakt.RemovePlaybackParams) error {
	return getC().RemovePlayback(id, params)
}

func (c *Client) RemovePlayback(id int64, params *trakt.RemovePlaybackParams) error {
	path := trakt.FormatURLPath("/sync/playback/%s", params.Type)
	return c.b.Call(http.MethodDelete, path, params, nil)
}

func Collection(params *trakt.ListCollectionParams) trakt.CollectionIterator {
	return getC().Collection(params)
}

func (c *Client) Collection(params *trakt.ListCollectionParams) trakt.CollectionIterator {
	if params.Type == trakt.TypeMovie {
		return c.movieCollection(params)
	}

	return c.showCollection(params)
}

func (c *Client) movieCollection(params *trakt.ListCollectionParams) *collection {
	return c.newCollectionIterator(trakt.TypeMovie, params)
}

func (c *Client) showCollection(params *trakt.ListCollectionParams) *collection {
	return c.newCollectionIterator(trakt.TypeShow, params)
}

func (c *Client) newCollectionIterator(t trakt.Type, p *trakt.ListCollectionParams) *collection {
	return &collection{
		genericIterator: genericIterator{
			BasicIterator: c.b.NewSimulatedIteratorWithCondition(
				http.MethodGet, trakt.FormatURLPath("/sync/collection/%s", t.Plural()), p, func() error {
					return compareType(p.Type, t)
				},
			),
			typ: p.Type,
		},
	}
}

func AddToCollection(params *trakt.AddToCollectionParams) (*trakt.AddToCollectionResult, error) {
	return getC().AddToCollection(params)
}

func (c *Client) AddToCollection(params *trakt.AddToCollectionParams) (*trakt.AddToCollectionResult, error) {
	rcv := &trakt.AddToCollectionResult{}
	err := c.b.Call(http.MethodPost, "/sync/collection", params, &rcv)
	return rcv, err
}

func RemoveFromCollection(params *trakt.RemoveFromCollectionParams) (*trakt.RemoveFromCollectionResult, error) {
	return getC().RemoveFromCollection(params)
}

func (c *Client) RemoveFromCollection(params *trakt.RemoveFromCollectionParams) (*trakt.RemoveFromCollectionResult, error) {
	rcv := &trakt.RemoveFromCollectionResult{}
	err := c.b.Call(http.MethodPost, "/sync/collection/remove", params, &rcv)
	return rcv, err
}

func Watched(params *trakt.ListCollectionParams) trakt.WatchedIterator {
	return getC().Watched(params)
}

func (c *Client) Watched(params *trakt.ListCollectionParams) trakt.WatchedIterator {
	if params.Type == trakt.TypeMovie {
		return c.watchedMovies(params)
	}

	return c.watchedShows(params)
}

func (c *Client) watchedMovies(params *trakt.ListWatchedParams) *watched {
	return c.newWatchedIterator(trakt.TypeMovie, params)
}

func (c *Client) watchedShows(params *trakt.ListWatchedParams) *watched {
	return c.newWatchedIterator(trakt.TypeShow, params)
}

func (c *Client) newWatchedIterator(t trakt.Type, p *trakt.ListWatchedParams) *watched {
	return &watched{
		genericIterator: genericIterator{
			BasicIterator: c.b.NewSimulatedIteratorWithCondition(
				http.MethodGet, trakt.FormatURLPath("/sync/watched/%s", t.Plural()), p, func() error {
					return compareType(p.Type, t)
				},
			),
			typ: p.Type,
		},
	}
}

func History(params *trakt.ListHistoryParams) *trakt.HistoryIterator {
	return getC().History(params)
}

func (c *Client) History(params *trakt.ListHistoryParams) *trakt.HistoryIterator {
	path := trakt.FormatURLPath("/sync/history/%s/%s", params.Type, params.ID)
	return &trakt.HistoryIterator{Iterator: c.b.NewIterator(http.MethodGet, path, params)}
}

func AddToHistory(params *trakt.AddToHistoryParams) (*trakt.AddToHistoryResult, error) {
	return getC().AddToHistory(params)
}

func (c *Client) AddToHistory(params *trakt.AddToHistoryParams) (*trakt.AddToHistoryResult, error) {
	rcv := &trakt.AddToHistoryResult{}
	err := c.b.Call(http.MethodPost, "/sync/history", params, &rcv)
	return rcv, err
}

func RemoveFromHistory(params *trakt.RemoveFromHistoryParams) (*trakt.RemoveFromHistoryResult, error) {
	return getC().RemoveFromHistory(params)
}

func (c *Client) RemoveFromHistory(params *trakt.RemoveFromHistoryParams) (*trakt.RemoveFromHistoryResult, error) {
	rcv := &trakt.RemoveFromHistoryResult{}
	err := c.b.Call(http.MethodPost, "/sync/history/remove", params, &rcv)
	return rcv, err
}

func Ratings(params *trakt.ListRatingParams) *trakt.RatingIterator {
	return getC().Ratings(params)
}

func (c *Client) Ratings(params *trakt.ListRatingParams) *trakt.RatingIterator {
	path := trakt.FormatURLPath("/sync/ratings/%s/%s", params.Type.Plural(), params.Ratings)
	return &trakt.RatingIterator{Iterator: c.b.NewIterator(http.MethodGet, path, params)}
}

func AddRatings(params *trakt.AddRatingsParams) (*trakt.AddRatingsResult, error) {
	return getC().AddRatings(params)
}

func (c *Client) AddRatings(params *trakt.AddRatingsParams) (*trakt.AddRatingsResult, error) {
	rcv := &trakt.AddRatingsResult{}
	err := c.b.Call(http.MethodPost, "/sync/ratings", params, &rcv)
	return rcv, err
}

func RemoveRatings(params *trakt.RemoveRatingsParams) (*trakt.RemoveRatingsResult, error) {
	return getC().RemoveRatings(params)
}

func (c *Client) RemoveRatings(params *trakt.RemoveRatingsParams) (*trakt.RemoveRatingsResult, error) {
	rcv := &trakt.RemoveRatingsResult{}
	err := c.b.Call(http.MethodPost, "/sync/ratings/remove", params, &rcv)
	return rcv, err
}

func WatchList(params *trakt.ListWatchListParams) *trakt.WatchListEntryIterator {
	return getC().WatchList(params)
}

func (c *Client) WatchList(params *trakt.ListWatchListParams) *trakt.WatchListEntryIterator {
	path := trakt.FormatURLPath("/sync/watchlist/%s/%s", params.Type.Plural(), params.Sort)
	return &trakt.WatchListEntryIterator{Iterator: c.b.NewIterator(http.MethodGet, path, params)}
}

func AddToWatchList(params *trakt.AddToWatchListParams) (*trakt.AddToWatchListResult, error) {
	return getC().AddToWatchList(params)
}

func (c *Client) AddToWatchList(params *trakt.AddToWatchListParams) (*trakt.AddToWatchListResult, error) {
	rcv := &trakt.AddToWatchListResult{}
	err := c.b.Call(http.MethodPost, "/sync/watchlist", params, &rcv)
	return rcv, err
}

func RemoveFromWatchList(params *trakt.RemoveFromWatchListParams) (*trakt.RemoveFromWatchListResult, error) {
	return getC().RemoveFromWatchList(params)
}

func (c *Client) RemoveFromWatchList(params *trakt.RemoveFromWatchListParams) (*trakt.RemoveFromWatchListResult, error) {
	rcv := &trakt.RemoveFromWatchListResult{}
	err := c.b.Call(http.MethodPost, "/sync/watchlist/remove", params, &rcv)
	return rcv, err
}

type genericIterator struct {
	trakt.BasicIterator
	typ trakt.Type
}

func (g *genericIterator) Type() trakt.Type { return g.typ }

type collection struct{ genericIterator }

func (c *collection) Show() (*trakt.CollectedShow, error) {
	rcv := &trakt.CollectedShow{}
	return rcv, c.Scan(rcv)
}

func (c *collection) Movie() (*trakt.CollectedMovie, error) {
	rcv := &trakt.CollectedMovie{}
	return rcv, c.Scan(rcv)
}

type watched struct{ genericIterator }

func (c *watched) Show() (*trakt.WatchedShow, error) {
	rcv := &trakt.WatchedShow{}
	return rcv, c.Scan(rcv)
}

func (c *watched) Movie() (*trakt.WatchedMovie, error) {
	rcv := &trakt.WatchedMovie{}
	return rcv, c.Scan(rcv)
}

// compareType helper function to compare to types to see if they are equal.
func compareType(a, b trakt.Type) error {
	if a == b {
		return nil
	}

	return invalidTypeError
}

func getC() *Client { return &Client{trakt.NewClient(trakt.GetBackend())} }
