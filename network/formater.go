package network

import "github.com/eatbytes/fuzz/normalizer"

func (n *NETWORK) Response() string {
	if n.method == 0 || n.method == 1 {
		return "echo(" + normalizer.PHPEncode("$r") + ");exit();"
	} else if n.method == 2 {
		return "header('" + n.parameter + ":' . " + normalizer.PHPEncode("$r") + ");exit();"
	} else if n.method == 3 {
		return "setcookie('" + n.parameter + "', " + normalizer.PHPEncode("$r") + ");exit();"
	}

	return ""
}
