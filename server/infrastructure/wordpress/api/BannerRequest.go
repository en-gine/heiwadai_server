package api

import (
	"server/infrastructure/wordpress/types"
	"time"
)

var BANNERPOSTURL = "https://www.heiwadai-hotel.co.jp/wp-json/app/v1/slider"

func GetWPBanners() (*[]types.WPBanner, error) {
	var CacheKey = "wp_banners_cache"
	var CacheExpiry = 60 * time.Minute * 2

	posts, err := Request[[]types.WPBanner](WPPOSTURL, CacheKey, CacheExpiry)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
