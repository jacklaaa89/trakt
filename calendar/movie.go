package calendar

import (
	"net/http"

	"github.com/jackaaa89/trakt"
)

func Movies(params *trakt.BasicCalendarParams) *trakt.CalendarMovieIterator {
	return getC().Movies(params)
}

func (c *Client) Movies(params *trakt.BasicCalendarParams) *trakt.CalendarMovieIterator {
	return c.movies("movies", params)
}

func DVDs(params *trakt.BasicCalendarParams) *trakt.CalendarMovieIterator {
	return getC().DVDs(params)
}

func (c *Client) DVDs(params *trakt.BasicCalendarParams) *trakt.CalendarMovieIterator {
	return c.movies("dvds", params)
}

func (c *Client) movies(mediaType string, params *trakt.BasicCalendarParams) *trakt.CalendarMovieIterator {
	return c.generateMovieIterator(
		trakt.FormatURLPath("/calendars/all/%s", mediaType),
		&wrappedBasicCalendarParams{*params},
	)
}

func (c *Client) generateMovieIterator(path string, params *wrappedBasicCalendarParams) *trakt.CalendarMovieIterator {
	return &trakt.CalendarMovieIterator{
		Iterator: trakt.NewIterator(params, func(p trakt.ListParamsContainer) (trakt.IterationFrame, error) {
			list := make([]*trakt.CalendarMovie, 0)
			f := trakt.NewEmptyFrame(&list)
			err := c.B.CallWithFrame(http.MethodGet, formatPath(path, params), c.Key, p, f)
			return f, err
		}),
	}
}
