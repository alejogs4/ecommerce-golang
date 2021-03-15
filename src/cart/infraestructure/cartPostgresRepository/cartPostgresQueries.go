package cartpostgresrepository

import (
	"database/sql"
	"errors"

	"github.com/alejogs4/hn-website/src/cart/domain/cart"
	"github.com/alejogs4/hn-website/src/products/domain/product"
)

// CartQueriesPostgresRespository struct to hold and accomplish QueriesRepository interface
type CartQueriesPostgresRespository struct {
	db *sql.DB
}

// NewCartQueriesPostgresRespository returns a new instance of CartPostgresRespository
func NewCartQueriesPostgresRespository(db *sql.DB) CartQueriesPostgresRespository {
	return CartQueriesPostgresRespository{db: db}
}

// GetUserCart get all the information of the current inprogress user cart
func (cqr CartQueriesPostgresRespository) GetUserCart(userID string) (cart.Cart, error) {
	result := cqr.db.QueryRow("SELECT id, user_id, state, created_at FROM cart WHERE user_id=$1 AND state=$2", userID, cart.InProgress)

	var cartID string
	var cartUserID string
	var cartState string
	var cartCreatedAt string

	err := result.Scan(&cartID, &cartUserID, &cartState, &cartCreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return cart.Cart{}, cart.ErrNotExistingCart
		}
		return cart.Cart{}, err
	}

	rows, err := cqr.db.Query(`
  SELECT i.id, i.state, p.id, p.name, p.description, p.picture, p.state, p.quantity, p.price, p.created_at
  FROM cart_item AS i
  INNER JOIN products AS p
  ON i.product_id = p.id
  WHERE i.cart_id=$1 AND i.state=$2 AND p.state=$3 
  `, cartID, cart.ADDED, product.Active)

	if err != nil {
		return cart.Cart{}, err
	}
	defer rows.Close()

	var cartItems []cart.Item
	for rows.Next() {
		var itemID string
		var itemState string
		var productID string
		var productName string
		var productDescription string
		var productPicture string
		var productState string
		var productQuantity int
		var productPrice float64
		var createdAt string

		err := rows.Scan(
			&itemID, &itemState, &productID, &productName, &productDescription, &productPicture, &productState, &productQuantity, &productPrice, &createdAt,
		)
		if err != nil {
			return cart.Cart{}, err
		}

		cartItemProduct, err := product.FromPrimitives(productID, productName, productDescription, productPicture, productQuantity, productPrice, productState, createdAt)
		if err != nil {
			return cart.Cart{}, err
		}
		cartItem := cart.NewCartItem(itemID, cartItemProduct, cart.ItemState(itemState))
		cartItems = append(cartItems, cartItem)
	}

	return cart.FromPrimitives(cartID, cartUserID, cartState, cartCreatedAt, cartItems)
}

// ExistUserCart returns if a user currently has inProgress cart
func (cqr CartQueriesPostgresRespository) ExistUserCart(userID string) (bool, error) {
	result := cqr.db.QueryRow("SELECT id FROM cart WHERE user_id=$1 AND state=$2", userID, cart.InProgress)

	var id string
	if err := result.Scan(&id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// GetCartItemCount get how many of a specific product there are in a cart
func (cqr CartQueriesPostgresRespository) GetCartItemCount(productID, cartID string) (int, error) {
	result := cqr.db.QueryRow(
		"SELECT COUNT(product_id) FROM cart_item WHERE cart_id=$1 AND product_id=$2 AND state=$3",
		cartID,
		productID,
		cart.ADDED,
	)

	var count int
	err := result.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// GetCartProductsCount get ho many product there are in a cart
func (cqr CartQueriesPostgresRespository) GetCartProductsCount(cartID string) (int, error) {
	result := cqr.db.QueryRow("SELECT COUNT(id) FROM cart_item WHERE cart_id=$1 AND state=$2", cartID, cart.ADDED)

	var count int
	err := result.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
