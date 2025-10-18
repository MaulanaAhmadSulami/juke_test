package protocol

import (
	"net/http"
	"encoding/json"
)


func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func WriteJSONError(w http.ResponseWriter, status int, message string) error {
	type envelope struct {
		Error string `json:"error"`
	}
	return WriteJSON(w, status, &envelope{Error: message})
}

func InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	WriteJSONError(w, http.StatusInternalServerError, "Something went wrong")
}

func NotFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	WriteJSONError(w, http.StatusNotFound, "not found")
}

func BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	WriteJSONError(w, http.StatusBadRequest, err.Error())
}
