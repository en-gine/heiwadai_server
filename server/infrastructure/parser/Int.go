package parser

import (
	"strconv"
)

func ToIntPtr(s string) *int {
	intValue, _ := strconv.Atoi(s)
	return &intValue
}
