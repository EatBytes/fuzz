package ferror

import (
	"fmt"
	"net/http"
)

type NetworkError struct {
	msg  string
	resp *http.Response
	req  *http.Request
	body string
}

func (e NetworkError) Error() string {
	str := "%v \n\n" +
		"     request: %v\n\n" +
		"     response: %v\n" +
		"          body-> %v\n"

	return fmt.Sprintf(str, e.msg, e.req, e.resp, e.body)
}

func TestErr(r *http.Response, req *http.Request, b string) NetworkError {
	return NetworkError{
		msg:  "Error: Server doesn't respond well to test",
		resp: r,
		req:  req,
		body: b,
	}
}
