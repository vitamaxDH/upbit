package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"upbit/model"

	"io/ioutil"
	"net/http"

	"github.com/dgrijalva/jwt-go/v4"
)

const (
	// EXCHANGE API
	baseUrl            = "https://api.upbit.com"
	GetAssetsUrl       = baseUrl + "/v1/accounts"
	GetChanceUrl       = baseUrl + "/v1/orders/chance"
	GetOrderUrl        = baseUrl + "/v1/order"
	GetApiKeysUrl      = baseUrl + "/v1/api_keys"
	GetWalletStatusUrl = baseUrl + "/v1/status/wallet"

	// QUOTATION API
	GetMarketsUrl      = baseUrl + "/v1/market/all"
	GetCandleUrl       = baseUrl + "/v1/candles"
	GetMinuteCandleUrl = GetCandleUrl + "/minutes/"
	GetDayCandleUrl    = GetCandleUrl + "/days?"
	GetWeekCandleUrl   = GetCandleUrl + "/weeks?"
	GetMonthCandleUrl  = GetCandleUrl + "/months?"
)

var (
	Markets []model.MarketCode
)

func CallNoParamApi(url string) string {
	claim := model.NewClaim()
	return CallGetApi(url, claim)
}

func CallGetApi(url string, claim jwt.MapClaims) string {
	return CallApi("GET", url, claim)
}

func CallApi(method, url string, claim jwt.MapClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	bearerToken := model.GetBearerToken(token)

	req, _ := http.NewRequest(method, url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", bearerToken)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}

func PrettyString(str string) string {
	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, []byte(str), "", "    ")
	return prettyJSON.String()
}

func ValidateMarket(market string) error {
	exists := false
	for _, marketCode := range Markets {
		if market == marketCode.Market {
			exists = true
		}
	}
	if exists {
		return nil
	} else {
		return fmt.Errorf("given market %s does not exist", market)
	}
}
