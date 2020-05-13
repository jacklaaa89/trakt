package trakt

// All the all constant, this is shared between different types
// so cant be allocated to a specific filter.
const All string = `all`

type ExtendedType string

func (e ExtendedType) String() string { return string(e) }

const (
	GuestStars ExtendedType = `guest_stars`
	Episodes   ExtendedType = `episodes`
	Metadata   ExtendedType = `metadata`
	NoSeasons  ExtendedType = `noseasons`
	Vip        ExtendedType = `vip`
	Full       ExtendedType = `full`
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
