package book

type CancelBody struct {
	DeleteBookingWithCP DeleteBookingWithCP `xml:"naif:deleteBookingWithCP"`
}

type DeleteBookingWithCP struct {
	XMLNs                      string                     `xml:"xmlns:naif,attr"`
	DeleteBookingWithCPRequest DeleteBookingWithCPRequest `xml:"deleteBookingWithCPRequest"`
}

type DeleteBookingWithCPRequest struct {
	CommonRequest CommonRequest `xml:"commonRequest"`
	BookingInfo   BookingInfo   `xml:"bookingInfo"`
}

type BookingInfo struct {
	TllHotelCode       string `xml:"tllHotelCode"`
	TllBookingNumber   string `xml:"tllBookingNumber"`
	DataID             string `xml:"DataID"`
	CancellationCharge int    `xml:"CancellationCharge"`
	CancellationNotice string `xml:"CancellationNotice"`
}
