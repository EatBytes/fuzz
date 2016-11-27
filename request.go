package razboy

import "bytes"
import "net/http"

type _shellscope struct {
	Name    string
	Content []string
}

type HEADER struct {
	Key   string
	Value string
}

type REQUEST struct {
	Action  string
	Scope   string
	Upload  bool
	Headers []HEADER
	Buffer  *bytes.Buffer
	c       *Config
	cmd     string
	http    *http.Request
	setup   bool
}

func CreateRequest(action string, scope string, c *Config) *REQUEST {
	return &REQUEST{
		Action: action,
		Scope:  scope,
		c:      c,
	}
}

func (req REQUEST) IsProtected() bool {
	if req.c.Key != "" {
		return true
	}

	return false
}

func (req REQUEST) GetHTTP() *http.Request {
	return req.http
}

func (req REQUEST) GetConfig() *Config {
	return req.c
}
