package book

import (
	"encoding/xml"
	"time"

	"server/core/entity"
	"server/infrastructure/booking/util"

	"github.com/ktnyt/go-moji"
)

type EnvelopeRQ struct {
	XMLName xml.Name `xml:"soapenv:Envelope"`
	SoapEnv string   `xml:"xmlns:soapenv,attr"`
	Naif    string   `xml:"xmlns:naif,attr"`
	Header  string   `xml:"soapenv:Header"`
	Body    Body     `xml:"soapenv:Body"`
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
	TotalRoomCount             uint          `xml:"TotalRoomCount"`      // 利用客室合計数
	GrandTotalPaxCount         uint          `xml:"GrandTotalPaxCount"`  //	お客様総合計人数
	TotalPaxMaleCount          uint          `xml:"TotalPaxMaleCount"`   //	お客様総合計男性人数
	TotalPaxFemaleCount        uint          `xml:"TotalPaxFemaleCount"` //	お客様総合計女性人数（男女区別が無い場合は0）
	TotalChildA70Count         uint          `xml:"TotalChildA70Count"`  //	お客様総合計子供A（70%）人数
	PackagePlanCode            string        `xml:"PackagePlanCode"`     // 企画(パッケージ)コード
	MealCondition              MealCondition `xml:"MealCondition"`       // 食事条件
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
	OptionCount int    `xml:"OptionCount"`
	OptionRate  int    `xml:"OptionRate"`
}

type RoomInformationList struct {
	RoomAndGuestList []RoomAndGuest `xml:"RoomAndGuestList>RoomAndGuest"`
}

type RoomAndGuest struct {
	RoomInformation     RoomInformation     `xml:"RoomInformation"`
	RoomRateInformation RoomRateInformation `xml:"RoomRateInformation"`
}

type RoomInformation struct {
	RoomTypeCode             string `xml:"RoomTypeCode"`
	RoomTypeName             string `xml:"RoomTypeName"`
	RelationRoomCode         *int   `xml:"RelationRoomCode"`
	RelationRoomName         *int   `xml:"RelationRoomName"`
	PerRoomPaxCount          uint   `xml:"PerRoomPaxCount"`
	RepresentativePersonName *int   `xml:"RepresentativePersonName"`
}

type RoomRateInformation struct {
	RoomDate                util.YYYYMMDD `xml:"RoomDate"`
	PerPaxRate              int           `xml:"PerPaxRate"`
	PerPaxFemaleRate        int           `xml:"PerPaxFemaleRate"`
	PerChildA70Rate         int           `xml:"PerChildA70Rate"`
	PerChildA70Rate2        int           `xml:"PerChildA70Rate2"`
	PerChildB50Rate         int           `xml:"PerChildB50Rate"`
	PerChildB50Rate2        int           `xml:"PerChildB50Rate2"`
	PerChildC30Rate         int           `xml:"PerChildC30Rate"`
	PerChildDNoneRate       int           `xml:"PerChildDNoneRate"`
	RoomRatePaxMaleCount    int           `xml:"RoomRatePaxMaleCount"`
	RoomRatePaxFemaleCount  int           `xml:"RoomRatePaxFemaleCount"`
	RoomRateChildA70Count   int           `xml:"RoomRateChildA70Count"`
	RoomRateChildA70Count2  int           `xml:"RoomRateChildA70Count2"`
	RoomRateChildB50Count   int           `xml:"RoomRateChildB50Count"`
	RoomRateChildB50Count2  int           `xml:"RoomRateChildB50Count2"`
	RoomRateChildC30Count   int           `xml:"RoomRateChildC30Count"`
	RoomRateChildDNoneCount int           `xml:"RoomRateChildDNoneCount"`
	RoomPaxMaleRequest      string        `xml:"RoomPaxMaleRequest"`
	RoomPaxFemaleRequest    string        `xml:"RoomPaxFemaleRequest"`
	// 他のフィールドも同様に定義する
}

// bookID : CCYYMMDD+9桁連番（0埋め、データ毎に+1）
func NewBookingRQ(agentID string, agentPass string, bookData *entity.Booking) *EnvelopeRQ {
	bookID := time.Now().Format("20060102") + "000000001"
	plan := bookData.BookPlan
	store := bookData.BookPlan.Store
	guest := bookData.GuestData

	var mealCondition MealCondition

	switch plan.MealType.String() {
	case "朝食あり夕食あり":
		mealCondition = MealCondition1night2meals
	case "朝食あり夕食なし":
		mealCondition = MealCondition1nightBreakfast
	case "朝食なし夕食あり":
		mealCondition = MealCondition1nightBreakfast
	case "食事なし":
		mealCondition = MealConditionWithoutMeal
	default:
		mealCondition = MealConditionOther
	}

	guestNameKana := moji.Convert(guest.LastNameKana, moji.ZK, moji.HK) + " " + moji.Convert(guest.FirstNameKana, moji.ZK, moji.HK)
	return &EnvelopeRQ{
		SoapEnv: "http://schemas.xmlsoap.org/soap/envelope/",
		Naif:    "http://naifc3000.naifc30.nai.lincoln.seanuts.co.jp/",
		Header:  "",
		Body: Body{
			EntryBooking: EntryBooking{
				EntryBookingRequest: EntryBookingRequest{
					CommonRequest: CommonRequest{
						AgtID:       agentID,
						AgtPassword: agentPass,
						SystemDate:  time.Now().Format("2006-01-02T15:04:05"),
					},
					ExtendLincoln: ExtendLincoln{
						TllHotelCode: store.BookingSystemID,
						UseTllPlan:   1,
					},
					SendInformation: SendInformation{
						AssignDiv: 1, // 部屋割ありデフォルト
						GenderDiv: 0, // 男女区分なしデフォルト
					},
					AllotmentBookingReport: AllotmentBookingReport{
						TransactionType: TransactionType{
							DataFrom:           "FromTravelAgency",
							DataClassification: "NewBookReport",
							DataID:             bookID,
						},
						AccommodationInformation: AccommodationInformation{ // 宿泊施設情報
							AccommodationName: store.Name + *store.BranchName,
							AccommodationCode: store.BookingSystemID,
						},
						SalesOfficeInformation: SalesOfficeInformation{
							SalesOfficeCompanyName: "平和台ホテルアプリ",
							SalesOfficeName:        "",
							SalesOfficeCode:        "10000000",
						},
						BasicInformation: BasicInformation{
							TravelAgencyBookingNumber:  bookID,
							TravelAgencyBookingDate:    time.Now().Format("2006-01-02"),
							TravelAgencyBookingTime:    time.Now().Format("15:04:05"),
							GuestOrGroupNameSingleByte: guestNameKana,
							GuestOrGroupKanjiName:      guest.LastName + " " + guest.FirstName,
							GuestOrGroupPhoneNumber:    *guest.Tel,
							GuestOrGroupEmail:          guest.Mail,
							GuestOrGroupPostalCode:     *guest.ZipCode,
							GuestOrGroupAddress:        guest.Prefecture.String() + *guest.City + *guest.Address,
							CheckInDate:                util.YYYYMMDD(bookData.StayFrom.Format("2006-01-02")),
							CheckInTime:                string(bookData.CheckInTime),
							Nights:                     uint(bookData.StayTo.Sub(bookData.StayFrom).Hours() / 24),
							TotalRoomCount:             bookData.RoomCount,
							GrandTotalPaxCount:         bookData.Adult + bookData.Child,
							TotalPaxMaleCount:          bookData.Adult,
							TotalPaxFemaleCount:        0,
							TotalChildA70Count:         bookData.Child,
							PackagePlanCode:            bookData.BookPlan.ID,
							MealCondition:              mealCondition,
						},
						BasicRateInformation: BasicRateInformation{
							RoomRateOrPersonalRate:   RoomRateRoom,
							TaxServiceFee:            IncludingServiceAndTax,
							Payment:                  "Cash",
							SettlementDiv:            0, // 現地決済
							TotalAccommodationCharge: int(bookData.TotalCost),
							PointsDiscountList: PointsDiscountList{
								PointsDiscount: 0, // ポイント値引き
							},
						},
						RoomInformationList: RoomInformationList{
							RoomAndGuestList: []RoomAndGuest{
								{
									RoomInformation: RoomInformation{
										RoomTypeCode:    plan.RoomType.Code(),
										RoomTypeName:    plan.RoomType.String(),
										PerRoomPaxCount: bookData.Adult + bookData.Child,
									},
									RoomRateInformation: RoomRateInformation{
										RoomDate: util.YYYYMMDD(bookData.StayFrom.Format("2006-01-02")),
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
