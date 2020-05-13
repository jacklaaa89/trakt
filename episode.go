package trakt

import "time"

type Episode struct {
	commonElements `json:",inline"`

	Season     int64     `json:"season"`
	Number     int64     `json:"number"`
	Absolute   int64     `json:"absolute"`
	FirstAired time.Time `json:"first_aired"`
}
