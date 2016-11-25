package razboy

import (
	"errors"
)

func Check(req *REQUEST) error {
	if req == nil {
		return errors.New("Empty pointer")
	}

	return _checkSERVER(req)
}

func _checkSERVER(req *REQUEST) error {
	if req.Url == "" {
		return errors.New("REQUEST [url] should not be empty")
	}

	if req.Method != "GET" && req.Method != "POST" && req.Method != "HEADER" && req.Method != "COOKIE" {
		req.Method = "GET"
	}

	if req.Parameter == "" {
		req.Parameter = "razboynik"
	}

	return nil
}
