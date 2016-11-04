package network

import (
	"bytes"
	"net/http"
	"strings"

	"github.com/eatbytes/razboy/core"
	"github.com/eatbytes/razboy/ferror"
)

const (
	GET    = "GET"
	POST   = "POST"
	HEADER = "HEADER"
	COOKIE = "COOKIE"
	KEY    = "RAZBOYNIK_KEY"
	PARAM  = "razboynik"
)

type NETWORK struct {
	config *core.Config
	status bool

	response *Response
	request  *Request
}

func Create(config *core.Config) (*NETWORK, error) {
	var n *NETWORK

	config.Url = strings.TrimSpace(config.Url)
	config.Method = strings.TrimSpace(strings.ToUpper(config.Method))
	config.Parameter = strings.TrimSpace(config.Parameter)
	config.Key = strings.TrimSpace(config.Key)

	if config.Url == "" {
		return nil, ferror.Default("The url should be specified")
	}

	if !strings.Contains(config.Url, "http://") && !strings.Contains(config.Url, "https://") {
		config.Url = "http://" + config.Url
	}

	if config.Method == "" {
		config.Method = GET
	}

	if config.Method != GET && config.Method != POST && config.Method != HEADER && config.Method != COOKIE {
		return nil, ferror.Default("The method (" + config.Method + ") is not a valid one. Please choose between: GET, POST, HEADER or COOKIE.")
	}

	if config.Parameter == "" {
		config.Parameter = PARAM
	}

	config.Crypt = false

	n = &NETWORK{
		config: config,
		status: true,
	}

	return n, nil
}

func (n *NETWORK) PrepareUpload(buf *bytes.Buffer, bondary string) (*Request, error) {
	var req *Request
	var err error

	if !n.IsSetup() {
		return nil, ferror.SetupErr()
	}

	req = &Request{
		config: *n.config,
		status: true,
	}

	req.Http, err = http.NewRequest(POST, req.config.Url, buf)
	req.Http.Header.Set("Content-Type", bondary)

	if n.config.Key != "" {
		req.Http.Header.Add(KEY, req.config.Key)
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

	req, err = n.CreateRequest(r)

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
		Http:   resp,
		body:   nil,
		config: req.config,
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
