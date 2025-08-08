package main

import (
	"fmt"
	"sync"
	"time"

	"server/infrastructure/logger"
	"server/router/user"

	"github.com/google/uuid"
)

func main() {
	// StampCard()
	// UnlimitCheckin()
	// Checkin()
	ConcurrentCheckin()
}

func Checkin() {
	usecase := user.InitializeUserCheckinUsecase()

	userID, _ := uuid.Parse("6db2a0bb-844e-4344-bcf8-79d760aacbf6")
	limitQr, _ := uuid.Parse("9d5308bc-b202-44d9-b064-e48470008e4a")
	stampCard, coupon, err := usecase.Checkin(userID, limitQr)
	if err != nil {
		logger.Error(err.Error())
	}
	println(stampCard)
	println(coupon)
}

func UnlimitCheckin() {
	usecase := user.InitializeUserCheckinUsecase()

	userID, _ := uuid.Parse("6db2a0bb-844e-4344-bcf8-79d760aacbf6")
	unlimitQr, _ := uuid.Parse("9d5308bc-b202-44d9-b064-e48470008e4a")
	stampCard, coupon, err := usecase.Checkin(userID, unlimitQr)
	if err != nil {
		logger.Error(err.Error())
	}
	println(stampCard)
	println(coupon)
}

func StampCard() {
	usecase := user.InitializeUserCheckinUsecase()
	userID, _ := uuid.Parse("6db2a0bb-844e-4344-bcf8-79d760aacbf6")
	card, err := usecase.GetStampCard(userID)
	if err != nil {
		logger.Error(err.Error())
	}
	fmt.Println(card.Stamps[0])
}

func ConcurrentCheckin() {
	fmt.Println("Testing concurrent check-in handling...")
	
	usecase := user.InitializeUserCheckinUsecase()
	userID, _ := uuid.Parse("6db2a0bb-844e-4344-bcf8-79d760aacbf6")
	qrCode, _ := uuid.Parse("9d5308bc-b202-44d9-b064-e48470008e4a")
	
	const numRequests = 5
	var wg sync.WaitGroup
	wg.Add(numRequests)
	
	successCount := 0
	processingCount := 0
	errorCount := 0
	var mu sync.Mutex
	
	startTime := time.Now()
	
	// Execute multiple check-in requests concurrently
	for i := 0; i < numRequests; i++ {
		go func(index int) {
			defer wg.Done()
			
			fmt.Printf("Request %d starting...\n", index)
			stampCard, coupon, err := usecase.Checkin(userID, qrCode)
			
			mu.Lock()
			defer mu.Unlock()
			
			if err == nil {
				successCount++
				fmt.Printf("Request %d succeeded! StampCard: %v, Coupon: %v\n", index, stampCard != nil, coupon != nil)
			} else {
				if err.Error() == "チェックイン処理中です。しばらくお待ちください。" {
					processingCount++
					fmt.Printf("Request %d blocked: %s\n", index, err.Error())
				} else {
					errorCount++
					fmt.Printf("Request %d error: %s\n", index, err.Error())
				}
			}
		}(i)
	}
	
	wg.Wait()
	
	elapsed := time.Since(startTime)
	fmt.Printf("\nTest completed in %v\n", elapsed)
	fmt.Printf("Success: %d, Processing: %d, Errors: %d\n", successCount, processingCount, errorCount)
	
	if successCount == 1 && processingCount == numRequests-1 {
		fmt.Println("✅ Concurrent check-in handling working correctly!")
	} else {
		fmt.Println("❌ Concurrent check-in handling has issues!")
	}
}
