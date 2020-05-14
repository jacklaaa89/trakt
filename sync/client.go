package sync

import (
	"net/http"

	"github.com/jackaaa89/trakt"
)

type Client struct{ b *trakt.BaseClient }

func LastActivities(params *trakt.Params) (*trakt.LastActivity, error) {
	return getC().LastActivities(params)
}

func (c *Client) LastActivities(params *trakt.Params) (*trakt.LastActivity, error) {
	l := &trakt.LastActivity{}
	err := c.b.Call(http.MethodGet, "/sync/last_activities", params, l)
	return l, err
}

func Playback(params *trakt.ListPlaybackParams) *trakt.PlaybackIterator {
	return getC().Playback(params)
}

func (c *Client) Playback(params *trakt.ListPlaybackParams) *trakt.PlaybackIterator {
	cl := make([]*trakt.Playback, 0)
	path := trakt.FormatURLPath("/sync/playback/%s", params.Type)
	return &trakt.PlaybackIterator{BasicIterator: c.b.NewSimulatedIterator(http.MethodGet, path, params, &cl)}
}

func RemovePlayback(id int64, params *trakt.RemovePlaybackParams) error {
	return getC().RemovePlayback(id, params)
}

func (c *Client) RemovePlayback(id int64, params *trakt.RemovePlaybackParams) error {
	path := trakt.FormatURLPath("/sync/playback/%s", params.Type)
	return c.b.Call(http.MethodDelete, path, params, nil)
}

func Collection(params *trakt.GetCollectionParams) trakt.CollectionIterator {
	return getC().Collection(params)
}

func (c *Client) Collection(params *trakt.GetCollectionParams) trakt.CollectionIterator {
	if params.Type == trakt.TypeMovie {
		return c.movieCollection(params)
	}

	return c.showCollection(params)
}

func (c *Client) movieCollection(params *trakt.GetCollectionParams) *collectedMovieIterator {
	cl := make([]*trakt.CollectedMovie, 0)
	return &collectedMovieIterator{
		BasicIterator: c.b.NewSimulatedIterator(http.MethodGet, "/sync/collection/movies", params, &cl),
	}
}

func (c *Client) showCollection(params *trakt.GetCollectionParams) *collectedShowIterator {
	cl := make([]*trakt.CollectedShow, 0)
	return &collectedShowIterator{
		BasicIterator: c.b.NewSimulatedIterator(http.MethodGet, "/sync/collection/shows", params, &cl),
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

type collectedMovieIterator struct{ trakt.BasicIterator }

func (c *collectedMovieIterator) Type() trakt.Type           { return trakt.TypeMovie }
func (c *collectedMovieIterator) Show() *trakt.CollectedShow { return nil }
func (c *collectedMovieIterator) Movie() *trakt.CollectedMovie {
	return c.Current().(*trakt.CollectedMovie)
}

type collectedShowIterator struct{ trakt.BasicIterator }

func (c *collectedShowIterator) Type() trakt.Type             { return trakt.TypeShow }
func (c *collectedShowIterator) Movie() *trakt.CollectedMovie { return nil }
func (c *collectedShowIterator) Show() *trakt.CollectedShow {
	return c.Current().(*trakt.CollectedShow)
}

func getC() *Client { return &Client{trakt.NewClient(trakt.GetBackend())} }
