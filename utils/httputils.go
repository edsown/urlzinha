package utils

import (
	"net/url"
	"strconv"
	"strings"
)

func GetValueFromQueryInt64(values url.Values, key string) *int64 {
	value := values[key]
	if len(value) > 0 && value[0] != "" {
		value, _ := strconv.ParseInt(value[0], 10, 64)
		return &value
	}
	return nil
}

func GetValueFromQueryInt64Slice(values url.Values, key string) []int64 {
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

func GetValueFromQueryStr(values url.Values, key string) *string {
	value := values[key]
	if len(value) > 0 && value[0] != "" {
		return &value[0]
	}
	return nil
}
