package model

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type ProductsRepositoryInterface interface {
	AddProduct(product *Product) error
	FindProductByName(name string) (*Product, error)
	FindProductById(id string) (*Product, error)
	FindAllProducts() ([]Product, error)
	DeleteProduct(id string) error
}

type Product struct {
	Base        `valid:"required"`
	Name        string  `json:"name" gorm:"type:varchar(255)" valid:"required"`
	Description string  `json:"description" gorm:"type:varchar(255)" valid:"required"`
	Price       float64 `json:"price" gorm:"type:float" valid:"required"`
}

func (product *Product) isValid() error {
	_, err := govalidator.ValidateStruct(product)
	if err != nil {
		return err
	}
	return nil
}

func NewProduct(name string, description string, price float64) (*Product, error) {
	product := Product{
		Name:        name,
		Description: description,
		Price:       price,
	}

	product.ID = uuid.NewV4().String()

	err := product.isValid()
	if err != nil {
		return nil, err
	}

	return &product, nil
}
