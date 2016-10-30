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
	KEY    = "RAZBOYNIK_KEY"
	PARAM  = "razboy"
)

type NETWORK struct {
	config *core.Config
	status bool

	response *Response
	request  *Request
}

func (n *NETWORK) PrepareUpload(buf *bytes.Buffer, bondary string) (*Request, error) {
	var req *Request
	var err error

	req = &Request{
		url:    n.config.Url,
		status: true,
	}

	if !n.IsSetup() {
		return nil, ferror.SetupErr()
	}

	req.Http, err = http.NewRequest(POST, n.config.Url, buf)
	req.Http.Header.Set("Content-Type", bondary)

	if n.config.Key != "" {
		req.Http.Header.Add(KEY, n.config.Key)
	}

	if err != nil {
		return nil, ferror.BuildRequestErr(err)
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
