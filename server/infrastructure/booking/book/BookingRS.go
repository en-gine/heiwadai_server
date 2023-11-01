package book

import (
	"encoding/xml"
	"time"
)

type EnvelopeRS struct {
	XMLName xml.Name `xml:"S:Envelope"`
	Body    BodyRS   `xml:"S:Body"`
}

type BodyRS struct {
	EntryBookingResponse EntryBookingResponse `xml:"ns2:entryBookingResponse"`
}

type EntryBookingResponse struct {
	XMLName            xml.Name           `xml:"ns2:entryBookingResponse"`
	EntryBookingResult EntryBookingResult `xml:"entryBookingResult"`
}

type EntryBookingResult struct {
	CommonResponse CommonResponse `xml:"commonResponse"`
	ExtendLincoln  ExtendLincoln  `xml:"extendLincoln"`
}

type CommonResponse struct {
	SystemDate time.Time   `xml:"systemDate"`
	ResultCode string      `xml:"resultCode"`
	ErrorInfos *ErrorInfos `xml:"errorInfos"`
}

type ErrorInfos struct {
	ErrorCode string `xml:"errorCode"`
	ErrorMsg  string `xml:"errorMsg"`
}
