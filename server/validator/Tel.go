package validator

import (
	"errors"
	"regexp"
)

type TelValidator struct {
}

func (v *TelValidator) IsValid(tel string) (bool, error) {
	// 半角ハイフンを含んだ1〜4桁・1〜4桁・3〜4桁の半角数字の形式で入力されているか
	re := regexp.MustCompile(`^0\d{1,4}-\d{1,4}-\d{3,4}$`)

	if re.MatchString(tel) {
		return true, nil
	} else {
		return false, errors.New("電話番号の形式が正しくありません。数字とハイフンの組み合わせである必要があります。")
	}

}
