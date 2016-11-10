package network

import (
	"io/ioutil"
	"net/http"

	"net/url"

	"github.com/eatbytes/razboy/core"
	"github.com/eatbytes/razboy/normalizer"
)

type Response struct {
	Http   *http.Response
	body   []byte
	config core.Config
}

func (resp *Response) GetBody() []byte {
	var (
		buffer []byte
		err    error
	)

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
	return resp.Http.Header.Get(resp.config.Parameter)
}

func (resp *Response) GetCookieStr() string {
	var (
		str     string
		cookies []*http.Cookie
		cookie  *http.Cookie
	)

	cookies = resp.Http.Cookies()

	for _, cookie = range cookies {
		if cookie.Name == resp.config.Parameter {
			str, _ = url.QueryUnescape(cookie.Value)
			return str
		}
	}

	return ""
}

func (resp *Response) GetResultStrByMethod(m string) string {
	if m == HEADER {
		return resp.GetHeaderStr()
	}

	if m == COOKIE {
		return resp.GetCookieStr()
	}

	return resp.GetBodyStr()
}

func (resp *Response) GetResultStr() string {
	return resp.GetResultStrByMethod(resp.config.Method)
}

func (resp *Response) GetResult() string {
	var str string

	str = resp.GetResultStr()

	if !resp.config.Raw {
		str, _ = normalizer.Decode(str)
	}

	return str
}
