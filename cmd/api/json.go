package main

import (
	"encoding/json"
	"net/http"
)



func WriteJSON(w http.ResponseWriter, status int, data any) error{
	w.Header().Set("content-type","application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func ReadJSON(w http.ResponseWriter, r *http.Request, data any) error{
	maxByteReader := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxByteReader))
	decoder := json.NewDecoder(r.Body)

	decoder.DisallowUnknownFields()
	return decoder.Decode(data)
}



func WriteJSONError(w http.ResponseWriter, r *http.Request, status int, message string) error{
	type errorStuct struct{
		Error string `json:"error"`
	}
	return WriteJSON(w, status, &errorStuct{message})
}