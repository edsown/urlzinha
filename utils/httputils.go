package utils

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func ExtractIntegerParamsFromRequest(values url.Values, key string) []int64 {
	param := values[key]
	if param != nil && param[0] != "" {
		var result []int64
		for _, v := range strings.Split(param, ",") {
			v, _ := strconv.ParseInt(v, 10, 64)
			result = append(result, v)
		}
	}
}
