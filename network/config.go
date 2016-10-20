package network

import (
	"bytes"
	"net/url"

	"github.com/eatbytes/fuzz/core"
	"github.com/eatbytes/fuzz/normalizer"
)

func (n *NETWORK) postConfig(r string) *core.Config {
	var request string
	var form url.Values
	var data *bytes.Buffer

	n.status = true

	request = normalizer.Encode(r)
	n.cmd = request

	form = url.Values{}
	form.Set(n.parameter, request)
	n._body = form

	data = bytes.NewBufferString(form.Encode())

	return &core.Config{
		Url:    n.host,
		Method: "POST",
		Form:   data,
	}
}

func (n *NETWORK) getConfig(r string) *core.Config {
	var request string
	var url string

	n.status = true

	request = normalizer.Encode(r)
	n.cmd = request

	url = n.host + "?" + n.parameter + "=" + request

	return &core.Config{
		Url:    url,
		Method: "GET",
		Form:   nil,
	}
}

func (n *NETWORK) headerConfig(r string) *core.Config {
	var request string

	n.status = true

	request = normalizer.Encode(r)
	n.cmd = request

	return &core.Config{
		Url:    n.host,
		Method: "GET",
		Form:   nil,
	}
}

func (n *NETWORK) cookieConfig(r string) {
}
