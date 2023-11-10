package book

import (
	"encoding/xml"
	"time"
)

type EnvelopeRS[TBody any] struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    TBody    `xml:"Body"`
}

type BodyRS struct {
	EntryBookingResponse EntryBookingResponse `xml:"entryBookingResponse"`
}

type EntryBookingResponse struct {
	XMLName            xml.Name           `xml:"entryBookingResponse"`
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
