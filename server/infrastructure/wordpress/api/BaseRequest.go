package api

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"server/infrastructure/logger"
	"server/infrastructure/redis"
)

var rdb, _ = redis.NewMemoryRepository()

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

	var data T
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func Request[T any](APIURL string, cacheKey string, cacheExpire time.Duration) (*T, error) {
	var data *T

	cache, err := rdb.Get(cacheKey)
	if err != nil {
		return nil, err
	}
	// キャッシュがあればそれを返す
	if cache != nil {
		err = json.Unmarshal(*cache, &data)
		if err != nil {
			return nil, err
		}
		return data, nil
	} else {
		data, err = FetchJSONData[T](APIURL)
		if err != nil {
			return nil, err
		}
		// rdbにキャッシュする
		cacheData, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		err = rdb.Set(cacheKey, cacheData, cacheExpire)
		if err != nil {
			logger.Errorf("Failed to set cache : %s", err)
			return nil, err
		}
		return data, nil
	}
}
