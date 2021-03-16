package cartpostgresrepository

import (
	"database/sql"

	"github.com/alejogs4/hn-website/src/cart/domain/cart"
)

type CartCommandsPostgresRespository struct {
	db *sql.DB
}

func NewCartCommandsPostgresRespository(db *sql.DB) CartCommandsPostgresRespository {
	return CartCommandsPostgresRespository{db: db}
}

func (ccr CartCommandsPostgresRespository) CreateCart(newCart cart.Cart) error {
	_, err := ccr.db.Exec("INSERT INTO cart(id, user_id, state) VALUES($1, $2, $3)", newCart.GetID(), newCart.GetUserID(), newCart.GetState())
	return err
}

func (ccr CartCommandsPostgresRespository) AddCartItem(item cart.Item, cartID string) error {
	_, err := ccr.db.Exec(
		"INSERT INTO cart_item(id, cart_id, product_id, state) VALUES($1, $2, $3, $4)",
		item.ID(),
		cartID,
		item.ProductID(),
		item.State(),
	)

	return err
}

func (ccr CartCommandsPostgresRespository) RemoveCartItem(cartItemID string) error {
	_, err := ccr.db.Exec(
		"UPDATE cart_item SET state=$1 WHERE id=$2",
		cart.REMOVED, cartItemID,
	)

	return err
}

func (ccr CartCommandsPostgresRespository) RemoveCart(cartID string) error {
	_, err := ccr.db.Exec("UPDATE cart SET state=$1 WHERE id=$2", cart.Removed, cartID)
	return err
}

func (ccr CartCommandsPostgresRespository) BuyCart(boughtCart cart.Cart) error {
	_, err := ccr.db.Exec("UPDATE cart SET state=$1 WHERE id=$2", cart.Ordered, boughtCart.GetID())
	return err
}
