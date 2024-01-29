package main

import (
	SendMailAction "server/infrastructure/action"
	"server/infrastructure/env"
)

func main() {
	// Search()
	action := SendMailAction.NewSendMailAction()
	To := env.GetEnv(env.TestAdminMail)
	CC := ""
	From := env.GetEnv(env.MailFrom)
	Title := "タイトル"
	Body := "本文"
	err := action.Send(To, CC, From, Title, Body)
	if err != nil {
		panic(err)
	}
}
