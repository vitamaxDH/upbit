package quotation

import (
	"encoding/json"
	"net/url"
	"strconv"
	"upbit/api"
	"upbit/model"
)

func GetMinuteCandle(market string, unit, count int) []model.MinuteCandel {
	queries := url.Values{
		"market": []string{market},
		"count":  []string{strconv.Itoa(count)},
	}
	claim := model.NewClaimForQuery(queries)
	apiUrl := api.GetMinutCandleUrl + strconv.Itoa(unit) + "?" + queries.Encode()
	result := api.CallGetApi(apiUrl, claim)
	var minuteCandels []model.MinuteCandel
	json.Unmarshal([]byte(result), &minuteCandels)
	return minuteCandels
}
