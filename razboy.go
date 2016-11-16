package razboy

import "github.com/eatbytes/razboy/razerror"

type REQUEST struct {
	Type  string
	PHPc  PHPCONFIG
	SHLc  SHELLCONFIG
	SRVc  SERVERCONFIG
	final string
}

type PHPCONFIG struct {
	IsRaw    bool
	IsUpload bool
}

type SHELLCONFIG struct {
	Method  string
	Context string
}

type SERVERCONFIG struct {
	Url       string
	Method    string
	Parameter string
	Key       string
	IsRaw     bool
}

type razRequest struct {
}

type razResponse struct {
}

func CreatePHP() {

}

func createSHELL() {

}

func Send(req *REQUEST) error {
	var err error

	err = _check(req)

	if err != nil {
		return err
	}

	request, err := _prepare(req)
	response, err := _send(request)

	return nil
}

func _check(req *REQUEST) error {
	if req.Type == "PHP" {

	} else if req.Type == "SHELL" {

	} else {
		return razerror.ConfigErrInvalidType()
	}

	return nil
}

func _prepare(req *REQUEST) (*razRequest, error) {

}

func _send(rz *razRequest) (*razResponse, error) {

}
