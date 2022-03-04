package controller

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/pobyzaarif/pos_lite/app/main/common"
)

type (
	transactionByDate struct {
		Date string `query:"date" validate:"required,datetime=2006-01-02"`
	}

	sendTransactionByDate struct {
		Date string `json:"date" validate:"required,datetime=2006-01-02"`
		URL  string `json:"url" validate:"required,url"`
	}
)

func (ctrl *Controller) GetTransactionReportByDate(c echo.Context) error {
	responseTransactionSummary := map[string]int{"total_transaction": 0, "total_products_sold": 0, "total_gross_earnings": 0}
	request := new(transactionByDate)
	if err := c.Bind(request); err != nil {
		response := common.NewResponse(http.StatusText(http.StatusBadRequest), responseTransactionSummary)
		return c.JSON(http.StatusBadRequest, response)
	}

	err := validator.New().Struct(request)
	if err != nil {
		response := common.NewResponse(err.Error(), responseTransactionSummary)
		return c.JSON(http.StatusBadRequest, response)
	}

	trxSummary, err := ctrl.transactionService.GetSummaryTransaction(request.Date)
	if err != nil {
		response := common.NewResponse(http.StatusText(http.StatusInternalServerError), responseTransactionSummary)
		return c.JSON(http.StatusInternalServerError, response)
	}
	responseTransactionSummary["total_transaction"] = trxSummary.TotalTransaction
	responseTransactionSummary["total_products_sold"] = trxSummary.TotalProductsSold
	responseTransactionSummary["total_gross_earnings"] = trxSummary.TotalGrossEarnings

	return c.JSON(http.StatusOK, responseTransactionSummary)
}

func (ctrl *Controller) SendTransactionReportByDate(c echo.Context) error {
	request := new(sendTransactionByDate)
	if err := c.Bind(request); err != nil {
		response := common.NewResponse(http.StatusText(http.StatusBadRequest), common.EmptyObject)
		return c.JSON(http.StatusBadRequest, response)
	}

	err := validator.New().Struct(request)
	if err != nil {
		response := common.NewResponse(err.Error(), common.EmptyObject)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = ctrl.transactionService.SendSummaryTransaction(request.Date, request.URL)
	if err != nil {
		response := common.NewResponse(err.Error(), common.EmptyObject)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := common.NewResponse(http.StatusText(http.StatusOK), common.EmptyObject)

	return c.JSON(http.StatusOK, response)
}
