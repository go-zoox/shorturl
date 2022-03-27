package shorturl

import "github.com/go-zoox/crypto/hash"

func ShortURL(longURL string) string {
	return hash.MurmurHash(longURL)
}
