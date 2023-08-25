package search

import (
	"strings"
)

type OTAHotelSearchRQ struct {
	Version       string `xml:"Version,attr"`
	PrimaryLangID string `xml:"PrimaryLangID,attr,omitempty"`
	ResponseType  string `xml:"ResponseType,attr"`
	Criteria      struct {
		Criterion Criterion `xml:"Criterion"`
	} `xml:"Criteria"`
}

func NewOTAHotelSearchRQ() OTAHotelSearchRQ {
	return OTAHotelSearchRQ{
		Version:       "1.0",
		PrimaryLangID: "jpn",
		ResponseType:  "PropertyList",
		Criteria:      Criteria{},
	}
}

type Criteria struct {
	Criterion Criterion `xml:"Criterion"`
}

type Criterion struct {
	Address   *Address   `xml:"Address,omitempty"`
	Telephone *Telephone `xml:"Telephone,omitempty"`
	HotelRef  *HotelRef  `xml:"HotelRef,omitempty"`
}

type Address struct {
	AddressLine *string      `xml:"AddressLine"`
	StateProv   *StateProv   `xml:"StateProv"`
	CountryName *CountryName `xml:"CountryName"`
}

type Telephone struct {
	PhoneTechType string `xml:"PhoneTechType,attr,omitempty"`
	PhoneNumber   string `xml:"PhoneNumber,attr,omitempty"`
}

func NewTelephone(phoneTechType string, phoneNumber string) Telephone {
	phoneNumber = strings.ReplaceAll(phoneNumber, "-", "")
	return Telephone{
		PhoneTechType: "1", //電話,
		PhoneNumber:   phoneNumber,
	}
}

type HotelRef struct {
	HotelCode        string `xml:"HotelCode,attr,omitempty"`
	HotelName        string `xml:"HotelName,attr,omitempty"`
	HotelCodeContext string `xml:"HotelCodeContext,attr,omitempty"`
}
