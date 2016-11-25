package razboy

import "github.com/eatbytes/razboy/core"

type REQUEST struct {
	Type   string
	Action string
	PHPc   *core.PHPCONFIG
	SHLc   *core.SHELLCONFIG
	SRVc   *core.SERVERCONFIG
	setup  bool
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
