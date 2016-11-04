package network

import (
	"bytes"
	"net/http"
	"net/url"

	"github.com/eatbytes/razboy/core"
	"github.com/eatbytes/razboy/normalizer"
)

type Request struct {
	Http   *http.Request
	cmd    string
	raw    string
	data   *url.Values
	status bool
	config core.Config
}

func (n *NETWORK) CreateRequest(str string) (*Request, error) {
	var req *Request
	var err error

	req = &Request{
		raw:    str,
		status: true,
		config: *n.config,
	}

	if !n.config.Raw {
		req.cmd = normalizer.Encode(req.raw)
	}

	switch req.config.Method {
	case POST:
		err = req.buildPOSTConfig()
		break
	case HEADER:
		err = req.buildHEADERConfig()
		break
	case COOKIE:
		break
	case GET:
		err = req.buildGETConfig()
		break
	default:
		req.status = false
	}

	if err != nil {
		return nil, err
	}

	return req, nil
}

func (req *Request) buildPOSTConfig() error {
	var form url.Values
	var data *bytes.Buffer
	var err error

	form = url.Values{}
	form.Set(req.config.Parameter, req.cmd)

	if req.config.Key != "" {
		form.Add(KEY, req.config.Key)
	}

	req.data = &form

	data = bytes.NewBufferString(form.Encode())
	req.Http, err = http.NewRequest(POST, req.config.Url, data)

	if err != nil {
		return err
	}

	req.Http.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return nil
}

func (req *Request) buildGETConfig() error {
	var err error

	req.config.Url = req.config.Url + "?" + req.config.Parameter + "=" + req.cmd

	if req.config.Key != "" {
		req.config.Url = req.config.Url + "&" + KEY + "=" + req.config.Key
	}

	req.Http, err = http.NewRequest(GET, req.config.Url, nil)

	if err != nil {
		return err
	}

	return nil
}

func (req *Request) buildHEADERConfig() error {
	var err error

	req.Http, err = http.NewRequest(GET, req.config.Url, nil)

	if err != nil {
		return err
	}

	req.Http.Header.Add(req.config.Parameter, req.cmd)

	if req.config.Key != "" {
		req.Http.Header.Add(KEY, req.config.Key)
	}

	return nil
}

func (req *Request) buildCOOKIEConfig() {
}
