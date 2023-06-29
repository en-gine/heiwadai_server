package parser

import "fmt"

// ToStringPtr converts any type to string pointer
func ToStringPtr(s interface{}) *string {
	if s == nil {
		return nil
	}
	strVal := fmt.Sprintf("%v", s)
	return &strVal
}
