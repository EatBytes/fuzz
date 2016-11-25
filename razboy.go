package razboy

import (
	"errors"
	"net/http"
)

func main() {}

func Send(req *REQUEST) (*RazResponse, error) {
	var (
		rzReq *RazRequest
		rzRes *RazResponse
		err   error
	)

	err = Check(req)

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

func Prepare(req *REQUEST) (*RazRequest, error) {
	var (
		rzReq *RazRequest
		err   error
	)

	err = Check(req)

	if err != nil {
		return nil, err
	}

	if req.Upload {
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
		req   *REQUEST
		rzRes *RazResponse
		err   error
	)

	//r = "$r=1;" + n.Response()
	req = &REQUEST{}

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
