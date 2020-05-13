package language

import (
	"net/http"

	"github.com/jackaaa89/trakt"
)

type Client struct {
	B   trakt.Backend
	Key string
}

func List(params *trakt.ListByTypeParams) *trakt.LanguageIterator {
	return getC().List(params)
}

func (c *Client) List(params *trakt.ListByTypeParams) *trakt.LanguageIterator {
	return &trakt.LanguageIterator{
		Iterator: trakt.NewSimulatedIterator(params, func(p trakt.ListParamsContainer) (trakt.IterationFrame, error) {
			rcv := make([]*trakt.Language, 0)
			f := trakt.NewEmptyFrame(&rcv)
			err := c.B.CallWithFrame(http.MethodGet, trakt.FormatURLPath("/languages/%s", params.Type), c.Key, p, f)
			return f, err
		}),
	}
}

func getC() *Client { return &Client{trakt.GetBackend(), trakt.Key} }
