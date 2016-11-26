package razboy

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/eatbytes/razboy/normalizer"
)

const KEY = "RAZBOYNIK_KEY"

func _createSimpleRequest(req *REQUEST) (*RazRequest, error) {
	var (
		rzReq *RazRequest
		proxy *url.URL
		err   error
	)

	switch req.Method {
	case "GET":
		rzReq, err = _buildGET(req)
		break
	case "POST":
		rzReq, err = _buildPOST(req)
		break
	case "HEADER":
		rzReq, err = _buildHEADER(req)
		break
	case "COOKIE":
		rzReq, err = _buildCOOKIE(req)
		break
	}

	if req.Proxy != "" {
		proxy, err = url.Parse("http://proxyIp:proxyPort")

		if err != nil {
			return nil, err
		}

		http.DefaultTransport = &http.Transport{
			Proxy: http.ProxyURL(proxy),
		}
	}

	if len(req.Headers) > 0 {
		for _, header := range req.Headers {
			rzReq.http.Header.Add(header.Key, header.Value)
		}
	}

	if err != nil {
		return nil, err
	}

	return rzReq, nil
}

func _createUploadRequest(req *REQUEST) (*RazRequest, error) {
	var (
		rzReq  *RazRequest
		writer *multipart.Writer
		data   *bytes.Buffer
		err    error
	)

	rzReq = _buildRzReqBase(req)

	data = req.Buffer

	writer = multipart.NewWriter(data)
	writer.WriteField(req.Parameter, rzReq.GetCMD())

	if req.IsProtected() {
		writer.WriteField(KEY, req.Key)
	}

	req.Buffer = data

	rzReq.http, err = http.NewRequest("POST", req.Url, data)

	if err != nil {
		return nil, err
	}

	err = writer.Close()

	if err != nil {
		return nil, err
	}

	fmt.Println(writer.FormDataContentType())

	rzReq.http.Header.Add("Content-Type", writer.FormDataContentType())

	return rzReq, nil
}

func _buildRzReqBase(req *REQUEST) *RazRequest {
	var rzReq *RazRequest

	rzReq = &RazRequest{
		url:       req.Url,
		parameter: req.Parameter,
		method:    req.Method,
		status:    true,
	}

	if !req.Raw {
		rzReq.cmd = normalizer.Encode(req.Action)
	}

	return rzReq
}

func _buildGET(req *REQUEST) (*RazRequest, error) {
	var (
		rzReq *RazRequest
		url   string
		err   error
	)

	rzReq = _buildRzReqBase(req)

	url = req.Url + "?" + req.Parameter + "=" + rzReq.GetCMD()

	if req.IsProtected() {
		url += "&" + KEY + "=" + req.Key
	}

	rzReq.http, err = http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	return rzReq, nil
}

func _buildPOST(req *REQUEST) (*RazRequest, error) {
	var (
		rzReq *RazRequest
		form  url.Values
		data  *bytes.Buffer
		err   error
	)

	rzReq = _buildRzReqBase(req)

	form = url.Values{}
	form.Set(req.Parameter, rzReq.GetCMD())

	if req.IsProtected() {
		form.Add(KEY, req.Key)
	}

	data = bytes.NewBufferString(form.Encode())

	rzReq.http, err = http.NewRequest("POST", req.Url, data)

	if err != nil {
		return nil, err
	}

	rzReq.http.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return rzReq, nil
}

func _buildHEADER(req *REQUEST) (*RazRequest, error) {
	var (
		rzReq *RazRequest
		err   error
	)

	rzReq = _buildRzReqBase(req)

	rzReq.http, err = http.NewRequest("GET", req.Url, nil)

	if err != nil {
		return nil, err
	}

	rzReq.http.Header.Add(req.Parameter, rzReq.GetCMD())

	if req.IsProtected() {
		rzReq.http.Header.Add(KEY, req.Key)
	}

	return rzReq, nil
}

func _buildCOOKIE(req *REQUEST) (*RazRequest, error) {
	var (
		rzReq           *RazRequest
		cookie, kcookie *http.Cookie
		err             error
	)

	rzReq = _buildRzReqBase(req)

	rzReq.http, err = http.NewRequest("GET", req.Url, nil)

	if err != nil {
		return nil, err
	}

	cookie = &http.Cookie{Name: req.Parameter, Value: rzReq.GetCMD()}
	rzReq.http.AddCookie(cookie)

	if req.IsProtected() {
		kcookie = &http.Cookie{Name: KEY, Value: req.Key}
		rzReq.http.AddCookie(kcookie)
	}

	return rzReq, nil
}
