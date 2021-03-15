package product

import (
	"time"

	"github.com/alejogs4/hn-website/src/shared/domain/valueobject"
)

// Product object entity
type Product struct {
	id          valueobject.MaybeEmpty
	name        valueobject.MaybeEmpty
	description valueobject.MaybeEmpty
	picture     valueobject.MaybeEmpty
	state       State
	quantity    int
	price       float64
	createdAt   time.Time
}

// NewProduct create a new product making the business checks, returning errors if some of these checks fails
func NewProduct(id, name, description, picture string, quantity int, price float64) (Product, error) {
	productID := valueobject.NewMaybeEmpty(id)
	productName := valueobject.NewMaybeEmpty(name)
	productDescription := valueobject.NewMaybeEmpty(description)
	productPicture := valueobject.NewMaybeEmpty(picture)

	if productID.IsEmpty() || productName.IsEmpty() || productDescription.IsEmpty() || productPicture.IsEmpty() {
		return Product{}, ErrBadProductData
	}

	if price <= 0 {
		return Product{}, ErrZeroOrNegativeProductPrice
	}

	if quantity < 0 {
		return Product{}, ErrProductQuantity
	}

	productState := Active
	if quantity == 0 {
		productState = UnAvailable
	}

	return Product{
		id:          productID,
		name:        productName,
		description: productDescription,
		picture:     productPicture,
		price:       price,
		quantity:    quantity,
		state:       productState,
		createdAt:   time.Now(),
	}, nil
}

// FromPrimitives retrieves a project from its primitives values
func FromPrimitives(id, name, description, picture string, quantity int, price float64, productState, createdAt string) (Product, error) {
	productID := valueobject.NewMaybeEmpty(id)
	productName := valueobject.NewMaybeEmpty(name)
	productDescription := valueobject.NewMaybeEmpty(description)
	productPicture := valueobject.NewMaybeEmpty(picture)

	if productID.IsEmpty() || productName.IsEmpty() || productDescription.IsEmpty() || productPicture.IsEmpty() {
		return Product{}, ErrBadProductData
	}

	if price <= 0 {
		return Product{}, ErrZeroOrNegativeProductPrice
	}

	if quantity < 0 {
		return Product{}, ErrProductQuantity
	}

	validState, err := NewState(productState)
	if err != nil {
		return Product{}, err
	}

	createdAtTime, err := time.Parse(time.RFC3339, createdAt)
	if err != nil {
		return Product{}, err
	}

	return Product{
		id:          productID,
		name:        productName,
		description: productDescription,
		picture:     productPicture,
		state:       validState,
		quantity:    quantity,
		price:       price,
		createdAt:   createdAtTime,
	}, nil
}

// RetrieveProductToUpdate get an existing product putting default quantity that will no be used in the future
func RetrieveProductToUpdate(id, name, description, picture string, price float64) (Product, error) {
	return NewProduct(id, name, description, picture, 10, price)
}

// IsUnAvailable check if product is not currently unavailable
func (p *Product) IsUnAvailable() bool {
	return p.state == UnAvailable
}

// SetProductUnAvailable .
func (p *Product) SetProductUnAvailable() {
	p.quantity = 0
	p.state = UnAvailable
}

// BuyProduct set new products quantity and its state if necessary
func (p *Product) BuyProduct(boughtProducts int) error {
	newProductQuantity := p.quantity - boughtProducts
	if newProductQuantity < 0 {
		return ErrNoEnoughQuantity
	}

	if newProductQuantity == 0 {
		p.SetProductUnAvailable()
		return nil
	}

	p.quantity = newProductQuantity
	return nil
}

// ID getter
func (p *Product) ID() string {
	return p.id.String()
}

// Name getter
func (p *Product) Name() string {
	return p.name.String()
}

// Description getter
func (p *Product) Description() string {
	return p.description.String()
}

// Picture getter
func (p *Product) Picture() string {
	return p.picture.String()
}

// State getter
func (p *Product) State() string {
	return p.state.String()
}

// Quantity getter
func (p *Product) Quantity() int {
	return p.quantity
}

// Price getter
func (p *Product) Price() float64 {
	return p.price
}

func (p *Product) CreatedAt() time.Time {
	return p.createdAt
}
