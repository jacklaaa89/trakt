package trakt

import (
	"encoding/json"
	"strings"
)

type Action string

const (
	Start    Action = "start"
	Pause    Action = "pause"
	Scrobble Action = "scrobble"
)

type ScrobbleParams struct {
	Params

	Type    Type                  `json:"-" url:"-"`
	Element *GenericElementParams `json:"-" url:"-"`

	AppVersion string  `json:"app_version" url:"-"`
	AppDate    string  `json:"app_date" url:"-"`
	Progress   float64 `json:"progress" url:"-"`
}

func (s *ScrobbleParams) MarshalJSON() ([]byte, error) {
	m := marshalToMap(s)
	m[strings.ToLower(s.Type.String())] = s.Element
	return json.Marshal(m)
}

type ScrobbleEvent struct {
	basePlaybackItem

	Action   Action  `json:"action"`
	Progress float64 `json:"progress"`
}
