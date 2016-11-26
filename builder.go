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

func _createSimpleRequest(req *REQUEST) error {
	var (
		proxy *url.URL
		err   error
	)

	switch req.c.Method {
	case "GET":
		err = _buildGET(req)
		break
	case "POST":
		err = _buildPOST(req)
		break
	case "HEADER":
		err = _buildHEADER(req)
		break
	case "COOKIE":
		err = _buildCOOKIE(req)
		break
	}

	if req.c.Proxy != "" {
		proxy, err = url.Parse(req.c.Proxy)

		if err != nil {
			return err
		}

		http.DefaultTransport = &http.Transport{
			Proxy: http.ProxyURL(proxy),
		}
	}

	if len(req.Headers) > 0 {
		for _, header := range req.Headers {
			req.http.Header.Add(header.Key, header.Value)
		}
	}

	if err != nil {
		return err
	}

	return nil
}

func _createUploadRequest(req *REQUEST) error {
	var (
		writer *multipart.Writer
		data   *bytes.Buffer
		err    error
	)

	_buildRzReqBase(req)

	data = req.Buffer

	writer = multipart.NewWriter(data)
	writer.WriteField(req.c.Parameter, req.cmd)

	if req.IsProtected() {
		writer.WriteField(KEY, req.c.Key)
	}

	req.Buffer = data

	req.http, err = http.NewRequest("POST", req.c.Url, data)

	if err != nil {
		return err
	}

	err = writer.Close()

	if err != nil {
		return err
	}

	fmt.Println(writer.FormDataContentType())

	req.http.Header.Add("Content-Type", writer.FormDataContentType())

	return nil
}

func _buildRzReqBase(req *REQUEST) {
	req.setup = true

	if !req.c.Raw {
		req.cmd = normalizer.Encode(req.Action)
	}
}

func _buildGET(req *REQUEST) error {
	var (
		url string
		err error
	)

	_buildRzReqBase(req)

	url = req.c.Url + "?" + req.c.Parameter + "=" + req.cmd

	if req.IsProtected() {
		url += "&" + KEY + "=" + req.c.Key
	}

	req.http, err = http.NewRequest("GET", url, nil)

	return err
}

func _buildPOST(req *REQUEST) error {
	var (
		form url.Values
		data *bytes.Buffer
		err  error
	)

	_buildRzReqBase(req)

	form = url.Values{}
	form.Set(req.c.Parameter, req.cmd)

	if req.IsProtected() {
		form.Add(KEY, req.c.Key)
	}

	data = bytes.NewBufferString(form.Encode())

	req.http, err = http.NewRequest("POST", req.c.Url, data)

	if err != nil {
		return err
	}

	req.http.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return nil
}

func _buildHEADER(req *REQUEST) error {
	var err error

	_buildRzReqBase(req)

	req.http, err = http.NewRequest("GET", req.c.Url, nil)

	if err != nil {
		return err
	}

	req.http.Header.Add(req.c.Parameter, req.cmd)

	if req.IsProtected() {
		req.http.Header.Add(KEY, req.c.Key)
	}

	return nil
}

func _buildCOOKIE(req *REQUEST) error {
	var (
		cookie, kcookie *http.Cookie
		err             error
	)

	_buildRzReqBase(req)

	req.http, err = http.NewRequest("GET", req.c.Url, nil)

	if err != nil {
		return err
	}

	cookie = &http.Cookie{Name: req.c.Parameter, Value: req.cmd}
	req.http.AddCookie(cookie)

	if req.IsProtected() {
		kcookie = &http.Cookie{Name: KEY, Value: req.c.Key}
		req.http.AddCookie(kcookie)
	}

	return nil
}
