// Code generated by MockGen. DO NOT EDIT.
// Source: core/domain/stockProducts.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/marc/workerRabbitMQ-example/core/domain"
	dto "github.com/marc/workerRabbitMQ-example/core/dto"
)

// MockIStockProductUseCase is a mock of IStockProductUseCase interface.
type MockIStockProductUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockIStockProductUseCaseMockRecorder
}

// MockIStockProductUseCaseMockRecorder is the mock recorder for MockIStockProductUseCase.
type MockIStockProductUseCaseMockRecorder struct {
	mock *MockIStockProductUseCase
}

// NewMockIStockProductUseCase creates a new mock instance.
func NewMockIStockProductUseCase(ctrl *gomock.Controller) *MockIStockProductUseCase {
	mock := &MockIStockProductUseCase{ctrl: ctrl}
	mock.recorder = &MockIStockProductUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIStockProductUseCase) EXPECT() *MockIStockProductUseCaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIStockProductUseCase) Create(stockProduct *dto.StockProductDTO) (*domain.StockProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", stockProduct)
	ret0, _ := ret[0].(*domain.StockProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIStockProductUseCaseMockRecorder) Create(stockProduct interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIStockProductUseCase)(nil).Create), stockProduct)
}

// MockIStockProductRepository is a mock of IStockProductRepository interface.
type MockIStockProductRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIStockProductRepositoryMockRecorder
}

// MockIStockProductRepositoryMockRecorder is the mock recorder for MockIStockProductRepository.
type MockIStockProductRepositoryMockRecorder struct {
	mock *MockIStockProductRepository
}

// NewMockIStockProductRepository creates a new mock instance.
func NewMockIStockProductRepository(ctrl *gomock.Controller) *MockIStockProductRepository {
	mock := &MockIStockProductRepository{ctrl: ctrl}
	mock.recorder = &MockIStockProductRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIStockProductRepository) EXPECT() *MockIStockProductRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIStockProductRepository) Create(StockProduct *dto.StockProductDTO) (*domain.StockProduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", StockProduct)
	ret0, _ := ret[0].(*domain.StockProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIStockProductRepositoryMockRecorder) Create(StockProduct interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIStockProductRepository)(nil).Create), StockProduct)
}
