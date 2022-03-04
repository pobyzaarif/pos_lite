package reportapi

import "time"

type API struct{}

func NewAPI() *API {
	return &API{}
}

func (r *API) SendReport(payload map[string]interface{}, toURL string, timeout time.Duration) (httpcode int, err error) {
	return 0, nil
}
