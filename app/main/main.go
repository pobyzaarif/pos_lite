package main

import (
	"context"
	"log"
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
	reportAPIModule "github.com/pobyzaarif/pos_lite/modules/reportAPI"
	transactionModule "github.com/pobyzaarif/pos_lite/modules/transaction"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	gommonLog "github.com/labstack/gommon/log"
	"gorm.io/gorm"

	goutilAppName "github.com/pobyzaarif/goutil/appname"
	goutilLogger "github.com/pobyzaarif/goutil/logger"
	goutilEchoMiddlerware "github.com/pobyzaarif/goutil/webframework/echo/v4/middleware"
)

var logger = goutilLogger.NewLog("main")

func newUserService(db *gorm.DB) (userService user.Service) {
	userRepo := userModule.NewGormRepository(db)
	userService = user.NewService(userRepo)

	return
}

func newTransactionService(db *gorm.DB) (transactionService transaction.Service) {
	transactionRepo := transactionModule.NewGormRepository(db)
	reportAPIRepo := reportAPIModule.NewAPI()
	transactionService = transaction.NewService(transactionRepo, reportAPIRepo)

	return
}

func main() {

	conf := config.GetAPPConfig()
	db := conf.GetDatabaseConnection()

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(goutilEchoMiddlerware.ServiceRequestTime)
	e.Use(goutilEchoMiddlerware.ServiceTrackerID)
	e.Use(middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
		Handler: goutilEchoMiddlerware.APILogHandler,
		Skipper: middleware.DefaultSkipper,
	}))

	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 4 << 10,
		LogLevel:  gommonLog.ERROR,
	}))

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
			log.Fatal("failed on http server" + conf.AppMainPort)
		}
	}()

	logger.SetTrackerID("main")
	logger.Info(goutilAppName.GetAPPName() + " service running in " + address)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// a timeout of 10 seconds to shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatalf("failed to shutting down echo server %v", err)
	} else {
		log.Print("successfully shutting down echo server")
	}
}
