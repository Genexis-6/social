package main

import (
	"fmt"
	"net/http"

	"github.com/Genexis-6/social/internal/models"
)

func (app *application) createCommentHandler(w http.ResponseWriter, r *http.Request) {

	var comment models.CommentModel

	if err := ReadJSON(w, r, &comment); err != nil {
		WriteJSONError(w, r, http.StatusBadRequest, fmt.Sprintf("An error occured while decoding json due to: %v", err.Error()))
		return 
	}

	if res := app.config.store.Comment.Create(r.Context(), &comment); res != nil{
		WriteJSONError(w, r, http.StatusBadRequest, fmt.Sprintf("An error occured while creating post due to: %v", res.Error()))
		return 
	}

	WriteJSON(w, 201, "post was successfully created")


}
