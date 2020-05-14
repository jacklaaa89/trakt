package trakt

import (
	"encoding/json"
	"time"
)

type GetCollectionParams struct {
	ListParams

	Type     Type         `json:"-" url:"-"`
	Extended ExtendedType `url:"extended" json:"-"`
}

// Metadata to assign to the collection object.
type Metadata struct {
	Type          MediaType  `json:"media_type,omitempty"`
	Resolution    Resolution `json:"resolution,omitempty"`
	HDR           HDR        `json:"hdr,omitempty"`
	Audio         Audio      `json:"audio,omitempty"`
	AudioChannels string     `json:"audio_channels,omitempty"`
	ThreeD        bool       `json:"3d,omitempty"`
}

type MediaCollectionParams struct {
	Metadata

	IDs         MediaIDs  `json:"ids,omitempty" url:"-"`
	Title       string    `json:"title,omitempty" url:"-"`
	Year        int64     `json:"year,omitempty" url:"-"`
	CollectedAt time.Time `json:"collected_at,omitempty" url:"-"`
}

type EpisodeCollectionParams struct {
	Metadata

	Number      int64     `json:"number"`
	CollectedAt time.Time `json:"collected_at,omitempty" url:"-"`
}

type SeasonCollectionParams struct {
	Metadata

	Number      int64                      `json:"number"`
	CollectedAt time.Time                  `json:"collected_at,omitempty" url:"-"`
	Episodes    []*EpisodeCollectionParams `json:"episodes,omitempty"`
}

type ShowCollectionParams struct {
	Metadata

	IDs         MediaIDs                  `json:"ids,omitempty" url:"-"`
	Title       string                    `json:"title,omitempty" url:"-"`
	Year        int64                     `json:"year,omitempty" url:"-"`
	CollectedAt time.Time                 `json:"collected_at,omitempty" url:"-"`
	Seasons     []*SeasonCollectionParams `json:"seasons,omitempty" url:"-"`
}

type AddToCollectionParams struct {
	Params

	Movies   []*MediaCollectionParams `json:"movies,omitempty" url:"-"`
	Seasons  []*MediaCollectionParams `json:"seasons,omitempty" url:"-"`
	Episodes []*MediaCollectionParams `json:"episodes,omitempty" url:"-"`
	Shows    []*ShowCollectionParams  `json:"shows,omitempty" url:"-"`
}

type MediaRemovalParams struct {
	IDs   MediaIDs `json:"ids,omitempty" url:"-"`
	Title string   `json:"title,omitempty" url:"-"`
	Year  int64    `json:"year,omitempty" url:"-"`
}

type SeasonRemovalParams struct {
	Number   int64   `json:"number" url:"-"`
	Episodes []int64 `json:"-" url:"-"`
}

func (s *SeasonRemovalParams) MarshalJSON() ([]byte, error) {
	type A SeasonRemovalParams
	type B struct {
		Number int64 `json:"number"`
	}
	type C struct {
		A
		Episodes []*B `json:"episodes,omitempty"`
	}

	var c = &C{A: A{Number: s.Number}, Episodes: make([]*B, len(s.Episodes))}
	for idx, epNo := range s.Episodes {
		c.Episodes[idx] = &B{epNo}
	}

	return json.Marshal(c)
}

type ShowRemovalParams struct {
	IDs     MediaIDs               `json:"ids,omitempty" url:"-"`
	Title   string                 `json:"title,omitempty" url:"-"`
	Year    int64                  `json:"year,omitempty" url:"-"`
	Seasons []*SeasonRemovalParams `json:"seasons,omitempty"`
}

type RemoveFromCollectionParams struct {
	Params

	Movies   []*MediaRemovalParams `json:"movies,omitempty" url:"-"`
	Seasons  []*MediaRemovalParams `json:"seasons,omitempty" url:"-"`
	Episodes []*MediaRemovalParams `json:"episodes,omitempty" url:"-"`
	Shows    []*ShowRemovalParams  `json:"shows,omitempty" url:"-"`
}

type commonMediaActivity struct {
	LastRated       time.Time `json:"rated_at"`
	LastWatchListed time.Time `json:"watchlisted_at"`
	LastCommented   time.Time `json:"commented_at"`
}

type ShowActivity struct {
	commonMediaActivity
	LastHidden time.Time `json:"hidden_at"`
}

type SeasonActivity = ShowActivity

type EpisodeActivity struct {
	commonMediaActivity
	LastWatched   time.Time `json:"watched_at"`
	LastCollected time.Time `json:"collected_at"`
	LastPaused    time.Time `json:"paused_at"`
}

type MovieActivity struct {
	EpisodeActivity
	LastHidden time.Time `json:"hidden_at"`
}

type CommentActivity struct {
	LastLiked time.Time `json:"liked_at"`
}

type ListActivity struct {
	CommentActivity
	LastUpdated   time.Time `json:"updated_at"`
	LastCommented time.Time `json:"commented_at"`
}

type AccountActivity struct {
	LastUpdatedSettings time.Time `json:"settings_at"`
}

type LastActivity struct {
	LastUpdated time.Time        `json:"all"`
	Account     *AccountActivity `json:"account"`
	Lists       *ListActivity    `json:"lists"`
	Comments    *CommentActivity `json:"comments"`
	Seasons     *SeasonActivity  `json:"seasons"`
	Shows       *ShowActivity    `json:"shows"`
	Episodes    *EpisodeActivity `json:"episodes"`
	Movies      *MovieActivity   `json:"movies"`
}

type MediaType string

const (
	DigitalMedia MediaType = "digital"
	Bluray       MediaType = "bluray"
	HDDVD        MediaType = "hddvd"
	VCD          MediaType = "vcd"
	VHS          MediaType = "vhs"
	DVD          MediaType = "dvd"
	BetaMax      MediaType = "betamax"
	LaserDisc    MediaType = "laserdisc"
)

type Resolution string

const (
	UHD     Resolution = "uhd_4k"
	HD1080p Resolution = "hd_1080p"
	HD1080i Resolution = "hd_1080i"
	HD720p  Resolution = "hd_720p"
	SD480p  Resolution = "sd_480p"
	SD480i  Resolution = "sd_480i"
	SD576p  Resolution = "sd_576p"
	SD576i  Resolution = "sd_576i"
)

type HDR string

const (
	DolbyVision HDR = "dolby_vision"
	HDR10       HDR = "hdr10"
	HDR10Plus   HDR = "hdr10_plus"
	HLG         HDR = "hlg"
)

type Audio string

const (
	LPCM             Audio = "lpcm"
	MP3              Audio = "mp3"
	AAC              Audio = "acc"
	OGG              Audio = "ogg"
	WMA              Audio = "wma"
	DTS              Audio = "dts"
	DTSMA            Audio = "dts_ma"
	DTSHR            Audio = "dts_hr"
	DTSX             Audio = "dts_x"
	Auro3D           Audio = "auro_3d"
	DolbyDigital     Audio = "dolby_digital"
	DolbyDigitalPlus Audio = "dolby_digital_plus"
	DolbyAtmos       Audio = "dolby_atmos"
	DolbyTrueHD      Audio = "dolby_truehd"
	DolbyPrologic    Audio = "dolby_prologic"
)

type CollectedMovie struct {
	Movie       *Movie    `json:"movie"`
	CollectedAt time.Time `json:"collected_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Metadata    *Metadata `json:"metadata"`
}

type CollectionIterator interface {
	BasicIterator

	Type() Type
	Show() *CollectedShow
	Movie() *CollectedMovie
}

type CollectedEpisode struct {
	Number      int64     `json:"number"`
	CollectedAt time.Time `json:"collected_at"`
	Metadata    *Metadata `json:"metadata"`
}

type CollectedSeason struct {
	Number   int64 `json:"number"`
	Episodes []*CollectedEpisode
}

type CollectedShow struct {
	LastCollectedAt time.Time          `json:"last_collected_at"`
	LastUpdatedAt   time.Time          `json:"last_updated_at"`
	Show            *Show              `json:"show"`
	Seasons         []*CollectedSeason `json:"seasons"`
}

type CollectedShowIterator struct{ BasicIterator }

func (c *CollectedShowIterator) CollectedShow() *CollectedShow { return c.Current().(*CollectedShow) }

type NotFound struct {
	Movies   []*GenericElementParams `json:"movies"`
	Shows    []*GenericElementParams `json:"shows"`
	Seasons  []*GenericElementParams `json:"seasons"`
	Episodes []*GenericElementParams `json:"episodes"`
}

type ChangeSet struct {
	Movies   int64 `json:"movies"`
	Episodes int64 `json:"episodes"`
}

type AddToCollectionResult struct {
	Added    *ChangeSet `json:"added"`
	Updated  *ChangeSet `json:"updated"`
	Existing *ChangeSet `json:"existing"`
	NotFound *NotFound  `json:"not_found"`
}

type RemoveFromCollectionResult struct {
	Deleted  *ChangeSet `json:"deleted"`
	NotFound *NotFound  `json:"not_found"`
}
