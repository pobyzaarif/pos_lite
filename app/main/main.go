package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/pobyzaarif/pos_lite/app/main/controller"
	"github.com/pobyzaarif/pos_lite/app/main/router"
	"github.com/pobyzaarif/pos_lite/config"

	"github.com/pobyzaarif/pos_lite/business/auth"

	"github.com/pobyzaarif/pos_lite/business/user"
	userModule "github.com/pobyzaarif/pos_lite/modules/user"

	"github.com/pobyzaarif/pos_lite/business/transaction"
	transactionModule "github.com/pobyzaarif/pos_lite/modules/transaction"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"

	goLoggerAppName "github.com/pobyzaarif/go-logger/appname"
	goLogger "github.com/pobyzaarif/go-logger/logger"
	goLoggerEchoMiddlerware "github.com/pobyzaarif/go-logger/rest/framework/echo/v4/middleware"
)

var logger = goLogger.NewLog("MAIN")

func newUserService(db *gorm.DB) (userService user.Service) {
	userRepo := userModule.NewGormRepository(db)
	userService = user.NewService(userRepo)

	return
}

func newTransactionService(db *gorm.DB) (transactionService transaction.Service) {
	transactionRepo := transactionModule.NewGormRepository(db)
	transactionService = transaction.NewService(transactionRepo)

	return
}

func main() {

	conf := config.GetAPPConfig()
	db := conf.GetDatabaseConnection()

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(goLoggerEchoMiddlerware.ServiceRequestTime)
	e.Use(goLoggerEchoMiddlerware.ServiceTrackerID)
	e.Use(middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
		Handler: goLoggerEchoMiddlerware.APILogHandler,
		Skipper: goLoggerEchoMiddlerware.DefaultSkipper,
	}))

	e.Use(goLoggerEchoMiddlerware.Recover())

	userService := newUserService(db)
	authService := auth.NewService(userService)
	transactionService := newTransactionService(db)

	controllerAPP := controller.NewController(
		conf,
		authService,
		userService,
		transactionService,
	)

	router.RegisterPath(
		e,
		controllerAPP,
	)

	address := "0.0.0.0:" + conf.AppMainPort
	go func() {
		if err := e.Start(address); err != http.ErrServerClosed {
			logger.Fatal("failed on http server " + conf.AppMainPort)
		}
	}()

	logger.SetTrackerID("main")
	logger.Info(goLoggerAppName.GetAPPName() + " service running in " + address)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// a timeout of 10 seconds to shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		logger.Fatal(fmt.Sprintf("failed to shutting down echo server %v", err))
	} else {
		logger.Info("successfully shutting down echo server")
	}
}
