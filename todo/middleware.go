package todo

import (
	"github.com/gorilla/mux"
	"github.com/xkamail/too-dule-app/member"
	"github.com/xkamail/too-dule-app/pkg/utils"
	"net/http"
)

// it should move to middleware pkg ?
func (t Service) IsOwnerOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		todo, err := t.repo.FindByID(r.Context(), vars["todo_id"])
		// not found
		if err != nil {
			_ = utils.JSONError(w, "todo is not found", http.StatusNotFound)
			return
		}
		user, err := member.GetMemberFromContext(r.Context())
		if err != nil {
			_ = utils.JSONError(w, "failed to get member from context", http.StatusNotFound)
			return
		}
		// check owner_id != user.id
		if todo.OwnerID != user.ID {
			_ = utils.JSONError(w, "you cannot update this todo", http.StatusUnauthorized)
			return
		}
		// set current to-do into a context
		ctx := NewFromContext(r.Context(), *todo)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// set to-do entity from middleware
func (t Service) SetWithParams(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		todo, err := t.repo.FindByID(r.Context(), vars["todo_id"])
		// not found
		if err != nil {
			_ = utils.JSONError(w, "todo is not found", http.StatusNotFound)
			return
		}
		// set current to-do into a context
		ctx := NewFromContext(r.Context(), *todo)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
