package exchange

import "upbit/api"

func GetApiKeys() {
	api.CallNoParamApi(api.GetApiKeysUrl)
}
