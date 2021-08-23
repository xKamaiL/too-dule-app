package app

import (
	"github.com/moonrhythm/hime"
	"github.com/xkamail/too-dule-app/pkg/todo"
)

type todoWrap struct {
	t *todo.Service
}

func newTodoWrap() todoWrap {
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
