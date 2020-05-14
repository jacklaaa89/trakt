package trakt

import (
	"encoding/json"
	"strings"
	"time"
)

type StartCheckinParams struct {
	Params

	Type    Type                  `json:"-"`
	Element *GenericElementParams `json:"-"`

	Message    string         `json:"message,omitempty"`
	VenueID    string         `json:"venue_id,omitempty"`
	VenueName  string         `json:"venue_name,omitempty"`
	Sharing    *SharingParams `json:"sharing,omitempty"`
	AppVersion string         `json:"app_version,omitempty"`
	AppDate    string         `json:"app_date,omitempty"`
}

func (s *StartCheckinParams) MarshalJSON() ([]byte, error) {
	m := marshalToMap(s)
	m[strings.ToLower(string(s.Type))] = s.Element
	return json.Marshal(m)
}

type Checkin struct {
	basePlaybackItem
	WatchedAt time.Time `json:"watched_at"`
}

func (c *Checkin) UnmarshalJSON(bytes []byte) error {
	type A Checkin
	var a = new(A)
	err := json.Unmarshal(bytes, a)
	if err != nil {
		return err
	}

	switch {
	case a.Episode != nil:
		a.Type = TypeEpisode
	case a.Movie != nil:
		a.Type = TypeMovie
	}

	*c = Checkin(*a)
	return nil
}
