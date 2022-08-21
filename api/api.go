package api

import (
	"bytes"
	"encoding/json"
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
	GetMarketsUrl     = baseUrl + "/v1/market/all"
	GetMinutCandleUrl = baseUrl + "/v1/candles/minutes/"
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
