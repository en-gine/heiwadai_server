package booking

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"server/infrastructure/logger"
)

func Request[TRequestType any, TResultType any](reqBody *TRequestType) (*TResultType, error) {

	out, _ := xml.MarshalIndent(reqBody, " ", "  ")
	body := xml.Header + string(out)
	url := "https://test472.tl-lincoln.net/agtapi/v1/crs/CrsAvailableInquiryService"
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(body)))

	if err != nil {
		logger.Error(err.Error())
		return nil, errors.New("予約サーバーへの接続に失敗しました。")
	}
	req.Header.Add("Content-Type", "text/xml; charset=utf-8")
	req.Header.Add("Accept-Encoding", "gzip")

	fmt.Print(req.Body)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		logger.Error(err.Error())
		return nil, errors.New("予約サーバーへのリクエストに失敗しました。")
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		logger.Errorf("http status code: %d", res.StatusCode)
		return nil, errors.New("予約サーバーがエラーレスポンスを返しました。")
	}

	result := new(TResultType)
	content, _ := io.ReadAll(res.Body)

	err = xml.Unmarshal(content, result)

	if err != nil {
		logger.Errorf("XML Unmarshal error: %s", err)
		return nil, errors.New("予約サーバーからのレスポンスの解析に失敗しました。")
	}
	fmt.Print(result)

	return result, nil
}
