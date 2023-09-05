package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	userRouter "server/router/user"
)

func OnlyServer() {
	// start server
	http.HandleFunc("/", Handler)
	fmt.Println("Server started on :3000")

	log.Fatal(http.ListenAndServe(":3000", nil))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// GETパラメーターを取得
	values := r.URL.Query()

	// 取得したGETパラメーターを表示
	for key, value := range values {
		fmt.Printf("Key: %s, Value: %s\n", key, value[0])
		w.Write([]byte(key + ":" + value[0]))
		if key == "access_token" && value[0] != "" {
			token = value[0]
			go SetPassword(token)
		}
	}
}

func SetPassword(token string) {
	authUsecase := userRouter.InitializeAuthUsecase()
	password := os.Getenv("TEST_USER_PASS")
	err := authUsecase.UpdatePassword(password, token)
	if err != nil {
		panic(err)
	}

}
