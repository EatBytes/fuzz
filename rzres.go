package razboy

import (
	"net/http"
)

type razResponse struct {
	http *http.Response
	body []byte
}
