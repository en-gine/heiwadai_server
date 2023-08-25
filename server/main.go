package main

import (
	"fmt"
	"os"

	adminRouter "server/router/admin"
	userRouter "server/router/user"
)

func main() {
	// booking.PlanQueryTest()
	userRouter.NewUserServer()
	adminRouter.NewAdminServer()

	msg := os.ExpandEnv("${ENV} mode run! port: ${SERVER_PORT}")
	fmt.Println(msg)

}
