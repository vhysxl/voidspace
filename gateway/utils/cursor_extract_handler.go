package utils

import (
	"strconv"
	"time"
)

func ExtractCursor(cursor string, cursorID string) (time.Time, int) {

	var cursorTime time.Time
	var cursorIDInt int

	// Parse cursor if provided
	if cursor != "" {
		// Try parsing as RFC3339 time string
		if parsedTime, err := time.Parse(time.RFC3339, cursor); err == nil {
			cursorTime = parsedTime
		} else {
			// If not valid RFC3339, try parsing as Unix timestamp
			if timestamp, err := strconv.ParseInt(cursor, 10, 64); err == nil {
				parsedTime := time.Unix(timestamp, 0)
				cursorTime = parsedTime
			}
		}
	}

	// Parse cursorID if provided
	if cursorID != "" {
		if id, err := strconv.Atoi(cursorID); err == nil {
			cursorIDInt = id
		}
	}

	return cursorTime, cursorIDInt
}
