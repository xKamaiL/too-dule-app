package app

import (
	"github.com/acoshift/header"
	mw "github.com/acoshift/middleware"
	"github.com/xkamail/too-dule-app/middlewares"

	"github.com/gorilla/mux"
	"github.com/moonrhythm/hime"
	"net/http"
)

/*
	New App...
*/
func New(app *hime.App) http.Handler {

	// router
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.NotFoundHandler = hime.Handler(middlewares.NotFoundHandler)

	// set prefix to /api
	apiRoute := r.PathPrefix("/api").Subrouter()
	apiRoute.Use(middlewares.Logging)
	
	{
		memberRouter := apiRoute.PathPrefix("/member").Subrouter()
		memberRouter.Get("/list")
	}

	{
		todoRouter := apiRoute.PathPrefix("/todo").Subrouter()
		todoRouter.Get("/list")
	}

	return mw.Chain(
		middlewares.PanicRecovery,
		middlewares.NoCORS,
	)(r)
}

func defaultCacheControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(header.CacheControl, "no-cache, no-store, must-revalidate")
		h.ServeHTTP(w, r)
	})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello world."))
}
