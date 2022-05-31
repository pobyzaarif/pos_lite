package transaction

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	goLoggerHTTPClient "github.com/pobyzaarif/go-logger/http/client"
	"github.com/pobyzaarif/pos_lite/business"
)

type (
	service struct {
		repository Repository
	}

	Service interface {
		GetSummaryTransaction(ic business.InternalContext, date string) (summaryTransaction TransactionSummaryByDate, err error)

		SendSummaryTransaction(ic business.InternalContext, date string, to string) (err error)
	}
)

func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (s *service) GetSummaryTransaction(ic business.InternalContext, date string) (summaryTransaction TransactionSummaryByDate, err error) {
	return s.repository.GetSummaryTransaction(ic, date)
}

func (s *service) SendSummaryTransaction(ic business.InternalContext, date string, toURL string) (err error) {
	trxSummary, err := s.repository.GetSummaryTransaction(ic, date)
	if err != nil {
		return
	}

	payload := new(bytes.Buffer)
	err = json.NewEncoder(payload).Encode(trxSummary)
	if err != nil {
		return
	}

	// construct request
	request, err := http.NewRequest(http.MethodPost, toURL, payload)
	if err != nil {
		return
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")

	timeout := time.Second * 5

	var response string
	httpCode, err := goLoggerHTTPClient.Call(
		ic.ToContext(),
		request,
		timeout,
		goLoggerHTTPClient.RawResponseBodyFormat,
		&response,
		nil,
	)

	if err != nil {
		err = fmt.Errorf("error processing request to " + toURL)
		return
	}

	if httpCode != 200 {
		err = fmt.Errorf("error response HTTP code is not 200 " + toURL)
	}

	return

}
