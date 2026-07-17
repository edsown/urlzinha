package utils

import (
	"net/url"
	"strconv"
	"strings"
)

func ExtractIntegerParamsFromRequest(values url.Values, key string) []int64 {
	param := values[key]
	var result []int64
	if param != nil && param[0] != "" {
		for _, v := range strings.Split(param[0], ",") {
			v, _ := strconv.ParseInt(v, 10, 64)
			result = append(result, v)
		}
	}
	return result
}
