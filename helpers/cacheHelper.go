package helpers

import (
	"crypto/sha1"
	"fmt"
)

func GenerateETag(body []byte, weak bool) string {
	hash := sha1.Sum(body)
	etag := fmt.Sprintf("\"%d-%x\"", int(len(hash)), hash)

	if weak {
		etag = "W/" + etag
	}

	return etag
}
