// Package calender provides functionality to query a calendar or schedule
// for movies and shows.
//
// By default, the calendar will return all shows or movies for the specified time period
// and can be global or user specific. The start_date defaults to today and days to 7.
// The maximum amount of days you can send is 31. All dates (including the start_date and first_aired)
// are in UTC, so it's up to your app to handle any offsets based on the user's time zone.
// The my calendar displays episodes for all shows that have been watched, collected,
// or watch-listed plus individual episodes on the watchlist. It will remove any shows that have been
// hidden from the calendar. The all calendar displays info for all shows airing during the specified period.
package calendar
