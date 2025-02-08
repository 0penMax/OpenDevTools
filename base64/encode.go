package base64

import "encoding/base64"

func Encode(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}
