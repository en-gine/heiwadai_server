package action

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"mime"
	"mime/quotedprintable"
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
	MAILFROM = env.GetEnv(env.MailFrom)
)

func NewSendMailAction() action.ISendMailAction {
	return &SendMailAction{}
}

type SendMailAction struct{}

func (s *SendMailAction) SendAll(mails *[]string, Title string, Body string) error {
	To := "no-reply@heiwadai-hotel.app" // 一斉送信の場合ダミー
	err := s.SendMail(To, Title, Body, mails, action.SendStylePlainText)
	if err != nil {
		return err
	}
	return nil
}

func (s *SendMailAction) Send(To string, Title string, Body string, sendStyle action.SendStyle) error {
	err := s.SendMail(To, Title, Body, nil, sendStyle)
	if err != nil {
		return err
	}
	return nil
}

func (s *SendMailAction) SendMail(To string, Title string, Body string, BulkTo *[]string, sendStyle action.SendStyle) error {
	var encodedBody string
	header := make(map[string]string)
	header["From"] = "平和台ホテル送信専用<" + MAILFROM + ">"
	header["To"] = To
	header["Subject"] = mime.QEncoding.Encode("UTF-8", Title)
	header["MIME-Version"] = "1.0"
	switch sendStyle {
	case action.SendStyleHTML:
		header["Content-Type"] = "text/html; charset=\"utf-8\""
		header["Content-Transfer-Encoding"] = "quoted-printable"
		encodedBody = quoteString(Body)
	case action.SendStylePlainText:
		fallthrough
	default:
		header["Content-Type"] = "text/plain; charset=\"utf-8\""
		header["Content-Transfer-Encoding"] = "base64"
		encodedBody = base64.StdEncoding.EncodeToString([]byte(Body))
	}

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
	message += "\r\n" + encodedBody

	// メール送信
	if MAILHOST == "" || MAILPORT == "" || MAILUSER == "" || MAILPASS == "" {
		logger.Errorf("smtp error: %s", "環境変数が設定されていません")
		return errors.New("環境変数が設定されていません")
	}
	err := smtp.SendMail(MAILHOST+":"+MAILPORT,
		smtp.PlainAuth("", MAILUSER, MAILPASS, MAILHOST),
		MAILFROM, []string{To}, []byte(message))
	if err != nil {
		logger.Errorf("smtp error: %s", err)
		return err
	}

	return nil
}

func quoteString(s string) string {
	var buf bytes.Buffer
	w := quotedprintable.NewWriter(&buf)
	_, err := w.Write([]byte(s))
	if err != nil {
		logger.Errorf("quotedprintable error: %s", err)
	}
	w.Close()
	return buf.String()
}
