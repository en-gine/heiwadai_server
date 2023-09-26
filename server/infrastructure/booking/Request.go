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

var user = os.Getenv("TLBOOKING_USERNAME")
var pass = os.Getenv("TLBOOKING_PASSWORD")
var envMode = os.Getenv("ENV_MODE")

func Request[TRequestType any, TResultType any](reqBody *TRequestType) (*TResultType, error) {
	if user == "" || pass == "" {
		return nil, errors.New("予約サーバーの認証情報が設定されていません。")
	}

	reqEnv := common.NewEnvelope[TRequestType](*reqBody, user, pass)

	out, _ := xml.MarshalIndent(reqEnv, " ", "  ")
	body := xml.Header + string(out)
	url := "https://test472.tl-lincoln.net/agtapi/v1/crs/CrsAvailableInquiryService"
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

	if envMode == "debug" {
		fmt.Print(string(content))
	}

	if res.StatusCode != http.StatusCreated {
		logger.Errorf("http status code: %d", res.StatusCode)
		logger.Error(string(content))
		return nil, errors.New("予約サーバーがエラーレスポンスを返しました。")
	}

	err = xml.Unmarshal(content, result)

	if err != nil {
		logger.Errorf("XML Unmarshal error: %s", err)
		return nil, errors.New("予約サーバーからのレスポンスの解析に失敗しました。")
	}

	return &result, nil
}
