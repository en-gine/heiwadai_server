package main

import (
	"fmt"

	"server/infrastructure/logger"
	"server/router/user"
)

func main() {
	Post()
	// PostAll()
	// Banner()
}

func PostAll() {
	usecase := user.InitializePostUsecase()

	posts, err := usecase.GetList()
	if err != nil {
		logger.Error(err.Error())
	}

	for _, post := range posts {
		fmt.Printf("%+v\n", post)
	}

	fmt.Println(posts[0].ID)
	// post, err := usecase.GetByID(uint32(posts[0].ID))
	// if err != nil {
	// 	logger.Error(err.Error())
	// }
	// fmt.Printf("%+v\n", post)
}

func Post() {
	usecase := user.InitializePostUsecase()

	post, err := usecase.GetByID(uint32(3076))
	if err != nil {
		logger.Error(err.Error())
	}
	fmt.Printf("%+v\n", post)
}

func Banner() {
	usecase := user.InitializeBannerUsecase()

	banners, err := usecase.GetList()
	if err != nil {
		logger.Error(err.Error())
	}

	for _, banner := range banners {
		fmt.Printf("%+v\n", banner)
	}
}
