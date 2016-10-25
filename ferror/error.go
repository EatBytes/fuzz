package ferror

import (
	"fmt"
	"net/http"

	"github.com/eatbytes/fuzz/core"
)

type FuzzerError struct {
	code     int
	msg      string
	bag      error
	conf     *core.Config
	resp     *http.Response
	respBody string
}

func (e FuzzerError) Error() string {
	return fmt.Sprintf("%v: %v \n\nbag: %v \nconf: %v \nresponse: %v \n    body-> %v\n\n", e.code, e.msg, e.bag, e.conf, e.resp, e.respBody)
}

func SetupErr() FuzzerError {
	return FuzzerError{
		msg: "Error: You havn't setup the required information, please refer to srv config.",
	}
}

func RequestErr(err error, c int) FuzzerError {
	return FuzzerError{
		code: c,
		msg:  "Error: Impossible to send request.",
		bag:  err,
	}
}

func BuildRequestErr(err error, c *core.Config) FuzzerError {
	return FuzzerError{
		msg:  "Error: Impossible to create request with config",
		bag:  err,
		conf: c,
	}
}

func NoMethodFoundErr() FuzzerError {
	return FuzzerError{
		msg: "Error: No method was find for the req to prepare",
	}
}

func FileErr(err error) FuzzerError {
	return FuzzerError{
		msg: "Error: Encounter a problem with file",
		bag: err,
	}
}

func PartErr(err error) FuzzerError {
	return FuzzerError{
		msg: "Error: Can't create part",
		bag: err,
	}
}

func NormalizeErr(err error) FuzzerError {
	return FuzzerError{
		msg: "Error: Impossible to normalize the string",
		bag: err,
	}
}

func TestErr(r *http.Response, b string) FuzzerError {
	return FuzzerError{
		msg:      "Error: Server doesn't respond well to test",
		resp:     r,
		respBody: b,
	}
}
