package exchange

import (
	"encoding/json"
	"log"
	"upbit/api"
	"upbit/model"
)

func GetAssets() {
	result := api.CallNoParamApi(api.GetAssetsUrl)
	var assets []model.Asset
	json.Unmarshal([]byte(result), &assets)
	log.Println(assets)
}
