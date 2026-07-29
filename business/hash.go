package business

import "strings"

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func StripOriginalUrl(url string) (strippedUrl string) {
	return strings.TrimPrefix(url, "http")
}

func Encode(id int64) string {
	if id == 0 {
		return "0"
	}

	chars := []byte{}
	for id > 0 {
		remainder := id % 62
		chars = append(chars, alphabet[remainder])
		id = id / 62
	}
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}

	return string(chars)
}

func Decode(str string) int64 {
	id := 0
	for _, char := range str {
		id = id*62 + strings.IndexRune(alphabet, char)
	}
	return int64(id)
}
