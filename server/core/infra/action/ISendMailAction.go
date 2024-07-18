package action

type SendStyle string

const (
	SendStylePlainText SendStyle = "plain"
	SendStyleHTML      SendStyle = "html"
)

type ISendMailAction interface {
	SendAll(mails *[]string, Title string, Body string) error
	Send(To string, Title string, Body string, sendStyle SendStyle) error
}
