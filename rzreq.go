package razboy

import "net/http"

type razRequest struct {
	raw    string
	cmd    string
	status bool
	http   *http.Request
}

func (rzReq razRequest) GetRaw() string {
	return rzReq.raw
}

func (rzReq razRequest) GetCMD() string {
	return rzReq.cmd
}

func (rzReq razRequest) GetStatus() bool {
	return rzReq.status
}

func (rzReq razRequest) GetHTTP() *http.Request {
	return rzReq.http
}
