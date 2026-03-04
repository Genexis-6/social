package main

import (
	"errors"
	"log"
	"net/http"

	"time"

	"github.com/Genexis-6/social/internal/db"
	"github.com/Genexis-6/social/internal/env"
	"github.com/Genexis-6/social/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

type application struct {
	config Config
}

type Config struct {
	addr  string
	store store.Storage
	db    dbConfig
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(errors.New("No env file founded"))
	}
	dbpool, err := db.DbPool(25,25, time.Duration(15*60),)
	if err != nil{
		panic(err)
	
	}
	
	defer dbpool.Close()
	log.Println("db connection created")
	return &Config{
		addr:  env.GetEnvString("ADDRESS", ":5000"),
		store: *store.NewStorage(dbpool),
		db: dbConfig{
			addr: env.GetEnvString("ADDRESS", ":5000"),
		},
		

	}
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdelConns int
	maxIdleTime  time.Duration
}

func (app *application) mount() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))
	router.Get("/health", app.health)

	router.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.health)
	})

	return router
}

func (app *application) runApp(handler http.Handler) error {

	server := &http.Server{
		Addr:         app.config.addr,
		Handler:      handler,
		IdleTimeout:  time.Minute,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
	}

	log.Printf("sever statrted at port %s\n", app.config.addr)
	return server.ListenAndServe()
}
