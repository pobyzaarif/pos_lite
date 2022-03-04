package reportapi

import "time"

type APIRespository interface {
	SendReport(payload map[string]interface{}, toURL string, timeout time.Duration) (httpcode int, err error)
}
