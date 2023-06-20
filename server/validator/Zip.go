package validator

import (
	"errors"
	"regexp"
)

type ZipValidator struct {
}

func (v *ZipValidator) IsValid(zip string) (bool, error) {
	re := regexp.MustCompile(`^[0-9]{3}-[0-9]{4}$`)

	if re.MatchString(zip) {
		return true, nil
	} else {
		return false, errors.New("郵便番号の形式が正しくありません")
	}

}
