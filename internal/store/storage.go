package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Genexis-6/social/internal/models"
)


var (
	NoResourceFound = errors.New("No resouces found")
)

type Storage struct {
	Users interface {
		Create(context.Context, *models.UserModel) error
		
	}
	Posts interface {
		Create(context.Context, *models.PostModel) error
		GetMultiplePosts(context.Context, int64) ([]models.PostSummary, error)
		GetPostById(context.Context, int64)(*models.PostModel, error)
	}
}


func NewStorage(db *sql.DB) *Storage{
	return &Storage{
		Users: &UserStore{db: db},
		Posts: &PostStore{db: db},
	}
}