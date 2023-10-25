package action

type ISendMailAction interface {
	SendAll(mails *[]string) error
	Send(To string, CC string, From string, Title string, Body string) error
}
