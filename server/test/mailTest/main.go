package main

import (
	"fmt"

	"server/core/entity"
	SendMailAction "server/infrastructure/action"
	"server/infrastructure/env"
	repository "server/infrastructure/repository"
)

func main() {
	Send()
	GetMailOKUser()
}

func Send() {
	action := SendMailAction.NewSendMailAction()
	To := env.GetEnv(env.TestAdminMail)
	Title := "タイトル"
	Body := "本文"
	err := action.Send(To, Title, Body)
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
