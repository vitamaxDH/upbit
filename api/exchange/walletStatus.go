package exchange

import "upbit/api"

func GetWalletStatus() {
	api.CallNoParamApi(api.GetWalletStatusUrl)
}
