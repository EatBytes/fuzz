package razboy

import "net/http"

type RazRequest struct {
	url       string
	method    string
	parameter string
	cmd       string
	IsRaw     bool
	status    bool
	http      *http.Request
}

func (rzReq RazRequest) GetCMD() string {
	return rzReq.cmd
}

func (rzReq RazRequest) GetParameter() string {
	return rzReq.parameter
}

func (rzReq RazRequest) GetMethod() string {
	return rzReq.method
}

func (rzReq RazRequest) GetStatus() bool {
	return rzReq.status
}

func (rzReq RazRequest) GetHTTP() *http.Request {
	return rzReq.http
}
