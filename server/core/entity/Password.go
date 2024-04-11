package entity

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"server/core/errors"
	validator "server/validator"
)

type Password string

func NewPassword(password string) (*Password, *errors.DomainError) {
	v := validator.PasswordValidator{}

	_, err := v.IsValid(password)
	if err != nil {
		return nil, errors.NewDomainError(errors.InvalidParameter, err.Error())
	}

	pass := Password(password)
	return &pass, nil
}

func (p *Password) String() string {
	return string(*p)
}

func GenerateRandomPassword() (*Password, *errors.DomainError) {
	const length = 8
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const minUpper = 1 // 最低限必要な大文字の数

	var upperCount int
	var result []byte
	for i := 0; i < length; i++ {
		charIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return nil, errors.NewDomainError(errors.InternalError, err.Error())
		}
		char := charset[charIndex.Int64()]
		if i >= length-minUpper-upperCount && upperCount < minUpper && char >= 'A' && char <= 'Z' {
			// 必要な大文字の数を満たしていない場合、大文字を追加
			result = append(result, char)
			upperCount++
		} else if char >= 'A' && char <= 'Z' {
			// 大文字が選択された場合
			result = append(result, char)
			upperCount++
		} else if i < length-minUpper-upperCount {
			// それ以外の場合、通常の文字を追加
			result = append(result, char)
		} else {
			// 大文字が必要な場合はループを繰り返す
			i--
		}
	}
	pass, err := NewPassword(string(result))
	fmt.Println(string(result))
	if err != nil {
		return nil, err
	}
	return pass, nil
}
