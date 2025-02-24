package shortner

import (
	"github.com/gitshubham45/urlShortner/encoding"
)

func ShortenUrl(url string) string {
	encodedUrl := encoding.EncodeBase62(url)

	return encodedUrl

}
