package models

import "time"

type CreateCommentReq struct {
	Content string `json:"content" validate:"required,min=1,max=100"`
	PostID  int32  `json:"post_id" validate:"required"`
}

type Comments struct {
	CommentID int32     `json:"comment_id"`
	PostID    int32     `json:"post_id"`
	Content   string    `json:"content"`
	Author    *User     `json:"author"`
	CreatedAt time.Time `json:"created_at"`
}
