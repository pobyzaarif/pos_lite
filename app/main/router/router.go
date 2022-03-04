package router

import (
	"github.com/labstack/echo/v4"

	"github.com/pobyzaarif/pos_lite/app/main/controller"
)

var apiVersion = "v1"

func RegisterPath(
	e *echo.Echo,
	controller *controller.Controller,
) {
	// user
	user := e.Group(apiVersion + "/user")
	user.POST("/login", controller.UserLogin)
	user.POST("/panic", controller.Panic)

	// transaction
	transaction := e.Group(apiVersion + "/transaction")
	transaction.GET("/summary", controller.GetTransactionReportByDate)
	transaction.POST("/send/summary", controller.SendTransactionReportByDate)
}
