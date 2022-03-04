package controller

import (
	"github.com/pobyzaarif/pos_lite/business/auth"
	"github.com/pobyzaarif/pos_lite/business/transaction"
	"github.com/pobyzaarif/pos_lite/business/user"
	"github.com/pobyzaarif/pos_lite/config"
)

type Controller struct {
	appConfig          *config.Config
	authService        auth.Service
	userService        user.Service
	transactionService transaction.Service
}

func NewController(
	appConfig *config.Config,
	authService auth.Service,
	userService user.Service,
	transactionService transaction.Service,
) *Controller {
	return &Controller{
		appConfig,
		authService,
		userService,
		transactionService,
	}
}
