package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSONError(w http.ResponseWriter, message string, statusCode int) {
	resp := map[string]interface{}{}
	resp["message"] = message
	w.WriteHeader(statusCode)
	if len(w.Header().Get("Content-Type")) == 0 {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	}
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Println(err)
	}
}
