package productRepository_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/marc/workerRabbitMQ-example/adapter/postgres/productRepository"
	"github.com/marc/workerRabbitMQ-example/core/domain"
	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/require"
)

func setupFindById() ([]string, domain.Product, pgxmock.PgxPoolIface) {

	cols := []string{"id", "name", "price", "description"}

	fakeProductDBResponse := domain.Product{}
	faker.FakeData(&fakeProductDBResponse)

	mock, _ := pgxmock.NewPool()

	return cols, fakeProductDBResponse, mock
}

func TestFindById(t *testing.T) {
	cols, fakeProductDBResponse, mock := setupFindById()
	fakeProductDBResponse.ID = 1
	defer mock.Close()

	rows := pgxmock.NewRows(cols).AddRow(
		fakeProductDBResponse.ID,
		fakeProductDBResponse.Name,
		fakeProductDBResponse.Price,
		fakeProductDBResponse.Description,
	)

	mock.ExpectQuery(regexp.QuoteMeta("select id, name, price, description from product where id =$1")).
		WithArgs(fakeProductDBResponse.ID).
		WillReturnRows(rows)

	sut := productRepository.New(mock)
	product, err := sut.FindById(fakeProductDBResponse.ID)

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	require.Nil(t, err)
	require.Equal(t, product.ID, fakeProductDBResponse.ID)
	require.Equal(t, product.Name, fakeProductDBResponse.Name)
	require.Equal(t, product.Description, fakeProductDBResponse.Description)

}

func TestFindById_QueryError(t *testing.T) {
	_, _, mock := setupFindById()
	defer mock.Close()

	mock.ExpectQuery("select id, name, price, description from product where id = \\$1").
		WithArgs(1).
		WillReturnError(fmt.Errorf("ANY QUERY ERROR"))

	sut := productRepository.New(mock)
	_, err := sut.FindById(1)

	require.NotNil(t, err)

}
