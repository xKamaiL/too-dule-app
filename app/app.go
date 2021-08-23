package app

import (
	"database/sql"
	"github.com/acoshift/header"
	mw "github.com/acoshift/middleware"
	"github.com/acoshift/pgsql/pgctx"
	"github.com/gorilla/mux"
	"github.com/moonrhythm/hime"
	"github.com/xkamail/too-dule-app/middlewares"
	"net/http"
)

/*
	New App...
*/
func New(app *hime.App, db *sql.DB) http.Handler {
	app.ETag = true
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
		t := newTodoWrap()
		todoRouter := apiRoute.PathPrefix("/todo").Subrouter()
		todoRouter.Handle("/list", hime.Handler(t.getTodo)).Methods("GET")
	}

	return mw.Chain(
		middlewares.PanicRecovery,
		middlewares.NoCORS,
		pgctx.Middleware(db),
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