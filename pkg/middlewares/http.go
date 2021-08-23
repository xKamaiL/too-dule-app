package middlewares

import (
	"github.com/moonrhythm/hime"
	"github.com/xkamail/too-dule-app/pkg/utils"
	"log"
	"net/http"
	"runtime/debug"
)

func NotFoundHandler(ctx *hime.Context) error {
	res := make(map[string]string)
	res["message"] = "Not found"
	return ctx.JSON(res)
}

func NoCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func PanicRecovery(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				utils.JSONError(w, "Internal Server Error", http.StatusInternalServerError)
				log.Println(err)
				debug.PrintStack()
			}
		}()
		h.ServeHTTP(w, r)
	})
}
