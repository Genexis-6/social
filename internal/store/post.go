package store

import (
	"context"
	"database/sql"
	"github.com/Genexis-6/social/internal/models"
	"github.com/lib/pq"
)

type PostStore struct {
	db *sql.DB
}

func (p *PostStore) Create(ctx context.Context, post *models.PostModel) error {
	q := `
	INSERT INTO post_model (content, tags, user_id, title) 
	VALUES ($1, $2, $3, $4) RETURNING id, created_ar, updated_at
	`
	res := p.db.QueryRowContext(ctx, q, post.Content, pq.Array(post.Tags), 
								post.UserId, post.Title).Scan(&post.ID, &post.Created_At, &post.Updated_At)

	return res
}
