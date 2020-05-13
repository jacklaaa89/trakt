package trakt

import (
	"strconv"
	"time"
)

type Status string

const (
	StatusReturningSeries Status = `returning series`
	StatusInProduction    Status = `in production`
	StatusPlanned         Status = `planned`
	StatusCancelled       Status = `canceled`
	StatusEnded           Status = `ended`
	StatusPostProduction  Status = `post production`
	StatusRumored         Status = `rumored`
	StatusReleased        Status = `released`
)

// ExtendedParams params which can be used when a request supports
// asking for additional information.
type ExtendedParams struct {
	BasicParams

	Extended ExtendedType `url:"extended"`
}

// ExtendedParams params which can be used when a request supports
// asking for additional information.
type ExtendedListParams struct {
	BasicListParams

	Extended ExtendedType `url:"extended"`
}

type SearchID interface {
	id() string
	path() string
}

type ID int64

func (i ID) id() string   { return strconv.Itoa(int(i)) }
func (i ID) path() string { return "trakt" }

type Slug string

func (s Slug) id() string   { return string(s) }
func (s Slug) path() string { return "trakt" }

type IMDB string

func (i IMDB) id() string   { return string(i) }
func (i IMDB) path() string { return "imdb" }

type TVDB int64

func (t TVDB) id() string   { return strconv.Itoa(int(t)) }
func (t TVDB) path() string { return "tvdb" }

type TMDB int64

func (t TMDB) id() string   { return strconv.Itoa(int(t)) }
func (t TMDB) path() string { return "tmdb" }

type baseIDs struct {
	Slug Slug `json:"slug"`
}

type objectIds struct {
	baseIDs
	Trakt ID `json:"trakt"`
}

type mediaIDs struct {
	Slug   Slug `json:"slug,omitempty"`
	Trakt  ID   `json:"trakt,omitempty"`
	TVDB   TVDB `json:"tvdb,omitempty"`
	IMDB   IMDB `json:"imdb,omitempty"`
	TMDB   TMDB `json:"tmdb,omitempty"`
	TVRage int  `json:"tvrage,omitempty"`
}

type Type string

func (t Type) String() string { return string(t) }

const (
	TypeMovie   Type = `movie`
	TypeShow    Type = `show`
	TypeSeason  Type = `season`
	TypeEpisode Type = `episode`
	TypeList    Type = `list`
	TypePerson  Type = `person`
)

type SharingParams struct {
	Twitter bool `json:"twitter"`
	Tumblr  bool `json:"tumblr"`
	Medium  bool `json:"medium"`
}

type GenericElementParams struct {
	mediaIDs `json:"ids"`

	Title string `json:"title,omitempty"`
	Year  int64  `json:"year,omitempty"`
}

type commonElements struct {
	mediaIDs `json:"ids"`

	Title                 string    `json:"title"`
	Overview              string    `json:"overview"`
	Rating                float64   `json:"rating"`
	Votes                 int64     `json:"votes"`
	Language              string    `json:"language"`
	AvailableTranslations []string  `json:"available_translations"`
	Runtime               int64     `json:"runtime"`
	UpdatedAt             time.Time `json:"updated_at"`
	Comments              int64     `json:"comment_count"`
}

type topLevelMediaElement struct {
	Type Type `json:"type"`

	Show  *Show  `json:"show"`
	Movie *Movie `json:"movie"`
}

type GenericMediaElement struct {
	topLevelMediaElement
	Episode *Episode `json:"episode"`
}

type GenericMediaElementIterator struct{ Iterator }

func (li *GenericMediaElementIterator) Type() Type {
	return li.Current().(*GenericMediaElement).Type
}

func (li *GenericMediaElementIterator) Show() *Show {
	return li.Current().(*GenericMediaElement).Show
}

func (li *GenericMediaElementIterator) Movie() *Movie {
	return li.Current().(*GenericMediaElement).Movie
}

type GenericElement struct {
	GenericMediaElement

	List *List `json:"list"`
}

type ListByTypeParams struct {
	BasicParams

	Type Type `json:"type"`
}

// IDPath helper function to format a search id into a search URL.
func IDPath(id SearchID) string { return "/search/" + id.path() + "/%s" }
