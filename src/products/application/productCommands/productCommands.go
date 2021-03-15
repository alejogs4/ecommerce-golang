package productcommands

import (
	"github.com/alejogs4/hn-website/src/products/domain/product"
	"github.com/google/uuid"
)

// ProductCommandUseCases handle the use cases which will have a direct effect in product information
type ProductCommandUseCases struct {
	productCommands product.CommandsRepository
}

// NewProductCommandsUseCases returns a new instance of ProductCommandsUseCases with productCommands repository
func NewProductCommandsUseCases(productCommands product.CommandsRepository) ProductCommandUseCases {
	return ProductCommandUseCases{productCommands: productCommands}
}

// CreateNewProduct create and handle
func (p *ProductCommandUseCases) CreateNewProduct(name, description, picture string, quantity int, price float64) (string, error) {
	productID := uuid.New().String()
	createdProduct, err := product.NewProduct(productID, name, description, picture, quantity, price)
	if err != nil {
		return "", err
	}

	err = p.productCommands.CreateProduct(createdProduct)
	if err != nil {
		return "", err
	}

	return productID, nil
}

// UpdateProduct execute update operations on an existing product
func (p *ProductCommandUseCases) UpdateProduct(id, name, description, picture string, price float64) error {
	existingProduct, err := product.RetrieveProductToUpdate(id, name, description, picture, price)
	if err != nil {
		return err
	}

	return p.productCommands.UpdateProduct(existingProduct)
}
