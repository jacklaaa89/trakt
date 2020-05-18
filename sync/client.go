package sync

import (
	"net/http"

	"github.com/jacklaaa89/trakt"
)

// client represents a sync client which gives us access to functions to sync
// trakt with one or more media centres.
type client struct{ b trakt.BaseClient }

// LastActivities returns all of the dates of the latest activity for a user.
//
// This method is a useful first step in the syncing process. We recommended caching these dates locally,
// then you can compare to know exactly what data has changed recently. This can greatly optimize your syncs
// so you don't pull down a ton of data only to see nothing has actually changed.
//
//  - OAuth Required
func LastActivities(params *trakt.Params) (*trakt.LastActivity, error) {
	return getC().LastActivities(params)
}

// LastActivities returns all of the dates of the latest activity for a user.
//
// This method is a useful first step in the syncing process. We recommended caching these dates locally,
// then you can compare to know exactly what data has changed recently. This can greatly optimize your syncs
// so you don't pull down a ton of data only to see nothing has actually changed.
//
//  - OAuth Required
func (c *client) LastActivities(params *trakt.Params) (*trakt.LastActivity, error) {
	l := &trakt.LastActivity{}
	err := c.b.Call(http.MethodGet, "/sync/last_activities", params, l)
	return l, err
}

// Playbacks returns a list of stored paused playbacks.
//
// Whenever a scrobble is paused, the playback progress is saved. Use this progress to sync up
// playback across different media centers or apps. For example, you can start watching a movie in a
// media center, stop it, then resume on your tablet from the same spot. Each item will have the progress
// percentage between 0 and 100.You can optionally specify a type to only get movies or episodes.
//
// By default, all results will be returned. However, you can send a limit if you only need a few recent
// results for something like an "on deck" feature.
//
//  - OAuth Required
func Playbacks(params *trakt.ListPlaybackParams) *trakt.PlaybackIterator {
	return getC().Playbacks(params)
}

// Playbacks returns a list of stored paused playbacks.
//
// Whenever a scrobble is paused, the playback progress is saved. Use this progress to sync up
// playback across different media centers or apps. For example, you can start watching a movie in a
// media center, stop it, then resume on your tablet from the same spot. Each item will have the progress
// percentage between 0 and 100.You can optionally specify a type to only get movies or episodes.
//
// By default, all results will be returned. However, you can send a limit if you only need a few recent
// results for something like an "on deck" feature.
//
//  - OAuth Required
func (c *client) Playbacks(params *trakt.ListPlaybackParams) *trakt.PlaybackIterator {
	path := trakt.FormatURLPath("/sync/playback/%s", params.Type)
	return &trakt.PlaybackIterator{BasicIterator: c.b.NewSimulatedIterator(http.MethodGet, path, params)}
}

func RemovePlayback(id int64, params *trakt.RemovePlaybackParams) error {
	return getC().RemovePlayback(id, params)
}

func (c *client) RemovePlayback(id int64, params *trakt.RemovePlaybackParams) error {
	path := trakt.FormatURLPath("/sync/playback/%s", params.Type)
	return c.b.Call(http.MethodDelete, path, params, nil)
}

func Collection(params *trakt.ListCollectionParams) trakt.CollectionIterator {
	return getC().Collection(params)
}

func (c *client) Collection(params *trakt.ListCollectionParams) trakt.CollectionIterator {
	if params.Type == trakt.TypeMovie {
		return c.movieCollection(params)
	}

	return c.showCollection(params)
}

func AddToCollection(params *trakt.AddToCollectionParams) (*trakt.AddToCollectionResult, error) {
	return getC().AddToCollection(params)
}

func (c *client) AddToCollection(params *trakt.AddToCollectionParams) (*trakt.AddToCollectionResult, error) {
	rcv := &trakt.AddToCollectionResult{}
	err := c.b.Call(http.MethodPost, "/sync/collection", params, &rcv)
	return rcv, err
}

func RemoveFromCollection(params *trakt.RemoveFromCollectionParams) (*trakt.RemoveFromCollectionResult, error) {
	return getC().RemoveFromCollection(params)
}

func (c *client) RemoveFromCollection(params *trakt.RemoveFromCollectionParams) (*trakt.RemoveFromCollectionResult, error) {
	rcv := &trakt.RemoveFromCollectionResult{}
	err := c.b.Call(http.MethodPost, "/sync/collection/remove", params, &rcv)
	return rcv, err
}

func Watched(params *trakt.ListCollectionParams) trakt.WatchedIterator {
	return getC().Watched(params)
}

func (c *client) Watched(params *trakt.ListCollectionParams) trakt.WatchedIterator {
	if params.Type == trakt.TypeMovie {
		return c.watchedMovies(params)
	}

	return c.watchedShows(params)
}

func History(params *trakt.ListHistoryParams) *trakt.HistoryIterator {
	return getC().History(params)
}

func (c *client) History(params *trakt.ListHistoryParams) *trakt.HistoryIterator {
	path := trakt.FormatURLPath("/sync/history/%s/%s", params.Type, params.ID)
	return &trakt.HistoryIterator{Iterator: c.b.NewIterator(http.MethodGet, path, params)}
}

func AddToHistory(params *trakt.AddToHistoryParams) (*trakt.AddToHistoryResult, error) {
	return getC().AddToHistory(params)
}

func (c *client) AddToHistory(params *trakt.AddToHistoryParams) (*trakt.AddToHistoryResult, error) {
	rcv := &trakt.AddToHistoryResult{}
	err := c.b.Call(http.MethodPost, "/sync/history", params, &rcv)
	return rcv, err
}

func RemoveFromHistory(params *trakt.RemoveFromHistoryParams) (*trakt.RemoveFromHistoryResult, error) {
	return getC().RemoveFromHistory(params)
}

func (c *client) RemoveFromHistory(params *trakt.RemoveFromHistoryParams) (*trakt.RemoveFromHistoryResult, error) {
	rcv := &trakt.RemoveFromHistoryResult{}
	err := c.b.Call(http.MethodPost, "/sync/history/remove", params, &rcv)
	return rcv, err
}

func Ratings(params *trakt.ListRatingParams) *trakt.RatingIterator {
	return getC().Ratings(params)
}

func (c *client) Ratings(params *trakt.ListRatingParams) *trakt.RatingIterator {
	path := trakt.FormatURLPath("/sync/ratings/%s/%s", params.Type.Plural(), params.Ratings)
	return &trakt.RatingIterator{Iterator: c.b.NewIterator(http.MethodGet, path, params)}
}

func AddRatings(params *trakt.AddRatingsParams) (*trakt.AddRatingsResult, error) {
	return getC().AddRatings(params)
}

func (c *client) AddRatings(params *trakt.AddRatingsParams) (*trakt.AddRatingsResult, error) {
	rcv := &trakt.AddRatingsResult{}
	err := c.b.Call(http.MethodPost, "/sync/ratings", params, &rcv)
	return rcv, err
}

func RemoveRatings(params *trakt.RemoveRatingsParams) (*trakt.RemoveRatingsResult, error) {
	return getC().RemoveRatings(params)
}

func (c *client) RemoveRatings(params *trakt.RemoveRatingsParams) (*trakt.RemoveRatingsResult, error) {
	rcv := &trakt.RemoveRatingsResult{}
	err := c.b.Call(http.MethodPost, "/sync/ratings/remove", params, &rcv)
	return rcv, err
}

func WatchList(params *trakt.ListWatchListParams) *trakt.WatchListEntryIterator {
	return getC().WatchList(params)
}

func (c *client) WatchList(params *trakt.ListWatchListParams) *trakt.WatchListEntryIterator {
	path := trakt.FormatURLPath("/sync/watchlist/%s/%s", params.Type.Plural(), params.Sort)
	return &trakt.WatchListEntryIterator{Iterator: c.b.NewIterator(http.MethodGet, path, params)}
}

func AddToWatchList(params *trakt.AddToWatchListParams) (*trakt.AddToWatchListResult, error) {
	return getC().AddToWatchList(params)
}

func (c *client) AddToWatchList(params *trakt.AddToWatchListParams) (*trakt.AddToWatchListResult, error) {
	rcv := &trakt.AddToWatchListResult{}
	err := c.b.Call(http.MethodPost, "/sync/watchlist", params, &rcv)
	return rcv, err
}

func RemoveFromWatchList(params *trakt.RemoveFromWatchListParams) (*trakt.RemoveFromWatchListResult, error) {
	return getC().RemoveFromWatchList(params)
}

func (c *client) RemoveFromWatchList(params *trakt.RemoveFromWatchListParams) (*trakt.RemoveFromWatchListResult, error) {
	rcv := &trakt.RemoveFromWatchListResult{}
	err := c.b.Call(http.MethodPost, "/sync/watchlist/remove", params, &rcv)
	return rcv, err
}

// movieCollection generates an iterator for collected movies.
func (c *client) movieCollection(params *trakt.ListCollectionParams) *collection {
	return c.newCollectionIterator(trakt.TypeMovie, params)
}

// showCollection generates an iterator for collected shows.
func (c *client) showCollection(params *trakt.ListCollectionParams) *collection {
	return c.newCollectionIterator(trakt.TypeShow, params)
}

// newWatchedIterator generates an iterator for either watched shows or movies
// based on the type.
func (c *client) newCollectionIterator(t trakt.Type, p *trakt.ListCollectionParams) *collection {
	path := trakt.FormatURLPath("/sync/collection/%s", t.Plural())
	return &collection{
		genericIterator: genericIterator{
			BasicIterator: c.b.NewSimulatedIteratorWithCondition(
				http.MethodGet, path, p, func() error {
					return compareType(path, p.Type, t)
				},
			),
			typ: p.Type,
		},
	}
}

// watchedMovies generates an iterator for watched movies.
func (c *client) watchedMovies(params *trakt.ListWatchedParams) *watched {
	return c.newWatchedIterator(trakt.TypeMovie, params)
}

// watchedShows generates an iterator for watched shows.
func (c *client) watchedShows(params *trakt.ListWatchedParams) *watched {
	return c.newWatchedIterator(trakt.TypeShow, params)
}

// newWatchedIterator generates an iterator for either watched shows or movies
// based on the type.
func (c *client) newWatchedIterator(t trakt.Type, p *trakt.ListWatchedParams) *watched {
	path := trakt.FormatURLPath("/sync/watched/%s", t.Plural())
	return &watched{
		genericIterator: genericIterator{
			BasicIterator: c.b.NewSimulatedIteratorWithCondition(
				http.MethodGet, path, p, func() error {
					return compareType(path, p.Type, t)
				},
			),
			typ: p.Type,
		},
	}
}

// genericIterator this is a generic iterator for listing
// both watched and collected movies and shows.
type genericIterator struct {
	trakt.BasicIterator

	// typ the type of object which this iterator represents
	// can either be show or movie.
	typ trakt.Type
}

// Type implements both WatchedIterator and CollectionIterator interfaces.
// returns the type so the user knows which entry to use.
func (g *genericIterator) Type() trakt.Type { return g.typ }

// collection an implementation of a CollectionIterator
// uses the internal iterator defined on the genericIterator
// to attempt to scan and cast to either a collected show or movie
// based on the type.
type collection struct{ genericIterator }

// Show implements CollectionIterator interface.
func (c *collection) Show() (*trakt.CollectedShow, error) {
	rcv := &trakt.CollectedShow{}
	return rcv, c.Scan(rcv)
}

// Movie implements CollectionIterator interface.
func (c *collection) Movie() (*trakt.CollectedMovie, error) {
	rcv := &trakt.CollectedMovie{}
	return rcv, c.Scan(rcv)
}

// watched an implementation of a WatchedIterator
// uses the internal iterator defined on the genericIterator
// to attempt to scan and cast to either a watched show or movie
// based on the type.
type watched struct{ genericIterator }

// Show implements WatchedIterator interface.
func (c *watched) Show() (*trakt.WatchedShow, error) {
	rcv := &trakt.WatchedShow{}
	return rcv, c.Scan(rcv)
}

// Movie implements WatchedIterator interface.
func (c *watched) Movie() (*trakt.WatchedMovie, error) {
	rcv := &trakt.WatchedMovie{}
	return rcv, c.Scan(rcv)
}

// compareType helper function to compare to types to see if they are equal.
// the path is required to generate the standard error signature if the types do not
// match.
func compareType(path string, a, b trakt.Type) error {
	if a == b {
		return nil
	}

	return &trakt.Error{
		HTTPStatusCode: http.StatusUnprocessableEntity,
		Body:           "invalid type: only movie / show are applicable",
		Resource:       path,
		Code:           trakt.ErrorCodeValidationError,
	}
}

// getC initialises a new sync client using the currently configured backend.
func getC() *client { return &client{trakt.NewClient()} }
