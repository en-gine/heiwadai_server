package action

type ISendMailAction interface {
	SendAll(mails *[]string, Title string, Body string) error
	Send(To string, Title string, Body string) error
}
