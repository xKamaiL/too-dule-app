package middlewares

import (
	"github.com/xkamail/too-dule-app/pkg/utils"
	"net/http"
	"strings"
)

func NeededJSONBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut || r.Method == http.MethodPost {
			if !strings.Contains(r.Header.Get("Content-Type"), "application/json") {
				msg := "Content-Type header is not application/json"
				_ = utils.JSONError(w, msg, http.StatusUnsupportedMediaType)
				return
			}
			// Max size of json to 1MB
			r.Body = http.MaxBytesReader(w, r.Body, 1048576)

		}

		next.ServeHTTP(w, r)
	})
}
