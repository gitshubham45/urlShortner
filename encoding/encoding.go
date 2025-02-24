package encoding

import (
	"math/big"
	"strings"
)

// Base62 characters
const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const base = 62

// EncodeBase62 converts a string to a Base62 encoded string
func EncodeBase62(input string) string {
	// Convert string to a big integer representation
	num := new(big.Int)
	num.SetBytes([]byte(input))

	var encoded strings.Builder
	zero := big.NewInt(0)
	baseBig := big.NewInt(base)

	for num.Cmp(zero) > 0 {
		remainder := new(big.Int)
		num.DivMod(num, baseBig, remainder)
		encoded.WriteByte(base62Chars[remainder.Int64()])
	}

	return reverseString(encoded.String())
}

// DecodeBase62 converts a Base62 encoded string back to the original string
func DecodeBase62(encoded string) string {
	num := big.NewInt(0)
	baseBig := big.NewInt(base)

	for _, char := range encoded {
		num.Mul(num, baseBig)
		num.Add(num, big.NewInt(int64(strings.IndexRune(base62Chars, char))))
	}

	return string(num.Bytes())
}

// Helper function to reverse a string
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
