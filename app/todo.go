package app

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/moonrhythm/hime"
	"github.com/xkamail/too-dule-app/member"
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
		return utils.ValidatorError(ctx, err)
	}
	user, _ := member.GetMemberFromContext(ctx)
	// get from user
	p.OwnerID = user.ID
	_, err = w.t.Create(ctx, p)
	if err != nil {
		log.Println(err)
		return utils.JSONError(ctx.ResponseWriter(), fmt.Sprintf("create todo failed: %s", err.Error()), http.StatusBadRequest)
	}
	ctx.ResponseWriter().WriteHeader(http.StatusCreated)
	return nil
}

func (w todoWrap) UpdateTodo(ctx *hime.Context) error {
	var p todo.UpdateParam
	err := ctx.BindJSON(&p)
	if err != nil {
		return utils.JSONError(ctx.ResponseWriter(), fmt.Sprintf("json error %s", err.Error()), http.StatusInternalServerError)
	}

	err = validate.Struct(p)
	if err != nil {
		return utils.ValidatorError(ctx, err)
	}
	todoFromCtx := todo.GetFromContext(ctx)

	_, err = w.t.Update(ctx, todoFromCtx.ID, p)
	if err != nil {
		return err
	}

	return nil
}

// any member can change it
func (w todoWrap) ChangeStatusTodo(ctx *hime.Context) error {

	// quick validate
	var p struct {
		Status bool `json:"status" validate:"required"`
	}
	err := ctx.BindJSON(&p)
	if err != nil {
		return utils.JSONError(ctx.ResponseWriter(), fmt.Sprintf("json error %s", err.Error()), http.StatusInternalServerError)
	}
	err = validate.Struct(p)
	if err != nil {
		return utils.ValidatorError(ctx, err)
	}

	return w.t.ChangeStatus(ctx, mux.Vars(ctx.Request)["todo_id"], p.Status)
}

func (w todoWrap) MakeAssign(ctx *hime.Context) error {
	var assignMemberID string
	params := mux.Vars(ctx.Request)
	if params["member_id"] == "" {
		u, _ := member.GetMemberFromContext(ctx)
		assignMemberID = u.ID
	} else {
		assignMemberID = params["member_id"]
	}

	return w.t.AssignToMember(ctx, params["todo_id"], assignMemberID)
}

func (w todoWrap) RemoveAssign(ctx *hime.Context) error {
	params := mux.Vars(ctx.Request)
	return w.t.RemoveAssign(ctx, params["todo_id"])
}

func (w todoWrap) RemoveTodo(ctx *hime.Context) error {
	// params := mux.Vars(ctx.Request)
	// or get from context
	item := todo.GetFromContext(ctx)
	err := w.t.Delete(ctx, item.ID)
	if err != nil {
		return err
	}

	// write no content
	ctx.ResponseWriter().WriteHeader(http.StatusNoContent)

	return nil

}
