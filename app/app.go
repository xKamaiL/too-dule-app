package app

import (
	"github.com/acoshift/header"
	"github.com/acoshift/middleware"

	"github.com/gorilla/mux"
	"github.com/moonrhythm/hime"
	"log"
	"net/http"
	"runtime/debug"
)

/*
	New App...
*/
func New(app *hime.App) http.Handler {

	app.
		Routes(hime.Routes{
			"index":  "/",
			"create": "/create",
			"done":   "/done",
			"remove": "/remove",
		})
	// router
	r := mux.NewRouter()
	r.NotFoundHandler = hime.Handler(notFoundHandler)
	r.HandleFunc("/", homeHandler)
	return middleware.Chain(
		panicRecovery,
		noCORS,
	)(r)
}

func defaultCacheControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(header.CacheControl, "no-cache, no-store, must-revalidate")
		h.ServeHTTP(w, r)
	})
}

func notFoundHandler(ctx *hime.Context) error {
	res := make(map[string]string)
	res["message"] = "Not found"
	return ctx.JSON(res)
}

func noCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			http.Error(w, "Forbidded", http.StatusForbidden)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func panicRecovery(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				log.Println(err)
				debug.PrintStack()
			}
		}()
		h.ServeHTTP(w, r)
	})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello world."))
}
