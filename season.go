package trakt

import "time"

type Season struct {
	MediaIDs `json:"ids"`

	Title      string    `json:"title"`
	Overview   string    `json:"overview"`
	Number     int64     `json:"number"`
	Rating     float64   `json:"rating"`
	Votes      int64     `json:"votes"`
	Episodes   int64     `json:"episodes"`
	Aired      int64     `json:"aired"`
	FirstAired time.Time `json:"first_aired"`
	Network    string    `json:"network"`
}

type SeasonIterator struct{ Iterator }

func (s *SeasonIterator) Season() *Season { return s.Current().(*Season) }
