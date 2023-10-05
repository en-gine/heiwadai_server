package avail

import (
	"encoding/xml"

	"server/infrastructure/booking/util"
)

type OTAHotelAvailRQ struct {
	XMLName              xml.Name             `xml:"ns:OTA_HotelAvailRQ"`
	AvailRequestSegments AvailRequestSegments `xml:"ns:AvailRequestSegments"`
	Version              string               `xml:"Version,attr"`
	PrimaryLangID        string               `xml:"PrimaryLangID,attr"`  // デフォルトは"jpn"
	SummaryOnly          *bool                `xml:"SummaryOnly,attr"`    // サマリ情報のみを返すフラグ。デフォルトはfalse
	AvailRatesOnly       *bool                `xml:"AvailRatesOnly,attr"` // 販売している部屋とプランのみを返すフラグ。デフォルトはTrue
	HotelStayOnly        *bool                `xml:"HotelStayOnly,attr"`  // ホテル情報のみを返すフラグ。デフォルトはfalse
	RateDetailsInd       *bool                `xml:"RateDetailsInd,attr"` // 料金の詳細情報を返すフラグ。デフォルトはTrue
	PricingMethod        PricingMethod        `xml:"PricingMethod,attr"`  // デフォルトは"None"
}

type PricingMethod string

const (
	PricingMethodAverage        PricingMethod = "Average"
	PricingMethodLowest         PricingMethod = "Lowest"           // 最安値
	PricingMethodHeighest       PricingMethod = "Heighest"         // 最高値
	PricingMethodLowestperstay  PricingMethod = "Lowest per stay"  // 連泊期間内の料金合計の最安値
	PricingMethodHighestperstay PricingMethod = "Highest per stay" // 連泊期間内の料金合計の最高値
	PricingMethodNone           PricingMethod = "None"             // 指定なし
)

type AvailRequestSegments struct {
	AvailReqType        *AvailReqType       `xml:"AvailReqType,attr"`
	AvailRequestSegment AvailRequestSegment `xml:"ns:AvailRequestSegment"`
}

type AvailRequestSegment struct {
	HotelSearchCriteria HotelSearchCriteria `xml:"ns:HotelSearchCriteria"`
}

type HotelSearchCriteria struct {
	Criterion Criterion `xml:"ns:Criterion"`
}

type Criterion struct {
	HotelRef           []HotelRef          `xml:"ns:HotelRef"`           // ホテルコード
	StayDateRange      *StayDateRange      `xml:"ns:StayDateRange"`      // 宿泊期間
	RateRange          *RateRange          `xml:"ns:RateRange"`          // 価格帯
	RatePlanCandidates *RatePlanCandidates `xml:"ns:RatePlanCandidates"` // プラン検索条件
	RoomStayCandidates *RoomStayCandidates `xml:"ns:RoomStayCandidates"` // 部屋検索条件
}

type HotelRef struct {
	HotelCode string `xml:"HotelCode,attr"`
}

type StayDateRange struct {
	Start *util.YYYYMMDD `xml:"Start,attr"` // 宿泊開始日
	End   *util.YYYYMMDD `xml:"End,attr"`   // 宿泊終了日
}

type RateRange struct {
	MinRate *int `xml:"MinRate,attr"` // 最低価格
	MaxRate *int `xml:"MaxRate,attr"` // 最高価格
}

type RatePlanCandidates struct {
	RatePlanCandidate []RatePlanCandidate `xml:"ns:RatePlanCandidate"`
}

type RatePlanCandidate struct {
	MealsIncluded *MealsIncluded `xml:"ns:MealsIncluded"`
}

type MealsIncluded struct {
	Breakfast *bool `xml:"Breakfast,attr"`
	Lunch     *bool `xml:"Lunch,attr"`
	Dinner    *bool `xml:"Dinner,attr"`
}

type RoomStayCandidates struct {
	RoomStayCandidate []RoomStayCandidate `xml:"ns:RoomStayCandidate"`
}

type RoomStayCandidate struct {
	RoomTypeCode  *string        `xml:"RoomTypeCode,attr"`
	BedTypeCode   *BedTypeCode   `xml:"BedTypeCode,attr"`
	NonSmoking    *bool          `xml:"NonSmoking,attr"` // True：禁煙、False：喫煙、省略：条件指定なし
	Quantity      *int           `xml:"Quantity,attr"`   // 利用部屋数
	EffectiveDate *util.YYYYMMDD `xml:"EffectiveDate,attr"`
	ExpireDate    *util.YYYYMMDD `xml:"ExpireDate,attr"`
	GuestCounts   *GuestCounts   `xml:"ns:GuestCounts"`
}

type GuestCounts struct {
	GuestCount []GuestCount `xml:"ns:GuestCount"`
}

type GuestCount struct {
	AgeQualifyingCode AgeQualifyingCode `xml:"AgeQualifyingCode,attr"`
	Count             int               `xml:"Count,attr"`
}

type BedTypeCode string

const (
	BedTypeDouble     BedTypeCode = "1"
	BedTypeTatami     BedTypeCode = "7"
	BedTypeTwin       BedTypeCode = "8"
	BedTypeSingle     BedTypeCode = "9"
	BedTypeSemiDouble BedTypeCode = "14"
	BedTypeTriple     BedTypeCode = "15"
	BedTypeFour       BedTypeCode = "16"
	BedTypeOther      BedTypeCode = "18"
)

type AgeQualifyingCode string

const (
	AgeQualifyingAdult AgeQualifyingCode = "10"
	AgeQualifyingChild AgeQualifyingCode = "8"
)

type AvailReqType string

const (
	AvailReqTypeRoom    AvailReqType = "Room"
	AvailReqTypeNonRoom AvailReqType = "NonRoom" // プランのみ検索
	AvailReqTypeBoth    AvailReqType = "Both"
)
