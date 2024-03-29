package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/moroz/shigoto-server/controllers"
)

func MustGetenv(name string) string {
	value := os.Getenv(name)
	if value != "" {
		return value
	}
	msg := fmt.Sprintf("Environment variable %s is not set!", name)
	panic(msg)
}

func initDB(connString string) *sqlx.DB {
	return sqlx.MustConnect("pgx", connString)
}

func main() {
	db := initDB(MustGetenv("DATABASE_URL"))

	uc := controllers.Users{DB: db}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)

	r.Get("/users", uc.List)
	r.Post("/users", uc.Create)

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", r)
}
