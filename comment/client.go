package comment

import (
	"net/http"

	"github.com/jackaaa89/trakt"
)

type Client struct {
	B   trakt.Backend
	Key string
}

// Get attempts to retrieve a comment by its id.
func Get(id string, params *trakt.BasicParams) (*trakt.Comment, error) {
	return getC().Get(id, params)
}

// Get attempts to retrieve a comment by its id.
func (c *Client) Get(id string, params *trakt.BasicParams) (*trakt.Comment, error) {
	path := trakt.FormatURLPath("/comments/%s", id)
	com := &trakt.Comment{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, com)
	return com, err
}

// Likes generates an iterator to retrieve the users which liked a comment by id.
func Likes(id string, params *trakt.BasicListParams) *trakt.UserLikeIterator {
	return getC().Likes(id, params)
}

// Likes generates an iterator to retrieve the users which liked a comment by id.
func (c *Client) Likes(id string, params *trakt.BasicListParams) *trakt.UserLikeIterator {
	return &trakt.UserLikeIterator{
		Iterator: trakt.NewIterator(params, func(p trakt.ListParamsContainer) (trakt.IterationFrame, error) {
			list := make([]*trakt.UserLike, 0)
			f := trakt.NewEmptyFrame(&list)
			err := c.B.CallWithFrame(http.MethodGet, trakt.FormatURLPath("/comments/%s/likes", id), c.Key, p, f)
			return f, err
		}),
	}
}

// Replies retrieves a list of replies attached to a comment by id.
func Replies(id string, params *trakt.ListParams) *trakt.CommentIterator {
	return getC().Replies(id, params)
}

// Replies retrieves a list of replies attached to a comment by id.
func (c *Client) Replies(id string, params *trakt.ListParams) *trakt.CommentIterator {
	return &trakt.CommentIterator{
		Iterator: trakt.NewIterator(params, func(p trakt.ListParamsContainer) (trakt.IterationFrame, error) {
			list := make([]*trakt.Comment, 0)
			f := trakt.NewEmptyFrame(&list)
			err := c.B.CallWithFrame(http.MethodGet, trakt.FormatURLPath("/comments/%s/replies", id), c.Key, p, f)
			return f, err
		}),
	}
}

// Item attempts to retrieve the item associated with a comment.
func Item(id string, params *trakt.ExtendedParams) (*trakt.GenericElement, error) {
	return getC().Item(id, params)
}

// Item attempts to retrieve the item associated with a comment.
func (c *Client) Item(id string, params *trakt.ExtendedParams) (*trakt.GenericElement, error) {
	path := trakt.FormatURLPath("/comments/%s/item", id)
	com := &trakt.GenericElement{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, com)
	return com, err
}

// Trending attempts to retrieve the currently trending comments.
func Trending(params *trakt.TrendingCommentParams) *trakt.CommentWithMediaElementIterator {
	return getC().Trending(params)
}

// Trending attempts to retrieve the currently trending comments.
func (c *Client) Trending(params *trakt.TrendingCommentParams) *trakt.CommentWithMediaElementIterator {
	return c.generateIterator(`trending`, params)
}

// Recent attempts to retrieve the most recent comments.
func Recent(params *trakt.RecentCommentParams) *trakt.CommentWithMediaElementIterator {
	return getC().Recent(params)
}

// Recent attempts to retrieve the most recent comments.
func (c *Client) Recent(params *trakt.RecentCommentParams) *trakt.CommentWithMediaElementIterator {
	return c.generateIterator(`recent`, params)
}

// Updates attempts to retrieve the most recently updated comments.
func Updates(params *trakt.UpdatedCommentParams) *trakt.CommentWithMediaElementIterator {
	return getC().Updates(params)
}

// Updates attempts to retrieve the most recently updated comments.
func (c *Client) Updates(params *trakt.UpdatedCommentParams) *trakt.CommentWithMediaElementIterator {
	return c.generateIterator(`updates`, params)
}

// Post attempts to post a new comment on an item.
func Post(params *trakt.PostCommentParams) (*trakt.Comment, error) {
	return getC().Post(params)
}

// Post attempts to post a new comment on an item.
func (c *Client) Post(params *trakt.PostCommentParams) (*trakt.Comment, error) {
	com := &trakt.Comment{}
	err := c.B.Call(http.MethodPost, "/comments", c.Key, &wrappedPostCommentParams{*params}, &com)
	return com, err
}

// Update attempts to update a comment on an item.
func Update(id string, params *trakt.UpdateCommentParams) (*trakt.Comment, error) {
	return getC().Update(id, params)
}

// Update attempts to update a comment on an item.
func (c *Client) Update(id string, params *trakt.UpdateCommentParams) (*trakt.Comment, error) {
	com := &trakt.Comment{}
	err := c.B.Call(
		http.MethodPut,
		trakt.FormatURLPath("/comments/%s", id),
		c.Key,
		&wrappedUpdateCommentParams{*params},
		&com,
	)
	return com, err
}

// Remove attempts to remove a comment by id.
func Remove(id string, params *trakt.Params) error {
	return getC().Remove(id, params)
}

// Remove attempts to remove a comment by id.
func (c *Client) Remove(id string, params *trakt.Params) error {
	return c.B.Call(
		http.MethodDelete, trakt.FormatURLPath("/comment/%s", id),
		c.Key, &wrappedRemoveCommentParams{*params}, nil,
	)
}

// AddReply attempts to add a reply to a comment by id.
func AddReply(id string, params *trakt.AddReplyParams) (*trakt.Comment, error) {
	return getC().AddReply(id, params)
}

// AddReply attempts to add a reply to a comment by id.
func (c *Client) AddReply(id string, params *trakt.AddReplyParams) (*trakt.Comment, error) {
	com := &trakt.Comment{}
	err := c.B.Call(
		http.MethodPost,
		trakt.FormatURLPath("/comments/%s/replies", id),
		c.Key,
		&wrappedUpdateCommentParams{*params},
		&com,
	)
	return com, err
}

// AddLike attempts to add a like to a comment.
func AddLike(id string, params *trakt.Params) error {
	return getC().AddLike(id, params)
}

// AddLike attempts to add a like to a comment.
func (c *Client) AddLike(id string, params *trakt.Params) error {
	return c.B.Call(http.MethodPost, trakt.FormatURLPath("/comments/%s/like", id), c.Key, params, nil)
}

// RemoveLike attempts to remove as like on a comment.
func RemoveLike(id string, params *trakt.Params) error {
	return getC().RemoveLike(id, params)
}

// RemoveLike attempts to remove as like on a comment.
func (c *Client) RemoveLike(id string, params *trakt.Params) error {
	return c.B.Call(http.MethodDelete, trakt.FormatURLPath("/comments/%s/like", id), c.Key, params, nil)
}

type wrappedPostCommentParams struct{ trakt.PostCommentParams }

func (wrappedPostCommentParams) Code(statusCode int) trakt.ErrorCode {
	return commentErrorHandler(statusCode)
}

type wrappedUpdateCommentParams struct{ trakt.UpdateCommentParams }

func (wrappedUpdateCommentParams) Code(statusCode int) trakt.ErrorCode {
	return commentErrorHandler(statusCode)
}

type wrappedRemoveCommentParams struct{ trakt.Params }

func (wrappedRemoveCommentParams) Code(statusCode int) trakt.ErrorCode {
	return commentErrorHandler(statusCode)
}

func commentErrorHandler(statusCode int) trakt.ErrorCode {
	switch statusCode {
	case http.StatusUnauthorized:
		return trakt.ErrorCodePostInvalidUser
	case http.StatusNotFound:
		return trakt.ErrorCodePostInvalidItem
	case http.StatusConflict:
		return trakt.ErrorCodeCommentCannotBeRemoved
	}

	return trakt.DefaultErrorHandler.Code(statusCode)
}

// generateIterator generates an iterator for the functions
// - Trending
// - Recent
// - Updates
// as the only thing that changes is the action that is called in terms of arguments.
func (c *Client) generateIterator(action string, params *trakt.TrendingCommentParams) *trakt.CommentWithMediaElementIterator {
	return &trakt.CommentWithMediaElementIterator{
		GenericMediaElementIterator: trakt.GenericMediaElementIterator{
			Iterator: trakt.NewIterator(params, func(p trakt.ListParamsContainer) (trakt.IterationFrame, error) {
				list := make([]*trakt.CommentWithMediaElement, 0)
				var ct, mt = trakt.All, trakt.All
				if params.MediaType != "" {
					mt = string(params.MediaType)
				}
				if params.CommentType != "" {
					mt = string(params.CommentType)
				}

				f := trakt.NewEmptyFrame(&list)
				path := trakt.FormatURLPath("/comments/%s/%s/%s", action, ct, mt)
				err := c.B.CallWithFrame(http.MethodGet, path, c.Key, p, f)

				return f, err
			}),
		},
	}
}

func getC() *Client {
	return &Client{B: trakt.GetBackend(), Key: trakt.Key}
}
