package helper

import "encoding/base64"

func Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func Decode(s string) (string, error) {
	output, err := base64.StdEncoding.DecodeString(s)
	return string(output), err
}
