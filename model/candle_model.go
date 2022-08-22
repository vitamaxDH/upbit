package model

type Candle struct {
	Market               string  `json:"market"`
	CandleDateTimeUTC    string  `json:"candle_date_time_utc"`
	CandleDateTimeKST    string  `json:"candle_date_time_kst"`
	OpeningPrice         float64 `json:"opening_price"`
	HighPrice            float64 `json:"high_price"`
	LowPrice             float64 `json:"low_price"`
	TradePrice           float64 `json:"trade_price"`
	Timestamp            int64   `json:"timestamp"`
	CandleAccTradePrice  float64 `json:"candle_acc_trade_price"`
	CandleAccTradeVolume float64 `json:"candle_acc_trade_volume"`
}

type MinuteCandle struct {
	Candle
	Unit int32 `json:"unit"`
}

type DayCandle struct {
	Candle
	PrevClosingPrice    float64 `json:"prev_closing_price"`
	ChangePrice         float64 `json:"change_price"`
	ChangeRate          float64 `json:"change_rate"`
	ConvertedTradePrice float64 `json:"converted_trade_price"`
}

type WeekCandle struct {
	Candle
	FirstDayOfPeriod string `json:"first_day_of_period"`
}

type MonthCandle struct {
	Candle
	FirstDayOfPeriod string `json:"first_day_of_period"`
}
