package booking

import (
	"bytes"
	"compress/gzip"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"

	"server/infrastructure/env"
	"server/infrastructure/logger"
)

var envMode = env.GetEnv(env.EnvMode)

func Request[TRequestType any, TResultType any](url string, reqBody *TRequestType) (*TResultType, error) {
	out, err := xml.MarshalIndent(reqBody, " ", "  ")
	if err != nil {
		logger.Error(err.Error())
		return nil, errors.New("リクエストの解析に失敗しました。")
	}

	// planeXml := "<soapenv:Envelope xmlns:soapenv='http://schemas.xmlsoap.org/soap/envelope/' xmlns:naif='http://naifc3000.naifc30.nai.lincoln.seanuts.co.jp/'> <soapenv:Header/> <soapenv:Body> <naif:entryBooking> <entryBookingRequest> <commonRequest> <agtId>N0000011</agtId> <agtPassword>????????</agtPassword> <systemDate>2016-12-05T18:06:13</systemDate> </commonRequest> <extendLincoln> <tllHotelCode>J48613</tllHotelCode> <useTllPlan>0</useTllPlan> </extendLincoln> <SendInformation> <assignDiv>1</assignDiv> <genderDiv>1</genderDiv> </SendInformation> <AllotmentBookingReport> <TransactionType> <DataFrom>FromTravelAgency</DataFrom> <DataClassification>NewBookReport</DataClassification> <DataID>20161205000000016</DataID> </TransactionType> <AccommodationInformation> <AccommodationArea></AccommodationArea> <AccommodationName>テスト施設</AccommodationName> <AccommodationCode>1001</AccommodationCode> <ChainName></ChainName> </AccommodationInformation> <SalesOfficeInformation> <SalesOfficeCompanyName>シーナッツトラベル</SalesOfficeCompanyName> <SalesOfficeName>テストシーナッツ支店</SalesOfficeName> <SalesOfficeCode>test472</SalesOfficeCode> <SalesOfficePersonInCharge></SalesOfficePersonInCharge> <SalesOfficeEmail></SalesOfficeEmail> <SalesOfficePhoneNumber>03-5404-6703</SalesOfficePhoneNumber> <SalesOfficeFaxNumber>03-5404-6706</SalesOfficeFaxNumber> </SalesOfficeInformation> <BasicInformation> <TravelAgencyBookingNumber>10001</TravelAgencyBookingNumber> <TravelAgencyBookingDate>2016-12-05</TravelAgencyBookingDate> <TravelAgencyBookingTime>18:06:13</TravelAgencyBookingTime> <GuestOrGroupMiddleName></GuestOrGroupMiddleName> <GuestOrGroupNameSingleByte>ｼｲﾅ ﾀﾛｳ</GuestOrGroupNameSingleByte> <GuestOrGroupNameDoubleByte></GuestOrGroupNameDoubleByte> <GuestOrGroupKanjiName>椎名 太郎</GuestOrGroupKanjiName> <GuestOrGroupContactDiv>1</GuestOrGroupContactDiv> <GuestOrGroupCellularNumber></GuestOrGroupCellularNumber> <GuestOrGroupOfficeNumber></GuestOrGroupOfficeNumber> <GuestOrGroupPhoneNumber>03-1111-2222</GuestOrGroupPhoneNumber> <GuestOrGroupEmail>shina@xxx.ne.jp</GuestOrGroupEmail> <GuestOrGroupPostalCode>105-0021</GuestOrGroupPostalCode> <GuestOrGroupAddress>東京都港区東新橋2-3-3 ルオーゴ汐留8F</GuestOrGroupAddress> <GroupNameWelcomeBoard></GroupNameWelcomeBoard> <GuestGenderDiv></GuestGenderDiv> <GuestGeneration></GuestGeneration> <GuestAge></GuestAge> <CheckInDate>2017-06-05</CheckInDate> <CheckInTime>17:00</CheckInTime> <CheckOutDate>2017-06-06</CheckOutDate> <CheckOutTime>10:00</CheckOutTime> <Nights>1</Nights> <Transportaion>Car</Transportaion> <TotalRoomCount>1</TotalRoomCount> <GrandTotalPaxCount>8</GrandTotalPaxCount> <TotalPaxMaleCount>1</TotalPaxMaleCount> <TotalPaxFemaleCount>1</TotalPaxFemaleCount> <TotalChildA70Count>1</TotalChildA70Count> <TotalChildA70Count2>1</TotalChildA70Count2> <TotalChildB50Count>1</TotalChildB50Count> <TotalChildB50Count2>1</TotalChildB50Count2> <TotalChildC30Count>1</TotalChildC30Count> <TotalChildDNoneCount>1</TotalChildDNoneCount> <TypeOfGroupDoubleByte></TypeOfGroupDoubleByte> <PackageType>Package</PackageType> <PackagePlanName>基本プラン（２食付き）</PackagePlanName> <PackagePlanCode>16143474</PackagePlanCode> <PackagePlanContent></PackagePlanContent> <MealCondition>1night2meals</MealCondition> <SpecificMealCondition>IncludingBreakfastAndDinner</SpecificMealCondition> <ModificationPoint></ModificationPoint> <SpecialServiceRequest>駐車場利用します。</SpecialServiceRequest> <OtherServiceInformation></OtherServiceInformation> <SalesOfficeComment></SalesOfficeComment> <QuestionAndAnswerList> <FromHotelQuestion>夕食のご希望時間をお知らせください。</FromHotelQuestion> <ToHotelAnswer>19時頃からでお願いします。</ToHotelAnswer> </QuestionAndAnswerList> </BasicInformation> <BasicRateInformation> <RoomRateOrPersonalRate>PersonalRate</RoomRateOrPersonalRate> <TaxServiceFee>IncludingServiceAndTax</TaxServiceFee> <Payment></Payment> <SettlementDiv>2</SettlementDiv> <TotalAccommodationCharge>74000</TotalAccommodationCharge> <TotalAccommodationConsumptionTax></TotalAccommodationConsumptionTax> <TotalAccommodationHotSpringTax></TotalAccommodationHotSpringTax> <TotalAccomodationServiceCharge></TotalAccomodationServiceCharge> <TotalAccommodationDiscountPoints>72000</TotalAccommodationDiscountPoints> <TotalAccommodationConsumptionTaxAfterDiscountPoints></TotalAccommodationConsumptionTaxAfterDiscountPoints> <AmountClaimed>0</AmountClaimed> <PointsDiscountList> <PointsDiv>2</PointsDiv> <PointsDiscountName>ポイント利用</PointsDiscountName> <PointsDiscount>2000</PointsDiscount> </PointsDiscountList> </BasicRateInformation> <MemberInformation> <MemberName>ｼｲﾅ ｲﾁﾛｳ</MemberName> <MemberKanjiName>椎名 一郎</MemberKanjiName> <MemberMiddleName></MemberMiddleName> <MemberDateOfBirth></MemberDateOfBirth> <MemberEmergencyNumber></MemberEmergencyNumber> <MemberOccupation></MemberOccupation> <MemberOrganization></MemberOrganization> <MemberOrganizationKana></MemberOrganizationKana> <MemberOrganizationCode></MemberOrganizationCode><MemberPosition></MemberPosition> <MemberOfficePostalCode></MemberOfficePostalCode> <MemberOfficeAddress></MemberOfficeAddress> <MemberOfficeTelephoneNumber></MemberOfficeTelephoneNumber> <MemberOfficeFaxNumber></MemberOfficeFaxNumber> <MemberGenderDiv></MemberGenderDiv> <MemberClass></MemberClass> <CurrentPoints></CurrentPoints> <MailDemandDiv></MailDemandDiv> <PamphletDemandDiv></PamphletDemandDiv> <MemberID>43256</MemberID> <MemberPhoneNumber>03-5555-6666</MemberPhoneNumber> <MemberEmail>ichiro@xxx.ne.jp</MemberEmail> <MemberPostalCode></MemberPostalCode> <MemberAddress></MemberAddress> </MemberInformation> <OptionInformation> <OptionList> <OptionDate>2017-06-05</OptionDate> <OptionCode>001</OptionCode> <Name>追加料理</Name> <NameRequest>舟盛り</NameRequest> <OptionCount>2</OptionCount> <OptionRate>5000</OptionRate> </OptionList> </OptionInformation> <RoomInformationList> <RoomAndGuestList> <RoomAndGuest> <RoomInformation> <RoomTypeCode>2</RoomTypeCode> <RoomTypeName>和室12畳</RoomTypeName> <PerRoomPaxCount>8</PerRoomPaxCount> </RoomInformation> <RoomRateInformation> <RoomDate>2017-06-05</RoomDate> <PerPaxRate>15000</PerPaxRate> <PerPaxFemaleRate>15000</PerPaxFemaleRate> <PerChildA70Rate>10000</PerChildA70Rate> <PerChildA70Rate2>9000</PerChildA70Rate2> <PerChildB50Rate>7000</PerChildB50Rate> <PerChildB50Rate2>6000</PerChildB50Rate2> <PerChildC30Rate>2000</PerChildC30Rate> <PerChildDNoneRate>0</PerChildDNoneRate> <RoomRatePaxMaleCount>1</RoomRatePaxMaleCount> <RoomRatePaxFemaleCount>1</RoomRatePaxFemaleCount> <RoomRateChildA70Count>1</RoomRateChildA70Count> <RoomRateChildA70Count2>1</RoomRateChildA70Count2> <RoomRateChildB50Count>1</RoomRateChildB50Count> <RoomRateChildB50Count2>1</RoomRateChildB50Count2> <RoomRateChildC30Count>1</RoomRateChildC30Count> <RoomRateChildDNoneCount>1</RoomRateChildDNoneCount> <RoomPaxMaleRequest></RoomPaxMaleRequest> <RoomPaxFemaleRequest></RoomPaxFemaleRequest> <RoomChildA70Request>小学校高学年</RoomChildA70Request> <RoomChildA70Request2>小学校低学年</RoomChildA70Request2> <RoomChildB50Request>幼児（食事・布団あり）</RoomChildB50Request> <RoomChildB50Request2>幼児（食事あり）</RoomChildB50Request2> <RoomChildC30Request>幼児（布団あり）</RoomChildC30Request> <RoomChildDNoneRequest>幼児（食事・布団なし）</RoomChildDNoneRequest> <TotalPerRoomRate></TotalPerRoomRate> </RoomRateInformation> </RoomAndGuest> </RoomAndGuestList> </RoomInformationList> </AllotmentBookingReport> </entryBookingRequest> </naif:entryBooking> </soapenv:Body></soapenv:Envelope>"
	// out = []byte(planeXml)

	if envMode == "dev" {
		fmt.Print(string(out))
	}

	body := xml.Header + string(out)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		logger.Error(err.Error())
		return nil, errors.New("予約サーバーへの接続に失敗しました。")
	}
	req.Header.Add("Content-Type", "text/xml; charset=utf-8")
	req.Header.Add("Accept-Encoding", "gzip")
	if envMode == "debug" {
		fmt.Print(req.Body)
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		logger.Error(err.Error())
		return nil, errors.New("予約サーバーへのリクエストに失敗しました。")
	}
	defer res.Body.Close()

	var reader io.Reader

	var result TResultType
	encoding := res.Header.Get("Content-Encoding")

	if encoding == "gzip" {
		reader, err = gzip.NewReader(res.Body)
		if err != nil {
			logger.Error(err.Error())
			return nil, errors.New("gzipデータのデコードに失敗しました。")
		}
	} else {
		reader = res.Body
	}

	content, err := io.ReadAll(reader)
	if err != nil {
		logger.Error(err.Error())
		return nil, errors.New("ストリームをデコード出来ませんでした。")
	}

	if envMode == "dev" {
		fmt.Print(string(content))
	}

	if res.StatusCode != http.StatusOK {
		logger.Errorf("http status code: %d", res.StatusCode)
		logger.Error(string(content))
		return nil, errors.New("予約サーバーがエラーレスポンスを返しました。")
	}

	err = xml.Unmarshal(content, &result)

	if err != nil {
		logger.Errorf("XML Unmarshal error: %s", err)
		logger.Error(string(content))
		return nil, errors.New("予約サーバーからのレスポンスの解析に失敗しました。")
	}
	// fmt.Printf("result: %+v", result)

	return &result, nil
}
