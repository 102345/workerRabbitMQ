// Code generated by MockGen. DO NOT EDIT.
// Source: core/domain/product.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/marc/workerRabbitMQ-example/core/domain"
)

// MockIProductUseCase is a mock of IProductUseCase interface.
type MockIProductUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockIProductUseCaseMockRecorder
}

// MockIProductUseCaseMockRecorder is the mock recorder for MockIProductUseCase.
type MockIProductUseCaseMockRecorder struct {
	mock *MockIProductUseCase
}

// NewMockIProductUseCase creates a new mock instance.
func NewMockIProductUseCase(ctrl *gomock.Controller) *MockIProductUseCase {
	mock := &MockIProductUseCase{ctrl: ctrl}
	mock.recorder = &MockIProductUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIProductUseCase) EXPECT() *MockIProductUseCaseMockRecorder {
	return m.recorder
}

// FindById mocks base method.
func (m *MockIProductUseCase) FindById(id int64) (domain.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", id)
	ret0, _ := ret[0].(domain.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockIProductUseCaseMockRecorder) FindById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockIProductUseCase)(nil).FindById), id)
}

// MockIProductRepository is a mock of IProductRepository interface.
type MockIProductRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIProductRepositoryMockRecorder
}

// MockIProductRepositoryMockRecorder is the mock recorder for MockIProductRepository.
type MockIProductRepositoryMockRecorder struct {
	mock *MockIProductRepository
}

// NewMockIProductRepository creates a new mock instance.
func NewMockIProductRepository(ctrl *gomock.Controller) *MockIProductRepository {
	mock := &MockIProductRepository{ctrl: ctrl}
	mock.recorder = &MockIProductRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIProductRepository) EXPECT() *MockIProductRepositoryMockRecorder {
	return m.recorder
}

// FindById mocks base method.
func (m *MockIProductRepository) FindById(id int64) (domain.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", id)
	ret0, _ := ret[0].(domain.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockIProductRepositoryMockRecorder) FindById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockIProductRepository)(nil).FindById), id)
}
