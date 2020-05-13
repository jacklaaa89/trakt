package trakt

import "encoding/json"

type basePlaybackItem struct {
	GenericMediaElement
	ID      int64          `json:"id"`
	Sharing *SharingParams `json:"sharing"`
}

func (b *basePlaybackItem) UnmarshalJSON(bytes []byte) error {
	type A basePlaybackItem
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

	*b = basePlaybackItem(*a)
	return nil
}
