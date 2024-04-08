package validator

import (
	"errors"
	"regexp"
)

type PasswordValidator struct {
}

func (v *PasswordValidator) IsValid(pass string) (bool, error) {
	// 正規表現による形式チェック
	// 各種文字クラスが少なくとも1つずつ含まれていることを確認する正規表現
	reLower := regexp.MustCompile("[a-z]")
	reUpper := regexp.MustCompile("[A-Z]")
	reDigit := regexp.MustCompile("[0-9]")
	// reSpecial := regexp.MustCompile("[^a-zA-Z0-9]")

	if len(pass) < 8 {
		return false, errors.New("パスワードは8文字以上である必要があります")
	}

	// 全ての文字条件を満たしているかチェック
	if reLower.MatchString(pass) &&
		reUpper.MatchString(pass) &&
		reDigit.MatchString(pass) {
		// reSpecial.MatchString(pass) {
		return true, nil
	} else {
		return false, errors.New("パスワードには小文字大文字のアルファベット及び数字を入れてください。")
	}

}
