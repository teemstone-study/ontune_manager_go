package app

import (
	"fmt"
	"time"
)

func ConsumerProcessing(results chan<- string, interval int) {
	ticker := time.NewTicker(time.Second * 1)
	for range ticker.C {
		results <- fmt.Sprintf("data: %d", time.Now().Unix())
		time.Sleep(time.Second * time.Duration(interval))
	}
}
