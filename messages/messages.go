package messages

import (
	"encoding/json"
	"net/http"
)

func Ok(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

func BadRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
}

//? send message
func SendMessage(w http.ResponseWriter, r *http.Request, message string) {
	json.NewEncoder(w).Encode(map[string]string{"message": message})
}

//? error message
func ErrorMessage(w http.ResponseWriter, r *http.Request, message string) {
	BadRequest(w)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
