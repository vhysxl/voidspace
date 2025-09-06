package models

import "time"

type GetFollowFeedReq struct {
	UserIDs      []int32   `json:"userid" validate:"required"`
	CursorUserID int       `json:"cursorid"`
	Cursor       time.Time `json:"cursor"`
}

type GetGlobalFeedReq struct {
	Cursor       time.Time `json:"cursor"`
	CursorUserID int       `json:"cursorid"`
}

type GetFeedResponse struct {
	Posts   []Post `json:"posts"`
	HasMore bool   `json:"hasmore"`
}
