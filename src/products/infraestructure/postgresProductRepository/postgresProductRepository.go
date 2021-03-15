package postgresproductrepository

import (
	"database/sql"

	"github.com/alejogs4/hn-website/src/products/domain/product"
)

// PostgresProductCommandsRepository .
type PostgresProductCommandsRepository struct {
	db *sql.DB
}

// NewPostgresProductCommandsRepository returns a new instance of PostgresProductCommandsRepository
func NewPostgresProductCommandsRepository(db *sql.DB) PostgresProductCommandsRepository {
	return PostgresProductCommandsRepository{db: db}
}

// CreateProduct insert product into the products table
func (ppp PostgresProductCommandsRepository) CreateProduct(newProduct product.Product) error {
	_, err := ppp.db.Exec(`
		INSERT INTO products(id, name, description, picture, state, quantity, price)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`,
		newProduct.ID(),
		newProduct.Name(),
		newProduct.Description(),
		newProduct.Picture(),
		newProduct.State(),
		newProduct.Quantity(),
		newProduct.Price(),
	)

	return err
}

// UpdateProduct .
func (ppp PostgresProductCommandsRepository) UpdateProduct(updatedProduct product.Product) error {
	_, err := ppp.db.Exec(`
		UPDATE products SET name=$1, description=$2, picture=$3, state=$4, quantity=$5, price=$6 WHERE id=$7
	`,
		updatedProduct.Name(),
		updatedProduct.Description(),
		updatedProduct.Picture(),
		updatedProduct.State(),
		updatedProduct.Quantity(),
		updatedProduct.Price(),
		updatedProduct.ID(),
	)

	return err
}
