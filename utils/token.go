package utils

import "encoding/base64"

func Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func Decode(token string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(token)
}
