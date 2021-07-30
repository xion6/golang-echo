package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/xion6/golang-echo/pkg/models"
)

func (a App) Register(c echo.Context) error {
	newUser, err := models.ValidateRegisterRequest(c)
	if err != nil {
		return c.JSON(err.Code, err)
	}

	err = a.userSvc.CreateAccount(newUser)

	if err != nil {
		return c.JSON(err.Code, err)
	}

	return c.String(http.StatusCreated, "")
}

func (a App) Login(c echo.Context) error {
	loginRequest, err := models.ValidateLoginRequest(c)
	if err != nil {
		return c.JSON(err.Code, err)
	}

	token, err := a.userSvc.Login(loginRequest)

	if err != nil {
		return c.JSON(err.Code, err)
	}

	response := &models.LoginResponse{Token: token}

	return c.JSON(http.StatusOK, response)
}
