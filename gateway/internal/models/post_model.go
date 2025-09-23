package models

import "time"

type PostRequest struct {
	Content    string   `json:"content" validate:"max=240"`
	PostImages []string `json:"post_images" validate:"omitempty,max=5,dive,url"`
}

type GetPostRequest struct {
	ID int `json:"id" validate:"required,gt=0"`
}

type Post struct {
	ID            int       `json:"id"`
	Content       string    `json:"content"`
	UserID        int       `json:"user_id"`
	PostImages    []string  `json:"post_images"`
	LikesCount    int       `json:"likes_count"`
	CommentsCount int       `json:"comments_count"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Author        *User     `json:"author"`
	IsLiked       bool      `json:"is_liked"`
}
