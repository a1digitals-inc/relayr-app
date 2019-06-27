package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

//ToError help function to handle error
func ToError(w http.ResponseWriter, err error, statusCode int) {
	ToJSON(w, struct {
		Message string `json:"message"`
	}{
		Message: err.Error(),
	}, statusCode)
}

// ToJSON is a help function to encode data to json
func ToJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	CheckError(err)
}

//CheckError help function to validate error
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
