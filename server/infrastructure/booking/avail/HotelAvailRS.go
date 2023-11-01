package avail

import (
	"encoding/xml"

	"server/infrastructure/booking/util"
)

type EnvelopeRS struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	OTA_HotelAvailRS OTA_HotelAvailRS `xml:"OTA_HotelAvailRS"`
}

type OTA_HotelAvailRS struct {
	Success    Success    `xml:"Success"`
	HotelStays HotelStays `xml:"HotelStays"`
	RoomStays  RoomStays  `xml:"RoomStays"`
	Criteria   Criteria   `xml:"Criteria"`
	Version    string     `xml:"Version,attr"`
	Errors     *Errors    `xml:"Errors"`
}

type Success struct{}

type HotelStays struct {
	HotelStay []HotelStay `xml:"HotelStay"`
}

type HotelStay struct {
	Availability      Availability      `xml:"Availability"`
	BasicPropertyInfo BasicPropertyInfo `xml:"BasicPropertyInfo"`
	Price             Price             `xml:"Price"`
	RoomStayRPH       string            `xml:"RoomStayRPH,attr"`
}

type Availability struct {
	End    util.YYYYMMDD `xml:"End,attr"`
	Start  util.YYYYMMDD `xml:"Start,attr"`
	Status Status        `xml:"Status,attr"`
}
type Status string

const (
	StatusOpen      Status = "Open"
	StatusClose     Status = "Close"
	StatusOnRequest Status = "OnRequest"
)

type BasicPropertyInfo struct {
	HotelCode string `xml:"HotelCode,attr"`
	HotelName string `xml:"HotelName,attr"`
}

type Price struct {
	AmountAfterTax int           `xml:"AmountAfterTax,attr"`
	CurrencyCode   string        `xml:"CurrencyCode,attr"`
	End            util.YYYYMMDD `xml:"End,attr"`
	Start          util.YYYYMMDD `xml:"Start,attr"`
}

type RoomStays struct {
	RoomStay []RoomStay `xml:"RoomStay"`
}

type RoomStay struct {
	RPH       string    `xml:"RPH,attr"`
	RoomTypes RoomTypes `xml:"RoomTypes"`
	RatePlans RatePlans `xml:"RatePlans"`
	RoomRates RoomRates `xml:"RoomRates"`
}

type RoomTypes struct {
	RoomType []RoomType `xml:"RoomType"`
}

type RoomType struct {
	BedTypeCode     BedTypeCode     `xml:"BedTypeCode,attr"`
	NonSmoking      bool            `xml:"NonSmoking,attr"`
	RoomTypeCode    string          `xml:"RoomTypeCode,attr"`
	RoomDescription RoomDescription `xml:"RoomDescription"`
	Amenities       Amenities       `xml:"Amenities"`
	Occupancy       Occupancy       `xml:"Occupancy"`
}

type RoomDescription struct {
	CreateDateTime     string `xml:"CreateDateTime,attr"`
	LastModifyDateTime string `xml:"LastModifyDateTime,attr"`
	Name               string `xml:"Name,attr"`
	Text               Text   `xml:"Text"`
	Image              Image  `xml:"Image"`
	URL                URL    `xml:"URL"`
}

type Text struct {
	Value string `xml:",chardata"`
}

type Image struct {
	Value string `xml:",chardata"`
}

type URL struct {
	Value string `xml:",chardata"`
}

type Amenities struct {
	Amenity []Amenity `xml:"Amenity"`
}

type Amenity struct {
	ExistsCode  ExistsCode  `xml:"ExistsCode,attr"`
	RoomAmenity RoomAmenity `xml:"RoomAmenity,attr"`
}

type RoomAmenity string

const (
	RoomAmenityPrivatebathroom    RoomAmenity = "85"
	RoomAmenityInternetInRoom     RoomAmenity = "207"
	RoomAmenitySeparateToiletArea RoomAmenity = "271" //（独立したトイレエリア）
)

type ExistsCode string

const (
	ExistsCodeYes ExistsCode = "Yes"
	ExistsCodeNo  ExistsCode = "No"
)

type Occupancy struct {
	MaxOccupancy string `xml:"MaxOccupancy,attr"`
	MinOccupancy string `xml:"MinOccupancy,attr"`
}

type RatePlans struct {
	RatePlan []RatePlan `xml:"RatePlan"`
}

type RatePlan struct {
	EffectiveDate       util.YYYYMMDD       `xml:"EffectiveDate,attr"`
	ExpireDate          util.YYYYMMDD       `xml:"ExpireDate,attr"`
	RatePlanCode        string              `xml:"RatePlanCode,attr"`
	RatePlanName        string              `xml:"RatePlanName,attr"`
	RatePlanDescription RatePlanDescription `xml:"RatePlanDescription"`
	RatePlanInclusions  RatePlanInclusions  `xml:"RatePlanInclusions"`
	MealsIncluded       MealsIncluded       `xml:"MealsIncluded"`
	AdditionalDetails   AdditionalDetails   `xml:"AdditionalDetails"`
}

type CancelPenalties struct {
	CancelPenalty []CancelPenalty `xml:"CancelPenalty"`
}

type CancelPenalty struct {
	ConfirmClassCode   *string              `xml:"ConfirmClassCode,attr"`
	AmountPercent      *AmountPercent       `xml:"AmountPercent"`
	Deadline           *Deadline            `xml:"Deadline"`
	PenaltyDescription []PenaltyDescription `xml:"PenaltyDescription"`
}

type AmountPercent struct {
	Percent      *string `xml:"Percent,attr"`
	Amount       *string `xml:"Amount,attr"`
	CurrencyCode *string `xml:"CurrencyCode,attr"`
}

type Deadline struct {
	OffsetDropTime       *string `xml:"OffsetDropTime,attr"`
	OffsetTimeUnit       *string `xml:"OffsetTimeUnit,attr"`
	OffsetUnitMultiplier *string `xml:"OffsetUnitMultiplier,attr"`
}

type PenaltyDescription struct {
	Text Text `xml:"Text"`
}

type RatePlanDescription struct {
	CreateDateTime     string `xml:"CreateDateTime,attr"`
	LastModifyDateTime string `xml:"LastModifyDateTime,attr"`
	Text               Text   `xml:"Text"`
	Image              Image  `xml:"Image"`
	URL                URL    `xml:"URL"`
}

type RatePlanInclusions struct {
	ServiceFeeInclusive bool `xml:"ServiceFeeInclusive,attr"`
	TaxInclusive        bool `xml:"TaxInclusive,attr"`
}

type AdditionalDetails struct {
	AdditionalDetail []AdditionalDetail `xml:"AdditionalDetail"`
}

type AdditionalDetail struct {
	Type              string            `xml:"Type,attr"`
	DetailDescription DetailDescription `xml:"DetailDescription"`
}

type DetailDescription struct {
	Text Text `xml:"Text"`
}

type RoomRates struct {
	RoomRate []RoomRate `xml:"RoomRate"`
}

type RoomRate struct {
	AvailabilityStatus string      `xml:"AvailabilityStatus,attr"`
	NumberOfUnits      string      `xml:"NumberOfUnits,attr"`
	RatePlanCode       string      `xml:"RatePlanCode,attr"`
	RoomTypeCode       string      `xml:"RoomTypeCode,attr"`
	Rates              Rates       `xml:"Rates"`
	Total              Total       `xml:"Total"`
	GuestCounts        GuestCounts `xml:"GuestCounts"`
}
type AvailabilityStatus string

const (
	AvailableForSale   AvailabilityStatus = "AvailableForSale" // 販売中
	AvailableOnRequest AvailabilityStatus = "OnRequest"        // 残りわずか
	AvailableClosedOut AvailabilityStatus = "ClosedOut"        // 売り切れ
)

type Rates struct {
	Rate Rate `xml:"Rate"`
}

type Rate struct {
	AgeQualifyingCode string `xml:"AgeQualifyingCode,attr"`
	RoomPricingType   string `xml:"RoomPricingType,attr"` // "Per night"：通常プラン
	Base              Base   `xml:"Base"`
}

type Base struct {
	AmountAfterTax string `xml:"AmountAfterTax,attr"`
	CurrencyCode   string `xml:"CurrencyCode,attr"`
	Type           string `xml:"Type,attr"`
}

type Total struct {
	AmountAfterTax string `xml:"AmountAfterTax,attr"`
	CurrencyCode   string `xml:"CurrencyCode,attr"`
}

type Criteria struct {
	Criterion Criterion `xml:"Criterion"`
}

type Errors struct {
	Error []Error `xml:"Error"`
}

type Error struct {
	Type      string `xml:"Type,attr"`
	ShortText string `xml:"ShortText,attr"`
	Code      string `xml:"Code,attr"`
}
