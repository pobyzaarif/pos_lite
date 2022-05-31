package controller

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	goLogger "github.com/pobyzaarif/go-logger/logger"
	"github.com/pobyzaarif/pos_lite/app/main/common"
	"github.com/pobyzaarif/pos_lite/business"
)

var tclogger = goLogger.NewLog("TRANSACTION_CONTROLLER")

type (
	transactionByDate struct {
		Date string `query:"date" validate:"required,datetime=2006-01-02"`
	}

	sendTransactionByDate struct {
		Date string `json:"date" validate:"required,datetime=2006-01-02,numeric,min=11,max=14"`
		URL  string `json:"url" validate:"required,url"`
	}
)

func (ctrl *Controller) GetTransactionReportByDate(c echo.Context) error {
	trackerID, _ := c.Get("tracker_id").(string)
	tclogger.SetTrackerID(trackerID)
	ic := business.NewInternalContext(trackerID)

	responseTransactionSummary := map[string]int{"total_transaction": 0, "total_products_sold": 0, "total_gross_earnings": 0}
	request := new(transactionByDate)
	if err := c.Bind(request); err != nil {
		tclogger.Error(common.ErrorBindingRequest.String(), err)
		response := common.NewResponse(http.StatusText(http.StatusBadRequest), responseTransactionSummary)
		return c.JSON(http.StatusBadRequest, response)
	}

	err := validator.New().Struct(request)
	if err != nil {
		tclogger.Error(common.ErrorValidationRequest.String(), err)
		response := common.NewResponse(err.Error(), responseTransactionSummary)
		return c.JSON(http.StatusBadRequest, response)
	}

	trxSummary, err := ctrl.transactionService.GetSummaryTransaction(ic, request.Date)
	if err != nil {
		tclogger.ErrorWithData(common.ErrorGeneral.String(), map[string]interface{}{
			"trx_summary": trxSummary,
		}, err)
		response := common.NewResponse(http.StatusText(http.StatusInternalServerError), responseTransactionSummary)
		return c.JSON(http.StatusInternalServerError, response)
	}
	responseTransactionSummary["total_transaction"] = trxSummary.TotalTransaction
	responseTransactionSummary["total_products_sold"] = trxSummary.TotalProductsSold
	responseTransactionSummary["total_gross_earnings"] = trxSummary.TotalGrossEarnings

	tclogger.InfoWithData("ok", map[string]interface{}{"response": responseTransactionSummary})
	return c.JSON(http.StatusOK, responseTransactionSummary)
}

func (ctrl *Controller) SendTransactionReportByDate(c echo.Context) error {
	trackerID, _ := c.Get("tracker_id").(string)
	tclogger.SetTrackerID(trackerID)
	ic := business.NewInternalContext(trackerID)

	responseData := make(map[string]interface{})
	request := new(sendTransactionByDate)
	if err := c.Bind(request); err != nil {
		tclogger.Error(common.ErrorBindingRequest.String(), err)
		response := common.NewResponse(http.StatusText(http.StatusBadRequest), responseData)
		return c.JSON(http.StatusBadRequest, response)
	}

	err := validator.New().Struct(request)
	if err != nil {
		tclogger.Error(common.ErrorValidationRequest.String(), err)
		response := common.NewResponse(err.Error(), responseData)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = ctrl.transactionService.SendSummaryTransaction(ic, request.Date, request.URL)
	if err != nil {
		tclogger.Error(common.ErrorGeneral.String(), err)
		response := common.NewResponse(err.Error(), responseData)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := common.NewResponse(http.StatusText(http.StatusOK), responseData)

	tclogger.InfoWithData("ok", map[string]interface{}{"response": response})
	return c.JSON(http.StatusOK, response)
}
