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
	Language      string    `json:"language"`
}

type ShowIterator struct{ Iterator }

func (li *ShowIterator) Show() *Show { return li.Current().(*Show) }

type TrendingShow struct {
	Show     *Show `json:"show"`
	Watchers int64 `json:"watchers"`
}

type TrendingShowIterator struct{ Iterator }

func (m *TrendingShowIterator) Trending() *TrendingShow { return m.Current().(*TrendingShow) }

type RecentlyUpdatedShow struct {
	Show      *Show     `json:"show"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RecentlyUpdatedShowIterator struct{ Iterator }

func (m *RecentlyUpdatedShowIterator) Show() *RecentlyUpdatedShow {
	return m.Current().(*RecentlyUpdatedShow)
}

type ShowWithStatistics struct {
	statistics
	Show *Show `json:"movie"`
}

type ShowWithStatisticsIterator struct{ Iterator }

func (m *ShowWithStatisticsIterator) Show() *ShowWithStatistics {
	return m.Current().(*ShowWithStatistics)
}

type AnticipatedShow struct {
	ListCount int64  `json:"list_count"`
	Movie     *Movie `json:"movie"`
}

type AnticipatedShowIterator struct{ Iterator }

func (m *AnticipatedMovieIterator) Show() *AnticipatedShow { return m.Current().(*AnticipatedShow) }
