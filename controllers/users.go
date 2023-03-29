package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/moroz/shigoto-server/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Users struct {
	DB boil.ContextExecutor
}

func (u *Users) List(w http.ResponseWriter, r *http.Request) {
	users, err := models.Users().All(context.Background(), u.DB)
	if err != nil {
		log.Fatal(err)
	}

	json, _ := json.Marshal(users)
	fmt.Fprint(w, string(json))
}
