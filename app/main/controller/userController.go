package controller

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/pobyzaarif/pos_lite/app/main/common"
	"github.com/pobyzaarif/pos_lite/business"

	goutilLogger "github.com/pobyzaarif/goutil/logger"
)

var uclogger = goutilLogger.NewLog("USER_CONTROLLER")

type (
	userLoginRequest struct {
		Email    string `json:"email" validate:"required,email,lowercase"`
		Password string `json:"password" validate:"required,min=8,max=32"`
	}
)

func (ctrl *Controller) UserLogin(c echo.Context) error {
	trackerID, _ := c.Get("tracker_id").(string)
	uclogger.SetTrackerID(trackerID)
	ic := business.NewInternalContext(trackerID)

	request := new(userLoginRequest)
	if err := c.Bind(request); err != nil {
		uclogger.Error(common.ErrorBindingRequest.String(), err)
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(common.EmptyObject))
	}

	err := validator.New().Struct(request)
	if err != nil {
		uclogger.Error(common.ErrorValidationRequest.String(), err)
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(common.EmptyObject))
	}

	user, err := ctrl.userService.FindByEmail(ic, request.Email)
	if err != nil || user.ID == 0 {
		uclogger.ErrorWithData(common.ErrorGeneral.String(), map[string]interface{}{
			"user": user,
		}, err)
		return c.JSON(http.StatusUnauthorized, common.NewUnauthorizedResponse(common.EmptyObject))
	}

	if !ctrl.authService.VerifyLogin(ic, request.Email, ctrl.appConfig.AppUserPasswordSalt, request.Password) {
		uclogger.ErrorWithData(common.ErrorGeneral.String(), map[string]interface{}{
			"user": user,
		}, err)
		return c.JSON(http.StatusUnauthorized, common.NewUnauthorizedResponse(common.EmptyObject))
	}

	token, err := ctrl.authService.GenerateToken(ctrl.appConfig.AppJWTSign, user.ID, user.Role)
	if err != nil {
		uclogger.ErrorWithData(common.ErrorGeneral.String(), map[string]interface{}{
			"user": user,
		}, err)
		return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse(common.EmptyObject))
	}

	responseToken := map[string]string{"token": token}
	response := common.NewResponse(http.StatusText(http.StatusOK), responseToken)

	uclogger.InfoWithData("ok", map[string]interface{}{"response": response})
	return c.JSON(http.StatusOK, response)
}

func (ctrl *Controller) Panic(c echo.Context) error {
	trackerID, _ := c.Get("tracker_id").(string)
	uclogger.SetTrackerID(trackerID)
	uclogger.Info("panic_controller")
	panic("panic")
}
