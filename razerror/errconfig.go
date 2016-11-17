package razerror

import "fmt"

const (
	InvalidTypeError = "Invalid Type"
	NilConfigError   = "Config nil"
)

type RazErrorConfig struct {
	Type string
	code int
	Msg  string
	bag  error
}

func (e RazErrorConfig) Error() string {
	str := "%v \n\n" +
		"     bag: %v \n"

	return fmt.Sprintf(str)
}

func ConfigErrInvalidType() RazErrorConfig {
	return RazErrorConfig{
		Type: InvalidTypeError,
		Msg:  "",
	}
}

func ConfigErrNilConfig() RazErrorConfig {
	return RazErrorConfig{
		Type: NilConfigError,
		Msg:  "",
	}
}
