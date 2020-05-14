package trakt

import (
	"fmt"
	"net/url"
)

// All the all constant, this is shared between different types
// so cant be allocated to a specific filter.
const All string = `all`

type ExtendedType string

func (e ExtendedType) String() string { return string(e) }

const (
	GuestStars         ExtendedType = `guest_stars`
	Episodes           ExtendedType = `episodes`
	CollectionMetadata ExtendedType = `metadata`
	NoSeasons          ExtendedType = `noseasons`
	Vip                ExtendedType = `vip`
	Full               ExtendedType = `full`
)

type TimePeriod string

func (t TimePeriod) String() string { return string(t) }

const (
	Weekly        TimePeriod = "weekly"
	Monthly       TimePeriod = "monthly"
	Yearly        TimePeriod = "yearly"
	TimePeriodAll            = TimePeriod(All)
)

type SortType string

func (s SortType) String() string { return string(s) }

const (
	// these sort types are used for comment sorting.
	Newest  SortType = "newest"
	Oldest  SortType = "oldest"
	Replies SortType = "replies"
	Highest SortType = "highest"
	Lowest  SortType = "lowest"
	Plays   SortType = "plays"

	// these sort types are used for list sorting.
	Popular  SortType = "popular"
	Likes    SortType = "likes"
	Comments SortType = "comments"
	Items    SortType = "items"
	Added    SortType = "added"
	Updated  SortType = "updated"
)

// Range represents a range type filter.
type Range struct {
	To   int64
	From int64
}

// EncodeValues implements Encoder interface
// takes a range and encodes it as a string in the form "<from>-<to>"
func (r *Range) EncodeValues(key string, v *url.Values) error {
	v.Set(key, fmt.Sprintf(`%d-%d`, r.From, r.To))
	return nil
}

// Filters represents the set of filters which can be applied to certain API
// endpoints.
type Filters struct {
	// Common filters.
	Query     string   `url:"query,omitempty" json:"-"`
	Years     []int64  `url:"years,comma,omitempty" json:"-"`
	Genres    []string `url:"genres,comma,omitempty" json:"-"`
	Languages []string `url:"languages,comma,omitempty" json:"-"`
	Countries []string `url:"countries,comma,omitempty" json:"-"`
	Runtime   *Range   `url:"runtimes,omitempty" json:"-"`
	Rating    *Range   `url:"ratings,comma,omitempty" json:"-"`

	// filters specific for movies and shows
	Certifications []string `url:"certifications,comma,omitempty" json:"-"`

	// filters specific for shows
	Networks []string `url:"networks,comma,omitempty" json:"-"`
	Statuses []Status `url:"statuses,comma,omitempty" json:"-"`
}
