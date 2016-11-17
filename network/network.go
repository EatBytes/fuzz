package network

import (
	"strings"

	"github.com/eatbytes/razboy/core"
	"github.com/eatbytes/razboy/ferror"
)

func Create(config *core.Config) (*NETWORK, error) {
	var n *NETWORK

	config.Url = strings.TrimSpace(config.Url)
	config.Method = strings.TrimSpace(strings.ToUpper(config.Method))
	config.Parameter = strings.TrimSpace(config.Parameter)
	config.Key = strings.TrimSpace(config.Key)

	if config.Url == "" {
		return nil, ferror.Default("The url should be specified")
	}

	if !strings.Contains(config.Url, "http://") && !strings.Contains(config.Url, "https://") {
		config.Url = "http://" + config.Url
	}

	if config.Method == "" {
		config.Method = GET
	}

	if config.Method != GET && config.Method != POST && config.Method != HEADER && config.Method != COOKIE {
		return nil, ferror.Default("The method (" + config.Method + ") is not a valid one. Please choose between: GET, POST, HEADER or COOKIE.")
	}

	if config.Parameter == "" {
		config.Parameter = PARAM
	}

	config.Crypt = false

	n = &NETWORK{
		config: config,
		status: true,
	}

	return n, nil
}
