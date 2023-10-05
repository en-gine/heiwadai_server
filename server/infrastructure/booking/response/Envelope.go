package response

import (
	"encoding/xml"
)

type Envelope[T any] struct {
	XMLName xml.Name `xml:"Envelope"`
	S       string   `xml:"xmlns:S,attr"`
	SoapEnv string   `xml:"xmlns:SOAP-ENV,attr"`
	Body    Body[T]  `xml:"S:Body"`
}

type Body[T any] struct {
	XMLName xml.Name `xml:"S:Body"`
	Content T        `xml:",innerxml"`
}

func NewEnvelope[T any](XMLBody T) Envelope[T] {
	return Envelope[T]{
		S:       "http://schemas.xmlsoap.org/soap/envelope/",
		SoapEnv: "http://schemas.xmlsoap.org/soap/envelope/",
		Body: Body[T]{
			Content: XMLBody,
		},
	}
}
