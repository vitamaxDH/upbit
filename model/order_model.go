package model

type Order struct {
	Uuid            string  `json:"uuid"`
	Side            string  `json:"side"`
	OrdType         string  `json:"ord_type"`
	Price           float64 `json:"price"`
	State           string  `json:"state"`
	Market          string  `json:"market"`
	CreatedAt       string  `json:"created_at"`
	Volume          float64 `json:"volume"`
	RemainingVolume float64 `json:"remaining_volume"`
	ReservedFee     float64 `json:"reserved_fee"`
	RemainingFee    float64 `json:"remining_fee"`
	PaidFee         float64 `json:"paid_fee"`
	Locked          float64 `json:"locked"`
	ExecutedVolume  float64 `json:"executed_volume"`
	TradesCount     int     `json:"trades_count"`
	Trades          []Trade `json:"trades"`
}

type Trade struct {
	Market    string  `json:"market"`
	Uuid      string  `json:"uuid"`
	Price     float64 `json:"price"`
	Volume    float64 `json:"volume"`
	Funds     float64 `json:"funds"`
	Side      string  `json:"side"`
	CreatedAt string  `json:"created_at"`
}
