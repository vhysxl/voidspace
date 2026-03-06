package models

import "time"

type CreateCommentRequest struct {
	Content string `json:"content" validate:"required,min=1,max=100"`
	PostID  int    `json:"post_id" validate:"required"`
}

type Comment struct {
	CommentID int       `json:"comment_id"`
	PostID    int       `json:"post_id"`
	Content   string    `json:"content"`
	Author    *User     `json:"author"`
	CreatedAt time.Time `json:"created_at"`
}
