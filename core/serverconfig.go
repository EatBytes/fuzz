package core

type HEADER struct {
	Key   string
	Value string
}

type SERVERCONFIG struct {
	Url       string
	Method    string
	Parameter string
	Key       string
	Headers   []HEADER
	Raw       bool
}

func (srv SERVERCONFIG) IsProtected() bool {
	if srv.Key != "" {
		return true
	}

	return false
}
