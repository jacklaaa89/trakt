package authorization

import (
	"context"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"

	"github.com/jackaaa89/trakt"
)

func TestNewCode(t *testing.T) {
	trakt.Key = "5500e6a5bcddb29ba9844ac3124e9c7a1516f46e52aec9563941987b02112690"
	spew.Dump(NewCode(&trakt.BasicParams{Context: context.Background()}))
}

// 4e4bd8adeac2892c360bcdbb00aa17e3e407c90d7e4f917947fc6b8c62f945b7

func TestPollAsync(t *testing.T) {
	trakt.Key = "5500e6a5bcddb29ba9844ac3124e9c7a1516f46e52aec9563941987b02112690"
	ch := PollAsync(&trakt.PollCodeParams{
		BasicParams: trakt.BasicParams{
			Context: context.Background(),
		},
		Code:         "697bab1f19c198b94a25d72b898810c8bf8723cb9b4d47cd8487661214ea2acf",
		ClientSecret: "ab7471f942f034096818c866aaea9cb9454aa19fcd14ad26f744bedbf5d08d10",
		Interval:     5 * time.Second,
		ExpiresIn:    10 * time.Minute,
	})

	res := <-ch
	spew.Dump(res)
}
