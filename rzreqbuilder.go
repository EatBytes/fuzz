package razboy

import (
	"bytes"
	"net/http"
	"net/url"

	"github.com/eatbytes/razboy/normalizer"
	"github.com/eatbytes/razboy/razerror"
)

const KEY = "RAZBOYNIK_KEY"

func _createSimpleRequest(req *REQUEST) (*razRequest, error) {
	var (
		rzReq *razRequest
		raw   string
		err   error
	)

	switch req.SRVc.GetMethod() {
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
	default:
		err = razerror.ConfigErrInvalidType()
	}

	if err != nil {
		return nil, err
	}

	return rzReq, nil
}

func _createUploadRequest(req *REQUEST) (*razRequest, error) {
	var (
		rzReq *razRequest
		err   error
	)

	rzReq = _buildRzReqBase(req)
	rzReq.http, err = http.NewRequest("POST", req.SRVc.GetUrl(), req.PHPc.buffer)

	if err != nil {
		return nil, err
	}

	rzReq.http.Header.Add("Content-Type", req.PHPc.bondary)

	if req.SRVc.IsProtected() {
		//Write into buffer
	}

	return rzReq, nil
}

func _buildRzReqBase(req *REQUEST) *razRequest {
	var rzReq *razRequest

	rzReq = &razRequest{
		raw:    req.raw,
		status: true,
	}

	if !req.SRVc.IsRaw() {
		rzReq.cmd = normalizer.Encode(req.raw)
	}

	return rzReq
}

func _buildGET(req *REQUEST) (*razRequest, error) {
	var (
		rzReq *razRequest
		url   string
		err   error
	)

	rzReq = _buildRzReqBase(req)

	url = req.SRVc.GetUrl() + "?" + req.SRVc.GetParameter() + "=" + rzReq.GetCMD()

	if req.SRVc.IsProtected() {
		url = url + "&" + KEY + "=" + req.SRVc.GetKey()
	}

	rzReq.http, err = http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	return rzReq, nil
}

func _buildPOST(req *REQUEST) (*razRequest, error) {
	var (
		rzReq *razRequest
		form  url.Values
		data  *bytes.Buffer
		err   error
	)

	rzReq = _buildRzReqBase(req)

	form = url.Values{}
	form.Set(req.SRVc.GetParameter(), rzReq.GetCMD())

	if req.SRVc.IsProtected() {
		form.Add(KEY, req.SRVc.GetKey())
	}

	//req.data = &form

	data = bytes.NewBufferString(form.Encode())

	rzReq.http, err = http.NewRequest("POST", req.SRVc.GetUrl(), data)

	if err != nil {
		return nil, err
	}

	rzReq.http.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return rzReq, nil
}

func _buildHEADER(req *REQUEST) (*razRequest, error) {
	var (
		rzReq *razRequest
		err   error
	)

	rzReq = _buildRzReqBase(req)

	rzReq.http, err = http.NewRequest("GET", req.SRVc.GetUrl(), nil)

	if err != nil {
		return nil, err
	}

	rzReq.http.Header.Add(req.SRVc.GetParameter(), rzReq.GetCMD())

	if req.SRVc.IsProtected() {
		rzReq.http.Header.Add(KEY, req.SRVc.GetKey())
	}

	return rzReq, nil
}

func _buildCOOKIE(req *REQUEST) (*razRequest, error) {
	var (
		rzReq           *razRequest
		cookie, kcookie *http.Cookie
		err             error
	)

	rzReq = _buildRzReqBase(req)

	rzReq.http, err = http.NewRequest("GET", req.SRVc.GetUrl(), nil)

	if err != nil {
		return nil, err
	}

	cookie = &http.Cookie{Name: req.SRVc.GetParameter(), Value: rzReq.GetCMD()}
	rzReq.http.AddCookie(cookie)

	if req.SRVc.IsProtected() {
		kcookie = &http.Cookie{Name: KEY, Value: req.SRVc.GetKey()}
		rzReq.http.AddCookie(kcookie)
	}

	return rzReq, nil
}
