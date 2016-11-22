package core

type REQUEST struct {
	Type   string
	Action string
	PHPc   *PHPCONFIG
	SHLc   *SHELLCONFIG
	SRVc   *SERVERCONFIG
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
