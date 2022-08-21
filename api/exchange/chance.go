package exchange

import (
	"encoding/json"
	"net/url"
	"upbit/api"
	"upbit/model"
)

func GetChance(market string) string {
	queries := url.Values{
		"market": []string{market},
	}
	claim := model.NewClaimForQuery(queries)
	apiUrl := api.GetChanceUrl + "?" + queries.Encode()
	result := api.CallGetApi(apiUrl, claim)
	var orderChance model.OrderChance
	json.Unmarshal([]byte(result), &orderChance)
	return result
}
