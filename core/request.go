package core

import "github.com/eatbytes/razboy/normalizer"

type REQUEST struct {
	Type  string
	Raw   string
	PHPc  PHPCONFIG
	SHLc  SHELLCONFIG
	SRVc  SERVERCONFIG
	setup bool
}

func (r REQUEST) IsSetup() bool {
	return r.setup
}

func (r REQUEST) IsPHP() bool {
	if r.Type == "PHP" {
		return true
	}

	return false
}

func (r REQUEST) IsSHELL() bool {
	if r.Type == "SHELL" {
		return true
	}

	return false
}

func (r *REQUEST) Build(w ...bool) {
	if r.Type == "SHELL" {
		r.Raw = r.SHLc.Cmd
	} else if r.Type == "PHP" {
		r.Raw = r.PHPc.Cmd
	}

	if len(w) > 0 && !w[0] {
		return
	}

	r.AddResponseLogic()
}

func (r *REQUEST) AddResponseLogic() {
	if r.SRVc.Method == "HEADER" {
		r.Raw = r.Raw + "header('" + r.SRVc.Parameter + ":' . " + normalizer.PHPEncode("$r") + ");exit();"
		return
	}

	if r.SRVc.Method == "COOKIE" {
		r.Raw = r.Raw + "setcookie('" + r.SRVc.Parameter + "', " + normalizer.PHPEncode("$r") + ");exit();"
		return
	}

	r.Raw = r.Raw + "echo(" + normalizer.PHPEncode("$r") + ");exit();"
}
