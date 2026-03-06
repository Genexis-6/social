package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"strconv"

	"github.com/Genexis-6/social/internal/models"
	"github.com/Genexis-6/social/internal/store"
	"github.com/go-chi/chi/v5"
)

type CreatePostPayload struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
	UserId  int64    `json:"user_id"`
}

type postKey string

const fetchedPost postKey = "fetchedPost"

func (app *application) createUserPost(w http.ResponseWriter, r *http.Request) {
	// userId := 1

	var userPOST CreatePostPayload

	if err := ReadJSON(w, r, &userPOST); err != nil {
		WriteJSONError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	post := &models.PostModel{
		Title:   userPOST.Title,
		Content: userPOST.Content,
		Tags:    userPOST.Tags,
		UserId:  userPOST.UserId,
	}
	if err := app.config.store.Posts.Create(r.Context(), post); err != nil {
		WriteJSONError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if err := WriteJSON(w, http.StatusCreated, post); err != nil {
		WriteJSONError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	post := getPostFromCtx(r)

	comments, err := app.config.store.Comment.GetCommentsByPostId(r.Context(), int64(post.ID))
	if err != nil {
		comments = []*models.CommentModel{}
		WriteJSONError(w, r, http.StatusBadRequest, err.Error())

	}
	post.Comments = comments
	WriteJSON(w, 200, post)
}

func (app *application) deletePostHandler(w http.ResponseWriter, r *http.Request) {
	query := chi.URLParam(r, "postID")

	postId, err := strconv.Atoi(query)
	if err != nil {
		WriteJSONError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	err = app.config.store.Posts.DeletePostById(r.Context(), int64(postId))
	switch {
	case errors.Is(err, nil):
		WriteJSON(w, 200, "post have been updated")
		return

	case errors.Is(err, store.NoRecordDeleted):
		WriteJSONError(w, r, http.StatusBadRequest, "unable to delete post")
		return

	default:
		WriteJSONError(w, r, 400, fmt.Sprintf("error occured while deleting post due to %v", err.Error()))

	}
}

type UpdatePost struct {
	Title   *string `json:"title"`
	Content *string `json:"content"`
}

func (app *application) updatePostHandler(w http.ResponseWriter, r *http.Request) {
	post := getPostFromCtx(r)

	var updatePost UpdatePost

	
	if err := ReadJSON(w, r, &updatePost); err != nil {
		WriteJSONError(w, r, http.StatusBadRequest, fmt.Sprintf("error decoding json due to %v", err.Error()))
		return
	}

	if updatePost.Content != nil{
		post.Content = *updatePost.Content
	}
	if updatePost.Title != nil{
		post.Title = *updatePost.Title
	}

	err := app.config.store.Posts.UpdatePostById(r.Context(), post)

	switch {
	case errors.Is(err, nil):
		WriteJSON(w, 200, "post have been updated")
		return

	case errors.Is(err, store.NoUpdateMode):
		WriteJSONError(w, r, http.StatusBadRequest, "unable to update post")
		return

	default:
		WriteJSONError(w, r, http.StatusInternalServerError, err.Error())

	}

}

// func (app *application) getMultiplePostHandler(w http.ResponseWriter, r *http.Request) {
// 	query := chi.URLParam(r, "postID")

// 	postId, err := strconv.Atoi(query)
// 	if err != nil {
// 		WriteJSONError(w, r, http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	posts, err := app.config.store.Posts.GetMultiplePosts(r.Context(), int64(postId))
// 	if err != nil {
// 		WriteJSONError(w, r, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	WriteJSON(w, http.StatusOK, posts)
// }

func (app *application) postContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := chi.URLParam(r, "postID")

		postId, err := strconv.Atoi(query)
		if err != nil {
			WriteJSONError(w, r, http.StatusBadRequest, err.Error())
			return
		}

		res, err := app.config.store.Posts.GetPostById(r.Context(), int64(postId))

		if err != nil {
			switch {
			case errors.Is(err, store.NoResourceFound):
				WriteJSONError(w, r, http.StatusNotFound, err.Error())
				return

			default:
				WriteJSONError(w, r, http.StatusBadRequest, err.Error())
				return
			}

		}
		ctx := r.Context()

		ctx = context.WithValue(ctx, fetchedPost, res)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}

func getPostFromCtx(r *http.Request) *models.PostModel {
	post := r.Context().Value(fetchedPost)
	return post.(*models.PostModel)
}
