package main

import (
	"fmt"

	"server/infrastructure/logger"
	"server/router/user"
)

func main() {
	Message()
}

func Message() {
	usecase := user.InitializeMessageUsecase()

	// msgID := uuid.MustParse("72d973f1-c42d-4490-aded-ba454884acf6")
	messages, err := usecase.GetAfter(nil)
	if err != nil {
		logger.Error(err.Error())
	}

	for _, msg := range messages {
		fmt.Printf("%+v\n", msg)
	}
}
