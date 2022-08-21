package exchange

import (
	"fmt"
	"net/url"
	"upbit/api"
	"upbit/model"
)

func GetOrder(values url.Values) string {
	at := model.NewClaimForQuery(values)
	result := api.CallGetApi(api.GetOrderUrl, at)
	fmt.Println(api.PrettyString(result))
	return result
}
