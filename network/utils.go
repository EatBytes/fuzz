package network

import (
	"strings"

	"github.com/eatbytes/razboy/core"
	"github.com/eatbytes/razboy/ferror"
	"github.com/eatbytes/razboy/normalizer"
)

func (n *NETWORK) IsSetup() bool {
	return n.status
}

func (n *NETWORK) Setup(cf *core.Config) error {
	var ferr ferror.FuzzerError

	cf.Url = strings.TrimSpace(cf.Url)
	cf.Method = strings.TrimSpace(strings.ToUpper(cf.Method))
	cf.Parameter = strings.TrimSpace(cf.Parameter)
	cf.Key = strings.TrimSpace(cf.Key)

	if cf.Url == "" {
		ferr = ferror.Default("The url should be specified")
		return ferr
	}

	if !strings.Contains(cf.Url, "http://") && !strings.Contains(cf.Url, "https://") {
		cf.Url = "http://" + cf.Url
	}

	if cf.Method != GET && cf.Method != POST && cf.Method != HEADER && cf.Method != COOKIE && cf.Method != "" {
		ferr = ferror.Default("The method (" + cf.Method + ") is not a valid one. Please choose between: GET, POST, HEADER or COOKIE.")
		return ferr
	}

	if cf.Method == "" {
		cf.Method = GET
	}

	if cf.Parameter == "" {
		cf.Parameter = PARAM
	}

	cf.Crypt = false
	n.config = cf
	n.status = true

	return nil
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

	r = resp.GetResultStr()

	if r != normalizer.Encode("1") {
		return false, ferror.TestErr(resp.Http, r)
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
