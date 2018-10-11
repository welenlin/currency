package model

import (
	"time"
)

type ExchangeRate struct {
	ID          int64     `json:"id" db:"id"`
	From        string    `json:"from" db:"from_currency"`
	To          string    `json:"to" db:"to_currency"`
	Date        time.Time `json:"date" db:"rate_date"`
	Rate        float64   `json:"rate" db:"rate"`
	OneWeekRate float64   `json:"week_rate"`
	Variance    float64   `json:"variance"`
}

type ExchangeList struct {
	ID   int64  `json:"id" db:"id"`
	From string `json:"from" db:"from_currency"`
	To   string `json:"to" db:"to_currency"`
}

type ExchangeRates []ExchangeRate

type Currency struct {
	Code  string
	Title string
}

type Currencies []Currency
