package models

import "time"

type UserModel struct {
	ID       int64  `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`

	CreatedAt time.Time `json:"created_at"`
}
