package trakt

import (
	"encoding/json"
	"time"
)

type CalendarParams struct {
	ListParams

	StartDate time.Time    `url:"-"`
	Days      int64        `url:"-"`
	Extended  ExtendedType `url:"extended"`
}

type BasicCalendarParams struct {
	BasicListParams

	StartDate time.Time    `url:"-"`
	Days      int64        `url:"-"`
	Extended  ExtendedType `url:"extended"`
}

type CalendarShow struct {
	Show       *Show     `json:"show"`
	Episode    *Episode  `json:"episode"`
	FirstAired time.Time `json:"first_aired"`
}

type CalendarMovie struct {
	Movie    *Movie    `json:"movie"`
	Released time.Time `json:"-"`
}

func (c *CalendarMovie) UnmarshalJSON(bytes []byte) error {
	type B CalendarMovie
	type A struct {
		B
		Released string `json:"released"`
	}

	var a = new(A)
	err := json.Unmarshal(bytes, a)
	if err != nil {
		return err
	}

	a.B.Released, err = time.Parse(`2006-01-02`, a.Released)
	*c = CalendarMovie(a.B)
	return err
}

type CalendarShowIterator struct{ Iterator }

func (li *CalendarShowIterator) Entry() *CalendarShow { return li.Current().(*CalendarShow) }

type CalendarMovieIterator struct{ Iterator }

func (li *CommentIterator) Entry() *CalendarMovie { return li.Current().(*CalendarMovie) }
