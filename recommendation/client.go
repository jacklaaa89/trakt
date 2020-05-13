package recommendation

import (
	"net/http"

	"github.com/jackaaa89/trakt"
)

type Client struct{ b *trakt.BaseClient }

func Movies(params *trakt.RecommendationListParams) *trakt.MovieIterator {
	return getC().Movies(params)
}

func (c *Client) Movies(params *trakt.RecommendationListParams) *trakt.MovieIterator {
	l := make([]*trakt.Movie, 0)
	return &trakt.MovieIterator{
		Iterator: c.b.NewIterator(http.MethodGet, "/recommendations/movies", params, &l),
	}
}

func Shows(params *trakt.RecommendationListParams) *trakt.ShowIterator {
	return getC().Shows(params)
}

func (c *Client) Shows(params *trakt.RecommendationListParams) *trakt.ShowIterator {
	l := make([]*trakt.Show, 0)
	return &trakt.ShowIterator{
		Iterator: c.b.NewIterator(http.MethodGet, "/recommendations/shows", params, &l),
	}
}

func HideShow(id trakt.SearchID, params *trakt.Params) error {
	return getC().HideShow(id, params)
}

func (c *Client) HideShow(id trakt.SearchID, params *trakt.Params) error {
	return c.b.Call(http.MethodDelete, trakt.FormatURLPath("/recommendations/shows/%s", id), params, nil)
}

func HideMovie(id trakt.SearchID, params *trakt.Params) error {
	return getC().HideMovie(id, params)
}

func (c *Client) HideMovie(id trakt.SearchID, params *trakt.Params) error {
	return c.b.Call(http.MethodDelete, trakt.FormatURLPath("/recommendations/movies/%s", id), params, nil)
}

func getC() *Client { return &Client{trakt.NewClient(trakt.GetBackend())} }
