package currency

import (
	"context"
	"currency/models"
	"time"
)

const (
	USD = "USD"
)

type ExchangeRepository interface {
	GetExchangeRate(ctx context.Context, from, to string, date time.Time) (*model.ExchangeRate, error)
	GetExchangeRateWeekAverage(ctx context.Context, from, to string, date time.Time) (*model.ExchangeRate, error)
	Store(ctx context.Context, c model.ExchangeRate) (int64, error)
}

type ExchangeListRepository interface {
	GetExchangeLists(ctx context.Context) ([]model.ExchangeList, error)
	Store(ctx context.Context, el model.ExchangeList) (int64, error)
	Delete(ctx context.Context, id int64) (bool, error)
}
