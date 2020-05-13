package calendar

import (
	"net/http"
	"strconv"
	"time"

	"github.com/jackaaa89/trakt"
)

type Client struct {
	B   trakt.Backend
	Key string
}

func MyShows(params *trakt.CalendarParams) *trakt.CalendarShowIterator {
	return getC().MyShows(params)
}

func (c *Client) MyShows(params *trakt.CalendarParams) *trakt.CalendarShowIterator {
	return c.shows("my", &wrappedCalendarParams{*params})
}

func MyNewShows(params *trakt.CalendarParams) *trakt.CalendarShowIterator {
	return getC().MyNewShows(params)
}

func (c *Client) MyNewShows(params *trakt.CalendarParams) *trakt.CalendarShowIterator {
	return c.newShows("my", &wrappedCalendarParams{*params})
}

func MySeasonPremieres(params *trakt.CalendarParams) *trakt.CalendarShowIterator {
	return getC().MySeasonPremieres(params)
}

func (c *Client) MySeasonPremieres(params *trakt.CalendarParams) *trakt.CalendarShowIterator {
	return c.seasonPremieres("my", &wrappedCalendarParams{*params})
}

func Shows(params *trakt.BasicCalendarParams) *trakt.CalendarShowIterator {
	return getC().Shows(params)
}

func (c *Client) Shows(params *trakt.BasicCalendarParams) *trakt.CalendarShowIterator {
	return c.shows("all", &wrappedBasicCalendarParams{*params})
}

func NewShows(params *trakt.BasicCalendarParams) *trakt.CalendarShowIterator {
	return getC().NewShows(params)
}

func (c *Client) NewShows(params *trakt.BasicCalendarParams) *trakt.CalendarShowIterator {
	return c.newShows("all", &wrappedBasicCalendarParams{*params})
}

func SeasonPremieres(params *trakt.BasicCalendarParams) *trakt.CalendarShowIterator {
	return getC().SeasonPremieres(params)
}

func (c *Client) SeasonPremieres(params *trakt.BasicCalendarParams) *trakt.CalendarShowIterator {
	return c.seasonPremieres("all", &wrappedBasicCalendarParams{*params})
}

func (c *Client) shows(action string, params calendarParams) *trakt.CalendarShowIterator {
	return c.generateShowIterator(
		trakt.FormatURLPath("/calendars/%s/shows", action),
		params,
	)
}

func (c *Client) newShows(action string, params calendarParams) *trakt.CalendarShowIterator {
	return c.generateShowIterator(
		trakt.FormatURLPath("/calendars/%s/shows/new", action),
		params,
	)
}

func (c *Client) seasonPremieres(action string, params calendarParams) *trakt.CalendarShowIterator {
	return c.generateShowIterator(
		trakt.FormatURLPath("/calendars/%s/shows/premieres", action),
		params,
	)
}

type calendarParams interface {
	trakt.ListParamsContainer
	startDate() time.Time
	days() int
}

type wrappedCalendarParams struct{ trakt.CalendarParams }

func (w *wrappedCalendarParams) startDate() time.Time { return w.StartDate }
func (w *wrappedCalendarParams) days() int            { return int(w.Days) }

type wrappedBasicCalendarParams struct{ trakt.BasicCalendarParams }

func (w *wrappedBasicCalendarParams) startDate() time.Time { return w.StartDate }
func (w *wrappedBasicCalendarParams) days() int            { return int(w.Days) }

func formatPath(path string, c calendarParams) string {
	var days = c.days()
	if days == 0 {
		days = 7
	}

	return trakt.FormatURLPath(path+"/%s/%s", c.startDate().Format("2006-01-02"), strconv.Itoa(days))
}

func (c *Client) generateShowIterator(path string, params calendarParams) *trakt.CalendarShowIterator {
	return &trakt.CalendarShowIterator{
		Iterator: trakt.NewIterator(params, func(p trakt.ListParamsContainer) (trakt.IterationFrame, error) {
			list := make([]*trakt.CalendarShow, 0)
			f := trakt.NewEmptyFrame(&list)
			err := c.B.CallWithFrame(http.MethodGet, formatPath(path, params), c.Key, p, f)
			return f, err
		}),
	}
}

func getC() *Client {
	return &Client{
		B:   trakt.GetBackend(),
		Key: trakt.Key,
	}
}
