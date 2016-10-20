package network

import (
	"bytes"
	"net/http"
	"net/url"

	"github.com/eatbytes/fuzz/core"
	"github.com/eatbytes/fuzz/ferror"
)

type NETWORK struct {
	host      string
	method    int
	parameter string
	crypt     bool
	status    bool
	cmd       string

	_body         url.Values
	_respBody     []byte
	_lastResponse *http.Response
}

func (n *NETWORK) PrepareUpload(bytes *bytes.Buffer, bondary string) (*http.Request, error) {
	var ferr ferror.FuzzerError
	var c core.Config
	var req *http.Request
	var err error

	if n.status != true {
		ferr = ferror.SetupErr()
		return nil, ferr
	}

	n.status = true

	c = core.Config{
		Url:    n.host,
		Method: "POST",
		Form:   bytes,
	}

	req, err = http.NewRequest(c.Method, c.Url, bytes)
	req.Header.Set("Content-Type", bondary)

	if err != nil {
		ferr := ferror.BuildRequestErr(err, &c)
		return nil, ferr
	}

	return req, nil
}

func (n *NETWORK) Prepare(r string) (*http.Request, error) {
	var ferr ferror.FuzzerError
	var config *core.Config
	var req *http.Request
	var err error

	ferr = ferror.NoMethodFoundErr()

	if n.status != true {
		ferr = ferror.SetupErr()
		return nil, ferr
	}

	if n.method == 0 {
		config = n.getConfig(r)
		req, err = http.NewRequest(config.Method, config.Url, nil)

		if err != nil {
			ferr = ferror.BuildRequestErr(err, config)
			return nil, ferr
		}

		return req, nil
	}

	if n.method == 1 {
		config = n.postConfig(r)
		req, err = http.NewRequest(config.Method, config.Url, config.Form)

		if err != nil {
			ferr = ferror.BuildRequestErr(err, config)
			return nil, ferr
		}

		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		return req, nil
	}

	if n.method == 2 {
		config = n.headerConfig(r)
		req, err = http.NewRequest(config.Method, config.Url, nil)

		if err != nil {
			ferr = ferror.BuildRequestErr(err, config)
			return nil, ferr
		}

		req.Header.Add(n.parameter, n.cmd)

		return req, nil
	}

	if n.method == 3 {
	}

	return req, ferr
}

func (n *NETWORK) Send(req *http.Request) (*http.Response, error) {
	var ferr ferror.FuzzerError
	var client *http.Client
	var resp *http.Response
	var err error
	var status int

	if n.status != true {
		ferr = ferror.SetupErr()
		return nil, ferr
	}

	n._respBody = nil

	client = &http.Client{}
	resp, err = client.Do(req)

	if err != nil {
		status = 500

		if resp != nil {
			status = resp.StatusCode
		}

		ferr = ferror.RequestErr(err, status)
		return nil, ferr
	}

	n._lastResponse = resp

	return resp, nil
}
