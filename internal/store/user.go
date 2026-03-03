package store

import (
	"context"
	"database/sql"

	"github.com/Genexis-6/social/internal/models"
)


type UserStore struct{
	db *sql.DB
}


func (user *UserStore) Create(ctx context.Context, userModel *models.UserModel) error{
	q := `
	INSERT INTO user_model (username, email, password) VALUES ($1, $2, $3) RETURNING id, created_at
	`
	return user.db.QueryRowContext(ctx, q, userModel.UserName, userModel.Email, userModel.Password).Scan(&userModel.ID)
}