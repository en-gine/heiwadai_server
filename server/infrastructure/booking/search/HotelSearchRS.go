package search

import (
	"encoding/xml"
	"server/infrastructure/booking/errors"
)

type OTAHotelSearchRS struct {
	XMLName    xml.Name      `xml:"OTA_HotelSearchRS"`
	Success    string        `xml:"Success"`
	Version    string        `xml:"Version,attr"`
	Properties Properties    `xml:"Properties"`
	Errors     errors.Errors `xml:"Errors"`
}

type Properties struct {
	Property []Property `xml:"Property"`
}

type Property struct {
	Address          Address        `xml:"Address"`
	ContactNumbers   ContactNumbers `xml:"ContactNumbers"`
	HotelCode        string         `xml:"HotelCode,attr"`
	HotelName        string         `xml:"HotelName,attr"`
	HotelCodeContext string         `xml:"HotelCodeContext,attr"`
	CurrencyCode     string         `xml:"CurrencyCode,attr"`
}

type StateProv struct {
	StateCode *string `xml:"StateCode,attr"`
	Value     *string `xml:",chardata"`
}

type CountryName struct {
	Code  string `xml:"Code,attr"`
	Value string `xml:",chardata"`
}

type ContactNumbers struct {
	ContactNumber []ContactNumber `xml:"ContactNumber"`
}

type ContactNumber struct {
	PhoneTechType string `xml:"PhoneTechType,attr"`
	PhoneNumber   string `xml:"PhoneNumber,attr"`
}
