package request

type ExchangeRequest struct {
	Date string  `json:"date"`
	From string  `json:"from"`
	To   string  `json:"to"`
	Rate float64 `json:"rate"`
}

type ExchangeListRequest struct {
	From string `json:"from"`
	To   string `json:"to"`
}
