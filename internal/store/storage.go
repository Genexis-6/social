package store

import (
	"context"
	"database/sql"

	"github.com/Genexis-6/social/internal/models"
)

type Storage struct {
	users interface {
		Create(context.Context, *models.UserModel) error
	}
	posts interface {
		Create(context.Context, *models.PostModel) error
	}
}


func NewStorage(db *sql.DB) *Storage{
	return &Storage{
		users: &UserStore{db: db},
		posts: &PostStore{db: db},
	}
}