package trakt

// All the all constant, this is shared between different types
// so cant be allocated to a specific filter.
const All string = `all`

type ExtendedType string

func (e ExtendedType) String() string { return string(e) }

const (
	ExtendedTypeGuestStars ExtendedType = `guest_stars`
	ExtendedTypeEpisodes   ExtendedType = `episodes`
	ExtendedTypeMetadata   ExtendedType = `metadata`
	ExtendedTypeNoSeasons  ExtendedType = `noseasons`
	ExtendedTypeVip        ExtendedType = `vip`
	ExtendedTypeFull       ExtendedType = `full`
)

type TimePeriod string

func (t TimePeriod) String() string { return string(t) }

const (
	TimePeriodWeekly  TimePeriod = "weekly"
	TimePeriodMonthly TimePeriod = "monthly"
	TimePeriodYearly  TimePeriod = "yearly"
	TimePeriodAll                = TimePeriod(All)
)

type CommentSortType string

func (c CommentSortType) String() string { return string(c) }

const (
	SortTypeNewest  CommentSortType = "newest"
	SortTypeOldest  CommentSortType = "oldest"
	SortTypeReplies CommentSortType = "replies"
	SortTypeHighest CommentSortType = "highest"
	SortTypeLowest  CommentSortType = "lowest"
	SortTypePlays   CommentSortType = "plays"
)
