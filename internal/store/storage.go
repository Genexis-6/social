package store

import (
	"context"
	"database/sql"
	"errors"


	"github.com/Genexis-6/social/internal/models"
)


var (
	NoResourceFound = errors.New("No resouces found")
	NoUpdateMode = errors.New("No Update made")
	NoRecordDeleted = errors.New("No Record deleted")
)

type Storage struct {
	Users interface {
		Create(context.Context, *models.UserModel) error
		
	}
	Posts interface {
		Create(context.Context, *models.PostModel) error
		// GetMultiplePosts(context.Context, int64) ([]models.PostSummary, error)
		GetPostById(context.Context, int64)(*models.PostModel, error)
		DeletePostById(context.Context, int64)error
		UpdatePostById(context.Context, *models.PostModel) error
	}
	Comment interface {
		GetCommentsByPostId(context.Context, int64)([]*models.CommentModel, error)
		Create(context.Context, *models.CommentModel)error
	}
}


func NewStorage(db *sql.DB) *Storage{
	return &Storage{
		Users: &UserStore{db: db},
		Posts: &PostStore{db: db},
		Comment: &CommentStore{db: db},
	}
}