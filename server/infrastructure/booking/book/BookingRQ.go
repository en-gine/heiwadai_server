package book

import (
	"encoding/xml"

	"server/infrastructure/booking/util"
)

type EnvelopeRQ[TBody any] struct {
	XMLName xml.Name `xml:"soapenv:Envelope"`
	SoapEnv string   `xml:"xmlns:soapenv,attr"`
	Naif    string   `xml:"xmlns:naif,attr"`
	Header  string   `xml:"soapenv:Header"`
	Body    TBody    `xml:"soapenv:Body"`
}

type Body struct {
	EntryBooking EntryBooking `xml:"naif:entryBooking"`
}

type EntryBooking struct {
	EntryBookingRequest EntryBookingRequest `xml:"entryBookingRequest"`
}

type EntryBookingRequest struct {
	CommonRequest          CommonRequest          `xml:"commonRequest"`
	ExtendLincoln          ExtendLincoln          `xml:"extendLincoln"`
	SendInformation        SendInformation        `xml:"SendInformation"`
	AllotmentBookingReport AllotmentBookingReport `xml:"AllotmentBookingReport"`
}

type CommonRequest struct {
	AgtID       string `xml:"agtId"`
	AgtPassword string `xml:"agtPassword"`
	SystemDate  string `xml:"systemDate"`
}

type ExtendLincoln struct {
	TllHotelCode     string  `xml:"tllHotelCode"`
	UseTllPlan       int     `xml:"useTllPlan"`       // リンカーンプラン0:使用しない、1:使用する
	TllBookingNumber *string `xml:"tllBookingNumber"` // 予約番号（変更ありの場合に入れる）
	TllCharge        *int    `xml:"tllCharge"`        // TLリンカーンの手数料を取る場合
}

type SendInformation struct {
	AssignDiv int `xml:"assignDiv"`
	GenderDiv int `xml:"genderDiv"`
}

type AllotmentBookingReport struct {
	TransactionType          TransactionType          `xml:"TransactionType"`
	AccommodationInformation AccommodationInformation `xml:"AccommodationInformation"`
	SalesOfficeInformation   SalesOfficeInformation   `xml:"SalesOfficeInformation"`
	BasicInformation         BasicInformation         `xml:"BasicInformation"`
	BasicRateInformation     BasicRateInformation     `xml:"BasicRateInformation"`
	MemberInformation        MemberInformation        `xml:"MemberInformation"`
	OptionInformation        OptionInformation        `xml:"OptionInformation"`
	RoomInformationList      RoomInformationList      `xml:"RoomInformationList"`
}

type TransactionType struct {
	DataFrom           string `xml:"DataFrom"`
	DataClassification string `xml:"DataClassification"`
	DataID             string `xml:"DataID"`
}

type AccommodationInformation struct {
	AccommodationArea *string `xml:"AccommodationArea"`
	AccommodationName string  `xml:"AccommodationName"`
	AccommodationCode string  `xml:"AccommodationCode"`
	ChainName         *string `xml:"ChainName"`
}

type SalesOfficeInformation struct {
	SalesOfficeCompanyName    string  `xml:"SalesOfficeCompanyName"`    // 旅行社名
	SalesOfficeName           string  `xml:"SalesOfficeName"`           // 旅行社営業所名
	SalesOfficeCode           string  `xml:"SalesOfficeCode"`           // 旅行社営業所コード
	SalesOfficePersonInCharge *string `xml:"SalesOfficePersonInCharge"` // 旅行社担当者名
	SalesOfficeEmail          *string `xml:"SalesOfficeEmail"`
	SalesOfficePhoneNumber    *string `xml:"SalesOfficePhoneNumber"`
	SalesOfficeFaxNumber      *string `xml:"SalesOfficeFaxNumber"`
}

type BasicInformation struct {
	TravelAgencyBookingNumber  string        `xml:"TravelAgencyBookingNumber"`
	TravelAgencyBookingDate    string        `xml:"TravelAgencyBookingDate"`
	TravelAgencyBookingTime    string        `xml:"TravelAgencyBookingTime"`
	GuestOrGroupMiddleName     *string       `xml:"GuestOrGroupMiddleName"`
	GuestOrGroupNameSingleByte string        `xml:"GuestOrGroupNameSingleByte"`
	GuestOrGroupNameDoubleByte *string       `xml:"GuestOrGroupNameDoubleByte"`
	GuestOrGroupKanjiName      string        `xml:"GuestOrGroupKanjiName"`
	GuestOrGroupContactDiv     *string       `xml:"GuestOrGroupContactDiv"`
	GuestOrGroupCellularNumber *string       `xml:"GuestOrGroupCellularNumber"`
	GuestOrGroupOfficeNumber   *string       `xml:"GuestOrGroupOfficeNumber"`
	GuestOrGroupPhoneNumber    string        `xml:"GuestOrGroupPhoneNumber"`
	GuestOrGroupEmail          string        `xml:"GuestOrGroupEmail"`
	GuestOrGroupPostalCode     string        `xml:"GuestOrGroupPostalCode"`
	GuestOrGroupAddress        string        `xml:"GuestOrGroupAddress"`
	CheckInDate                util.YYYYMMDD `xml:"CheckInDate"`
	CheckInTime                string        `xml:"CheckInTime"`
	Nights                     uint          `xml:"Nights"`
	TotalRoomCount             uint          `xml:"TotalRoomCount"`        // 利用客室合計数
	GrandTotalPaxCount         uint          `xml:"GrandTotalPaxCount"`    //	お客様総合計人数
	TotalPaxMaleCount          uint          `xml:"TotalPaxMaleCount"`     //	お客様総合計男性人数
	TotalPaxFemaleCount        uint          `xml:"TotalPaxFemaleCount"`   //	お客様総合計女性人数（男女区別が無い場合は0）
	TotalChildA70Count         uint          `xml:"TotalChildA70Count"`    //	お客様総合計子供A（70%）人数
	MealCondition              MealCondition `xml:"MealCondition"`         // 食事条件
	PackageType                string        `xml:"PackageType"`           // Package固定
	PackagePlanCode            string        `xml:"PackagePlanCode"`       // 企画(パッケージ)コード
	PackagePlanName            string        `xml:"PackagePlanName"`       // 企画(パッケージ)名
	PackagePlanContent         string        `xml:"PackagePlanContent"`    // 企画(パッケージ)内容
	SpecialServiceRequest      string        `xml:"SpecialServiceRequest"` // お客様からの要望
}
type MealCondition string

const (
	MealCondition1night2meals    MealCondition = "1night2meals"    // 一泊二食
	MealCondition1nightBreakfast MealCondition = "1nightBreakfast" // 一泊朝食付
	MealConditionWithoutMeal     MealCondition = "WithoutMeal"     // 食事なし
	MealConditionOther           MealCondition = "Other"           // 他
)

type BasicRateInformation struct {
	RoomRateOrPersonalRate                              RoomRateType       `xml:"RoomRateOrPersonalRate"`
	TaxServiceFee                                       TaxServiceFeeType  `xml:"TaxServiceFee"`
	Payment                                             string             `xml:"Payment"`
	SettlementDiv                                       int                `xml:"SettlementDiv"`
	TotalAccommodationCharge                            int                `xml:"TotalAccommodationCharge"`
	TotalAccommodationConsumptionTax                    string             `xml:"TotalAccommodationConsumptionTax"`
	TotalAccommodationHotSpringTax                      string             `xml:"TotalAccommodationHotSpringTax"`
	TotalAccomodationServiceCharge                      string             `xml:"TotalAccomodationServiceCharge"`
	TotalAccommodationDiscountPoints                    int                `xml:"TotalAccommodationDiscountPoints"`
	TotalAccommodationConsumptionTaxAfterDiscountPoints string             `xml:"TotalAccommodationConsumptionTaxAfterDiscountPoints"`
	AmountClaimed                                       int                `xml:"AmountClaimed"`
	PointsDiscountList                                  PointsDiscountList `xml:"PointsDiscountList"`
}

type RoomRateType = string

const (
	RoomRatePersonal RoomRateType = "PersonalRate"
	RoomRateRoom     RoomRateType = "RoomRate"
)

type TaxServiceFeeType = string

const (
	IncludingServiceWithoutTax TaxServiceFeeType = "IncludingServiceWithoutTax" // サ込税別
	IncludingServiceAndTax     TaxServiceFeeType = "IncludingServiceAndTax"     // サ込税込
	WithoutServiceAndTax       TaxServiceFeeType = "WithoutServiceAndTax"       // サ別税別
	WithoutServiceIncludingTax TaxServiceFeeType = "WithoutServiceIncludingTax" // サ別税込

)

type PointsDiscountList struct {
	PointsDiv          int    `xml:"PointsDiv"`
	PointsDiscountName string `xml:"PointsDiscountName"`
	PointsDiscount     int    `xml:"PointsDiscount"`
}

type MemberInformation struct {
	MemberName                  string `xml:"MemberName"`
	MemberKanjiName             string `xml:"MemberKanjiName"`
	MemberMiddleName            string `xml:"MemberMiddleName"`
	MemberDateOfBirth           string `xml:"MemberDateOfBirth"`
	MemberEmergencyNumber       string `xml:"MemberEmergencyNumber"`
	MemberOccupation            string `xml:"MemberOccupation"`
	MemberOrganization          string `xml:"MemberOrganization"`
	MemberOrganizationKana      string `xml:"MemberOrganizationKana"`
	MemberOrganizationCode      string `xml:"MemberOrganizationCode"`
	MemberPosition              string `xml:"MemberPosition"`
	MemberOfficePostalCode      string `xml:"MemberOfficePostalCode"`
	MemberOfficeAddress         string `xml:"MemberOfficeAddress"`
	MemberOfficeTelephoneNumber string `xml:"MemberOfficeTelephoneNumber"`
	MemberOfficeFaxNumber       string `xml:"MemberOfficeFaxNumber"`
	MemberGenderDiv             string `xml:"MemberGenderDiv"`
	MemberClass                 string `xml:"MemberClass"`
	CurrentPoints               string `xml:"CurrentPoints"`
	MailDemandDiv               string `xml:"MailDemandDiv"`
	PamphletDemandDiv           string `xml:"PamphletDemandDiv"`
	MemberID                    int    `xml:"MemberID"`
	MemberPhoneNumber           string `xml:"MemberPhoneNumber"`
	MemberEmail                 string `xml:"MemberEmail"`
	MemberPostalCode            string `xml:"MemberPostalCode"`
	MemberAddress               string `xml:"MemberAddress"`
}

type OptionInformation struct {
	OptionList OptionList `xml:"OptionList"`
}

type OptionList struct {
	OptionDate  string `xml:"OptionDate"`
	OptionCode  string `xml:"OptionCode"`
	Name        string `xml:"Name"`
	NameRequest string `xml:"NameRequest"`
	OptionCount string `xml:"OptionCount"`
	OptionRate  string `xml:"OptionRate"`
}

type RoomInformationList struct {
	RoomAndGuestList []RoomAndGuest `xml:"RoomAndGuestList>RoomAndGuest"`
}

type RoomAndGuest struct {
	RoomInformation     RoomInformation     `xml:"RoomInformation"`
	RoomRateInformation RoomRateInformation `xml:"RoomRateInformation"`
}

type RoomInformation struct {
	RoomTypeCode             int    `xml:"RoomTypeCode"`
	RoomTypeName             string `xml:"RoomTypeName"`
	RelationRoomCode         *int   `xml:"RelationRoomCode"`
	RelationRoomName         *int   `xml:"RelationRoomName"`
	PerRoomPaxCount          uint   `xml:"PerRoomPaxCount"`
	RepresentativePersonName *int   `xml:"RepresentativePersonName"`
}

type RoomRateInformation struct {
	RoomDate                util.YYYYMMDD `xml:"RoomDate"`
	PerPaxRate              *int          `xml:"PerPaxRate"`
	PerPaxFemaleRate        *int          `xml:"PerPaxFemaleRate"`
	PerChildA70Rate         *int          `xml:"PerChildA70Rate"`
	PerChildA70Rate2        *int          `xml:"PerChildA70Rate2"`
	PerChildB50Rate         *int          `xml:"PerChildB50Rate"`
	PerChildB50Rate2        *int          `xml:"PerChildB50Rate2"`
	PerChildC30Rate         *int          `xml:"PerChildC30Rate"`
	PerChildDNoneRate       *int          `xml:"PerChildDNoneRate"`
	RoomRatePaxMaleCount    *int          `xml:"RoomRatePaxMaleCount"`
	RoomRatePaxFemaleCount  *int          `xml:"RoomRatePaxFemaleCount"`
	RoomRateChildA70Count   *int          `xml:"RoomRateChildA70Count"`
	RoomRateChildA70Count2  *int          `xml:"RoomRateChildA70Count2"`
	RoomRateChildB50Count   *int          `xml:"RoomRateChildB50Count"`
	RoomRateChildB50Count2  *int          `xml:"RoomRateChildB50Count2"`
	RoomRateChildC30Count   *int          `xml:"RoomRateChildC30Count"`
	RoomRateChildDNoneCount *int          `xml:"RoomRateChildDNoneCount"`
	RoomPaxMaleRequest      string        `xml:"RoomPaxMaleRequest"`
	RoomPaxFemaleRequest    string        `xml:"RoomPaxFemaleRequest"`
	// 他のフィールドも同様に定義する
}
