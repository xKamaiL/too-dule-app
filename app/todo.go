package app

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/moonrhythm/hime"
	"github.com/xkamail/too-dule-app/pkg/utils"
	"github.com/xkamail/too-dule-app/todo"
	"log"
	"net/http"
)

var validate *validator.Validate

type todoWrap struct {
	t *todo.Service
}

func newTodoWrap() todoWrap {
	validate = validator.New()
	return todoWrap{
		t: todo.NewTodo(),
	}
}

func (w todoWrap) getTodo(ctx *hime.Context) error {
	resp, err := w.t.List(ctx, todo.ListParam{})
	if err != nil {
		return ctx.Error(err.Error())
	}

	return ctx.JSON(resp)
}
func (w todoWrap) createNewTodo(ctx *hime.Context) error {
	var p todo.CreateParam
	err := ctx.BindJSON(&p)
	if err != nil {
		return utils.JSONError(ctx.ResponseWriter(), fmt.Sprintf("json error %s", err.Error()), http.StatusInternalServerError)
	}

	err = validate.Struct(p)
	if err != nil {
		log.Println(err)
		return utils.ValidatorError(ctx, err)
	}

	_, err = w.t.Create(ctx, p)
	if err != nil {
		log.Println(err)
		return utils.JSONError(ctx.ResponseWriter(), fmt.Sprintf("create todo failed: %s", err.Error()), http.StatusBadRequest)
	}
	ctx.Status(http.StatusCreated)
	return nil
}
