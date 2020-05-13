package trakt

import (
	"encoding/json"
	"time"
)

type Person struct {
	MediaIDs `json:"ids"`

	Name       string    `json:"name"`
	Biography  string    `json:"biography"`
	Birthday   time.Time `json:"-"`
	Death      time.Time `json:"-"`
	Homepage   string    `json:"homepage"`
	Birthplace string    `json:"birthplace"`
}

func (p *Person) UnmarshalJSON(bytes []byte) error {
	type B Person
	type A struct {
		B
		Birthday string `json:"birthday"`
		Death    string `json:"death"`
	}

	var a = new(A)
	err := json.Unmarshal(bytes, a)
	if err != nil {
		return err
	}

	var pErr error
	if len(a.Birthday) > 0 {
		a.B.Birthday, pErr = time.Parse(`2006-01-02`, a.Birthday)
		if pErr != nil {
			return pErr
		}
	}

	if len(a.Death) > 0 {
		a.B.Death, pErr = time.Parse(`2006-01-02`, a.Death)
		if pErr != nil {
			return pErr
		}
	}

	*p = Person(a.B)
	return nil
}
