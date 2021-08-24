package app

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/moonrhythm/hime"
	"github.com/xkamail/too-dule-app/member"
	"github.com/xkamail/too-dule-app/pkg/utils"
	"log"
	"net/http"
)

type memberWrap struct {
	m *member.Service
}

func newMemberWrap() *memberWrap {
	validate = validator.New()

	return &memberWrap{m: member.New()}
}

func (m *memberWrap) getMe(ctx *hime.Context) error {
	mem, err := member.GetMemberFromContext(ctx)
	if err != nil {
		return utils.JSONError(ctx.ResponseWriter(), "Failed to get user", http.StatusUnauthorized)
	}
	return ctx.JSON(mem)
}

func (m *memberWrap) postMemberRegister(ctx *hime.Context) error {
	var p member.CreateMemberParam

	err := ctx.BindJSON(&p)
	if err != nil {
		return utils.JSONError(ctx.ResponseWriter(), fmt.Sprintf("json error %s", err.Error()), http.StatusInternalServerError)
	}

	err = validate.Struct(p)
	if err != nil {
		log.Println(err)
		return utils.ValidatorError(ctx, err)
	}

	accessToken, err := m.m.Create(ctx, p)
	if err != nil {
		return utils.JSONError(ctx.ResponseWriter(), err.Error(), http.StatusBadRequest)
	}
	return ctx.JSON(map[string]interface{}{"accessToken": accessToken})
}
func (m *memberWrap) postMemberSignIn(ctx *hime.Context) error {
	var p member.LoginParam

	err := ctx.BindJSON(&p)
	if err != nil {
		return utils.JSONError(ctx.ResponseWriter(), fmt.Sprintf("json error %s", err.Error()), http.StatusInternalServerError)
	}

	err = validate.Struct(p)
	if err != nil {
		log.Println(err)
		return utils.ValidatorError(ctx, err)
	}

	accessToken, err := m.m.SignIn(ctx, p)
	if err != nil {
		return utils.JSONError(ctx.ResponseWriter(), err.Error(), http.StatusBadRequest)
	}

	return ctx.JSON(map[string]interface{}{"accessToken": accessToken})
}
