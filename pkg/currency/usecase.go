package currency

import (
	"context"
	"currency/api"
	"currency/models"
)

var AvailableExchangeRates = map[string]string{
	"USD": "United States Dollar",
	"RON": "Romanian Leu",
	"RUB": "Russian Ruble",
	"BRL": "Brazilian Real",
	"TRY": "Turkish lira",
	"DKK": "Danish Krone",
	"KRW": "South Korean won",
	"JPY": "Japanese Yen",
	"HUF": "Hungarian Forint",
	"SGD": "Singapore Dollar",
	"PHP": "Philippine Piso",
	"CNY": "Chinese Yuan",
	"NOK": "Norwegian Krone",
	"SEK": "Swedish Krona",
	"MXN": "Mexican Peso",
	"GBP": "Pound sterling",
	"IDR": "Indonesian Rupiah",
	"HRK": "Croatian Kuna",
	"AUD": "Australian Dollar",
	"ZAR": "South African Rand",
	"BGN": "Bulgarian Lev",
	"CZK": "Czech Koruna",
	"NZD": "New Zealand Dollar",
	"INR": "Indian Rupee",
	"CAD": "Canadian Dollar",
	"THB": "Thai Baht",
	"HKD": "Hong Kong Dollar",
	"MYR": "Malaysian Ringgit",
	"PLN": "Poland złoty",
	"ILS": "Israeli New Shekel",
	"ISK": "Icelandic Króna",
	"CHF": "Swiss Franc",
	"EUR": "Euro",
}

type ExchangeUsecase interface {
	GetExchangeRate(ctx context.Context) (model.ExchangeRate, api.ApiError)
	GetCurrencyCodes(ctx context.Context) (model.Currencies, api.ApiError)
	GetExchangeLists(ctx context.Context) ([]model.ExchangeList, api.ApiError)
	StoreExchangeRate(ctx context.Context) (int64, api.ApiError)
	StoreExchangeList(ctx context.Context) (int64, api.ApiError)
	DeleteExchangeList(ctx context.Context, id int64) (bool, api.ApiError)
}
