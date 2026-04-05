package models

import "time"

type PostImage struct {
	ImageURL string `json:"image_url" validate:"required,url"`
	Order    int    `json:"order" validate:"required,gt=0"`
	Width    int    `json:"width" validate:"required,gt=0"`
	Height   int    `json:"height" validate:"required,gt=0"`
}

type CreatePostRequest struct {
	Content    string      `json:"content" validate:"max=240"`
	PostImages []PostImage `json:"post_images" validate:"omitempty,max=5,dive"`
}

type GetPostRequest struct {
	ID int `json:"id" validate:"required,gt=0"`
}

type Post struct {
	ID            int         `json:"id"`
	Content       string      `json:"content"`
	PostImages    []PostImage `json:"post_images"`
	LikesCount    int         `json:"likes_count"`
	CommentsCount int         `json:"comments_count"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
	Author        *User       `json:"author"`
	IsLiked       bool        `json:"is_liked"`
}
