package stockProductUseCase_test

import (
	"fmt"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/marc/workerRabbitMQ-example/core/domain"
	"github.com/marc/workerRabbitMQ-example/core/domain/mocks"
	"github.com/marc/workerRabbitMQ-example/core/dto"
	"github.com/marc/workerRabbitMQ-example/core/usecase/stockProductUseCase"
	"github.com/stretchr/testify/require"
)

func TestCreateStockProduct(t *testing.T) {
	fakeStockProductDTO := dto.StockProductDTO{}
	fakeStockProduct := domain.StockProduct{}
	faker.FakeData(&fakeStockProductDTO)
	faker.FakeData(&fakeStockProduct)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockStockProductRepository := mocks.NewMockIStockProductRepository(mockCtrl)
	mockStockProductRepository.EXPECT().Create(&fakeStockProductDTO).Return(&fakeStockProduct, nil)

	sut := stockProductUseCase.New(mockStockProductRepository)
	stockProductRet, err := sut.Create(&fakeStockProductDTO)

	require.Nil(t, err)
	require.NotEmpty(t, stockProductRet.ID)
	require.Equal(t, stockProductRet.ProductID, fakeStockProduct.ProductID)
	require.Equal(t, stockProductRet.Quantity, fakeStockProduct.Quantity)
	require.Equal(t, stockProductRet.Balance, fakeStockProduct.Balance)
}

func TestCreateStockProduct_Error(t *testing.T) {
	fakeStockProductDTO := dto.StockProductDTO{}
	faker.FakeData(&fakeStockProductDTO)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockStockProductRepository := mocks.NewMockIStockProductRepository(mockCtrl)
	mockStockProductRepository.EXPECT().Create(&fakeStockProductDTO).Return(nil, fmt.Errorf("ANY ERROR"))

	sut := stockProductUseCase.New(mockStockProductRepository)
	stockProductRet, err := sut.Create(&fakeStockProductDTO)

	require.NotNil(t, err)
	require.Nil(t, stockProductRet)
}
