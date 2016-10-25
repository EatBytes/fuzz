package network

import (
	"bytes"
	"net/http"
	"net/url"

	"github.com/eatbytes/fuzz/core"
	"github.com/eatbytes/fuzz/ferror"
)

type NETWORK struct {
	config *core.Config
	status bool
	cmd    string

	_body         url.Values
	_respBody     []byte
	_lastResponse *http.Response
}

func (n *NETWORK) PrepareUpload(bytes *bytes.Buffer, bondary string) (*http.Request, error) {
	var ferr ferror.FuzzerError
	var config netconfig
	var req *http.Request
	var err error

	if n.status != true {
		ferr = ferror.SetupErr()
		return nil, ferr
	}

	n.status = true

	config = netconfig{
		Url:    n.config.Url,
		Method: "POST",
		Form:   bytes,
	}

	req, err = http.NewRequest(config.Method, config.Url, config.Form)
	req.Header.Set("Content-Type", bondary)

	if err != nil {
		ferr := ferror.BuildRequestErr(err)
		return nil, ferr
	}

	return req, nil
}

func (n *NETWORK) Prepare(r string) (*http.Request, error) {
	var ferr ferror.FuzzerError
	var netconfig *netconfig
	var req *http.Request
	var err error

	if n.status != true {
		ferr = ferror.SetupErr()
		return nil, ferr
	}

	netconfig = n.GetConfig(r)

	if netconfig.Form != nil {
		req, err = http.NewRequest(netconfig.Method, netconfig.Url, netconfig.Form)
	} else {
		req, err = http.NewRequest(netconfig.Method, netconfig.Url, nil)
	}

	if err != nil {
		ferr = ferror.BuildRequestErr(err)
		return nil, ferr
	}

	if n.config.Method == "POST" {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	if n.config.Method == "HEADER" {
		req.Header.Add(n.config.Parameter, n.cmd)
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
