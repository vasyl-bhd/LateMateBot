package utils

import (
	"time"
)

func StartOfPastWeek() time.Time {
	// Get the current time
	now := time.Now()

	// Find the start of the current week (assuming the week starts on Monday)
	startOfWeek := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	for startOfWeek.Weekday() != time.Monday {
		startOfWeek = startOfWeek.AddDate(0, 0, -1)
	}

	// Subtract 7 days to get the start of the previous week
	startOfPastWeek := startOfWeek.AddDate(0, 0, -7)

	return startOfPastWeek
}
