package book

import (
	"encoding/xml"
)

type CancelBodyRS struct {
	DeleteBookingWithCPResponse DeleteBookingWithCPResponse `xml:"ns2:deleteBookingWithCPResponse"`
}

type DeleteBookingWithCPResponse struct {
	XMLName                   xml.Name                  `xml:"ns2:deleteBookingWithCPResponse"`
	XMLNs                     string                    `xml:"xmlns:ns2,attr"`
	DeleteBookingWithCPResult DeleteBookingWithCPResult `xml:"deleteBookingWithCPResult"`
}

type DeleteBookingWithCPResult struct {
	CommonResponse CommonResponse `xml:"commonResponse"`
	BookingInfo    BookingInfo    `xml:"bookingInfo"`
}
