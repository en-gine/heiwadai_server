package common

import (
	"encoding/xml"
)

type EnvelopeRQ[T any] struct {
	XMLName xml.Name  `xml:"soapenv:Envelope"`
	SoapEnv string    `xml:"xmlns:soapenv,attr"`
	Head    string    `xml:"xmlns:head,attr"`
	Ns      string    `xml:"xmlns:ns,attr"`
	Header  Header    `xml:"soapenv:Header"`
	Body    BodyRQ[T] `xml:"soapenv:Body"`
}

type BodyRQ[T any] struct {
	XMLName xml.Name `xml:"soapenv:Body"`
	Content T        `xml:",innerxml"`
}

type Header struct {
	Interface Interface `xml:"head:Interface"`
}

type Interface struct {
	PayloadInfo PayloadInfo `xml:"head:PayloadInfo"`
}

type PayloadInfo struct {
	CommDescriptor CommDescriptor `xml:"head:CommDescriptor"`
}

type CommDescriptor struct {
	Authentication Authentication `xml:"head:Authentication"`
}

type Authentication struct {
	Username string `xml:"head:Username"`
	Password string `xml:"head:Password"`
}

func NewEnvelopeRQ[T any](XMLBody T, UserName string, Password string) EnvelopeRQ[T] {
	return EnvelopeRQ[T]{
		SoapEnv: "http://schemas.xmlsoap.org/soap/envelope/",
		Head:    "http://www.seanuts.co.jp/ota/header",
		Ns:      "http://www.opentravel.org/OTA/2003/05",
		Header: Header{
			Interface: Interface{
				PayloadInfo: PayloadInfo{
					CommDescriptor: CommDescriptor{
						Authentication: Authentication{
							Username: UserName,
							Password: Password,
						},
					},
				},
			},
		},
		Body: BodyRQ[T]{
			Content: XMLBody,
		},
	}
}
