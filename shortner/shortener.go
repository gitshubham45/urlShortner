package shortner

import (
	"log"

	"github.com/gitshubham45/urlShortner/encoding"
)

func ShortenUrl(url string) string {
	encodedUrl := encoding.HashString(url)

	log.Println(encodedUrl)

	return encodedUrl
}
