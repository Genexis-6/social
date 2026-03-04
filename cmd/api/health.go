package main

import (
	"log"
	"net/http"
)


func (app *application) health(w http.ResponseWriter, r *http.Request){
	data := map[string]any{
		"api_version": "v1",
		"welcome_message":"welcome to my first go backend porject",
	}
	if err:= WriteJSON(w, http.StatusAccepted, data); err != nil{
		log.Fatalln(err)
	}
}