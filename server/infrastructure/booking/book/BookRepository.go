package book

import (
	"time"

	"server/core/entity"
	domainErr "server/core/errors"
	queryservice "server/core/infra/queryService"
	"server/core/infra/repository"
	"server/infrastructure/booking/util"
	"server/infrastructure/env"
	"server/infrastructure/logger"
)

type BookRepository struct {
	storeQuery queryservice.IStoreQueryService
	bookQuery  queryservice.IBookQueryService
}

var _ repository.IBookAPIRepository = &BookRepository{}
var (
	TLBookingUser = env.GetEnv(env.TlbookingUsername)
	TLBookingPass = env.GetEnv(env.TlbookingPassword)
	BookURL       = env.GetEnv(env.TlbookingBookingApiUrl)
	CancelURL     = env.GetEnv(env.TlbookingCancelApiUrl)
)

func NewBookRepository(storeQuery queryservice.IStoreQueryService, bookQuery queryservice.IBookQueryService) *BookRepository {
	return &BookRepository{
		storeQuery: storeQuery,
		bookQuery:  bookQuery,
	}
}

func (p BookRepository) Cancel(bookData *entity.Booking, newDataID string) (*domainErr.DomainError, error) {
	store, err := p.storeQuery.GetStayableByID(bookData.BookPlan.Plan.StoreID)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	if bookData.TlBookingNumber == nil {
		return domainErr.NewDomainError(domainErr.CancelButNeedFeedBack, "予約番号が存在しません"), nil
	}

	reqBody := NewCancelRQ(bookData, store, newDataID)
	res, err := util.Request[EnvelopeRQ[CancelBody], EnvelopeRS[CancelBodyRS]](CancelURL, reqBody)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	if res.Body.DeleteBookingWithCPResponse.DeleteBookingWithCPResult.CommonResponse.ResultCode == "False" {
		msg := res.Body.DeleteBookingWithCPResponse.DeleteBookingWithCPResult.CommonResponse.ErrorInfos.ErrorMsg
		code := res.Body.DeleteBookingWithCPResponse.DeleteBookingWithCPResult.CommonResponse.ErrorInfos.ErrorCode
		logger.Error(code + ":" + msg)
		return domainErr.NewDomainError(domainErr.CancelButNeedFeedBack, msg), nil
	}

	return nil, nil
}

func (p *BookRepository) Reserve(
	bookData *entity.Booking,
) (*string, *domainErr.DomainError, error) {
	store, err := p.storeQuery.GetStayableByID(bookData.BookPlan.Plan.StoreID)
	if err != nil {
		logger.Error(err.Error())
		return nil, nil, err
	}

	reqBody := NewBookingRQ(
		bookData,
		store,
	)

	res, err := util.Request[EnvelopeRQ[Body], EnvelopeRS[BodyRS]](BookURL, reqBody)
	if err != nil {
		logger.Error(err.Error())
		return nil, nil, err
	}
	if res.Body.EntryBookingResponse.EntryBookingResult.CommonResponse.ResultCode == "False" {
		msg := res.Body.EntryBookingResponse.EntryBookingResult.CommonResponse.ErrorInfos.ErrorMsg
		code := res.Body.EntryBookingResponse.EntryBookingResult.CommonResponse.ErrorInfos.ErrorCode
		logger.Error(code + ":" + msg)
		return nil, domainErr.NewDomainError(domainErr.CancelButNeedFeedBack, msg), nil
	}

	TlBookingNumber := res.Body.EntryBookingResponse.EntryBookingResult.ExtendLincoln.TllBookingNumber
	return TlBookingNumber, nil, nil
}

func NewBookingRQ(bookData *entity.Booking, store *entity.StayableStore) *EnvelopeRQ[Body] {
	plan := bookData.BookPlan.Plan
	guest := bookData.GuestData

	var mealCondition MealCondition
	var specificMealCondition SpecificMealCondition
	switch plan.MealType.String() {
	case "朝食あり夕食あり":
		mealCondition = MealCondition1night2meals
		specificMealCondition = SpecificMealConditionIncludingBreakfastAndDinner
	case "朝食あり夕食なし":
		mealCondition = MealCondition1nightBreakfast
		specificMealCondition = SpecificMealConditionIncludingBreakfast
	case "朝食なし夕食あり":
		mealCondition = MealConditionOther
		specificMealCondition = SpecificMealConditionIncludingDinner
	case "食事なし": // SpecificMealConditionはない
		mealCondition = MealConditionWithoutMeal
	default:
		mealCondition = MealConditionOther
	}

	var BookingSystemID string
	if env.GetEnv(env.TlbookingIsTest) == "true" {
		BookingSystemID = "E69502"
	} else {
		BookingSystemID = store.BookingSystemID
	}

	guestNameKana := util.HiraToHalfKana(guest.LastNameKana + " " + guest.FirstNameKana)
	dateInfos := *bookData.BookPlan.StayDateInfos
	var RoomAndGuestList []RoomAndGuest
	for _, dateInfo := range dateInfos {
		unitPrice := int(dateInfo.StayDateTotalPrice) / int(bookData.RoomCount) / int(bookData.Adult+bookData.Child)

		for i := 0; i < int(bookData.RoomCount); i++ {
			xml := RoomAndGuest{
				RoomInformation: RoomInformation{
					RoomTypeCode:    plan.TlBookingRoomTypeCode,
					RoomTypeName:    plan.TlBookingRoomTypeName,
					PerRoomPaxCount: (bookData.Adult + bookData.Child) / bookData.RoomCount, // 1部屋あたりの人数
				},
				RoomRateInformation: RoomRateInformation{
					RoomDate:             util.DateToStrDate(dateInfo.StayDate),
					PerPaxRate:           &unitPrice,
					RoomRatePaxMaleCount: bookData.Adult,
				},
			}
			RoomAndGuestList = append(RoomAndGuestList, xml)
		}
	}

	// for d := bookData.StayFrom; d.Unix() < bookData.StayTo.Unix(); d = d.AddDate(0, 0, 1) {
	// 	for i := 0; i < int(bookData.RoomCount); i++ {
	// 		xml := RoomAndGuest{
	// 			RoomInformation: RoomInformation{
	// 				RoomTypeCode:    plan.TlBookingRoomTypeCode,
	// 				RoomTypeName:    plan.RoomType.String(),
	// 				PerRoomPaxCount: (bookData.Adult + bookData.Child) / bookData.RoomCount, // 1部屋あたりの人数
	// 			},
	// 			RoomRateInformation: RoomRateInformation{
	// 				RoomDate:   util.DateToStrDate(d),
	// 				PerPaxRate: &unitPrice,
	// 			},
	// 		}

	// 		RoomAndGuestList = append(RoomAndGuestList, xml)
	// 	}
	// }

	return &EnvelopeRQ[Body]{
		SoapEnv: "http://schemas.xmlsoap.org/soap/envelope/",
		Naif:    "http://naifc3000.naifc30.nai.lincoln.seanuts.co.jp/",
		Header:  "",
		Body: Body{
			EntryBooking: EntryBooking{
				EntryBookingRequest: EntryBookingRequest{
					CommonRequest: CommonRequest{
						AgtID:       TLBookingUser,
						AgtPassword: TLBookingPass,
						SystemDate:  time.Now().Format("2006-01-02T15:04:05"),
					},
					ExtendLincoln: ExtendLincoln{
						TllHotelCode: BookingSystemID,
						UseTllPlan:   0,
					},
					SendInformation: SendInformation{
						AssignDiv: 1, // 部屋割ありデフォルト
						GenderDiv: 0, // 男女区分なしデフォルト
					},
					AllotmentBookingReport: AllotmentBookingReport{
						TransactionType: TransactionType{
							DataFrom:           "FromTravelAgency",
							DataClassification: "NewBookReport",
							DataID:             bookData.TlDataID,
						},
						AccommodationInformation: AccommodationInformation{ // 宿泊施設情報
							AccommodationName: store.Name + *store.BranchName,
							AccommodationCode: BookingSystemID,
						},
						SalesOfficeInformation: SalesOfficeInformation{
							SalesOfficeCompanyName: "平和台ホテルアプリ",
							SalesOfficeName:        "平和台ホテルアプリ",
							SalesOfficeCode:        "10000000",
						},
						BasicInformation: BasicInformation{
							TravelAgencyBookingNumber:  bookData.TlDataID,
							TravelAgencyBookingDate:    time.Now().Format("2006-01-02"),
							TravelAgencyBookingTime:    time.Now().Format("15:04:05"),
							GuestOrGroupNameSingleByte: guestNameKana,
							GuestOrGroupKanjiName:      guest.LastName + " " + guest.FirstName,
							GuestOrGroupPhoneNumber:    *guest.Tel,
							GuestOrGroupEmail:          guest.Mail,
							GuestOrGroupPostalCode:     *guest.ZipCode,
							GuestOrGroupAddress:        guest.Prefecture.String() + *guest.City + *guest.Address,
							CheckInDate:                util.DateToStrDate(bookData.StayFrom),
							CheckInTime:                string(bookData.CheckInTime),
							Nights:                     uint(bookData.StayTo.Sub(bookData.StayFrom).Hours() / 24),
							TotalRoomCount:             bookData.RoomCount,
							GrandTotalPaxCount:         bookData.Adult + bookData.Child,
							TotalPaxMaleCount:          bookData.Adult,
							TotalChildA70Count:         bookData.Child,
							MealCondition:              mealCondition,
							SpecificMealCondition:      specificMealCondition,
							PackageType:                "Package",
							PackagePlanCode:            plan.ID,
							PackagePlanName:            plan.Title,
							SpecialServiceRequest:      bookData.Note,
						},
						BasicRateInformation: BasicRateInformation{
							RoomRateOrPersonalRate:   RoomRatePersonal,
							TaxServiceFee:            IncludingServiceAndTax,
							Payment:                  "Cash", //客の宿泊施設に対する支払い方法　事前カード決済などの場合は省略
							SettlementDiv:            0,      // 現地決済
							TotalAccommodationCharge: int(bookData.TotalCost),
							PointsDiscountList: PointsDiscountList{
								PointsDiscount: 0, // ポイント値引き
							},
							AmountClaimed: int(bookData.TotalCost),
						},
						RoomInformationList: RoomInformationList{
							RoomAndGuestList: RoomAndGuestList,
						},
					},
				},
			},
		},
	}
}

func NewCancelRQ(bookData *entity.Booking, store *entity.StayableStore, dataID string) *EnvelopeRQ[CancelBody] {
	var BookingSystemID string
	if env.GetEnv(env.TlbookingIsTest) == "true" {
		BookingSystemID = "E69502"
	} else {
		BookingSystemID = store.BookingSystemID
	}
	return &EnvelopeRQ[CancelBody]{
		SoapEnv: "http://schemas.xmlsoap.org/soap/envelope/",
		Naif:    "http://naifc3000.naifc30.nai.lincoln.seanuts.co.jp/",
		Header:  "",
		Body: CancelBody{
			DeleteBookingWithCP: DeleteBookingWithCP{
				DeleteBookingWithCPRequest: DeleteBookingWithCPRequest{
					CommonRequest: CommonRequest{
						AgtID:       TLBookingUser,
						AgtPassword: TLBookingPass,
						SystemDate:  time.Now().Format("2006-01-02T15:04:05"),
					},
					BookingInfo: BookingInfo{
						TllHotelCode:     BookingSystemID,
						TllBookingNumber: *bookData.TlBookingNumber,
						DataID:           dataID,
					},
				},
			},
		},
	}
}
