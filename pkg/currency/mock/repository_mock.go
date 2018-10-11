// Code generated by MockGen. DO NOT EDIT.
// Source: C:/Users/Welen/go/src/currency/pkg/currency/repository.go

// Package mock_currency is a generated GoMock package.
package mock_currency

import (
	context "context"
	models "currency/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockExchangeRepository is a mock of ExchangeRepository interface
type MockExchangeRepository struct {
	ctrl     *gomock.Controller
	recorder *MockExchangeRepositoryMockRecorder
}

// MockExchangeRepositoryMockRecorder is the mock recorder for MockExchangeRepository
type MockExchangeRepositoryMockRecorder struct {
	mock *MockExchangeRepository
}

// NewMockExchangeRepository creates a new mock instance
func NewMockExchangeRepository(ctrl *gomock.Controller) *MockExchangeRepository {
	mock := &MockExchangeRepository{ctrl: ctrl}
	mock.recorder = &MockExchangeRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockExchangeRepository) EXPECT() *MockExchangeRepositoryMockRecorder {
	return m.recorder
}

// GetExchangeRate mocks base method
func (m *MockExchangeRepository) GetExchangeRate(ctx context.Context, from, to string, date time.Time) (*models.ExchangeRate, error) {
	ret := m.ctrl.Call(m, "GetExchangeRate", ctx, from, to, date)
	ret0, _ := ret[0].(*models.ExchangeRate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExchangeRate indicates an expected call of GetExchangeRate
func (mr *MockExchangeRepositoryMockRecorder) GetExchangeRate(ctx, from, to, date interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExchangeRate", reflect.TypeOf((*MockExchangeRepository)(nil).GetExchangeRate), ctx, from, to, date)
}

// GetExchangeRateWeekAverage mocks base method
func (m *MockExchangeRepository) GetExchangeRateWeekAverage(ctx context.Context, from, to string, date time.Time) (*models.ExchangeRate, error) {
	ret := m.ctrl.Call(m, "GetExchangeRateWeekAverage", ctx, from, to, date)
	ret0, _ := ret[0].(*models.ExchangeRate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExchangeRateWeekAverage indicates an expected call of GetExchangeRateWeekAverage
func (mr *MockExchangeRepositoryMockRecorder) GetExchangeRateWeekAverage(ctx, from, to, date interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExchangeRateWeekAverage", reflect.TypeOf((*MockExchangeRepository)(nil).GetExchangeRateWeekAverage), ctx, from, to, date)
}

// Store mocks base method
func (m *MockExchangeRepository) Store(ctx context.Context, c models.ExchangeRate) (int64, error) {
	ret := m.ctrl.Call(m, "Store", ctx, c)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Store indicates an expected call of Store
func (mr *MockExchangeRepositoryMockRecorder) Store(ctx, c interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockExchangeRepository)(nil).Store), ctx, c)
}

// MockExchangeListRepository is a mock of ExchangeListRepository interface
type MockExchangeListRepository struct {
	ctrl     *gomock.Controller
	recorder *MockExchangeListRepositoryMockRecorder
}

// MockExchangeListRepositoryMockRecorder is the mock recorder for MockExchangeListRepository
type MockExchangeListRepositoryMockRecorder struct {
	mock *MockExchangeListRepository
}

// NewMockExchangeListRepository creates a new mock instance
func NewMockExchangeListRepository(ctrl *gomock.Controller) *MockExchangeListRepository {
	mock := &MockExchangeListRepository{ctrl: ctrl}
	mock.recorder = &MockExchangeListRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockExchangeListRepository) EXPECT() *MockExchangeListRepositoryMockRecorder {
	return m.recorder
}

// GetExchangeLists mocks base method
func (m *MockExchangeListRepository) GetExchangeLists(ctx context.Context) ([]models.ExchangeList, error) {
	ret := m.ctrl.Call(m, "GetExchangeLists", ctx)
	ret0, _ := ret[0].([]models.ExchangeList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExchangeLists indicates an expected call of GetExchangeLists
func (mr *MockExchangeListRepositoryMockRecorder) GetExchangeLists(ctx interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExchangeLists", reflect.TypeOf((*MockExchangeListRepository)(nil).GetExchangeLists), ctx)
}

// Store mocks base method
func (m *MockExchangeListRepository) Store(ctx context.Context, el models.ExchangeList) (int64, error) {
	ret := m.ctrl.Call(m, "Store", ctx, el)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Store indicates an expected call of Store
func (mr *MockExchangeListRepositoryMockRecorder) Store(ctx, el interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockExchangeListRepository)(nil).Store), ctx, el)
}

// Delete mocks base method
func (m *MockExchangeListRepository) Delete(ctx context.Context, id int64) (bool, error) {
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockExchangeListRepositoryMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockExchangeListRepository)(nil).Delete), ctx, id)
}