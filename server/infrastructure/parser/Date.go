package parser

import (
	"time"
)

// ToStringPtr converts any type to string pointer
func ToDate(stringDate string) (time.Time, error) {
	return time.Parse("2006-01-02", stringDate)
}
