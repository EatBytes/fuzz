package network

import (
	"github.com/eatbytes/razboy/ferror"
	"github.com/eatbytes/razboy/normalizer"
)

func (n *NETWORK) IsSetup() bool {
	return n.status
}

func (n *NETWORK) Test() (bool, error) {
	var r string
	var resp *Response
	var err error

	r = "$r=1;" + n.Response()
	resp, err = n.PrepareSend(r)

	if err != nil {
		return false, err
	}

	r = resp.GetResult()

	if r != "1" {
		return false, ferror.TestErr(resp.Http, n.request.Http, r)
	}

	return true, nil
}

func (n *NETWORK) QuickSend(str string) (string, error) {
	var resp *Response
	var err error

	resp, err = n.PrepareSend(str)

	if err != nil {
		return "", err
	}

	return resp.GetResultStr(), nil
}

func (n *NETWORK) QuickProcess(str string) (string, error) {
	var resp string
	var result string
	var err error

	resp, err = n.QuickSend(str)

	if err != nil {
		return "", err
	}

	result, err = normalizer.Decode(resp)

	if err != nil {
		return "", err
	}

	return result, nil
}

func (n *NETWORK) PrepareSend(str string) (*Response, error) {
	var req *Request
	var resp *Response
	var err error

	req, err = n.Prepare(str)

	if err != nil {
		return nil, err
	}

	resp, err = n.Send(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (n *NETWORK) Response() string {
	switch n.config.Method {
	case HEADER:
		return "header('" + n.config.Parameter + ":' . " + normalizer.PHPEncode("$r") + ");exit();"
	case COOKIE:
		return "setcookie('" + n.config.Parameter + "', " + normalizer.PHPEncode("$r") + ");exit();"
	}

	return "echo(" + normalizer.PHPEncode("$r") + ");exit();"
}
