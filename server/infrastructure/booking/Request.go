package booking

import (
	"bytes"
	"compress/gzip"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"

	"server/infrastructure/booking/common"
	"server/infrastructure/env"
	"server/infrastructure/logger"
)

var (
	user    = env.GetEnv(env.TlbookingUsername)
	pass    = env.GetEnv(env.TlbookingPassword)
	envMode = env.GetEnv(env.EnvMode)
)

func Request[TRequestType any, TResultType any](url string, reqBody *TRequestType) (*TResultType, error) {
	if user == "" || pass == "" {
		return nil, errors.New("予約サーバーの認証情報が設定されていません。")
	}

	reqEnv := common.NewEnvelopeRQ[TRequestType](*reqBody, user, pass)
	out, err := xml.MarshalIndent(reqEnv, " ", "  ")
	if err != nil {
		logger.Error(err.Error())
		return nil, errors.New("リクエストの解析に失敗しました。")
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
		// fmt.Print(string(content))
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
