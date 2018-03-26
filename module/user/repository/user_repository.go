package repository

import (
	base "github.com/dafian47/go-freya-rest-api/module"
	model "github.com/dafian47/go-freya-rest-api/module/user"
	"github.com/dafian47/go-freya-rest-api/util"
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	UserLogin(user model.User) (model.User, error)
	UserRegister(user model.User) (model.User, error)
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{DB: db}
}

func (r *userRepository) UserLogin(user model.User) (model.User, error) {

	var auth model.User

	username := user.Username
	plainPassword := user.Password

	r.DB.Where(&model.User{Username: username}).First(&auth)

	if auth.UserID == "" {
		return auth, base.NOT_FOUND_ERROR
	}

	isMatch := util.MatchString(auth.Password, plainPassword)
	if !isMatch {
		return auth, base.CONFLIT_ERROR
	}

	auth.Password = ""

	return auth, nil
}

func (r *userRepository) UserRegister(user model.User) (model.User, error) {

	var auth model.User

	userID, err := util.GenerateUserID()
	if err != nil {
		return user, base.INTERNAL_SERVER_ERROR
	}

	hashPassword, err := util.HashString(user.Password)
	if err != nil {
		return user, base.INTERNAL_SERVER_ERROR
	}

	auth.UserID = userID
	auth.Username = user.Username
	auth.Password = hashPassword

	r.DB.Save(&auth)

	if auth.UserID == "" {
		return auth, base.FAILED_SAVE_ERROR
	}

	auth.Password = ""

	return auth, nil
}
