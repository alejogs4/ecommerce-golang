package product_test

import (
	"errors"
	"testing"

	"github.com/alejogs4/hn-website/src/products/domain/product"
	"github.com/icrowley/fake"
)

func TestNewProduct(t *testing.T) {
	testCases := []struct {
		Name               string
		ID                 string
		ProductName        string
		ProductDescription string
		ProductPicture     string
		Quantity           int
		Price              int
		State              product.State
		ExpectedError      error
	}{
		{
			Name:               "Should return error if id is empty",
			ID:                 "",
			ProductName:        fake.ProductName(),
			ProductDescription: fake.Product(),
			ProductPicture:     fake.TopLevelDomain(),
			Quantity:           10,
			Price:              10,
			ExpectedError:      product.ErrBadProductData,
		},
		{
			Name:               "Should return error if name is empty",
			ID:                 fake.IPv4(),
			ProductName:        "",
			ProductDescription: fake.Product(),
			ProductPicture:     fake.TopLevelDomain(),
			Quantity:           10,
			Price:              10,
			ExpectedError:      product.ErrBadProductData,
		},
		{
			Name:               "Should return error if description is empty",
			ID:                 fake.IPv4(),
			ProductName:        fake.ProductName(),
			ProductDescription: "",
			ProductPicture:     fake.TopLevelDomain(),
			Quantity:           10,
			Price:              10,
			ExpectedError:      product.ErrBadProductData,
		},
		{
			Name:               "Should return error if picture is empty",
			ID:                 fake.IPv4(),
			ProductName:        fake.ProductName(),
			ProductDescription: fake.Brand(),
			ProductPicture:     "",
			Quantity:           10,
			Price:              10,
			ExpectedError:      product.ErrBadProductData,
		},
		{
			Name:               "Should return error if quantity it's smaller than zero",
			ID:                 fake.IPv4(),
			ProductName:        fake.ProductName(),
			ProductDescription: fake.Brand(),
			ProductPicture:     fake.TopLevelDomain(),
			Quantity:           -10,
			Price:              10,
			ExpectedError:      product.ErrProductQuantity,
		},
		{
			Name:               "Should return error if price it's cheaper than zero",
			ID:                 fake.IPv4(),
			ProductName:        fake.ProductName(),
			ProductDescription: fake.Brand(),
			ProductPicture:     fake.TopLevelDomain(),
			Quantity:           10,
			Price:              -10,
			ExpectedError:      product.ErrZeroOrNegativeProductPrice,
		},
		{
			Name:               "Should return error if price it's equal to zero",
			ID:                 fake.IPv4(),
			ProductName:        fake.ProductName(),
			ProductDescription: fake.Brand(),
			ProductPicture:     fake.TopLevelDomain(),
			Quantity:           10,
			Price:              0,
			ExpectedError:      product.ErrZeroOrNegativeProductPrice,
		},
		{
			Name:               "Should return state in active if quantity is bigger than zero",
			ID:                 fake.IPv4(),
			ProductName:        fake.ProductName(),
			ProductDescription: fake.Brand(),
			ProductPicture:     fake.TopLevelDomain(),
			Quantity:           10,
			Price:              20,
			State:              product.Active,
			ExpectedError:      nil,
		},
		{
			Name:               "Should return state in unavailable if quantity is zero",
			ID:                 fake.IPv4(),
			ProductName:        fake.ProductName(),
			ProductDescription: fake.Brand(),
			ProductPicture:     fake.TopLevelDomain(),
			Quantity:           0,
			Price:              20,
			State:              product.UnAvailable,
			ExpectedError:      nil,
		},
	}

	for _, c := range testCases {
		t.Run(c.Name, func(t *testing.T) {
			gotProduct, err := product.NewProduct(
				c.ID,
				c.ProductName,
				c.ProductDescription,
				c.ProductPicture,
				c.Quantity,
				float64(c.Price),
			)

			if !errors.Is(err, c.ExpectedError) {
				t.Fatalf("Error: Expected error %v, got error %v", c.ExpectedError, err)
			}

			if err == nil {
				if gotProduct.ID() != c.ID {
					t.Fatalf("Error: Expected id was %s, got id was %s", c.ID, gotProduct.ID())
				}

				if gotProduct.Name() != c.ProductName {
					t.Fatalf("Error: Expected name was %s, got name was %s", c.ProductName, gotProduct.Name())
				}

				if gotProduct.Description() != c.ProductDescription {
					t.Fatalf("Error: Expected description was %s, got description was %s", c.ProductDescription, gotProduct.Description())
				}

				if gotProduct.Picture() != c.ProductPicture {
					t.Fatalf("Error: Expected picture was %s, got picture was %s", c.ProductPicture, gotProduct.Picture())
				}

				if gotProduct.Quantity() != c.Quantity {
					t.Fatalf("Error: Expected quantity was %d, got quantity was %d", c.Quantity, gotProduct.Quantity())
				}

				if gotProduct.Price() != float64(c.Price) {
					t.Fatalf("Error: Expected price was %x, got price was %x", c.Price, gotProduct.Picture())
				}

				if c.State.String() != "" {
					if c.State.String() != gotProduct.State() {
						t.Fatalf("Error: Expected state was %s, got state was %s", c.State, gotProduct.State())
					}
				}
			}
		})
	}
}

func TestBuyProduct(t *testing.T) {
	testsCases := []struct {
		Name                string
		QuantityToBuy       int
		ExpectedNewQuantity int
		ExpectedError       error
		ExpectedState       string
	}{
		{
			Name:                "Should return error if quantity to bought is greather than the current quantity",
			QuantityToBuy:       20,
			ExpectedNewQuantity: 10,
			ExpectedError:       product.ErrNoEnoughQuantity,
			ExpectedState:       string(product.Active),
		},
		{
			Name:                "Should change state to unavailable if remaining quantity is 0",
			QuantityToBuy:       10,
			ExpectedNewQuantity: 0,
			ExpectedError:       nil,
			ExpectedState:       string(product.UnAvailable),
		},
		{
			Name:                "Should change only quantity if quantity to buy is smaller than current quantity",
			QuantityToBuy:       4,
			ExpectedNewQuantity: 6,
			ExpectedError:       nil,
			ExpectedState:       string(product.Active),
		},
	}

	for _, c := range testsCases {
		t.Run(c.Name, func(t *testing.T) {
			product, _ := product.NewProduct(
				fake.IPv4(),
				fake.ProductName(),
				fake.Product(),
				fake.TopLevelDomain(),
				10,
				10,
			)

			err := product.BuyProduct(c.QuantityToBuy)
			if !errors.Is(err, c.ExpectedError) {
				t.Fatalf("Error: Expected error %v, Got error %v", c.ExpectedError, err)
			}

			if c.ExpectedState != product.State() {
				t.Fatalf("Error: Expected state %s, Got state %s", c.ExpectedState, product.State())
			}

			if c.ExpectedNewQuantity != product.Quantity() {
				t.Fatalf("Error: Expected quantiy %d, Got quantity %d", c.ExpectedNewQuantity, product.Quantity())
			}
		})
	}
}
