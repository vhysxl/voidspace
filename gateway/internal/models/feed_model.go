package models

import "time"

type GetFollowFeedReq struct {
	UserIDs      []int32   `json:"userid" validate:"required"`
	CursorUserID int       `json:"cursorid" validate:"required"`
	Cursor       time.Time `json:"cursor" validate:"required"`
}

type GetGlobalFeed struct {
	Cursor       time.Time `json:"cursor" validate:"required"`
	CursorUserID int       `json:"cursorid" validate:"required"`
}

type GetFeedResponse struct {
	Posts   []Post `json:"posts"`
	HasMore bool   `json:"hasmore"`
}
