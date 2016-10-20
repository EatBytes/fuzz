package core

import "bytes"

type Config struct {
	Url    string
	Method string
	Form   *bytes.Buffer
	Jar    []string
	Proxy  string
	File   bool
}
