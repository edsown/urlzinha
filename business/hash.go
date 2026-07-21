package business

import "strings"

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func StripOriginalUrl(url string) (strippedUrl string) {
	return strings.TrimPrefix(url, "http")
} // will this be used?
// and should it be in the hash.go file?

func Encode(id uint64) string {
	if id == 0 {
		return "0"
	}

	chars := []byte{}
	for id > 0 {
		remainder := id % 62
		chars = append(chars, alphabet[remainder])
		id = id / 62
	}
	// reversing the slice
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}

	return string(chars)
}

func Decode(str string) uint64 {
	id := 0
	for _, char := range str {
		id = id*62 + strings.IndexRune(alphabet, char)
	}
	return uint64(id)
}

func saveShortUrl() error {
	return nil
}
