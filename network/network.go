package network

import (
	"bytes"
	"net/http"

	"github.com/eatbytes/razboy/core"
	"github.com/eatbytes/razboy/ferror"
)

const (
	GET    = "GET"
	POST   = "POST"
	HEADER = "HEADER"
	COOKIE = "COOKIE"
)

type NETWORK struct {
	config *core.Config
	status bool

	response *Response
	request  *Request
}

func (n *NETWORK) PrepareUpload(bytes *bytes.Buffer, bondary string) (*Request, error) {
	var ferr ferror.FuzzerError
	var req *Request
	var err error

	if !n.IsSetup() {
		ferr = ferror.SetupErr()
		return nil, ferr
	}

	req.Http, err = http.NewRequest(POST, n.config.Url, bytes)
	req.Http.Header.Set("Content-Type", bondary)

	if err != nil {
		ferr := ferror.BuildRequestErr(err)
		return nil, ferr
	}

	return req, nil
}

func (n *NETWORK) Prepare(r string) (*Request, error) {
	var req *Request
	var err error

	if !n.IsSetup() {
		return nil, ferror.SetupErr()
	}

	req, err = n.buildRequest(r)

	if err != nil {
		return nil, ferror.BuildRequestErr(err)
	}

	if !req.status {
		return nil, ferror.SetupErr()
	}

	return req, nil
}

func (n *NETWORK) Send(req *Request) (*Response, error) {
	var client *http.Client
	var resp *http.Response
	var err error
	var status int

	if !n.IsSetup() {
		return nil, ferror.SetupErr()
	}

	n.request = req

	client = &http.Client{}
	resp, err = client.Do(req.Http)

	if err != nil {
		status = 500

		if resp != nil {
			status = resp.StatusCode
		}

		return nil, ferror.RequestErr(err, status)
	}

	response := &Response{
		resp,
		nil,
		n.config.Parameter,
		n.config.Method,
	}

	n.response = response

	return response, nil
}

func (n *NETWORK) GetUrl() string {
	return n.config.Url
}

func (n *NETWORK) GetMethod() string {
	return n.config.Method
}

func (n *NETWORK) GetParameter() string {
	return n.config.Parameter
}

func (n *NETWORK) GetResponse() *Response {
	return n.response
}

func (n *NETWORK) GetRequest() *Request {
	return n.request
}
