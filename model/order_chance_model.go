package model

type OrderChance struct {
	BidFee     int        `json:"bid_fee"`
	AskFee     int        `json:"ask_fee"`
	Market     Market     `json:"market"`
	BidAccount BidAccount `json:"bid_account"`
	AskAccount AskAccount `json:"ask_account"`
}

type Market struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	OrderTypes []string `json:"order_types"`
	OrderSides []string `json:"order_sides"`
	Bid        Bid      `json:"bid"`
	Ask        Ask      `json:"ask"`
	MaxTotal   int      `json:"max_total"`
	State      string   `json:"state"`
}

type Bid struct {
	Currency  string `json:"currency"`
	PriceUnit string `json:"price_unit"`
	MinTotal  int    `json:"min_total"`
}

type Ask struct {
	Currency  string `json:"currency"`
	PriceUnit string `json:"price_unit"`
	MinTotal  int    `json:"min_total"`
}

type BidAccount struct {
	Currency            string `json:"currency"`
	Balance             int    `json:"balance"`
	Locked              string `json:"locked"`
	AvgBuyPrice         string `json:"avg_buy_price"`
	AvgBuyPriceModified bool   `json:"avg_buy_price_modified"`
	UnitCurrency        string `json:"unit_currency"`
}

type AskAccount struct {
	Currency            string `json:"currency"`
	Balance             int    `json:"balance"`
	Locked              string `json:"locked"`
	AvgBuyPrice         string `json:"avg_buy_price"`
	AvgBuyPriceModified bool   `json:"avg_buy_price_modified"`
	UnitCurrency        string `json:"unit_currency"`
}
