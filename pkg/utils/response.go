package utils

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/moonrhythm/hime"
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

type ErrorResponse struct {
	Message string `json:"message"`
}

type validationErrorResponse struct {
	Errors []string `json:"errors"`
}

func ValidatorError(ctx *hime.Context, err error) error {
	fields := make([]string, 0)
	for _, err := range err.(validator.ValidationErrors) {
		fields = append(fields, fmt.Sprintf("%s %s", err.Field(), err.Value()))
	}
	return ctx.JSON(validationErrorResponse{Errors: fields})
}
