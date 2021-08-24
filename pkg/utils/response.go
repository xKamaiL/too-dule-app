package utils

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/moonrhythm/hime"
	"log"
	"net/http"
)

func JSONError(w http.ResponseWriter, message string, statusCode int) error {
	resp := map[string]interface{}{}
	resp["message"] = message
	w.WriteHeader(statusCode)
	if len(w.Header().Get("Content-Type")) == 0 {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	}
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type validationErrFields struct {
	Name    string `json:"name"`
	Tag     string `json:"tag"`
	Param   string `json:"param"`
	Message string `json:"message"`
}

type validationErrorResponse struct {
	Errors []validationErrFields `json:"errors"`
}

func ValidatorError(ctx *hime.Context, err error) error {
	fields := make([]validationErrFields, 0)
	for _, err := range err.(validator.ValidationErrors) {
		fields = append(fields, validationErrFields{
			Name:    err.Field(),
			Tag:     err.Tag(),
			Param:   err.Param(),
			Message: err.Error(),
		})
	}
	ctx.Status(http.StatusUnprocessableEntity)
	return ctx.JSON(validationErrorResponse{Errors: fields})
}
