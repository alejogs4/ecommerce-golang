package postgresproductrepository

import (
	"database/sql"

	"github.com/alejogs4/hn-website/src/products/domain/product"
)

// ProductQueriesPostgresRepository struct to query information about products
type ProductQueriesPostgresRepository struct {
	db *sql.DB
}

// NewProductQueriesPostgresRepository returns a new instance of ProductQueriesPostgresRepository
func NewProductQueriesPostgresRepository(db *sql.DB) ProductQueriesPostgresRepository {
	return ProductQueriesPostgresRepository{db: db}
}

// GetAll .
func (pqpr ProductQueriesPostgresRepository) GetAll() ([]product.Product, error) {
	rows, err := pqpr.db.Query(
		"SELECT id, name, description, picture, state, quantity, price, created_at FROM products WHERE state=$1",
		product.Active,
	)
	if err != nil {
		return []product.Product{}, err
	}

	defer rows.Close()
	var products []product.Product

	for rows.Next() {
		var id string
		var name string
		var description string
		var picture string
		var productState string
		var quantity int
		var price float64
		var createdAtDate string

		err := rows.Scan(
			&id,
			&name,
			&description,
			&picture,
			&productState,
			&quantity,
			&price,
			&createdAtDate,
		)

		if err != nil {
			return []product.Product{}, err
		}
		gotProduct, err := product.FromPrimitives(
			id,
			name,
			description,
			picture,
			quantity,
			price,
			productState,
			createdAtDate,
		)

		if err != nil {
			return []product.Product{}, err
		}
		products = append(products, gotProduct)
	}

	return products, nil
}

// GetByID .
func (pqpr ProductQueriesPostgresRepository) GetByID(id string) (product.Product, error) {
	result := pqpr.db.QueryRow(
		"SELECT name, description, picture, state, quantity, price, created_at FROM products WHERE id=$1 AND state=$2",
		id, product.Active,
	)

	var name string
	var description string
	var picture string
	var productState string
	var quantity int
	var price float64
	var createdAtDate string

	err := result.Scan(
		&name,
		&description,
		&picture,
		&productState,
		&quantity,
		&price,
		&createdAtDate,
	)

	if err != nil {
		return product.Product{}, err
	}

	return product.FromPrimitives(
		id,
		name,
		description,
		picture,
		quantity,
		price,
		productState,
		createdAtDate,
	)
}
