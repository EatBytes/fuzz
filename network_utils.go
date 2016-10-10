package fuzzcore

import (
	"io/ioutil"
	"net/http"
)

func (n *NETWORK) IsSetup() bool {
	return n.status
}

func (n *NETWORK) SetConfig(url string, method int, parameter string, crypt bool) {
	n.host = url
	n.method = method
	n.parameter = parameter
	n.crypt = crypt

	n.status = true
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
