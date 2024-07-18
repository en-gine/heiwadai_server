package main

import (
	"fmt"

	"server/core/entity"
	"server/core/infra/action"
	SendMailAction "server/infrastructure/action"
	"server/infrastructure/env"
	repository "server/infrastructure/repository"
)

func main() {
	Send()
	SendHTML()
	GetMailOKUser()
}
func Send() {
	act := SendMailAction.NewSendMailAction()
	To := env.GetEnv(env.TestAdminMail)
	Title := "タイトル"
	Body := "こんにちはこれはHTMLメールのテストです。"
	err := act.Send(To, Title, Body, action.SendStylePlainText)
	if err != nil {
		panic(err)
	}
}

func SendHTML() {
	act := SendMailAction.NewSendMailAction()
	To := env.GetEnv(env.TestAdminMail)
	Title := "タイトル"
	Body := "<html><body><h1>こんにちは</h1><p>これはHTMLメールのテストです。</p></body></html>"
	err := act.Send(To, Title, Body, action.SendStyleHTML)
	if err != nil {
		panic(err)
	}
}

func GetMailOKUser() {
	query := repository.NewUserQueryService()
	prefectures := &[]entity.Prefecture{}
	count, err := query.GetMailOKUserCount(prefectures)
	if err != nil {
		panic(err)
	}
	fmt.Println("メールOKユーザー: ", *count)
}
