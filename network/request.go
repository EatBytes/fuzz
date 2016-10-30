package network

import (
	"bytes"
	"net/http"
	"net/url"

	"github.com/eatbytes/razboy/normalizer"
)

type Request struct {
	Http   *http.Request
	cmd    string
	raw    string
	data   *url.Values
	status bool
	url    string
}

func (n *NETWORK) buildRequest(str string) (*Request, error) {
	var req *Request
	var err error

	req = &Request{
		raw: str,
	}

	req.status = true
	req.url = n.config.Url

	if n.config.Base64 {
		req.cmd = normalizer.Encode(req.raw)
	}

	switch n.config.Method {
	case POST:
		err = n.buildPOSTConfig(req)
		break
	case HEADER:
		err = n.buildHEADERConfig(req)
		break
	case COOKIE:
		break
	case GET:
		err = n.buildGETConfig(req)
		break
	default:
		req.status = false
	}

	if err != nil {
		return nil, err
	}

	if n.config.Key != "" {
		req.Http.Header.Add(KEY, n.config.Key)
	}

	return req, nil
}

func (n *NETWORK) buildPOSTConfig(r *Request) error {
	var form url.Values
	var data *bytes.Buffer
	var err error

	form = url.Values{}
	form.Set(n.config.Parameter, r.cmd)
	r.data = &form

	data = bytes.NewBufferString(form.Encode())
	r.Http, err = http.NewRequest(POST, r.url, data)

	if err != nil {
		return err
	}

	r.Http.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return nil
}

func (n *NETWORK) buildGETConfig(r *Request) error {
	var err error

	r.url = n.config.Url + "?" + n.config.Parameter + "=" + r.cmd
	r.Http, err = http.NewRequest(GET, r.url, nil)

	if err != nil {
		return err
	}

	return nil
}

func (n *NETWORK) buildHEADERConfig(r *Request) error {
	var err error

	r.Http, err = http.NewRequest(GET, r.url, nil)

	if err != nil {
		return err
	}

	r.Http.Header.Add(n.config.Parameter, r.cmd)

	return nil
}

func (n *NETWORK) buildCOOKIEConfig(r *Request) {
}
