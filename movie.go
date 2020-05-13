package trakt

import (
	"encoding/json"
	"time"
)

type ReleaseType string

const (
	ReleaseTypeUnknown    ReleaseType = "unknown"
	ReleaseTypePremiere   ReleaseType = "premiere"
	ReleaseTypeLimited    ReleaseType = "limited"
	ReleaseTypeTheatrical ReleaseType = "theatrical"
	ReleaseTypeDigital    ReleaseType = "digital"
	ReleaseTypePhysical   ReleaseType = "physical"
	ReleaseTypeTV         ReleaseType = "tv"
)

type TimePeriodListParams struct {
	BasicListParams

	Period   TimePeriod   `url:"-" json:"-"`
	Extended ExtendedType `url:"extended" json:"-"`
}

type BoxOfficeListParams = ExtendedParams

type RecentlyUpdatedListParams struct {
	BasicListParams

	StartDate time.Time    `json:"-" url:"-"`
	Extended  ExtendedType `url:"extended" json:"-"`
}

type ReleaseListParams struct {
	BasicListParams

	Country string `json:"-" url:"-"`
}

type Movie struct {
	CommonElements `json:",inline"`

	Year          uint      `json:"year"`
	Tagline       string    `json:"tagline"`
	Released      time.Time `json:"released"`
	Certification string    `json:"certification"`
	Country       string    `json:"country"`
	TrailerURL    string    `json:"trailer"`
	HomepageURL   string    `json:"homepage"`
	Status        Status    `json:"status"`
	Genres        []string  `json:"genre"`
}

func (m *Movie) UnmarshalJSON(bytes []byte) error {
	type B Movie
	type A struct {
		B
		Released string `json:"released"`
	}

	var a = new(A)
	err := json.Unmarshal(bytes, a)
	if err != nil {
		return err
	}

	if len(a.Released) > 0 {
		a.B.Released, err = time.Parse(`2006-01-02`, a.Released)
	}

	*m = Movie(a.B)
	return err
}

type MovieIterator struct{ Iterator }

func (m *MovieIterator) Movie() *Movie { return m.Current().(*Movie) }

type Release struct {
	Country       string      `json:"country"`
	Certification string      `json:"certification"`
	ReleaseDate   time.Time   `json:"-"`
	ReleaseType   ReleaseType `json:"release_type"`
	Note          string      `json:"note"`
}

func (r *Release) UnmarshalJSON(bytes []byte) error {
	type B Release
	type A struct {
		B
		Released string `json:"release_date"`
	}

	var a = new(A)
	err := json.Unmarshal(bytes, a)
	if err != nil {
		return err
	}

	if len(a.Released) > 0 {
		a.B.ReleaseDate, err = time.Parse(`2006-01-02`, a.Released)
	}

	*r = Release(a.B)
	return err
}

type ReleaseIterator struct{ Iterator }

func (m *ReleaseIterator) Release() *Release { return m.Current().(*Release) }

type TrendingMovie struct {
	Movie    *Movie `json:"movie"`
	Watchers int64  `json:"watchers"`
}

type TrendingMovieIterator struct{ Iterator }

func (m *TrendingMovieIterator) Trending() *TrendingMovie { return m.Current().(*TrendingMovie) }

type RecentlyUpdatedMovie struct {
	Movie     *Movie    `json:"movie"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RecentlyUpdatedMovieIterator struct{ Iterator }

func (m *RecentlyUpdatedMovieIterator) Movie() *RecentlyUpdatedMovie {
	return m.Current().(*RecentlyUpdatedMovie)
}

type MovieWithStatistics struct {
	BasicStatistics
	Movie *Movie `json:"movie"`
}

type MovieWithStatisticsIterator struct{ Iterator }

func (m *MovieWithStatisticsIterator) Movie() *MovieWithStatistics {
	return m.Current().(*MovieWithStatistics)
}

type AnticipatedMovie struct {
	ListCount int64  `json:"list_count"`
	Movie     *Movie `json:"movie"`
}

type AnticipatedMovieIterator struct{ Iterator }

func (m *AnticipatedMovieIterator) Movie() *AnticipatedMovie { return m.Current().(*AnticipatedMovie) }

type BoxOfficeMovie struct {
	Revenue int64  `json:"revenue"`
	Movie   *Movie `json:"movie"`
}

type BoxOfficeMovieIterator struct{ Iterator }

func (m *BoxOfficeMovieIterator) Movie() *BoxOfficeMovie { return m.Current().(*BoxOfficeMovie) }
