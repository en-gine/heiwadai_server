package booking

import (
	"bytes"
	"compress/gzip"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"server/infrastructure/booking/common"
	"server/infrastructure/logger"
)

var (
	user    = os.Getenv("TLBOOKING_USERNAME")
	pass    = os.Getenv("TLBOOKING_PASSWORD")
	envMode = os.Getenv("ENV_MODE")
)

func Request[TRequestType any, TResultType any](reqBody *TRequestType) (*TResultType, error) {
	if user == "" || pass == "" {
		return nil, errors.New("予約サーバーの認証情報が設定されていません。")
	}

	reqEnv := common.NewEnvelope[TRequestType](*reqBody, user, pass)

	out, _ := xml.MarshalIndent(reqEnv, " ", "  ")
	body := xml.Header + string(out)
	// body = "<?xml version='1.0' encoding='UTF-8'?><soapenv:Envelope xmlns:soapenv='http://schemas.xmlsoap.org/soap/envelope/' xmlns:head='http://www.seanuts.co.jp/ota/header' xmlns:ns='http://www.opentravel.org/OTA/2003/05'><soapenv:Header><head:Interface><head:PayloadInfo><head:CommDescriptor><head:Authentication><head:Username>B3959709</head:Username><head:Password>@.rhTm9sO6mp</head:Password></head:Authentication></head:CommDescriptor></head:PayloadInfo></head:Interface></soapenv:Header><soapenv:Body><ns:OTA_HotelAvailRQ Version='1.0' PrimaryLangID='jpn' RateDetailsInd='false'><ns:AvailRequestSegments><ns:AvailRequestSegment><ns:HotelSearchCriteria><ns:Criterion><ns:HotelRef HotelCode='E69502' /><ns:StayDateRange Start='2021-01-01' End='2021-01-02' /><ns:RateRange MinRate='3000' MaxRate='10000' /><ns:RatePlanCandidates><ns:RatePlanCandidate><ns:MealsIncluded Breakfast='true' Lunch='false' Dinner='false' /></ns:RatePlanCandidate></ns:RatePlanCandidates><ns:RoomStayCandidates><ns:RoomStayCandidate BedTypeCode='9' NonSmoking='true' Quantity='1'><ns:GuestCounts><ns:GuestCount AgeQualifyingCode='51' Count='1' /></ns:GuestCounts></ns:RoomStayCandidate></ns:RoomStayCandidates></ns:Criterion></ns:HotelSearchCriteria></ns:AvailRequestSegment></ns:AvailRequestSegments></ns:OTA_HotelAvailRQ></soapenv:Body></soapenv:Envelope>"
	// body = "<?xml version='1.0' encoding='UTF-8'?><soapenv:Envelope xmlns:soapenv='http://schemas.xmlsoap.org/soap/envelope/' xmlns:head='http://www.seanuts.co.jp/ota/header' xmlns:ns='http://www.opentravel.org/OTA/2003/05'><soapenv:Header><head:Interface><head:PayloadInfo><head:CommDescriptor><head:Authentication><head:Username>B3959709</head:Username><head:Password>@.rhTm9sO6mp</head:Password></head:Authentication></head:CommDescriptor></head:PayloadInfo></head:Interface></soapenv:Header><soapenv:Body><ns:OTA_HotelAvailRQ Version='1.0' PrimaryLangID='jpn' RateDetailsInd='false'><ns:AvailRequestSegments><ns:AvailRequestSegment><ns:HotelSearchCriteria><ns:Criterion><ns:HotelRef HotelCode='E69502' /><ns:StayDateRange Start='2021-01-01' End='2021-01-02' /><ns:RateRange MinRate='3000' MaxRate='10000' /><ns:RatePlanCandidates><ns:RatePlanCandidate><ns:MealsIncluded Breakfast='true' Dinner='false' /></ns:RatePlanCandidate></ns:RatePlanCandidates><ns:RoomStayCandidates><ns:RoomStayCandidate BedTypeCode='9' NonSmoking='true' Quantity='1'><ns:GuestCounts><ns:GuestCount AgeQualifyingCode='51' Count='1' /></ns:GuestCounts></ns:RoomStayCandidate></ns:RoomStayCandidates></ns:Criterion></ns:HotelSearchCriteria></ns:AvailRequestSegment></ns:AvailRequestSegments></ns:OTA_HotelAvailRQ></soapenv:Body></soapenv:Envelope>"
	url := "https://test472.tl-lincoln.net/agtapi/v1/crs/CrsAvailableInquiryService"
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		logger.Error(err.Error())
		return nil, errors.New("予約サーバーへの接続に失敗しました。")
	}
	req.Header.Add("Content-Type", "text/xml; charset=utf-8")
	req.Header.Add("Accept-Encoding", "gzip")
	if envMode == "dev" {
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

	result := *new(TResultType)
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

	content, _ := io.ReadAll(reader)

	if envMode == "dev" {
		fmt.Print(string(content))
	}

	if res.StatusCode != http.StatusOK {
		logger.Errorf("http status code: %d", res.StatusCode)
		logger.Error(string(content))
		return nil, errors.New("予約サーバーがエラーレスポンスを返しました。")
	}

	err = xml.Unmarshal(content, result)

	if err != nil {
		logger.Errorf("XML Unmarshal error: %s", err)
		logger.Error(string(content))
		return nil, errors.New("予約サーバーからのレスポンスの解析に失敗しました。")
	}

	return &result, nil
}
