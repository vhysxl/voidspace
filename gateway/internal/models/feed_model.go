package models

import "time"

type GetFollowFeedReq struct {
	CursorID int       `json:"cursorid"`
	Cursor   time.Time `json:"cursor"`
}

type GetGlobalFeedReq struct {
	Cursor   time.Time `json:"cursor"`
	CursorID int       `json:"cursorid"`
}

type GetFeedResponse struct {
	Posts   []Post `json:"posts"`
	HasMore bool   `json:"hasmore"`
}
