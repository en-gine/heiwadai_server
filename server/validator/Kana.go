package validator

import (
	"errors"
	"regexp"
)

type KanaValidator struct {
}

func (v *KanaValidator) IsValid(kana string) (bool, error) {
	// 正規表現による形式チェック
	re := regexp.MustCompile(`^[ァ-ヾ]+$`)

	if re.MatchString(kana) {
		return true, nil
	} else {
		return false, errors.New("カナの形式が正しくありません")
	}

}
