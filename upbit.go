package main

import (
	"fmt"
	"math/rand"
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
	minuteCandles := quotation.GetMinuteCandleTo(10, "KRW-QKC", "2022-08-22 13:00:01", 3)
	var firstTwoSum float64
	for i, minuteCandle := range minuteCandles {
		if i != 2 {
			firstTwoSum += minuteCandle.HighPrice
		} else {
			fmt.Println("first two avg", firstTwoSum/2)
		}
		fmt.Printf("KST: %v TradePrice: %.2f\n", minuteCandle.CandleDateTimeKST, minuteCandle.HighPrice)
	}
}
