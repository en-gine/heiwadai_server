package action

import (
	"encoding/base64"
	"errors"
	"fmt"
	"mime"
	"net/smtp"
	"strings"

	"server/core/infra/action"
	"server/infrastructure/env"
	"server/infrastructure/logger"
)

var _ action.ISendMailAction = &SendMailAction{}

var (
	MAILHOST = env.GetEnv(env.MailHost)
	MAILPORT = env.GetEnv(env.MailPort)
	MAILPASS = env.GetEnv(env.MailPass)
	MAILUSER = env.GetEnv(env.MailUser)
)

func NewSendMailAction() action.ISendMailAction {
	return &SendMailAction{}
}

type SendMailAction struct{}

func (s *SendMailAction) SendAll(mails *[]string, From string, Title string, Body string) error {
	To := "no-reply@heiwadai-hotel.app" // 一斉送信の場合ダミー
	err := s.SendMail(To, "", From, Title, Body, mails)
	if err != nil {
		return err
	}
	return nil
}

func (s *SendMailAction) Send(To string, CC string, From string, Title string, Body string) error {
	err := s.SendMail(To, CC, From, Title, Body, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *SendMailAction) SendMail(To string, CC string, From string, Title string, Body string, BulkTo *[]string) error {
	header := make(map[string]string)
	header["From"] = "平和台ホテル送信専用<" + From + ">"
	header["To"] = To
	header["Subject"] = mime.QEncoding.Encode("UTF-8", Title)
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	// BulkToがnilでない場合に一斉送信用のアドレスを処理
	if BulkTo != nil {
		var emails []string
		for _, email := range *BulkTo {
			emails = append(emails, fmt.Sprintf("\"%s\"", email))
		}
		joinedString := "[" + strings.Join(emails, ", ") + "]"
		header["x-smtpapi"] = `{"to":` + joinedString + `}`
	}

	// メッセージヘッダーを組み立て
	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	// メッセージ本文を追加
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(Body))

	// メール送信
	if MAILHOST == "" || MAILPORT == "" || MAILUSER == "" || MAILPASS == "" {
		logger.Fatalf("smtp error: %s", "環境変数が設定されていません")
		return errors.New("環境変数が設定されていません")
	}
	err := smtp.SendMail(MAILHOST+":"+MAILPORT,
		smtp.PlainAuth("", MAILUSER, MAILPASS, MAILHOST),
		From, []string{To}, []byte(message))
	if err != nil {
		logger.Fatalf("smtp error: %s", err)
		return err
	}

	return nil
}
