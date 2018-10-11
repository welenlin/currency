package usecase

import (
	"context"
	model "currency/models"
	"currency/pkg/currency/mock"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetExchangeLists(t *testing.T) {
	mockz := gomock.NewController(t)
	mockz2 := gomock.NewController(t)
	defer mockz.Finish()
	defer mockz2.Finish()
	mockExchangeRepo := mock_currency.NewMockExchangeRepository(mockz)
	mockExchangeListRepo := mock_currency.NewMockExchangeListRepository(mockz2)
	expected := []model.ExchangeList{
		{1, "USD", "IDR"},
		{2, "IDR", "USD"},
	}
	mockExchangeListRepo.EXPECT().GetExchangeLists(gomock.Any()).Return(expected, nil)

	// mockExchangeRepo.EXPECT().GetExchangeRate(gomock.Any()).Return(model.ExchangeRate{}, api.ApiError{})
	usecase := NewExchangeRateUsecase(mockExchangeRepo, mockExchangeListRepo)
	res, _ := usecase.GetExchangeLists(context.TODO())

	assert.Equal(t, expected, res)

	mockExchangeListRepo.EXPECT().GetExchangeLists(gomock.Any()).Return([]model.ExchangeList{}, errors.New("ERR"))

	// mockExchangeRepo.EXPECT().GetExchangeRate(gomock.Any()).Return(model.ExchangeRate{}, api.ApiError{})
	_, err := usecase.GetExchangeLists(context.TODO())
	assert.NotNil(t, err)

}
