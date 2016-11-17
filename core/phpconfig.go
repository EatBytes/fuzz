package core

import "bytes"
import "io"

type PHPCONFIG struct {
	Cmd    string
	Raw    bool
	Upload bool
	Buffer *bytes.Buffer
	Writer io.Writer
}
