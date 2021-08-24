package middlewares

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/xkamail/too-dule-app/member"
	"github.com/xkamail/too-dule-app/pkg/config"
	"github.com/xkamail/too-dule-app/pkg/utils"
	"net/http"
)

func MemberAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg := config.Load()
		token := r.Header.Get("X-Member-Token")
		if len(token) == 0 {
			_ = utils.JSONError(w, "No token", http.StatusUnauthorized)
			return
		}

		jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(cfg.JWTSecretKey), nil
		})
		if err != nil {
			_ = utils.JSONError(w, "Token is invalid.", http.StatusUnauthorized)

			return
		}

		if _, ok := jwtToken.Claims.(jwt.Claims); !ok && !jwtToken.Valid {
			_ = utils.JSONError(w, "Failed claims", http.StatusUnauthorized)

			return
		}
		meta, ok := jwtToken.Claims.(jwt.MapClaims)
		if !ok {
			_ = utils.JSONError(w, "Failed to extract data from jwt", http.StatusInternalServerError)
			return
		}
		memberID, ok := meta["user_id"].(string)
		if !ok {
			_ = utils.JSONError(w, "Failed to extract user_id from jwt", http.StatusInternalServerError)
			return
		}

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
