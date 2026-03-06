package store

import (
	"context"
	"database/sql"

	"github.com/Genexis-6/social/internal/models"
)


type CommentStore struct{
	db *sql.DB
}



func (c *CommentStore) GetCommentsByPostId(ctx context.Context, postId int64) ([]*models.CommentModel, error){
	q := `
	 SELECT comment_model.id, comment_model.content, comment_model.created_at, user_model.username, user_model.email FROM comment_model 
	 LEFT JOIN user_model ON user_model.id = comment_model.user_id 
	 WHERE comment_model.post_id = $1
	 ORDER BY comment_model.created_at DESC 
	
	`
	var comments []*models.CommentModel

	row, err := c.db.QueryContext(ctx, q, postId)
	

	if err != nil{
		return nil, err
	}

	defer row.Close()

	for row.Next(){
		var comment models.CommentModel;

		if err := row.Scan(&comment.ID, &comment.Content, &comment.CreatedAt, &comment.User.UserName, &comment.User.Email); err != nil{
			return nil , err;
		}

		comments = append(comments, &comment)
	}

	

	if row.Err() != nil{
		return  nil, err
	}

	return comments, nil

}


func (c *CommentStore) Create(ctx context.Context, comment *models.CommentModel)error{
	q:= `
	INSERT INTO comment_model (content, user_id, post_id) VALUES ($1, $2, $3) RETURNING created_at, id
	`

	return c.db.QueryRowContext(ctx, q, comment.Content, comment.UserID, comment.PostID).Scan(&comment.CreatedAt, &comment.ID)
}