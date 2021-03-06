package cart

import "github.com/alejogs4/hn-website/src/products/domain/product"

type ItemState string

const (
	ADDED   ItemState = "ADDED"
	REMOVED ItemState = "REMOVED"
)

// Item modeled the products added to a user cart
type Item struct {
	id      string
	product product.Product
	state   ItemState
}

func NewCartItem(id string, product product.Product, state ItemState) Item {
	return Item{id: id, product: product, state: state}
}

func (item *Item) ID() string {
	return item.id
}

func (item *Item) ProductID() string {
	return item.product.ID()
}

func (item *Item) Product() product.Product {
	return item.product
}

func (item *Item) State() string {
	return string(item.state)
}
