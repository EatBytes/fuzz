package razboy

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/eatbytes/razboy/normalizer"
)

type RazResponse struct {
	http  *http.Response
	rzReq *RazRequest
}

func (rzRes *RazResponse) GetHTTP() *http.Response {
	return rzRes.http
}

func (rzRes *RazResponse) GetRequest() *RazRequest {
	return rzRes.rzReq
}

func (rzRes *RazResponse) GetBody() []byte {
	var (
		buffer []byte
		err    error
	)

	defer rzRes.http.Body.Close()
	buffer, err = ioutil.ReadAll(rzRes.http.Body)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return buffer
}

func (rzRes *RazResponse) GetBodyStr() string {
	return string(rzRes.GetBody())
}

func (rzRes *RazResponse) GetHeaderStr() string {
	return rzRes.http.Header.Get(rzRes.rzReq.GetParameter())
}

func (rzRes *RazResponse) GetCookieStr() string {
	var (
		str     string
		cookies []*http.Cookie
		cookie  *http.Cookie
	)

	cookies = rzRes.http.Cookies()

	for _, cookie = range cookies {
		if cookie.Name == rzRes.rzReq.GetParameter() {
			str, _ = url.QueryUnescape(cookie.Value)
			return str
		}
	}

	return ""
}

func (rzRes *RazResponse) GetResultStrByMethod(m string) string {
	if m == "HEADER" {
		return rzRes.GetHeaderStr()
	}

	if m == "COOKIE" {
		return rzRes.GetCookieStr()
	}

	return rzRes.GetBodyStr()
}

func (rzRes *RazResponse) GetResultStr() string {
	return rzRes.GetResultStrByMethod(rzRes.rzReq.GetMethod())
}

func (rzRes *RazResponse) GetResult() string {
	var str string

	str = rzRes.GetResultStr()

	if !rzRes.rzReq.IsRaw {
		str, _ = normalizer.Decode(str)
	}

	return str
}
