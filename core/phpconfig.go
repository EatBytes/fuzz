package core

import "bytes"
import "io"

type PHPCONFIG struct {
	Raw    bool
	Upload bool
	Buffer *bytes.Buffer
	Writer io.Writer
}
