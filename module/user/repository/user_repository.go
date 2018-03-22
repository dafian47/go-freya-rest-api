package repository

import (
	model "github.com/dafian47/go-freya-rest-api/module/user"
)

type UserRepository interface {
	UserLogin() (*model.User, error)
	UserRegister() (*model.User, error)
}
