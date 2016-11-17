package razboy

import (
	"net/http"

	"github.com/eatbytes/razboy/ferror"
)

func CreatePHP() {

}

func createSHELL() {

}

func Send(req *REQUEST) (*razResponse, error) {
	var (
		rzReq *razRequest
		rzRes *razResponse
		err   error
	)

	err = _check(req)

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

func Prepare(req *REQUEST) (*razRequest, error) {
	var (
		rzReq *razRequest
		err   error
	)

	if !req.setup {
		return nil, ferror.SetupErr()
	}

	if req.PHPc.IsUpload() {
		rzReq, err = _createUploadRequest(req)
	} else {
		rzReq, err = _createSimpleRequest(req)
	}

	if err != nil {
		return nil, ferror.BuildRequestErr(err)
	}

	if !req.setup {
		return nil, ferror.SetupErr()
	}

	return rzReq, nil
}

func SendRequest(rzReq *razRequest) (*razResponse, error) {
	var (
		rzRes  *razResponse
		client *http.Client
		resp   *http.Response
		err    error
		status int
	)

	if !rzReq.status {
		return nil, ferror.SetupErr()
	}

	client = &http.Client{}
	resp, err = client.Do(rzReq.http)

	if err != nil {
		status = 500

		if resp != nil {
			status = resp.StatusCode
		}

		return nil, ferror.RequestErr(err, status)
	}

	rzRes = &razResponse{
		http: resp,
		body: nil,
	}

	return rzRes, nil
}
