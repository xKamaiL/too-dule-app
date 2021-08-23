package app

import (
	"github.com/go-playground/validator/v10"
	"github.com/moonrhythm/hime"
	"github.com/xkamail/too-dule-app/pkg/utils"
	"github.com/xkamail/too-dule-app/todo"
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
		return err
	}

	err = validate.Struct(p)
	if err != nil {
		return utils.ValidatorError(ctx, err)
	}

	_, err = w.t.Create(ctx, p)
	if err != nil {
		// should be error message
		return err
	}
	ctx.Status(http.StatusCreated)
	return nil
}
