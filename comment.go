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

type PostCommentParams struct {
	Params

	Type    Type                  `json:"-"`
	Element *GenericElementParams `json:"-"`

	Text    string         `json:"comment"`
	Spoiler bool           `json:"spoiler"`
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

	Sort CommentSortType `json:"-" url:"-"`
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

func (li *CommentIterator) Comment() *Comment { return li.Current().(*Comment) }

type UserLike struct {
	User    *User     `json:"user"`
	LikedAt time.Time `json:"liked_at"`
}

type UserLikeIterator struct{ Iterator }

func (li *UserLikeIterator) UserLike() *UserLike { return li.Current().(*UserLike) }

type CommentWithMediaElement struct {
	GenericMediaElement
	Comment *Comment `json:"comment"`
}

type CommentWithMediaElementIterator struct{ GenericMediaElementIterator }

func (li *UserLikeIterator) CommentW() *Comment {
	return li.Current().(*CommentWithMediaElement).Comment
}
