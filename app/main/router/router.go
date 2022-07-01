package router

import (
	"github.com/labstack/echo/v4"

	"github.com/pobyzaarif/pos_lite/app/main/controller"
	"github.com/pobyzaarif/pos_lite/app/main/middleware"
	"github.com/pobyzaarif/pos_lite/config"
)

var apiVersion = "v1"

func RegisterPath(
	e *echo.Echo,
	appConfig *config.Config,
	controller *controller.Controller,
) {
	// public
	e.GET(apiVersion+"/health", func(c echo.Context) error {
		return c.NoContent(200)
	})

	jwtMiddleware := middleware.JWTMiddleware(appConfig.AppJWTSign)

	// user
	user := e.Group(apiVersion + "/user")
	user.POST("/login", controller.UserLogin)

	// transaction
	transaction := e.Group(apiVersion+"/transaction", jwtMiddleware)
	transaction.GET("/summary", controller.GetTransactionReportByDate, middleware.RBACMiddleware("transaction", "read"))
	transaction.POST("/send/summary", controller.SendTransactionReportByDate, middleware.RBACMiddleware("transaction", "execute"))

	// other
	other := e.Group(apiVersion+"/other", jwtMiddleware)
	other.POST("/panic", controller.Panic, middleware.RBACMiddleware("other", "execute"))
}
