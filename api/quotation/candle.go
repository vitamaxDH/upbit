package quotation

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"upbit/api"
	"upbit/model"
	"upbit/util"
)

func GetMinuteCandle(unit int32, market string, count int32) ([]model.MinuteCandle, error) {
	return GetMinuteCandleTo(unit, market, "", count)
}

func GetMinuteCandleTo(unit int32, market, to string, count int32) ([]model.MinuteCandle, error) {
	// todo: parameter validation
	switch unit {
	case 1, 3, 5, 10, 15, 30, 60, 240:
	default:
		return nil, fmt.Errorf("unit should be one of 1, 3, 5, 10, 15, 30, 60, 240")
	}
	if err := api.ValidateMarket(market); err != nil {
		return nil, err
	}
	if count < 1 || count > 200 {
		return nil, fmt.Errorf("count should be greater than 0 and less than 201")
	}
	// to: yyyy-MM-dd'T'HH:mm:ss'Z' or yyyy-MM-dd HH:mm:ss
	// count: min 1, max 200
	queries := url.Values{
		"market": []string{market},
		"count":  []string{strconv.Itoa(int(count))},
	}
	util.AddQueryIfNotEmpty(&queries, "to", to)
	claim := model.NewClaimForQuery(queries)
	apiUrl := api.GetMinuteCandleUrl + strconv.Itoa(int(unit)) + "?" + queries.Encode()
	result := api.CallGetApi(apiUrl, claim)

	var minuteCandles []model.MinuteCandle
	json.Unmarshal([]byte(result), &minuteCandles)
	return minuteCandles, nil
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
