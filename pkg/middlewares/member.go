package middlewares

import (
	"github.com/xkamail/too-dule-app/member"
	"github.com/xkamail/too-dule-app/pkg/utils"
	"net/http"
)

func MemberAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Member-Token")
		if len(token) == 0 {
			_ = utils.JSONError(w, "No token", http.StatusUnauthorized)
			return
		}
		// TODO: implement JWT
		memberID := token

		memberRepo := member.Repository{}
		memberUser, err := memberRepo.FindByID(r.Context(), memberID)
		if err != nil {
			_ = utils.JSONError(w, "Authentication failed", http.StatusForbidden)
			return
		}
		newContext := member.NewMemberContext(r.Context(), *memberUser)
		next.ServeHTTP(w, r.WithContext(newContext))
	})
}
