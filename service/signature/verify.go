package signature

import (
	"crypto/hmac"
)

func Verify(signature string, params ...string) bool {
	return hmac.Equal(
		[]byte(signature),
		[]byte(Create(params...)),
	)
}
