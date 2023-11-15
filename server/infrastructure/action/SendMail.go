package action

import (
	"encoding/base64"
	"fmt"
	"mime"
	"net/smtp"
	"strings"

	"server/core/infra/action"
	"server/infrastructure/env"
	"server/infrastructure/logger"
)

var _ action.ISendMailAction = &SendMail{}

var (
	HOST = env.GetEnv(env.MailHost)
	PORT = env.GetEnv(env.MailPort)
	PASS = env.GetEnv(env.MailPass)
)

type SendMail struct{}

func (s *SendMail) SendAll(mails *[]string, From string, Title string, Body string) error {
	To := "no-reply@heiwadai-hotel.app" // 一斉送信の場合ダミー
	err := s.SendMail(To, "", From, Title, Body, mails)
	if err != nil {
		return err
	}
	return nil
}

func (s *SendMail) Send(To string, CC string, From string, Title string, Body string) error {
	err := s.SendMail(To, CC, From, Title, Body, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *SendMail) SendMail(To string, CC string, From string, Title string, Body string, BulkTo *[]string) error {
	header := make(map[string]string)
	header["From"] = From
	header["To"] = To
	header["Subject"] = mime.QEncoding.Encode("UTF-8", Title)
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"
	if BulkTo != nil {
		var emails []string
		for i, email := range *BulkTo {
			emails[i] = fmt.Sprintf("\"%s\"", email)
		}
		joinedString := "[" + strings.Join(emails, ", ") + "]"
		header["x-smtpapi"] = `{"to":` + joinedString + `}` // 一斉送信の場合
	}
	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(Body))

	err := smtp.SendMail(HOST+":"+PORT,
		smtp.PlainAuth("", From, PASS, HOST),
		From, []string{To}, []byte(message))
	if err != nil {
		logger.Fatalf("smtp error: %s", err)
		return err
	}

	return nil
}

func main() {
	BulkTo := []string{"test@test.jp"}
	BulkTo = append(BulkTo, "aaaa@bbbb.co.jp")
	fmt.Printf("%v", BulkTo)
}
