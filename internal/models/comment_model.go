package models

import (
	"time"

)



type CommentModel struct{
	ID int64 `json:"id"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UserID int64 `json:"user_id"`
	PostID int64 `json:"post_id"`
	User UserModel  `json:"user"`
}