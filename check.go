package razboy

import "github.com/eatbytes/razboy/razerror"

func _check(req *REQUEST) error {
	var err error

	err = nil

	if req.Type == "PHP" {
		err = _checkPHP(&req.PHPc)
	} else if req.Type == "SHELL" {
		err = _checkSHELL(&req.SHLc)
	} else {
		err = razerror.ConfigErrInvalidType()
	}

	if err == nil {
		err = _checkSERVER(&req.SRVc)
	}

	return err
}

func _checkPHP(c *PHPCONFIG) error {
	if c == nil {
		return razerror.ConfigErrNilConfig()
	}

	return nil
}

func _checkSHELL(c *SHELLCONFIG) error {
	if c == nil {
		return razerror.ConfigErrNilConfig()
	}

	return nil
}

func _checkSERVER(c *SERVERCONFIG) error {
	if c == nil {
		return razerror.ConfigErrNilConfig()
	}

	return nil
}
