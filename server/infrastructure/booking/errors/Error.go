package errors

type Errors struct {
	Error []Error `xml:"Error"`
}

type Error struct {
	Type      string `xml:"Type,attr"`
	ShortText string `xml:"ShortText,attr"`
	Code      string `xml:"Code,attr"`
}
