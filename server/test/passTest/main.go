package main

import (
	"server/core/entity"
)

func main() {
	for i := 0; i < 100; i++ {
		_, err := entity.GenerateRandomPassword()
		if err != nil {
			panic(err)
		}
	}

}
