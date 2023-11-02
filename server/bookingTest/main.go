package main

import (
	"fmt"
	"time"

	"server/core/entity"
	"server/infrastructure/booking"
	implements "server/infrastructure/repository"

	"github.com/google/uuid"
)

func main() {
	// Search()
	// Book()
}

func Book() {
	storeQuery := implements.NewStoreQueryService()
	stores, err := storeQuery.GetStayables()
	if err != nil {
		fmt.Print(err)
		return
	}
	p := booking.NewPlanRepository(storeQuery)
	tomorrow := time.Now().Add(2 * 24 * time.Hour)
	meal := entity.MealType{Morning: true, Dinner: true}
	smork := entity.SmokeTypeSmoking

	// ダミーデータ
	companyName := "株式会社サンプル"
	zipCode := "100-0001"
	city := "千代田区"
	address := "丸の内1-1-1"
	tel := "03-1234-5678"

	guest := &entity.GuestData{
		FirstName:     "太郎",
		LastName:      "田中",
		FirstNameKana: "タロウ",
		LastNameKana:  "タナカ",
		CompanyName:   &companyName,
		BirthDate:     time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
		ZipCode:       &zipCode,
		Prefecture:    entity.Prefecture(13),
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
		*&stores[0].ID,
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
		"駐車場利用します")
	err = p.Book(bookData, "20231102000000001")
	if err != nil {
		fmt.Print(err)
		return
	}
}

func Search() {
	storeQuery := implements.NewStoreQueryService()
	p := booking.NewPlanQuery(storeQuery)

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
		&smokes,
		&meal,
		&rooms,
	)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Print(plans)
}
