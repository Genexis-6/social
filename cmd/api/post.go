package main

import (
	"errors"
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
	WriteJSON(w, 200, res)
}

func (app *application) getMultiplePostHandler(w http.ResponseWriter, r *http.Request) {
	query := chi.URLParam(r, "postID")

	postId, err := strconv.Atoi(query)
	if err != nil {
		WriteJSONError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	posts, err := app.config.store.Posts.GetMultiplePosts(r.Context(), int64(postId))
	if err != nil {
		WriteJSONError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	WriteJSON(w, http.StatusOK, posts)
}
