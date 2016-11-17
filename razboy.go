package razboy

import (
	"errors"
	"net/http"

	"github.com/eatbytes/razboy/checker"
	"github.com/eatbytes/razboy/core"
)

func main() {}

func Send(req *core.REQUEST) (*RazResponse, error) {
	var (
		rzReq *RazRequest
		rzRes *RazResponse
		err   error
	)

	err = checker.Check(req)

	if err != nil {
		return nil, err
	}

	rzReq, err = Prepare(req)

	if err != nil {
		return nil, err
	}

	rzRes, err = SendRequest(rzReq)

	if err != nil {
		return nil, err
	}

	return rzRes, nil
}

func Prepare(req *core.REQUEST) (*RazRequest, error) {
	var (
		rzReq *RazRequest
		err   error
	)

	err = checker.Check(req)

	if err != nil {
		return nil, err
	}

	if req.PHPc.Upload {
		rzReq, err = _createUploadRequest(req)
	} else {
		rzReq, err = _createSimpleRequest(req)
	}

	if err != nil {
		return nil, err
	}

	return rzReq, nil
}

func SendRequest(rzReq *RazRequest) (*RazResponse, error) {
	var (
		rzRes  *RazResponse
		client *http.Client
		resp   *http.Response
		err    error
	)

	if !rzReq.status {
		return nil, errors.New("Problem with request")
	}

	client = &http.Client{}
	resp, err = client.Do(rzReq.http)

	if err != nil {
		return nil, err
	}

	rzRes = &RazResponse{
		http:  resp,
		rzReq: rzReq,
	}

	return rzRes, nil
}

func Test() (bool, error) {
	var (
		r     string
		req   *core.REQUEST
		rzRes *RazResponse
		err   error
	)

	//r = "$r=1;" + n.Response()
	req = &core.REQUEST{}

	rzRes, err = Send(req)

	if err != nil {
		return false, err
	}

	r = rzRes.GetResult()

	if r != "1" {
		return false, nil
	}

	return true, nil
}
