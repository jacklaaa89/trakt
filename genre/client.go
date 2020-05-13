package genre

import (
	"net/http"

	"github.com/jackaaa89/trakt"
)

type Client struct {
	B   trakt.Backend
	Key string
}

func List(params *trakt.ListByTypeParams) *trakt.CountryIterator {
	return getC().List(params)
}

func (c *Client) List(params *trakt.ListByTypeParams) *trakt.CountryIterator {
	return &trakt.CountryIterator{
		Iterator: trakt.NewSimulatedIterator(params, func(p trakt.ListParamsContainer) (trakt.IterationFrame, error) {
			gl := make([]*trakt.Genre, 0)
			f := trakt.NewEmptyFrame(&gl)
			err := c.B.CallWithFrame(http.MethodGet, trakt.FormatURLPath("/genres/%s", string(params.Type)), c.Key, p, f)
			return f, err
		}),
	}
}

func getC() *Client { return &Client{trakt.GetBackend(), trakt.Key} }
