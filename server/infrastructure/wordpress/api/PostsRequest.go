package api

import (
	"fmt"
	"strconv"
	"time"

	"server/infrastructure/wordpress/types"
)

var WPPOSTURL = "https://www.heiwadai-hotel.co.jp/wp-json/app/v1/all_posts/"

func GetWPPosts() (*[]types.WPALLPost, error) {
	CacheKey := "wp_posts_cache"
	APIURL := WPPOSTURL + "?_embed"
	CacheExpiry := 60 * time.Minute

	posts, err := Request[[]types.WPALLPost](APIURL, CacheKey, CacheExpiry)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func GetWPPost(id int) (*types.WPPost, error) {
	CacheKey := "wp_post_cache_" + strconv.Itoa(id)
	APIURL := WPPOSTURL + strconv.Itoa(id) + "?_embed"
	fmt.Println(APIURL)
	CacheExpiry := 60 * time.Minute

	post, err := Request[types.WPPost](APIURL, CacheKey, CacheExpiry)
	if err != nil {
		return nil, err
	}
	return post, nil
}
