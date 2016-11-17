package razboy

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/eatbytes/razboy/core"
	"github.com/eatbytes/razboy/normalizer"
)

const KEY = "RAZBOYNIK_KEY"

func _createSimpleRequest(req *core.REQUEST) (*RazRequest, error) {
	var (
		rzReq *RazRequest
		err   error
	)

	switch req.SRVc.Method {
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

	if len(req.SRVc.Headers) > 0 {
		for _, header := range req.SRVc.Headers {
			rzReq.http.Header.Add(header.Key, header.Value)
		}
	}

	if err != nil {
		return nil, err
	}

	return rzReq, nil
}

func _createUploadRequest(req *core.REQUEST) (*RazRequest, error) {
	var (
		rzReq  *RazRequest
		writer *multipart.Writer
		data   *bytes.Buffer
		err    error
	)

	rzReq = _buildRzReqBase(req)

	data = req.PHPc.Buffer

	writer = multipart.NewWriter(data)
	writer.WriteField(req.SRVc.Parameter, rzReq.GetCMD())

	if req.SRVc.IsProtected() {
		writer.WriteField(KEY, req.SRVc.Key)
	}

	req.PHPc.Buffer = data

	rzReq.http, err = http.NewRequest("POST", req.SRVc.Url, data)

	if err != nil {
		return nil, err
	}

	err = writer.Close()

	if err != nil {
		return nil, err
	}

	rzReq.http.Header.Add("Content-Type", writer.FormDataContentType())

	return rzReq, nil
}

func _buildRzReqBase(req *core.REQUEST) *RazRequest {
	var rzReq *RazRequest

	rzReq = &RazRequest{
		url:       req.SRVc.Url,
		parameter: req.SRVc.Parameter,
		method:    req.SRVc.Method,
		status:    true,
	}

	if !req.SRVc.Raw {
		rzReq.cmd = normalizer.Encode(req.Raw)
	}

	return rzReq
}

func _buildGET(req *core.REQUEST) (*RazRequest, error) {
	var (
		rzReq *RazRequest
		url   string
		err   error
	)

	rzReq = _buildRzReqBase(req)

	url = req.SRVc.Url + "?" + req.SRVc.Parameter + "=" + rzReq.GetCMD()

	if req.SRVc.IsProtected() {
		url = url + "&" + KEY + "=" + req.SRVc.Key
	}

	rzReq.http, err = http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	return rzReq, nil
}

func _buildPOST(req *core.REQUEST) (*RazRequest, error) {
	var (
		rzReq *RazRequest
		form  url.Values
		data  *bytes.Buffer
		err   error
	)

	rzReq = _buildRzReqBase(req)

	form = url.Values{}
	form.Set(req.SRVc.Parameter, rzReq.GetCMD())

	if req.SRVc.IsProtected() {
		form.Add(KEY, req.SRVc.Key)
	}

	data = bytes.NewBufferString(form.Encode())

	rzReq.http, err = http.NewRequest("POST", req.SRVc.Url, data)

	if err != nil {
		return nil, err
	}

	rzReq.http.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return rzReq, nil
}

func _buildHEADER(req *core.REQUEST) (*RazRequest, error) {
	var (
		rzReq *RazRequest
		err   error
	)

	rzReq = _buildRzReqBase(req)

	rzReq.http, err = http.NewRequest("GET", req.SRVc.Url, nil)

	if err != nil {
		return nil, err
	}

	rzReq.http.Header.Add(req.SRVc.Parameter, rzReq.GetCMD())

	if req.SRVc.IsProtected() {
		rzReq.http.Header.Add(KEY, req.SRVc.Key)
	}

	return rzReq, nil
}

func _buildCOOKIE(req *core.REQUEST) (*RazRequest, error) {
	var (
		rzReq           *RazRequest
		cookie, kcookie *http.Cookie
		err             error
	)

	rzReq = _buildRzReqBase(req)

	rzReq.http, err = http.NewRequest("GET", req.SRVc.Url, nil)

	if err != nil {
		return nil, err
	}

	cookie = &http.Cookie{Name: req.SRVc.Parameter, Value: rzReq.GetCMD()}
	rzReq.http.AddCookie(cookie)

	if req.SRVc.IsProtected() {
		kcookie = &http.Cookie{Name: KEY, Value: req.SRVc.Key}
		rzReq.http.AddCookie(kcookie)
	}

	return rzReq, nil
}
