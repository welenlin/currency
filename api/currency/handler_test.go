package apicurrency

import (
	"currency/api"
	model "currency/models"
	"currency/pkg/currency/mock"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetExchangeHandler(t *testing.T) {
	mockz := gomock.NewController(t)
	mockUsecase := mock_currency.NewMockExchangeUsecase(mockz)
	mockUsecase.EXPECT().GetExchangeRate(gomock.Any()).Return(model.ExchangeRate{}, api.ApiError{})

	defer mockz.Finish()

	f := GetExchangeHandler(mockUsecase)

	r := httptest.NewRequest(http.MethodPost, "https://site.com", nil)
	w := httptest.NewRecorder()

	f(w, r)
	mockUsecase.EXPECT().GetExchangeRate(gomock.Any()).Return(model.ExchangeRate{}, api.ApiError{Status: 1})
	f(w, r)
}

func TestGetCurrencyHandler(t *testing.T) {
	mockz := gomock.NewController(t)
	mockUsecase := mock_currency.NewMockExchangeUsecase(mockz)
	mockUsecase.EXPECT().GetCurrencyCodes(gomock.Any()).Return(model.Currencies{}, api.ApiError{})

	defer mockz.Finish()

	f := GetCurrencyHandler(mockUsecase)

	r := httptest.NewRequest(http.MethodPost, "https://site.com", nil)
	w := httptest.NewRecorder()

	f(w, r)
	mockUsecase.EXPECT().GetCurrencyCodes(gomock.Any()).Return(model.Currencies{}, api.ApiError{Status: 1})
	f(w, r)
}

func TestPostExchangeHandler(t *testing.T) {
	mockz := gomock.NewController(t)
	mockUsecase := mock_currency.NewMockExchangeUsecase(mockz)
	mockUsecase.EXPECT().StoreExchangeRate(gomock.Any()).Return(int64(0), api.ApiError{})

	defer mockz.Finish()

	f := PostExchangeHandler(mockUsecase)

	r := httptest.NewRequest(http.MethodPost, "https://site.com", nil)
	w := httptest.NewRecorder()

	f(w, r)
	mockUsecase.EXPECT().StoreExchangeRate(gomock.Any()).Return(int64(0), api.ApiError{Status: 1})
	f(w, r)
}

func TestPostExchangeListHandler(t *testing.T) {
	mockz := gomock.NewController(t)
	mockUsecase := mock_currency.NewMockExchangeUsecase(mockz)
	mockUsecase.EXPECT().StoreExchangeList(gomock.Any()).Return(int64(0), api.ApiError{})

	defer mockz.Finish()

	f := PostExchangeListHandler(mockUsecase)

	r := httptest.NewRequest(http.MethodPost, "https://site.com", nil)
	w := httptest.NewRecorder()

	f(w, r)
	mockUsecase.EXPECT().StoreExchangeList(gomock.Any()).Return(int64(0), api.ApiError{Status: 1})
	f(w, r)
}
