package language

import (
	"context"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/jackaaa89/trakt"
)

func TestList(t *testing.T) {
	trakt.Key = "5500e6a5bcddb29ba9844ac3124e9c7a1516f46e52aec9563941987b02112690"
	it := List(&trakt.ListByTypeParams{
		BasicListParams: trakt.BasicListParams{
			Context: context.Background(),
		},
		Type: trakt.TypeMovie,
	})

	for it.Next() {
		spew.Println(it.Language())
	}

	spew.Println(it.Err())
}
