package network

import (
	"io/ioutil"
	"net/http"

	"github.com/eatbytes/fuzz/ferror"
	"github.com/eatbytes/fuzz/normalizer"
)

func (n *NETWORK) IsSetup() bool {
	return n.status
}

func (n *NETWORK) GetUrl() string {
	return n.host
}

func (n *NETWORK) GetMethod() int {
	return n.method
}

func (n *NETWORK) GetMethodStr() string {
	if n.method == 0 {
		return "GET"
	}

	if n.method == 1 {
		return "POST"
	}

	if n.method == 3 {
		return "HEADER"
	}

	if n.method == 4 {
		return "COOKIE"
	}

	return "ERROR"
}

func (n *NETWORK) GetParameter() string {
	return n.parameter
}

func (n *NETWORK) GetBody(r *http.Response) []byte {
	if n._respBody != nil {
		return n._respBody
	}

	defer r.Body.Close()
	buffer, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	n._respBody = buffer

	return buffer
}

func (n *NETWORK) GetBodyStr(r *http.Response) string {
	buffer := n.GetBody(r)
	return string(buffer)
}

func (n *NETWORK) GetResponse() *http.Response {
	return n._lastResponse
}

func (n *NETWORK) GetRequest() *http.Request {
	n._lastResponse.Request.PostForm = n._body
	return n._lastResponse.Request
}

func (n *NETWORK) GetHeaderStr(r *http.Response) string {
	str := r.Header.Get(n.parameter)
	return str
}

func (n *NETWORK) GetResultStrByMethod(m int, r *http.Response) string {
	if m == 0 || m == 1 {
		return n.GetBodyStr(r)
	}

	if m == 2 {
		return n.GetHeaderStr(r)
	}

	if m == 3 {
	}

	return ""
}

func (n *NETWORK) GetResultStr(r *http.Response) string {
	return n.GetResultStrByMethod(n.method, r)
}

func (n *NETWORK) Setup(url, parameter string, method int) {
	n.host = url
	n.parameter = parameter
	n.method = method
	n.status = true
}

func (n *NETWORK) Test() (bool, error) {
	var r string
	var resp *http.Response
	var err error
	var ferr ferror.FuzzerError

	r = "$r=1;" + n.Response()
	resp, err = n.PrepareSend(r)

	if err != nil {
		return false, err
	}

	r = n.GetResultStr(resp)

	if r != normalizer.Encode("1") {
		ferr = ferror.TestErr(resp, r)
		return false, ferr
	}

	return true, nil
}

func (n *NETWORK) QuickSend(str string) (string, error) {
	var resp *http.Response
	var err error

	resp, err = n.PrepareSend(str)

	if err != nil {
		return "", err
	}

	return n.GetResultStr(resp), nil
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

func (n *NETWORK) PrepareSend(str string) (*http.Response, error) {
	var req *http.Request
	var resp *http.Response
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
