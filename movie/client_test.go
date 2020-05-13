package movie

import (
	"context"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/jackaaa89/trakt"
)

func TestClient_Popular(t *testing.T) {
	trakt.Key = "5500e6a5bcddb29ba9844ac3124e9c7a1516f46e52aec9563941987b02112690"
	c := getC()
	id := trakt.IMDB("tt8579674")

	it := c.Comments(
		id,
		&trakt.CommentListParams{
			BasicListParams: trakt.BasicListParams{Context: context.Background()},
			Sort:            trakt.SortTypeNewest,
		},
	)

	for it.Next() {
		spew.Dump(it.Comment().Text)
	}
}
