package models

import "time"

type PostModel struct {
	ID         int64           `json:"id"`
	Title      string          `json:"title"`
	Content    string          `json:"content"`
	Authur     string          `json:"authur"`
	Tags       []string        `json:"tags"`
	UserId     int64           `json:"user_id"`
	Created_At time.Time       `json:"created_at"`
	Version    int64           `json:"version"`
	Updated_At time.Time       `json:"updated_at"`
	Comments   []*CommentModel `json:"comments"`
}
