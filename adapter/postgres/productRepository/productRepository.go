package productRepository

import (
	"context"
	"log"

	"github.com/marc/workerRabbitMQ-example/adapter/postgres"
	"github.com/marc/workerRabbitMQ-example/core/domain"
)

type repository struct {
	db postgres.PoolInterface
}

// New returns contract implementation of ProductRepository
func New(db postgres.PoolInterface) domain.IProductRepository {
	return &repository{
		db: db,
	}
}

func (repository repository) FindById(id int64) (domain.Product, error) {

	ctx := context.Background()

	row, erro := repository.db.Query(
		ctx,
		"select id, name, price, description from product where id =$1",
		id,
	)

	if erro != nil {
		log.Printf("Error database : %s", erro)
		return domain.Product{}, erro
	}

	defer row.Close()

	var product domain.Product

	if row.Next() {
		if erro = row.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.Description,
		); erro != nil {
			return domain.Product{}, erro
		}
	}

	return product, nil
}
