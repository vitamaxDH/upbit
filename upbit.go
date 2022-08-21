package main

import (
	"log"
	"math/rand"
	"time"
	"upbit/api/quotation"
)

func main() {
	// api.GetAssets()
	rand.Seed(time.Now().UnixMilli())
	// markets := quotation.GetMarkets()
	// for _, market := range markets {
	// 	log.Println(market)
	// }
	log.Println()
	for _, minuteCandle := range quotation.GetMinuteCandle("KRW-MVL", 1, 5) {
		log.Println(minuteCandle)
	}

	// exchange.GetChance("KRW-MVL")
	// GetOrder(map[string]string{"uuid": uuid.NewString()})
	// GetApiKeys()
	// GetWalletStatus()
}
