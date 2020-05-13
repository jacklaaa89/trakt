package certification

import (
	"net/http"

	"github.com/jackaaa89/trakt"
)

type Client struct{ b *trakt.BaseClient }

func List(params *trakt.ListByTypeParams) *trakt.CertificationIterator {
	return getC().List(params)
}

func (c *Client) List(params *trakt.ListByTypeParams) *trakt.CertificationIterator {
	gl := make([]*trakt.Certification, 0)
	path := trakt.FormatURLPath("/certifications/%s", params.Type)
	return &trakt.CertificationIterator{Iterator: c.b.NewSimulatedIterator(http.MethodGet, path, params, &gl)}
}

func getC() *Client { return &Client{trakt.NewClient(trakt.GetBackend())} }
