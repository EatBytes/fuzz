package fuzzcore

import "encoding/base64"

func Encode(str string) string {
	sEnc := base64.StdEncoding.EncodeToString([]byte(str))

	return sEnc
}

func Decode(str string) (string, error) {
	sDec, err := base64.StdEncoding.DecodeString(str)

	if err != nil {
		ferr := NormalizeErr(err)
		return "", ferr
	}

	return string(sDec), nil
}

func PHPEncode(str string) string {
	return "base64_encode(" + str + ")"
}

func PHPDecode(str string) string {
	return "base64_decode(" + str + ")"
}
