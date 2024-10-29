package api

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"server/infrastructure/logger"
	"server/infrastructure/redis"
)

var memory = redis.NewMemoryRepository()

func FetchJSONData[T any](APIURL string) (*T, error) {
	res, err := http.Get(APIURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		// 外部APIからデータを取得できなかった場合エラーだとアプリケーションが停止するため空を返す
		logger.Warn("Failed to fetch data : " + string(body))
		return nil, nil
	}

	var data T
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func Request[T any](APIURL string, cacheKey string, cacheExpire time.Duration) (*T, error) {
	var data *T

	cache := memory.Get(cacheKey)
	// キャッシュがあればそれを返す
	if cache != nil {
		err := json.Unmarshal(*cache, &data)
		if err != nil {
			return nil, err
		}
		return data, nil
	} else {
		data, err := FetchJSONData[T](APIURL)
		if err != nil {
			return nil, err
		}
		// memoryにキャッシュする
		cacheData, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		memory.Set(cacheKey, cacheData, cacheExpire)

		return data, nil
	}
}
