package quotation

import (
	"encoding/json"
	"net/url"
	"strconv"
	"upbit/api"
	"upbit/model"
	"upbit/util"
)

func GetMinuteCandle(unit int32, market string, count int32) []model.MinuteCandle {
	return GetMinuteCandleTo(unit, market, "", count)
}

func GetMinuteCandleTo(unit int32, market, to string, count int32) []model.MinuteCandle {
	// todo: parameter validation
	// unit: 1, 3, 5, 15, 10, 30, 60, 240
	// market: check if all markets.contains(market)
	// to: yyyy-MM-dd'T'HH:mm:ss'Z' or yyyy-MM-dd HH:mm:ss
	// count: min 1, max 200
	queries := url.Values{
		"market": []string{market},
		"count":  []string{strconv.Itoa(int(count))},
	}
	util.AddQueryIfNotEmpty(&queries, "to", to)
	claim := model.NewClaimForQuery(queries)
	apiUrl := api.GetMinuteCandleUrl + strconv.Itoa(int(count)) + "?" + queries.Encode()
	result := api.CallGetApi(apiUrl, claim)

	var minuteCandles []model.MinuteCandle
	json.Unmarshal([]byte(result), &minuteCandles)
	return minuteCandles
}

func GetDayCandle(market string, count int32, convertingPriceUnit string) []model.DayCandle {
	return GetDayCandleTo(market, "", count, convertingPriceUnit)
}

func GetDayCandleTo(market, to string, count int32, convertingPriceUnit string) []model.DayCandle {
	queries := url.Values{
		"market": []string{market},
		"count":  []string{strconv.Itoa(int(count))},
	}
	util.AddQueryIfNotEmpty(&queries, "to", to)
	util.AddQueryIfNotEmpty(&queries, "convertingPriceUnit", convertingPriceUnit)
	claim := model.NewClaimForQuery(queries)
	apiUrl := api.GetDayCandleUrl + queries.Encode()
	result := api.CallGetApi(apiUrl, claim)

	var dayCandles []model.DayCandle
	json.Unmarshal([]byte(result), &dayCandles)
	return dayCandles
}
