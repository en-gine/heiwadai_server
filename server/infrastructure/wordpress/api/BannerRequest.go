package api

import (
	"time"

	"server/infrastructure/wordpress/types"
)

var BANNERPOSTURL = "https://www.heiwadai-hotel.co.jp/wp-json/app/v1/slider"

func GetWPBanners() (*[]types.WPBanner, error) {
	CacheKey := "wp_banners_cache"
	CacheExpiry := 60 * time.Minute * 2

	posts, err := Request[[]types.WPBanner](BANNERPOSTURL, CacheKey, CacheExpiry)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
