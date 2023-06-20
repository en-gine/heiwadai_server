package validator

import (
	"errors"
	"regexp"
)

type EmailValidator struct {
}

func (v *EmailValidator) IsValid(email string) (bool, error) {
	// 正規表現による形式チェック
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	if re.MatchString(email) {
		return true, nil
	} else {
		return false, errors.New("メールアドレスの形式が正しくありません")
	}

}
