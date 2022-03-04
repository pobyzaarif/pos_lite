package transaction

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	reportapi "github.com/pobyzaarif/pos_lite/business/transaction/reportAPI"
)

type (
	service struct {
		repository          Repository
		reportAPIRepository reportapi.APIRespository
	}

	Service interface {
		GetSummaryTransaction(date string) (summaryTransaction TransactionSummaryByDate, err error)

		SendSummaryTransaction(date string, to string) (err error)
	}
)

func NewService(repository Repository, reportAPIRepository reportapi.APIRespository) Service {
	return &service{
		repository,
		reportAPIRepository,
	}
}

func (s *service) GetSummaryTransaction(date string) (summaryTransaction TransactionSummaryByDate, err error) {
	return s.repository.GetSummaryTransaction(date)
}

func (s *service) SendSummaryTransaction(date string, toURL string) (err error) {
	trxSummary, err := s.repository.GetSummaryTransaction(date)
	if err != nil {
		return
	}

	payload := new(bytes.Buffer)
	err = json.NewEncoder(payload).Encode(trxSummary)
	if err != nil {
		return
	}

	rand.Seed(time.Now().UnixNano())
	timeDelay := rand.Intn(7000-1000) + 1000 // simulate delay response between 1000 until 7000 mili
	timeDelayString := fmt.Sprintf("%v", timeDelay)
	httpCodeResponseList := []string{"200", "400", "500"} // simulate http response code either 200, 400 or 500
	httpCodeResponse := httpCodeResponseList[rand.Intn(len(httpCodeResponseList))]

	// construct request
	request, err := http.NewRequest(http.MethodPost, toURL+"/"+httpCodeResponse+"?sleep="+timeDelayString, payload)
	if err != nil {
		return
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")

	var client http.Client
	client.Timeout = time.Second * 5

	res, err := client.Do(request)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		err = fmt.Errorf("response HTTP code is not 200")
	}

	return

}
