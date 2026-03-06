package models

import "time"

type GetFollowFeedRequest struct {
	CursorID int       `json:"cursor_id"`
	Cursor   time.Time `json:"cursor"`
}

type GetGlobalFeedRequest struct {
	CursorID int       `json:"cursor_id"`
	Cursor   time.Time `json:"cursor"`
}

type GetFeedResponse struct {
	Posts   []Post `json:"posts"`
	HasMore bool   `json:"has_more"`
}
