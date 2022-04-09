package utils

import "strconv"

const (
	prefix = "skm"
	quotes = "quotes"
)

func QuotesKey() string {
	return prefix + "_" + quotes
}

func GetKey(id int64) string {
	return strconv.FormatInt(id, 10)
}
