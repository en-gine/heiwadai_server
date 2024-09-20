package main

import (
	"fmt"
	adminRouter "server/router/admin"
)

func StoreSeeder() {
	storeUsecase := adminRouter.InitializeStoreUsecase()

	// 天神

	branchname := "天神"
	parking := "8台（予約制）"
	latitude := 33.5933853
	longitude := 130.3926853
	accessInfo := "天神駅徒歩8分"
	bookingSystemID := "BESS1"
	bookingSystemLoginID := "BESS1"
	bookingSystemPassword := "BESS1"
	restAPIURL := "https://www.heiwadai-hotel.co.jp/tenjin/wp-json/wp/v2/posts?_embed"
	store, error := storeUsecase.Create(
		"平和台ホテル",
		&branchname,
		"810-0054",
		"福岡県福岡市中央区舞鶴1-5-6",
		"092-737-1000",
		"https://www.heiwadai-hotel.co.jp/tenjin/",
		"https://chbqhfrawgjohpgennle.supabase.co/storage/v1/object/public/public/stores/stamp_tenjin.png?t=2023-08-01T06%3A45%3A05.255Z",
		true,
		&parking,
		&latitude,
		&longitude,
		&accessInfo,
		&restAPIURL,
		&bookingSystemID,
		&bookingSystemLoginID,
		&bookingSystemPassword,
	)
	if error != nil {
		panic(error)
	}
	if store != nil {
		fmt.Println(store)
	}

	// 5
	branchname = "5"
	parking = "8台（予約制）"
	latitude = 33.5890682
	longitude = 130.3685846
	accessInfo = "地下鉄唐人町駅すぐ"
	restAPIURL = "https://www.heiwadai-hotel.co.jp/five/wp-json/wp/v2/posts?_embed"
	bookingSystemID = "BESS1"
	bookingSystemLoginID = "BESS1"
	bookingSystemPassword = "BESS1"
	store, error = storeUsecase.Create(
		"平和台ホテル",
		&branchname,
		"810-0073",
		"福岡県福岡市中央区今川1丁目4-2",
		"092-732-5000",
		"https://www.heiwadai-hotel.co.jp/five/",
		"https://chbqhfrawgjohpgennle.supabase.co/storage/v1/object/public/public/stores/stamp_five.png?t=2023-08-01T06%3A45%3A05.255Z",
		true,
		&parking,
		&latitude,
		&longitude,
		&accessInfo,
		&restAPIURL,
		&bookingSystemID,
		&bookingSystemLoginID,
		&bookingSystemPassword,
	)
	if error != nil {
		panic(error)
	}
	if store != nil {
		fmt.Println(store)
	}

	branchname = "荒戸"
	parking = "6台（予約制）"
	latitude = 33.5920984
	longitude = 130.3790031
	accessInfo = "大濠公園駅徒歩約3分"
	restAPIURL = "https://www.heiwadai-hotel.co.jp/arato/wp-json/wp/v2/posts?_embed"
	bookingSystemID = "BESS1"
	bookingSystemLoginID = "BESS1"
	bookingSystemPassword = "BESS1"

	store, error = storeUsecase.Create(
		"平和台ホテル",
		&branchname,
		"810-0062",
		"福岡県福岡市中央区荒戸1丁目5-27",
		"092-761-1361",
		"https://www.heiwadai-hotel.co.jp/arato/",
		"https://chbqhfrawgjohpgennle.supabase.co/storage/v1/object/public/public/stores/stamp_five.png?t=2023-08-01T06%3A45%3A05.255Z",
		true,
		&parking,
		&latitude,
		&longitude,
		&accessInfo,
		&restAPIURL,
		&bookingSystemID,
		&bookingSystemLoginID,
		&bookingSystemPassword,
	)
	if error != nil {
		panic(error)
	}
	if store != nil {
		fmt.Println(store)
	}

	branchname = "大手門"
	parking = "4台（予約制）"
	latitude = 33.5889767
	longitude = 130.3838308
	accessInfo = "大濠公園駅徒歩約6分"
	restAPIURL = "https://www.heiwadai-hotel.co.jp/otemon/wp-json/wp/v2/posts?_embed"
	bookingSystemID = "BESS1"
	store, error = storeUsecase.Create(
		"平和台ホテル",
		&branchname,
		"810-0074",
		"福岡県福岡市中央区大手門1丁目5-4",
		"092-741-4422",
		"https://www.heiwadai-hotel.co.jp/otemon/",
		"https://chbqhfrawgjohpgennle.supabase.co/storage/v1/object/public/public/stores/stamp_five.png?t=2023-08-01T06%3A45%3A05.255Z",
		true,
		&parking,
		&latitude,
		&longitude,
		&accessInfo,
		&restAPIURL,
		&bookingSystemID,
		&bookingSystemLoginID,
		&bookingSystemPassword,
	)
	if error != nil {
		panic(error)
	}
	if store != nil {
		fmt.Println(store)
	}
}
