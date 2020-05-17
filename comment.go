package trakt

import (
	"encoding/json"
	"strings"
	"time"
)

type CommentType string

const (
	CommentTypeUnspecified CommentType = `unspecified`
	CommentTypeReview      CommentType = `review`
	CommentTypeShout       CommentType = `shout`
	CommentTypeAll                     = CommentType(All)
)

// PostCommentParams parameters in order to post a new comment.
type PostCommentParams struct {
	Params

	// Type the type of media element. Can either be TypeMovie or TypeEpisode.
	Type Type `json:"-"`
	// Element the actual element data. We can provide as much or as little as
	// we want about the item. The recommended values are either the trakt ID or slug.
	Element *GenericElementParams `json:"-"`

	// Text for the comment. This supports markdown and emojis.
	// Emojis are declared as short codes like :smiley: and :raised_hands:.
	Text string `json:"comment"`
	// Spoiler represents whether the comment is a spoiler to the attached item.
	Spoiler bool `json:"spoiler"`

	// The sharing object is optional and will apply the user's settings if not sent.
	// If sharing is sent, each key will override the user's setting for that social network.
	// Send true to post or false to not post on the indicated social network. You can see which
	// social networks a user has connected with the /users/settings method.
	Sharing *SharingParams `json:"sharing,omitempty"`
}

func (p *PostCommentParams) MarshalJSON() ([]byte, error) {
	m := marshalToMap(p)
	m[strings.ToLower(string(p.Type))] = p.Element
	return json.Marshal(m)
}

type UpdateCommentParams struct {
	Params

	Text    string `json:"comment"`
	Spoiler bool   `json:"spoiler"`
}

type AddReplyParams = UpdateCommentParams

type TrendingCommentParams struct {
	BasicListParams

	ExtendedType ExtendedType `url:"extended,omitempty" json:"-"`
	CommentType  CommentType  `url:"-"  json:"-"`
	MediaType    Type         `url:"-"  json:"-"`
}

type RecentCommentParams = TrendingCommentParams
type UpdatedCommentParams = TrendingCommentParams

type CommentListParams struct {
	BasicListParams

	Sort SortType `json:"-" url:"-"`
}

type Comment struct {
	ID        int       `json:"id"`
	Text      string    `json:"comment"`
	Spoiler   bool      `json:"spoiler"`
	Review    bool      `json:"review"`
	Parent    int       `json:"parent_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Replies   int       `json:"replies"`
	Likes     int       `json:"likes"`
	User      *User     `json:"user"`
}

type CommentIterator struct{ Iterator }

func (li *CommentIterator) Comment() (*Comment, error) {
	rcv := &Comment{}
	return rcv, li.Scan(rcv)
}

type UserLike struct {
	User    `json:"user"`
	LikedAt time.Time `json:"liked_at"`
}

type UserLikeIterator struct{ Iterator }

func (li *UserLikeIterator) UserLike() (*UserLike, error) {
	rcv := &UserLike{}
	return rcv, li.Scan(rcv)
}

type CommentWithMediaElement struct {
	GenericMediaElement
	Comment *Comment `json:"comment"`
}

type CommentWithMediaElementIterator struct{ Iterator }

func (li *CommentWithMediaElementIterator) CommentWithMediaElement() (*CommentWithMediaElement, error) {
	rcv := &CommentWithMediaElement{}
	return rcv, li.Scan(rcv)
}
