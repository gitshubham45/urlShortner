package encoding

import (
	"crypto/sha256"
	"math/big"
	"strings"
)

// Base62 characters
const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func HashString(url string) string {
	hasher := sha256.New()
	hasher.Write([]byte(url))
	hashBytes := hasher.Sum(nil)

	hashInt := new(big.Int).SetBytes(hashBytes)

	var base62Str strings.Builder
	base := big.NewInt(62)
	zero := big.NewInt(0)
	remainder := new(big.Int)

	for hashInt.Cmp(zero) > 0 {
		hashInt.DivMod(hashInt, base, remainder)
		base62Str.WriteByte(base62Chars[remainder.Int64()])
	}

	result := base62Str.String()
	if len(result) < 7 {
		result = strings.Repeat("0", 7-len(result)) + result
	} else if len(result) > 7 {
		result = result[:7] // Trim to 7 characters
	}

	return result
}
