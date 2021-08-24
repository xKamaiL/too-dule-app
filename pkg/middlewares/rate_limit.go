package middlewares

import (
	"fmt"
	"go.uber.org/ratelimit"
	"log"
	"net/http"
	"time"
)

func RateLimit(perSec int) func(h http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := r.Header.Get("X-Forwarded-For")
			fmt.Println("Rate-Limit-For-IP: ", ip)
			rl := ratelimit.New(perSec) // per second
			prev := time.Now()

			now := rl.Take()
			now.Sub(prev)
			prev = now

			log.Println(r.RequestURI)
			next.ServeHTTP(w, r)
		})
	}
}
