package business

import "strings"
import "crypto/sha256"
import "enconding"
import "fmt"

func stripOriginalUrl(url string) (strippedUrl string) {
	return strings.TrimPrefix(url, "http")
}

func hashBase62(url string) (hashUrl string, error error) {

	md5.Md5Base62()
	return

}
