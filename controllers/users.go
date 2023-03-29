package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/moroz/shigoto-server/models"
	"github.com/moroz/shigoto-server/services"
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

func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var p services.CreateUserParams
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := services.CreateUser(u.DB, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	json, _ := json.Marshal(user)
	fmt.Fprint(w, string(json))
}
