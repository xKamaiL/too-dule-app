package app

import (
	"database/sql"
	"github.com/acoshift/header"
	mw "github.com/acoshift/middleware"
	"github.com/acoshift/pgsql/pgctx"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/moonrhythm/hime"
	"github.com/xkamail/too-dule-app/pkg/config"
	"github.com/xkamail/too-dule-app/pkg/middlewares"
	pkgRedis "github.com/xkamail/too-dule-app/pkg/redis"
	"golang.org/x/time/rate"
	"net/http"
)

/*
	New App...
*/
func New(app *hime.App, db *sql.DB, redisClient *redis.Client) http.Handler {
	app.ETag = true
	cfg := config.Load()

	// set rate limit
	var limiter = middlewares.NewIPRateLimiter(rate.Limit(cfg.RateLimitAllow), 1)

	// router
	r := mux.NewRouter()

	r.HandleFunc("/", homeHandler)
	r.NotFoundHandler = hime.Handler(middlewares.NotFoundHandler)

	// set prefix to /api
	apiRoute := r.PathPrefix("/api").Subrouter()
	apiRoute.Use(middlewares.Logging)

	{
		t := newMemberWrap()
		memberRouter := apiRoute.PathPrefix("/member").Subrouter()
		memberRouter.Use(middlewares.NeededJSONBody)

		memberRouter.Handle("/sign-in", hime.Handler(t.postMemberSignIn)).Methods(http.MethodPost)

		memberRouter.Handle("/register", middlewares.RateLimit(hime.Handler(t.postMemberRegister))).Methods(http.MethodPost)

		authMemberRouter := memberRouter.PathPrefix("/").Subrouter()
		authMemberRouter.Use(middlewares.MemberAuthorization)
		authMemberRouter.Handle("/me", hime.Handler(t.getMe)).Methods(http.MethodGet)

		authMemberRouter.Handle("/list", hime.Handler(t.getMemberList)).Methods(http.MethodGet)

	}

	{
		t := newTodoWrap()
		todoRouter := apiRoute.PathPrefix("/todo").Subrouter()
		todoRouter.Use(middlewares.NeededJSONBody, middlewares.MemberAuthorization)
		todoRouter.Handle("/list", hime.Handler(t.getTodo)).Methods(http.MethodGet)

		todoRouter.Handle("", middlewares.RateLimit(hime.Handler(t.createNewTodo))).Methods(http.MethodPost)
		// update to-do
		todoRouter.Handle("/{todo_id}", t.t.IsOwnerOnly(hime.Handler(t.UpdateTodo))).Methods(http.MethodPut)

		// /{todo_id}/assign : Assign
		todoRouter.Handle("/{todo_id}", t.t.IsOwnerOnly(hime.Handler(t.MakeAssign))).Methods(http.MethodPut)

		// /{todo_id}/re-assign : Remove Assign

		// /{todo_id}/status : Change Status
		todoRouter.Handle("/{todo_id}/status", t.t.IsOwnerOnly(hime.Handler(t.ChangeStatusTodo))).Methods(http.MethodPut)

	}

	return mw.Chain(
		middlewares.PanicRecovery,
		middlewares.NoCORS,
		middlewares.WithRateLimitContext(limiter),
		pgctx.Middleware(db),
		pkgRedis.Middleware(redisClient),
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
