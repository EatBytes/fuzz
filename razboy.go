package razboy

import (
	"errors"
	"net/http"
)

func main() {}

func Send(req *REQUEST) (*RESPONSE, error) {
	var (
		res *RESPONSE
		err error
	)

	err = Check(req)

	if err != nil {
		return nil, err
	}

	err = Prepare(req)

	if err != nil {
		return nil, err
	}

	res, err = SendRequest(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func Prepare(req *REQUEST) error {
	var err error

	err = Check(req)

	if err != nil {
		return err
	}

	if req.Upload {
		err = _createUploadRequest(req)
	} else {
		err = _createSimpleRequest(req)
	}

	return err
}

func SendRequest(req *REQUEST) (*RESPONSE, error) {
	var (
		res    *RESPONSE
		client *http.Client
		resp   *http.Response
		err    error
	)

	if !req.setup {
		return nil, errors.New("Problem with request")
	}

	client = &http.Client{}
	resp, err = client.Do(req.http)

	if err != nil {
		return nil, err
	}

	res = &RESPONSE{
		http:    resp,
		request: req,
	}

	return res, nil
}

func Test() (bool, error) {
	var (
		r   string
		req *REQUEST
		res *RESPONSE
		err error
	)

	//r = "$r=1;" + n.Response()
	req = &REQUEST{}

	res, err = Send(req)

	if err != nil {
		return false, err
	}

	r = res.GetResult()

	if r != "1" {
		return false, nil
	}

	return true, nil
}
