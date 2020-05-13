package recommendation

import (
	"net/http"

	"github.com/jackaaa89/trakt"
)

type Client struct {
	B   trakt.Backend
	Key string
}

func Movies(params *trakt.RecommendationListParams) *trakt.MovieIterator {
	return getC().Movies(params)
}

func (c *Client) Movies(params *trakt.RecommendationListParams) *trakt.MovieIterator {
	return &trakt.MovieIterator{
		Iterator: trakt.NewIterator(params, func(p trakt.ListParamsContainer) (trakt.IterationFrame, error) {
			l := make([]*trakt.Movie, 0)
			f := trakt.NewEmptyFrame(&l)
			err := c.B.CallWithFrame(http.MethodGet, "/recommendations/movies", c.Key, p, f)
			return f, err
		}),
	}
}

func Shows(params *trakt.RecommendationListParams) *trakt.ShowIterator {
	return getC().Shows(params)
}

func (c *Client) Shows(params *trakt.RecommendationListParams) *trakt.ShowIterator {
	return &trakt.ShowIterator{
		Iterator: trakt.NewIterator(params, func(p trakt.ListParamsContainer) (trakt.IterationFrame, error) {
			l := make([]*trakt.Show, 0)
			f := trakt.NewEmptyFrame(&l)
			err := c.B.CallWithFrame(http.MethodGet, "/recommendations/shows", c.Key, p, f)
			return f, err
		}),
	}
}

func HideShow(id trakt.SearchID, params *trakt.Params) error {
	return getC().HideShow(id, params)
}

func (c *Client) HideShow(id trakt.SearchID, params *trakt.Params) error {
	return c.B.Call(http.MethodDelete, trakt.FormatURLPath("/recommendations/shows/%s", id), c.Key, params, nil)
}

func HideMovie(id trakt.SearchID, params *trakt.Params) error {
	return getC().HideMovie(id, params)
}

func (c *Client) HideMovie(id trakt.SearchID, params *trakt.Params) error {
	return c.B.Call(http.MethodDelete, trakt.FormatURLPath("/recommendations/movies/%s", id), c.Key, params, nil)
}

func getC() *Client {
	return &Client{B: trakt.GetBackend(), Key: trakt.Key}
}
