package middlewares

import (
	"github.com/xkamail/too-dule-app/pkg/utils"
	"net/http"
)

func MemberAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Member-Token")
		if len(token) == 0 {
			utils.JSONError(w, "No token", http.StatusUnauthorized)
			return
		}
		// TODO: implement JWT
		if token != "hahaha" {
			utils.JSONError(w, "Authentication failed", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
