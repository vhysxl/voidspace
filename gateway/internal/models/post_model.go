package models

import "time"

type PostRequest struct {
	Content    string   `json:"content" validate:"required,min=1,max=500"`
	PostImages []string `json:"post_images" validate:"omitempty,dive,url"`
}

type GetPostRequest struct {
	ID int `validate:"required,gt=0"`
}

type Post struct {
	ID         int       `json:"id"`
	Content    string    `json:"content"`
	UserID     int       `json:"user_id"`
	PostImages []string  `json:"post_images"`
	LikesCount int       `json:"likes_count"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Author     *User     `json:"author"`
}
