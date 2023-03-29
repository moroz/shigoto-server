package services

import (
	"context"

	valid "github.com/asaskevich/govalidator"
	"github.com/moroz/shigoto-server/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	valid.SetFieldsRequiredByDefault(true)
}

type CreateUserParams struct {
	Email    string `json:"email" valid:"email"`
	Password string `json:"password" valid:"stringlength(10|128)"`
}

func encryptPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func CreateUser(db boil.ContextExecutor, params CreateUserParams) (*models.User, error) {
	isValid, err := valid.ValidateStruct(params)
	if !isValid {
		return nil, err
	}
	user := models.User{Email: params.Email}
	user.PasswordHash = null.StringFrom(encryptPassword(params.Password))
	err = user.Insert(context.Background(), db, boil.Infer())
	if err != nil {
		return nil, err
	}
	return &user, nil
}
