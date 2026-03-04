package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Genexis-6/social/internal/models"
	"github.com/lib/pq"
)

type PostStore struct {
	db *sql.DB
}

func (p *PostStore) Create(ctx context.Context, post *models.PostModel) error {
	q := `
	INSERT INTO post_model (content, tags, user_id, title) 
	VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at
	`
	res := p.db.QueryRowContext(ctx, q, post.Content, pq.Array(post.Tags),
		post.UserId, post.Title).Scan(&post.ID, &post.Created_At, &post.Updated_At)

	return res
}

func (p *PostStore) GetMultiplePosts(ctx context.Context, userId int64) ([]models.PostSummary, error) {
	q := `SELECT title, tags, content, created_at FROM post_model WHERE user_id=$1`
	var posts []models.PostSummary

	res, err := p.db.QueryContext(ctx, q, userId)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	for res.Next() {
		var post models.PostSummary
		if err := res.Scan(&post.Title, pq.Array(&post.Tags), &post.Content, &post.Created_At); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err = res.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *PostStore) GetPostById(ctx context.Context, id int64) (*models.PostModel, error) {
	q := `
	SELECT id, title, content, tags, created_at, updated_at FROM post_model WHERE id=$1
	`
	var post models.PostModel

	err := p.db.QueryRowContext(ctx, q, id).Scan(&post.ID, &post.Title, &post.Content, pq.Array(&post.Tags), &post.Created_At, &post.Updated_At)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, NoResourceFound

	default:
		return &post, err
	}

}
