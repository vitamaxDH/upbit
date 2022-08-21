package quotation

import (
	"encoding/json"
	"upbit/api"
	"upbit/model"
)

func GetMarkets() []model.MarketCode {
	var markets []model.MarketCode
	result := api.CallNoParamApi(api.GetMarketsUrl)
	json.Unmarshal([]byte(result), &markets)
	return markets
}
