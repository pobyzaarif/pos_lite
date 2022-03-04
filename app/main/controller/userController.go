package controller

import (
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/pobyzaarif/pos_lite/app/main/common"

	goutilLogger "github.com/pobyzaarif/goutil/logger"
)

var logger = goutilLogger.NewLog("main") // sekarang jam berapa

type (
	userLoginRequest struct {
		Email    string `json:"email" validate:"required,email,lowercase"`
		Password string `json:"password" validate:"required,min=8,max=32"`
	}
)

func (ctrl *Controller) UserLogin(c echo.Context) error {
	c.Get("")
	request := new(userLoginRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(common.EmptyObject))
	}

	err := validator.New().Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(common.EmptyObject))
	}

	user, err := ctrl.userService.FindByEmail(request.Email)
	if err != nil {
		// return c.JSON(http.StatusBadRequest, common.NewNotFoundResponse(common.EmptyObject, []string{"email"}))
		return c.JSON(http.StatusUnauthorized, common.NewUnauthorizedResponse(common.EmptyObject))
	}

	if !ctrl.authService.VerifyLogin(request.Email, ctrl.appConfig.AppUserPasswordSalt, request.Password) {
		return c.JSON(http.StatusUnauthorized, common.NewUnauthorizedResponse(common.EmptyObject))
	}

	token, err := ctrl.authService.GenerateToken(ctrl.appConfig.AppJWTSign, user.ID, user.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse(common.EmptyObject))
	}

	responseToken := map[string]string{"token": token}
	response := common.NewResponse(http.StatusText(http.StatusOK), responseToken)

	time.Sleep(time.Second)
	return c.JSON(http.StatusOK, response)
}

func (ctrl *Controller) Panic(c echo.Context) error {
	panic("panic")
}
