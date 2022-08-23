package main

import (
	"log"
	"math/rand"
	"os"
	"time"
	"upbit/api"
	"upbit/api/quotation"
)

// todo: timestamp 관련 util 함수 만들기
//  docker 연동
//	db 연동
//  로컬 테스트
//  Dockerfile 만들기
//  ec2 배포
//	rds free tier 구매

func main() {
	rand.Seed(time.Now().UnixMilli())
	api.Markets = quotation.GetMarkets()
	time := "2022-08-23 01:00:01"
	minuteCandles, err := quotation.GetMinuteCandleTo(10, "KRW-ATOM", time, 3)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	var firstTwoSum float64
	var lastPrice float64
	candleSize := len(minuteCandles)
	for i := candleSize - 1; i >= 0; i-- {
		minuteCandle := minuteCandles[i]
		if i != 0 {
			firstTwoSum += minuteCandle.HighPrice
			log.Printf("i: %d, price: %.f", i, minuteCandle.HighPrice)
		} else {
			lastPrice = minuteCandle.HighPrice
			log.Printf("i: %d, lastPrice: %.f", i, lastPrice)
		}
	}
	firstTwoAvg := firstTwoSum / 2
	somePctOfAvg := (firstTwoAvg / 100) * 2
	log.Printf("%s firstTwo avg: %.f, lastPrice: %.f, somePrctOfAvg: %.f\n", time, firstTwoAvg, lastPrice, somePctOfAvg)
	diff := lastPrice - firstTwoAvg
	log.Printf("diff: %.f", diff)
	if diff > somePctOfAvg {
		log.Printf("firstTwo avg: %.f, lastPrice: %.f, somePrctOfAvg: %.f\n", firstTwoAvg, lastPrice, somePctOfAvg)
	}
}
