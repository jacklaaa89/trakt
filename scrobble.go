package trakt

import (
	"encoding/json"
	"strings"
)

type Action string

const (
	StartAction    Action = "start"
	PauseAction    Action = "pause"
	ScrobbleAction Action = "scrobble"
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

type Scrobble struct {
	basePlaybackItem

	Action   Action  `json:"action"`
	Progress float64 `json:"progress"`
}
