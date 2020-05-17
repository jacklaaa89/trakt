// Package list contains functions which are capable of retrieving
// the most popular lists.
package list

import (
	"net/http"

	"github.com/jacklaaa89/trakt"
)

// Client represents a list client.
type Client struct{ b trakt.BaseClient }

// Trending returns all lists with the most likes and comments over the last 7 days.
//
// - Pagination
func Trending(params *trakt.BasicListParams) *trakt.RecentListIterator {
	return getC().Trending(params)
}

// Trending returns all lists with the most likes and comments over the last 7 days.
//
// - Pagination
func (c *Client) Trending(params *trakt.BasicListParams) *trakt.RecentListIterator {
	return c.generateListIterator("trending", params)
}

// Popular returns the most popular lists. Popularity is calculated using total number of
// likes and comments.
//
// - Pagination
func Popular(params *trakt.BasicListParams) *trakt.RecentListIterator {
	return getC().Popular(params)
}

// Popular returns the most popular lists. Popularity is calculated using total number of
// likes and comments.
//
// - Pagination
func (c *Client) Popular(params *trakt.BasicListParams) *trakt.RecentListIterator {
	return c.generateListIterator("popular", params)
}

// generateListIterator generates an iterator which retrieves a set of lists by action.
func (c *Client) generateListIterator(action string, params *trakt.BasicListParams) *trakt.RecentListIterator {
	path := trakt.FormatURLPath("/lists/%s", action)
	return &trakt.RecentListIterator{Iterator: c.b.NewIterator(http.MethodGet, path, params)}
}

// getC initialises a new list client with the current backend configuration.
func getC() *Client { return &Client{trakt.NewClient()} }
