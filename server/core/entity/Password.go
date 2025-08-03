package entity

import (
	"math/rand"
	"strings"

	"server/core/errors"
	"server/infrastructure/encrypt"
	"server/infrastructure/logger"
	validator "server/validator"
)

type Password string

func NewPassword(encryptedPassword string) (*Password, *errors.DomainError) {
	v := validator.PasswordValidator{}

	pass := Password(encryptedPassword)
	decrypted, err := pass.DecriptedString()
	if err != nil {
		return nil, errors.NewDomainError(errors.InvalidParameter, err.Error())
	}
	_, err = v.IsValid(decrypted)
	if err != nil {
		return nil, errors.NewDomainError(errors.InvalidParameter, err.Error())
	}
	return &pass, nil
}

func (p *Password) string() string {
	return string(*p)
}

func (p *Password) DecriptedString() (string, error) {
	// フロントからパスワードを暗号化して投げるようにしたため、復号化した文字列を返します。
	return encrypt.Decrypt(p.string())
}

func (p *Password) EncriptedString() (*Password, error) {
	pass, err := encrypt.Encrypt(p.string())
	if err != nil {
		return nil, err
	}
	pswd, domaiErr := NewPassword(pass)
	if domaiErr != nil {
		return nil, err
	}
	return pswd, nil
}

func GenerateRandomPassword() (*Password, *errors.DomainError) {
	var password strings.Builder
	const (
		minLength = 8
		digits    = "0123456789"
		lowers    = "abcdefghijklmnopqrstuvwxyz"
		uppers    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	)

	// 大文字、小文字、数字をそれぞれ1文字以上含める
	password.WriteString(string(uppers[rand.Intn(len(uppers))]))
	password.WriteString(string(lowers[rand.Intn(len(lowers))]))
	password.WriteString(string(digits[rand.Intn(len(digits))]))

	// 残りの文字をランダムに選択
	remainingLength := minLength + rand.Intn(5) - 3
	for i := 0; i < remainingLength; i++ {
		charSet := uppers + lowers + digits
		password.WriteString(string(charSet[rand.Intn(len(charSet))]))
	}

	pass := Password(password.String())

	logger.Debugf("Generated password: %s", pass.string())
	encryptedPass, err := pass.EncriptedString()
	if err != nil {
		logger.Error(err.Error())
		return nil, errors.NewDomainError(errors.InvalidParameter, err.Error())
	}
	return encryptedPass, nil
}
