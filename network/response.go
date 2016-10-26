package network

import (
	"io/ioutil"
	"net/http"
)

type Response struct {
	Http      *http.Response
	body      []byte
	parameter string
	method    string
}

func (resp *Response) GetBody() []byte {
	var buffer []byte
	var err error

	if resp.body != nil {
		return resp.body
	}

	defer resp.Http.Body.Close()
	buffer, err = ioutil.ReadAll(resp.Http.Body)

	if err != nil {
		panic(err)
	}

	resp.body = buffer

	return buffer
}

func (resp *Response) GetBodyStr() string {
	return string(resp.GetBody())
}

func (resp *Response) GetHeaderStr() string {
	return resp.Http.Header.Get(resp.parameter)
}

func (resp *Response) GetResultStrByMethod(m string) string {
	if m == HEADER {
		return resp.GetHeaderStr()
	}

	if m == COOKIE {
	}

	return resp.GetBodyStr()
}

func (resp *Response) GetResultStr() string {
	return resp.GetResultStrByMethod(resp.method)
}
