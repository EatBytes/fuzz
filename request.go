package razboy

import "bytes"

const PHP = "PHP"
const SHELL = "SHELL"

type _shellscope struct {
	Name    string
	Content []string
}

type HEADER struct {
	Key   string
	Value string
}

type REQUEST struct {
	Action      string
	Url         string
	Method      string
	Parameter   string
	Key         string
	Shellmethod string
	Shellscope  string
	Raw         bool
	Upload      bool
	Headers     []HEADER
	Buffer      *bytes.Buffer
	setup       bool
}

func CreateRequest(srv [4]string, shl [2]string, php [2]bool) *REQUEST {
	return &REQUEST{
		Url:         srv[0],
		Method:      srv[1],
		Parameter:   srv[2],
		Key:         srv[3],
		Shellmethod: shl[0],
		Shellscope:  shl[1],
		Raw:         php[0],
		Upload:      php[1],
	}
}

func (r REQUEST) IsProtected() bool {
	if r.Key != "" {
		return true
	}

	return false
}

func (r REQUEST) IsSetup() bool {
	return r.setup
}
