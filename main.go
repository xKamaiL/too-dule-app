package main

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
	"github.com/moonrhythm/hime"
	"github.com/xkamail/too-dule-app/app"
	"github.com/xkamail/too-dule-app/pkg/config"
	"log"
	"runtime"
)

func main() {
	config.Init()
	cfg := config.Load()

	// loc, err := time.LoadLocation("Asia/Bangkok")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for tracing sql query
	// pgDriver, _ := ocsql.Register("postgres", ocsql.WithQuery(true))
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.Username, cfg.DB.Password, cfg.DB.Name))
	if err != nil {
		log.Fatalf("can not open db; %v", err)
	}

	maxConns := runtime.NumCPU() * 4
	db.SetMaxIdleConns(maxConns)
	db.SetMaxOpenConns(maxConns)


	// start redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Username: cfg.Redis.Username,
		Password: cfg.Redis.Password,
	})

	himeApp := hime.New()

	himeApp.GracefulShutdown()

	fmt.Println("Serve: 8080")

	err = himeApp.
		Handler(app.New(himeApp, db, redisClient)).
		Address(":8080").
		ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
