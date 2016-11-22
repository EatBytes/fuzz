package checker

import (
	"errors"

	"github.com/eatbytes/razboy/core"
)

func Check(req *core.REQUEST) error {
	var err error

	err = nil

	if req.Type == "PHP" {
		err = _checkPHP(req.PHPc)
	} else if req.Type == "SHELL" {
		err = _checkSHELL(req.SHLc)
	} else {
		err = errors.New("The request should have a valid type ('PHP' or 'SHELL')")
	}

	if err == nil {
		err = _checkSERVER(req.SRVc)
	}

	return err
}

func _checkPHP(c *core.PHPCONFIG) error {
	if c == nil {
		return errors.New("Empty pointer")
	}

	return nil
}

func _checkSHELL(c *core.SHELLCONFIG) error {
	if c == nil {
		return errors.New("Empty pointer")
	}

	return nil
}

func _checkSERVER(c *core.SERVERCONFIG) error {
	if c == nil {
		return errors.New("Empty pointer")
	}

	if c.Url == "" {
		return errors.New("REQUEST -> SERVERCONFIG [url] should not be empty")
	}

	if c.Method != "GET" && c.Method != "POST" && c.Method != "HEADER" && c.Method != "COOKIE" {
		c.Method = "GET"
	}

	return nil
}
