package certification

import (
	"net/http"

	"github.com/jackaaa89/trakt"
)

type Client struct {
	B   trakt.Backend
	Key string
}

func List(params *trakt.ListByTypeParams) *trakt.CertificationIterator {
	return getC().List(params)
}

func (c *Client) List(params *trakt.ListByTypeParams) *trakt.CertificationIterator {
	return &trakt.CertificationIterator{
		Iterator: trakt.NewSimulatedIterator(params, func(p trakt.ListParamsContainer) (trakt.IterationFrame, error) {
			gl := make([]*trakt.Certification, 0)
			f := trakt.NewEmptyFrame(&gl)
			err := c.B.CallWithFrame(http.MethodGet, trakt.FormatURLPath("/certifications/%s", string(params.Type)), c.Key, p, f)
			return f, err
		}),
	}
}

func getC() *Client { return &Client{trakt.GetBackend(), trakt.Key} }
