package network

import (
	"bytes"
	"net/url"

	"github.com/eatbytes/fuzz/normalizer"
)

type netconfig struct {
	Url    string
	Method string
	Form   *bytes.Buffer
	Jar    []string
	Proxy  string
	File   bool
}

func (n *NETWORK) GetConfig(r string) *netconfig {
	switch n.config.Method {
	case "POST":
		return n.GetPOSTConfig(r)
	case "HEADER":
		return n.GetHEADERConfig(r)
	case "COOKIE":
		break
	}

	return n.GetGETConfig(r)
}

func (n *NETWORK) normalizeRequest(r string) string {
	var request string

	request = normalizer.Encode(r)
	n.cmd = request

	return request
}

func (n *NETWORK) GetPOSTConfig(r string) *netconfig {
	var form url.Values
	var data *bytes.Buffer

	n.status = true

	form = url.Values{}
	form.Set(n.config.Parameter, n.normalizeRequest(r))
	n._body = form

	data = bytes.NewBufferString(form.Encode())

	return &netconfig{
		Url:    n.config.Url,
		Method: "POST",
		Form:   data,
	}
}

func (n *NETWORK) GetGETConfig(r string) *netconfig {
	var url string

	n.status = true
	url = n.config.Url + "?" + n.config.Parameter + "=" + n.normalizeRequest(r)

	return &netconfig{
		Url:    url,
		Method: "GET",
		Form:   nil,
	}
}

func (n *NETWORK) GetHEADERConfig(r string) *netconfig {
	n.status = true

	return &netconfig{
		Url:    n.config.Url,
		Method: "GET",
		Form:   nil,
	}
}

func (n *NETWORK) GetCOOKIEConfig(r string) {
}
