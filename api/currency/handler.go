package apicurrency

import (
	"currency/api"
	"currency/pkg/currency"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetExchangeHandler(usecase currency.ExchangeUsecase) api.HandlerFunc {
	return api.HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		ctx := api.GetAllContextFromRequest(r)
		exchangeRate, apiErr := usecase.GetExchangeRate(ctx)
		if apiErr.Status != 0 {
			jsonErr := api.JsonErrorResponse{Error: &apiErr}
			api.ErrorResponse(w, jsonErr)
			return nil
		}
		data := api.JsonResponse{}
		data.Data = exchangeRate
		api.ResponseJSON(w, data)
		return nil
	})
}

func GetCurrencyHandler(usecase currency.ExchangeUsecase) api.HandlerFunc {
	return api.HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		currencies, apiErr := usecase.GetCurrencyCodes(r.Context())
		if apiErr.Status != 0 {
			jsonErr := api.JsonErrorResponse{Error: &apiErr}
			api.ErrorResponse(w, jsonErr)
		}
		data := api.JsonResponse{}
		data.Data = currencies
		api.ResponseJSON(w, data)
		return nil
	})
}

func PostExchangeHandler(usecase currency.ExchangeUsecase) api.HandlerFunc {
	return api.HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		ctx := api.GetAllContextFromRequest(r)
		_, apiErr := usecase.StoreExchangeRate(ctx)

		if apiErr.Status != 0 {
			jsonErr := api.JsonErrorResponse{Error: &apiErr}
			api.ErrorResponse(w, jsonErr)
			return nil
		}

		data := api.JsonResponse{}
		data.Data = struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}{
			Status:  "OK",
			Message: "Data has been successfuly created",
		}
		api.ResponseJSON(w, data)
		return nil
	})
}

func PostExchangeListHandler(usecase currency.ExchangeUsecase) api.HandlerFunc {
	return api.HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		ctx := api.GetAllContextFromRequest(r)
		_, apiErr := usecase.StoreExchangeList(ctx)

		if apiErr.Status != 0 {
			jsonErr := api.JsonErrorResponse{Error: &apiErr}
			api.ErrorResponse(w, jsonErr)
			return nil
		}

		data := api.JsonResponse{}
		data.Data = struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}{
			Status:  "OK",
			Message: "Data has been successfuly created",
		}
		api.ResponseJSON(w, data)
		return nil
	})
}

func DeleteExchangeListHandler(usecase currency.ExchangeUsecase) api.HandlerFunc {
	return api.HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		vars := mux.Vars(r)
		val, ok := vars["id"]

		if !ok {
			jsonErr := api.JsonErrorResponse{Error: &api.ApiError{
				Status: http.StatusBadRequest,
				Title:  "Missing ID Value",
			}}
			api.ErrorResponse(w, jsonErr)
			return nil
		}
		id, err := strconv.ParseInt(val, 10, 64)

		if err != nil {
			jsonErr := api.JsonErrorResponse{Error: &api.ApiError{
				Status: http.StatusBadRequest,
				Title:  "Missing ID Value",
			}}
			api.ErrorResponse(w, jsonErr)
			return nil
		}

		ctx := api.GetAllContextFromRequest(r)
		_, apiErr := usecase.DeleteExchangeList(ctx, id)

		if apiErr.Status != 0 {
			jsonErr := api.JsonErrorResponse{Error: &apiErr}
			api.ErrorResponse(w, jsonErr)
			return nil
		}

		data := api.JsonResponse{}
		data.Data = struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}{
			Status:  "OK",
			Message: "Data has been successfuly deleted",
		}
		api.ResponseJSON(w, data)
		return nil
	})
}

func GetExchangeListsHandler(usecase currency.ExchangeUsecase) api.HandlerFunc {
	return api.HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		ctx := api.GetAllContextFromRequest(r)
		exchangeLists, apiErr := usecase.GetExchangeLists(ctx)
		if apiErr.Status != 0 {
			jsonErr := api.JsonErrorResponse{Error: &apiErr}
			api.ErrorResponse(w, jsonErr)
			return nil
		}
		data := api.JsonResponse{}
		data.Data = exchangeLists
		api.ResponseJSON(w, data)
		return nil
	})
}
