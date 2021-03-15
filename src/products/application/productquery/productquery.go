package productquery

import (
	productsdto "github.com/alejogs4/hn-website/src/products/application/productsDTO"
	"github.com/alejogs4/hn-website/src/products/domain/product"
)

// ProductQueriesUseCases struct to handle all product get operations
type ProductQueriesUseCases struct {
	productQueries product.QueriesRepository
}

// NewProductQueriesUseCases returns a new instance of ProductQueriesUseCases
func NewProductQueriesUseCases(productQueries product.QueriesRepository) ProductQueriesUseCases {
	return ProductQueriesUseCases{productQueries: productQueries}
}

// GetAllProducts .
func (p ProductQueriesUseCases) GetAllProducts() ([]productsdto.ProductDTO, error) {
	products, err := p.productQueries.GetAll()
	if err != nil {
		return []productsdto.ProductDTO{}, err
	}

	var productsDTO []productsdto.ProductDTO
	for _, storedProduct := range products {
		productsDTO = append(productsDTO, productsdto.FromEntity(storedProduct))
	}

	return productsDTO, nil
}

// GetProductByID .
func (p ProductQueriesUseCases) GetProductByID(id string) (productsdto.ProductDTO, error) {
	product, err := p.productQueries.GetByID(id)
	if err != nil {
		return productsdto.ProductDTO{}, err
	}

	return productsdto.FromEntity(product), nil
}
