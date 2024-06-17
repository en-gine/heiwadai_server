package main

import (
	"fmt"
	"time"

	"server/core/entity"
	"server/infrastructure/booking/avail"
	"server/infrastructure/booking/book"
	implements "server/infrastructure/repository"

	"github.com/google/uuid"
)

func main() {
	// Search()
	Book()
}

func Book() {
	storeQuery := implements.NewStoreQueryService()
	bookQuery := implements.NewBookQueryService()
	stores, err := storeQuery.GetStayables()
	if err != nil {
		fmt.Print(err)
		return
	}
	p := book.NewBookRepository(storeQuery, bookQuery)
	tomorrow := time.Now().Add(2 * 24 * time.Hour)
	meal := entity.MealType{Morning: true, Dinner: true}
	smork := entity.SmokeTypeSmoking

	// ダミーデータ
	companyName := "株式会社サンプル"
	zipCode := "100-0001"
	city := "千代田区"
	address := "丸の内1-1-1"
	tel := "03-1234-5678"
	pref := entity.Prefecture(13)

	guest := &entity.GuestData{
		FirstName:     "太郎",
		LastName:      "田中",
		FirstNameKana: "タロウ",
		LastNameKana:  "タナカ",
		CompanyName:   &companyName,
		ZipCode:       &zipCode,
		Prefecture:    &pref,
		City:          &city,
		Address:       &address,
		Tel:           &tel,
		Mail:          "tanaka.taro@example.com",
	}
	// プラン
	plan := entity.RegenPlan(
		"12354569",
		"いい温泉宿プラン",
		10000,
		"https://yahoo.co.jp",
		entity.RoomTypeDouble,
		meal,
		smork,
		"広々とした部屋です。",
		stores[0].ID,
		"1",
	)

	bookData := entity.CreateBooking(
		time.Now(),
		tomorrow,
		uint(1), // adult
		uint(1), // child
		uint(1), // roomCount
		"18:00",
		10000,
		guest,
		plan,
		uuid.New(),
		"駐車場利用します",
		"123456789",
		nil,
	)

	reserveID, domainErr, err := p.Reserve(bookData)
	if err != nil {
		fmt.Print(err)
		return
	}
	if domainErr != nil {
		fmt.Print(domainErr)
		return
	}
	fmt.Print(reserveID)
}

func Search() {
	storeQuery := implements.NewStoreQueryService()
	p := avail.NewPlanQuery(storeQuery)

	tomorrow := time.Now().Add(2 * 24 * time.Hour)
	single := entity.RoomTypeSingle
	double := entity.RoomTypeDouble
	rooms := []entity.RoomType{single, double}
	meal := entity.MealType{Morning: true, Dinner: true}
	smork := entity.SmokeTypeSmoking
	nonSmoke := entity.SmokeTypeNonSmoking
	smokes := []entity.SmokeType{smork, nonSmoke}
	plans, err := p.Search(
		nil,
		time.Now(),
		tomorrow,
		1,
		1,
		1,
		smokes,
		meal,
		rooms,
	)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Print(plans)
}
