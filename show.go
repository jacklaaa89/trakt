package trakt

import "time"

type Airs struct {
	Day      string
	Time     string
	Timezone string
}

type Show struct {
	commonElements `json:",inline"`

	Year          int64     `json:"year"`
	FirstAired    time.Time `json:"first_aired"`
	Airs          *Airs     `json:"airs"`
	Certification string    `json:"certification"`
	Country       string    `json:"country"`
	TrailerURL    string    `json:"trailer"`
	HomepageURL   string    `json:"homepage"`
	Status        Status    `json:"status"`
	Genres        []string  `json:"genre"`
	AiredEpisodes int64     `json:"aired_episodes"`
}

type ShowIterator struct{ Iterator }

func (li *ShowIterator) Show() *Show { return li.Current().(*Show) }
