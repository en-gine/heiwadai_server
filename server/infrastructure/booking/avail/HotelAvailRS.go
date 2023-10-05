package avail

import (
	"encoding/xml"
	"server/infrastructure/booking/util"
)

type OTAHotelAvailRS struct {
	XMLName    xml.Name      `xml:"ns:OTA_HotelAvailRS"`
	Version    *string       `xml:"Version,attr"`
	Success    *string       `xml:"Success"`
	HotelStays *[]HotelStays `xml:"HotelStays"`
	RoomStays  *RoomStays    `xml:"RoomStays"`
	Services   *Services     `xml:"Services"`
	Criteria   *Criteria     `xml:"Criteria"`
	Errors     *Errors       `xml:"Errors"`
}

type Success struct {
}

type HotelStays struct {
	HotelStay []HotelStay `xml:"HotelStay"`
}

type HotelStay struct {
	RoomStayRPH       *string              `xml:"RoomStayRPH,attr"`
	Availability      *[]Availability      `xml:"Availability"`
	BasicPropertyInfo *[]BasicPropertyInfo `xml:"BasicPropertyInfo"`
	Price             *[]Price             `xml:"Price"`
}

type Availability struct {
	Status *Status        `xml:"Status,attr"`
	Start  *util.YYYYMMDD `xml:"Start,attr"`
	End    *util.YYYYMMDD `xml:"End,attr"`
}

type Status string

const (
	StatusOpen      Status = "Open"
	StatusClose     Status = "Close"
	StatusOnRequest Status = "OnRequest"
)

type BasicPropertyInfo struct {
	HotelCode *string `xml:"HotelCode,attr"`
	HotelName *string `xml:"HotelName,attr"`
}

type Price struct {
	Start          *util.YYYYMMDD `xml:"Start,attr"`
	End            *util.YYYYMMDD `xml:"End,attr"`
	AmountAfterTax *int           `xml:"AmountAfterTax,attr"`
	CurrencyCode   *string        `xml:"CurrencyCode,attr"`
}

type RoomStays struct {
	RoomStay []RoomStay `xml:"RoomStay"`
}

type RoomStay struct {
	RPH       *string    `xml:"RPH,attr"`
	RoomTypes *RoomTypes `xml:"RoomTypes"`
	RatePlans *RatePlans `xml:"RatePlans"`
	RoomRates *RoomRates `xml:"RoomRates"`
	Total     *Total     `xml:"Total"`
}

type RoomTypes struct {
	RoomType []RoomType `xml:"RoomType"`
}

type RoomType struct {
	RoomTypeCode    *string          `xml:"RoomTypeCode,attr"`
	BedTypeCode     *BedTypeCode     `xml:"BedTypeCode,attr"`
	NonSmoking      *bool            `xml:"NonSmoking,attr"`
	RoomDescription *RoomDescription `xml:"RoomDescription"`
	Amenities       *Amenities       `xml:"Amenities"`
	Occupancy       *[]Occupancy     `xml:"Occupancy"`
}

type RoomDescription struct {
	Name               *string  `xml:"Name,attr"`
	CreateDateTime     *string  `xml:"CreateDateTime,attr"`
	LastModifyDateTime *string  `xml:"LastModifyDateTime,attr"`
	Text               *[]Text  `xml:"Text"`
	Image              *[]Image `xml:"Image"`
	URL                *[]URL   `xml:"URL"`
}

type Text struct {
	Value *string `xml:",chardata"`
}

type Image struct {
	Value *string `xml:",chardata"`
}

type URL struct {
	Value *string `xml:",chardata"`
}

type Amenities struct {
	Amenity []Amenity `xml:"Amenity"`
}

type Amenity struct {
	RoomAmenity *RoomAmenity `xml:"RoomAmenity,attr"`
	ExistsCode  *ExistsCode  `xml:"ExistsCode,attr"`
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
	MinOccupancy *string `xml:"MinOccupancy,attr"`
	MaxOccupancy *string `xml:"MaxOccupancy,attr"`
}

type RatePlans struct {
	RatePlan *[]RatePlan `xml:"RatePlan"`
}

type RatePlan struct {
	RatePlanCode        *string              `xml:"RatePlanCode,attr"`
	EffectiveDate       *string              `xml:"EffectiveDate,attr"`
	ExpireDate          *string              `xml:"ExpireDate,attr"`
	RatePlanType        *string              `xml:"RatePlanType,attr"`
	RatePlanName        *string              `xml:"RatePlanName,attr"`
	CancelPenalties     *CancelPenalties     `xml:"CancelPenalties"`
	RatePlanDescription *RatePlanDescription `xml:"RatePlanDescription"`
	RatePlanInclusions  *RatePlanInclusions  `xml:"RatePlanInclusions"`
	MealsIncluded       *MealsIncluded       `xml:"MealsIncluded"`
	AdditionalDetails   *AdditionalDetails   `xml:"AdditionalDetails"`
}

type CancelPenalties struct {
	CancelPenalty []CancelPenalty `xml:"CancelPenalty"`
}

type CancelPenalty struct {
	ConfirmClassCode   *string               `xml:"ConfirmClassCode,attr"`
	AmountPercent      *AmountPercent        `xml:"AmountPercent"`
	Deadline           *Deadline             `xml:"Deadline"`
	PenaltyDescription *[]PenaltyDescription `xml:"PenaltyDescription"`
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
	CreateDateTime     *string  `xml:"CreateDateTime,attr"`
	LastModifyDateTime *string  `xml:"LastModifyDateTime,attr"`
	Text               *[]Text  `xml:"Text"`
	Image              *[]Image `xml:"Image"`
	URL                *[]URL   `xml:"URL"`
}

type RatePlanInclusions struct {
	TaxInclusive        *string `xml:"TaxInclusive,attr"`
	ServiceFeeInclusive *string `xml:"ServiceFeeInclusive,attr"`
}

type AdditionalDetails struct {
	AdditionalDetail *[]AdditionalDetail `xml:"AdditionalDetail"`
}

type AdditionalDetail struct {
	Type              *string            `xml:"Type,attr"`
	DetailDescription *DetailDescription `xml:"DetailDescription"`
}

type DetailDescription struct {
	Text *[]Text `xml:"Text"`
}

type RoomRates struct {
	RoomRate []RoomRate `xml:"RoomRate"`
}

type RoomRate struct {
	EffectiveDate      *string            `xml:"EffectiveDate,attr"`
	ExpireDate         *string            `xml:"ExpireDate,attr"`
	RoomTypeCode       string             `xml:"RoomTypeCode,attr"`
	NumberOfUnits      string             `xml:"NumberOfUnits,attr"`
	RatePlanCode       string             `xml:"RatePlanCode,attr"`
	AvailabilityStatus AvailabilityStatus `xml:"AvailabilityStatus,attr"`
	Rates              *Rates             `xml:"Rates"`
	Total              *Total             `xml:"Total"`
	ServiceRPHs        *ServiceRPHs       `xml:"ServiceRPHs"`
	GuestCounts        *GuestCounts       `xml:"GuestCounts"`
}

type AvailabilityStatus string

const (
	AvailabilityStatusAvailableForSale AvailabilityStatus = "AvailableForSale"
	AvailabilityStatusOnRequest        AvailabilityStatus = "OnRequest"
	AvailabilityStatusClosedOut        AvailabilityStatus = "ClosedOut"
)

type Rates struct {
	Rate []Rate `xml:"Rate"`
}

type Rate struct {
	AgeQualifyingCode *string `xml:"AgeQualifyingCode,attr"`
	RoomPricingType   *string `xml:"RoomPricingType,attr"`
	Base              *Base   `xml:"Base"`
}

type Base struct {
	AmountAfterTax *string `xml:"AmountAfterTax,attr"`
	CurrencyCode   *string `xml:"CurrencyCode,attr"`
	Type           *string `xml:"Type,attr"`
}

type Total struct {
	AmountAfterTax *string `xml:"AmountAfterTax,attr"`
	CurrencyCode   *string `xml:"CurrencyCode,attr"`
}

type ServiceRPHs struct {
	ServiceRPH []ServiceRPH `xml:"ServiceRPH"`
}

type ServiceRPH struct {
	RPH string `xml:"RPH,attr"`
}

type Services struct {
	Service []Service `xml:"Service"`
}

type Service struct {
	ServiceRPH     *string           `xml:"ServiceRPH,attr"`
	RatePlanCode   *string           `xml:"RatePlanCode,attr"`
	Type           string            `xml:"Type,attr"`
	ID             string            `xml:"ID,attr"`
	ServicePrise   *[]ServicePrice   `xml:"Price"`
	ServiceDetails *[]ServiceDetails `xml:"ServiceDetails"`
	TPAExtensions  TPAExtensions     `xml:"TPA_Extensions"`
}

type ServicePrice struct {
	Base          Base   `xml:"Base"`
	EffectiveDate string `xml:"EffectiveDate,attr"`
	ExpireDate    string `xml:"ExpireDate,attr"`
}
type ServiceDetails struct {
	Timespan           *Timespan           `xml:"Timespan"`
	ServiceDescription *ServiceDescription `xml:"ServiceDescription"`
}

type Timespan struct {
	Start *util.YYYYMMDD `xml:"Start,attr"`
	End   *util.YYYYMMDD `xml:"End,attr"`
}

type ServiceDescription struct {
	Name  *string  `xml:"Name,attr"`
	Text  *[]Text  `xml:"Text"`
	Image *[]Image `xml:"Image"`
	URL   *[]URL   `xml:"URL"`
}

type TPAExtensions struct {
	AddtionalServiceDetails AddtionalServiceDetails `xml:"AddtionalServiceDetails"`
}

type AddtionalServiceDetails struct {
	ServiceSpecificableRange ServiceSpecificableRange `xml:"ServiceSpecificableRange"`
}

type ServiceSpecificableRange struct {
	MinCount int `xml:"MinCount,attr"`
	MaxCount int `xml:"MaxCount,attr"`
}

type Criteria struct {
	Criterion []Criterion `xml:"Criterion"`
}

type Errors struct {
	Error []Error `xml:"Error"`
}

type Error struct {
	Type      string `xml:"Type,attr"`
	ShortText string `xml:"ShortText,attr"`
	Code      string `xml:"Code,attr"`
}
