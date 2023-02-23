package domain

// Product is entity of table product database column
type Product struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

// ProductUseCase is a contract of business rule layer
type IProductUseCase interface {
	FindById(id int64) (Product, error)
}

// ProductRepository is a contract of database connection adapter layer
type IProductRepository interface {
	FindById(id int64) (Product, error)
}
