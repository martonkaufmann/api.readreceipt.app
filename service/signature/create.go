package signature

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"github.com/readreceipt/api/config"
)

func Create(params ...string) string {
	h := hmac.New(sha256.New, []byte(config.Secret()))
	h.Write([]byte(strings.Join(params, ".")))

	return hex.EncodeToString(h.Sum(nil))
}
