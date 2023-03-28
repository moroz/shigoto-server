package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/moroz/shigoto-server/models"
)

func main() {
	conn, err := sqlx.Connect("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	err = conn.Ping()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	users, err := models.Users().All(ctx, conn)
	if err != nil {
		panic(err)
	}
	json, _ := json.MarshalIndent(users, "", "  ")
	fmt.Println(string(json))
}
