package razboy

import (
	"bytes"
)

type REQUEST struct {
	Type  string
	PHPc  PHPCONFIG
	SHLc  SHELLCONFIG
	SRVc  SERVERCONFIG
	setup bool
}

type PHPCONFIG struct {
	upload  bool
	buffer  *bytes.Buffer
	bondary string
}

type SHELLCONFIG struct {
	method  string
	context string
}

type SERVERCONFIG struct {
	url       string
	method    string
	parameter string
	key       string
	raw       bool
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

func (php PHPCONFIG) IsUpload() bool {
	return php.upload
}

func (srv SERVERCONFIG) IsRaw() bool {
	return srv.raw
}

func (srv SERVERCONFIG) GetMethod() string {
	return srv.method
}

func (srv SERVERCONFIG) GetUrl() string {
	return srv.url
}

func (srv SERVERCONFIG) GetParameter() string {
	return srv.parameter
}

func (srv SERVERCONFIG) GetKey() string {
	return srv.key
}

func (srv SERVERCONFIG) IsProtected() bool {
	if srv.key != "" {
		return true
	}

	return false
}
