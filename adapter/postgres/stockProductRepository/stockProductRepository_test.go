package stockProductRepository_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/marc/workerRabbitMQ-example/adapter/postgres/stockProductRepository"
	"github.com/marc/workerRabbitMQ-example/core/domain"
	"github.com/marc/workerRabbitMQ-example/core/dto"
	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/require"
)

func setupCreateStockProduct() ([]string, dto.StockProductDTO, domain.StockProduct, pgxmock.PgxPoolIface) {
	cols := []string{"id", "productid", "quantity", "balance"}
	fakeStockProductDTO := dto.StockProductDTO{}
	fakeStockProduct := domain.StockProduct{}
	faker.FakeData(&fakeStockProductDTO)
	faker.FakeData(&fakeStockProduct)

	mock, _ := pgxmock.NewPool()

	return cols, fakeStockProductDTO, fakeStockProduct, mock
}

func TestCreateStockProduct(t *testing.T) {
	cols, fakeStockProductDTO, fakeStockProduct, mock := setupCreateStockProduct()
	defer mock.Close()

	mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO StockProduct (productid,quantity,balance) VALUES ($1, $2, $3) returning *")).WithArgs(
		fakeStockProductDTO.ProductID,
		fakeStockProductDTO.Quantity,
		fakeStockProductDTO.Balance,
	).WillReturnRows(pgxmock.NewRows(cols).AddRow(
		fakeStockProduct.ID,
		fakeStockProduct.ProductID,
		fakeStockProduct.Quantity,
		fakeStockProduct.Balance,
	))

	sut := stockProductRepository.New(mock)
	stockProductRet, err := sut.Create(&fakeStockProductDTO)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	require.Nil(t, err)
	require.NotEmpty(t, stockProductRet.ID)
	require.Equal(t, stockProductRet.ProductID, stockProductRet.ProductID)
	require.Equal(t, stockProductRet.Quantity, stockProductRet.Quantity)

}

func TestCreateStockProduct_DBError(t *testing.T) {
	_, fakeStockProductDTO, _, mock := setupCreateStockProduct()

	defer mock.Close()

	mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO StockProduct (productid,quantity,balance) VALUES ($1, $2, $3) returning *")).WithArgs(
		fakeStockProductDTO.ProductID,
		fakeStockProductDTO.Quantity,
		fakeStockProductDTO.Balance,
	).WillReturnError(fmt.Errorf("ANY DATABASE ERROR"))

	sut := stockProductRepository.New(mock)
	stockProductRet, err := sut.Create(&fakeStockProductDTO)

	require.NotNil(t, err)
	require.Nil(t, stockProductRet)
}
