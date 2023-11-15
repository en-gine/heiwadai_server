package action

type ISendMailAction interface {
	SendAll(mails *[]string, From string, Title string, Body string) error
	Send(To string, CC string, From string, Title string, Body string) error
}
