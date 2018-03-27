package handler

import (
	"github.com/labstack/echo"
	"net/http"

	base "github.com/dafian47/go-freya-rest-api/module"
	model "github.com/dafian47/go-freya-rest-api/module/user"
	"github.com/dafian47/go-freya-rest-api/module/user/repository"
)

type UserHandler interface {
	Login(c echo.Context) error
	Register(c echo.Context) error
}

type userHandler struct {
	repo repository.UserRepository
}

func NewUserHandler(e *echo.Echo, userRepo repository.UserRepository) {

	handler := &userHandler{repo: userRepo}

	e.POST("/login", handler.Login)
	e.POST("/register", handler.Register)
}

func (r *userHandler) Login(c echo.Context) error {

	var auth model.User

	err := c.Bind(&auth)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &base.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	result, err := r.repo.UserLogin(auth)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &base.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, &base.Response{
		Status:  http.StatusOK,
		Message: base.SUCCESS_LOGIN_DATA,
		Data:    result,
	})
}

func (r *userHandler) Register(c echo.Context) error {

	var user model.User

	err := c.Bind(&user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &base.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	result, err := r.repo.UserRegister(user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &base.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, &base.Response{
		Status:  http.StatusOK,
		Message: base.SUCCESS_REGISTER_DATA,
		Data:    result,
	})
}
