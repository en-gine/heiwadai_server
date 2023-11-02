package booking

import (
	"errors"
	"fmt"
	"time"

	"server/core/entity"
	queryservice "server/core/infra/queryService"
	"server/infrastructure/booking/book"
	"server/infrastructure/booking/util"
	"server/infrastructure/env"
	"server/infrastructure/logger"

	uuid "github.com/google/uuid"
	"github.com/ktnyt/go-moji"
)

type PlanRepository struct {
	storeQuery queryservice.IStoreQueryService
}

var BookURL = env.GetEnv(env.TlbookingBookingApiUrl)

func NewPlanRepository(storeQuery queryservice.IStoreQueryService) *PlanRepository {
	return &PlanRepository{
		storeQuery: storeQuery,
	}
}

func (p *PlanRepository) GetMyBooking(userID uuid.UUID) (*[]entity.Plan, error) {
	return nil, nil
}

func (p *PlanRepository) Book(
	bookData *entity.Booking,
	nextBookID string,
) error {
	store, err := p.storeQuery.GetStayableByID(bookData.BookPlan.StoreID)
	if err != nil {
		return err
	}

	reqBody := NewBookingRQ(
		"testID", "testPass", bookData, store, nextBookID,
	)

	res, err := Request[book.EnvelopeRQ, book.EnvelopeRS](BookURL, reqBody)
	if err != nil {
		return err
	}

	if res.Body.EntryBookingResponse.EntryBookingResult.CommonResponse.ResultCode == "False" {
		msg := res.Body.EntryBookingResponse.EntryBookingResult.CommonResponse.ErrorInfos.ErrorMsg
		code := res.Body.EntryBookingResponse.EntryBookingResult.CommonResponse.ErrorInfos.ErrorCode
		logger.Error(code + ":" + msg)
		return errors.New(msg)
	}

	bookSystemID := res.Body.EntryBookingResponse.EntryBookingResult.ExtendLincoln.TllBookingNumber
	fmt.Print(bookSystemID)
	return nil
}

// bookID : CCYYMMDD+9桁連番（0埋め、データ毎に+1）
func NewBookingRQ(agentID string, agentPass string, bookData *entity.Booking, store *entity.StayableStore, nextBookID string) *book.EnvelopeRQ {
	plan := bookData.BookPlan
	guest := bookData.GuestData

	var mealCondition book.MealCondition

	switch plan.MealType.String() {
	case "朝食あり夕食あり":
		mealCondition = book.MealCondition1night2meals
	case "朝食あり夕食なし":
		mealCondition = book.MealCondition1nightBreakfast
	case "朝食なし夕食あり":
		mealCondition = book.MealCondition1nightBreakfast
	case "食事なし":
		mealCondition = book.MealConditionWithoutMeal
	default:
		mealCondition = book.MealConditionOther
	}

	guestNameKana := moji.Convert(guest.LastNameKana, moji.ZK, moji.HK) + " " + moji.Convert(guest.FirstNameKana, moji.ZK, moji.HK)
	return &book.EnvelopeRQ{
		SoapEnv: "http://schemas.xmlsoap.org/soap/envelope/",
		Naif:    "http://naifc3000.naifc30.nai.lincoln.seanuts.co.jp/",
		Header:  "",
		Body: book.Body{
			EntryBooking: book.EntryBooking{
				EntryBookingRequest: book.EntryBookingRequest{
					CommonRequest: book.CommonRequest{
						AgtID:       agentID,
						AgtPassword: agentPass,
						SystemDate:  time.Now().Format("2006-01-02T15:04:05"),
					},
					ExtendLincoln: book.ExtendLincoln{
						TllHotelCode: store.BookingSystemID,
						UseTllPlan:   1,
					},
					SendInformation: book.SendInformation{
						AssignDiv: 1, // 部屋割ありデフォルト
						GenderDiv: 0, // 男女区分なしデフォルト
					},
					AllotmentBookingReport: book.AllotmentBookingReport{
						TransactionType: book.TransactionType{
							DataFrom:           "FromTravelAgency",
							DataClassification: "NewBookReport",
							DataID:             nextBookID,
						},
						AccommodationInformation: book.AccommodationInformation{ // 宿泊施設情報
							AccommodationName: store.Name + *store.BranchName,
							AccommodationCode: store.BookingSystemID,
						},
						SalesOfficeInformation: book.SalesOfficeInformation{
							SalesOfficeCompanyName: "平和台ホテルアプリ",
							SalesOfficeName:        "",
							SalesOfficeCode:        "10000000",
						},
						BasicInformation: book.BasicInformation{
							TravelAgencyBookingNumber:  nextBookID,
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
						BasicRateInformation: book.BasicRateInformation{
							RoomRateOrPersonalRate:   book.RoomRateRoom,
							TaxServiceFee:            book.IncludingServiceAndTax,
							Payment:                  "Cash",
							SettlementDiv:            0, // 現地決済
							TotalAccommodationCharge: int(bookData.TotalCost),
							PointsDiscountList: book.PointsDiscountList{
								PointsDiscount: 0, // ポイント値引き
							},
						},
						RoomInformationList: book.RoomInformationList{
							RoomAndGuestList: []book.RoomAndGuest{
								{
									RoomInformation: book.RoomInformation{
										RoomTypeCode:    plan.RoomType.Code(),
										RoomTypeName:    plan.RoomType.String(),
										PerRoomPaxCount: bookData.Adult + bookData.Child,
									},
									RoomRateInformation: book.RoomRateInformation{
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
