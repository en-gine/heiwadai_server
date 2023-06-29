package action

type Mail struct {
	From    string
	To      string
	Subject string
	Body    string
}

type IMailAction interface {
	Send(mail *Mail) error
	SendAll(mails *[]Mail) error
}
