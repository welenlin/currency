package usecase

import (
	"context"
	"currency/api"
	model "currency/models"
	"currency/pkg/currency"
	"currency/pkg/currency/request"
	"log"
	"net/http"
	"time"
)

type currencyUsecase struct {
	exRepo  currency.ExchangeRepository
	exlRepo currency.ExchangeListRepository
}

func NewExchangeRateUsecase(ex currency.ExchangeRepository, exl currency.ExchangeListRepository) currency.ExchangeUsecase {
	return &currencyUsecase{
		exRepo:  ex,
		exlRepo: exl,
	}
}

func (c *currencyUsecase) GetExchangeRate(ctx context.Context) (model.ExchangeRate, api.ApiError) {
	apiErr := api.ApiError{}

	var d = time.Now()

	dateInput, _ := api.DateKeyFromContext(ctx)
	if dateInput != "" {
		d, _ = time.Parse("2006-01-02", dateInput)
	}
	from, ok := api.FromKeyFromContext(ctx)
	if !ok {
		apiErr.Status = http.StatusInternalServerError
		apiErr.Title = "From Not Found"
		return model.ExchangeRate{}, apiErr
	}
	to, ok := api.ToKeyFromContext(ctx)
	if !ok {
		apiErr.Status = http.StatusInternalServerError
		apiErr.Title = "To Not Found"
		return model.ExchangeRate{}, apiErr
	}

	ex, err := c.exRepo.GetExchangeRate(ctx, from, to, d)
	if err != nil {
		apiErr.Status = http.StatusInternalServerError
		apiErr.Title = err.Error()
		return model.ExchangeRate{}, apiErr
	}
	return *ex, apiErr

}

func (c *currencyUsecase) GetCurrencyCodes(ctx context.Context) (model.Currencies, api.ApiError) {
	apiErr := api.ApiError{}
	currencies := model.Currencies{}
	for i := range currency.AvailableExchangeRates {
		c := model.Currency{
			Code:  i,
			Title: currency.AvailableExchangeRates[i],
		}
		currencies = append(currencies, c)
	}

	return currencies, apiErr
}
func (c *currencyUsecase) GetExchangeLists(ctx context.Context) ([]model.ExchangeList, api.ApiError) {
	apiErr := api.ApiError{}
	exl, err := c.exlRepo.GetExchangeLists(ctx)
	if err != nil {
		apiErr.Status = http.StatusInternalServerError
		apiErr.Title = err.Error()
		return []model.ExchangeList{}, apiErr
	}
	return exl, apiErr
}

func (c *currencyUsecase) StoreExchangeRate(ctx context.Context) (int64, api.ApiError) {
	apiErr := api.ApiError{}
	request := request.ExchangeRequest{}

	ok := api.JSONFromContext(ctx, &request)
	if !ok {
		apiErr.Status = http.StatusBadRequest
		apiErr.Title = "Wrong Json Format"
		return 0, apiErr
	}
	if valid := isExchangeRequestValid(request); !valid {
		apiErr.Status = http.StatusBadRequest
		apiErr.Title = "Wrong Value Format"
		return 0, apiErr
	}

	d, _ := time.Parse("2006-01-02", request.Date)

	e := model.ExchangeRate{
		Date: d,
		From: request.From,
		To:   request.To,
		Rate: request.Rate,
	}

	lastID, err := c.exRepo.Store(ctx, e)
	if err != nil {
		apiErr.Status = http.StatusInternalServerError
		apiErr.Title = err.Error()
		return 0, apiErr
	}

	return lastID, api.ApiError{}
}

func (c *currencyUsecase) StoreExchangeList(ctx context.Context) (int64, api.ApiError) {
	apiErr := api.ApiError{}
	request := request.ExchangeListRequest{}

	ok := api.JSONFromContext(ctx, &request)
	if !ok {
		apiErr.Status = http.StatusBadRequest
		apiErr.Title = "Wrong Json Format"
		return 0, apiErr
	}
	if valid := isExchangeListRequestValid(request); !valid {
		apiErr.Status = http.StatusBadRequest
		apiErr.Title = "Wrong Value Format"
		return 0, apiErr
	}

	e := model.ExchangeList{
		From: request.From,
		To:   request.To,
	}

	lastID, err := c.exlRepo.Store(ctx, e)

	if err != nil {
		apiErr.Status = http.StatusInternalServerError
		apiErr.Title = err.Error()
		return 0, apiErr
	}

	return lastID, api.ApiError{}
}

func isExchangeRequestValid(e request.ExchangeRequest) bool {
	// d, err := time.Parse("2006-01-02", e.Date)
	// if err != nil {
	// 	return false
	// }

	if !validateFromAndTo(e.From, e.To) {
		return false
	}

	return true

}

func isExchangeListRequestValid(e request.ExchangeListRequest) bool {

	if !validateFromAndTo(e.From, e.To) {
		return false
	}

	return true

}

func validateFromAndTo(from, to string) bool {
	if len(from) != 3 || len(to) != 3 {
		return false
	}

	if from == "" || to == "" {
		return false
	}
	return true
}

func (c *currencyUsecase) DeleteExchangeList(ctx context.Context, id int64) (bool, api.ApiError) {
	apiErr := api.ApiError{}
	ok, err := c.exlRepo.Delete(ctx, id)
	if err != nil {
		log.Println("func DeleteExchangeList", err)
		apiErr.Status = http.StatusInternalServerError
		apiErr.Title = err.Error()
		return false, apiErr
	}

	return ok, apiErr
}
