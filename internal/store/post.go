package store

import (
	"context"
	"database/sql"
	"errors"
	"log"

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

func (p *PostStore) GetPostById(ctx context.Context, id int64) (*models.PostModel, error) {
	q := `
	SELECT id, title, content, tags, created_at, updated_at, user_id, version FROM post_model WHERE id=$1
	`
	var post models.PostModel

	err := p.db.QueryRowContext(ctx, q, id).Scan(&post.ID, &post.Title, &post.Content, pq.Array(&post.Tags), &post.Created_At, &post.Updated_At, &post.UserId, &post.Version)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, NoResourceFound

	default:
		return &post, err
	}

}

func (p *PostStore) DeletePostById(ctx context.Context, id int64) error {
	q := `
	DELETE FROM post_model WHERE id=$1
	`
	row, err := p.db.ExecContext(ctx, q, id)
	if err != nil {
		return err
	}

	affected, err := row.RowsAffected()

	if affected == 0 {
		return NoRecordDeleted
	}

	return err
}

func (p *PostStore) UpdatePostById(ctx context.Context, post *models.PostModel) error {
	log.Println(post.UserId)
	q := `
	UPDATE post_model SET content=COALESCE($1, post_model.content), created_at=post_model.created_at, updated_at= CURRENT_TIMESTAMP, title=COALESCE($2, post_model.title), version=post_model.version+1  WHERE id=$3 AND user_id=$4 AND version=$5
	`
	row, err := p.db.ExecContext(ctx, q, post.Content, post.Title, post.ID, post.UserId, post.Version)

	if err != nil {
		return err
	}
	affected, err := row.RowsAffected()
	if err != nil {
		return err
	}
	// log.Println(affected)
	if affected == 0 {
		return NoUpdateMode
	}

	return nil
}
