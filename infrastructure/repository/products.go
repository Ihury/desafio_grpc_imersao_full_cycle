package repository

import (
	"desafiogrpc/domain/model"
	"fmt"
	"github.com/jinzhu/gorm"
)

type ProductsRepositoryDb struct {
	Db *gorm.DB
}

func (r ProductsRepositoryDb) AddProduct(product *model.Product) error {
	err := r.Db.Create(product).Error
	if err != nil {
		return err
	}
	return nil
}

func (r ProductsRepositoryDb) FindProductByName(name string) (*model.Product, error) {
	var product model.Product
	err := r.Db.First(&product, "name = ?", name).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, fmt.Errorf("no product found")
		}
		return nil, err
	}

	return &product, nil
}

func (r ProductsRepositoryDb) FindProductById(id string) (*model.Product, error) {
	var product model.Product
	err := r.Db.First(&product, "id = ?", id).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, fmt.Errorf("no product found")
		}
		return nil, err
	}

	return &product, nil
}

func (r ProductsRepositoryDb) FindAllProducts() ([]model.Product, error) {
	var products []model.Product
	err := r.Db.Find(&products).Error

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r ProductsRepositoryDb) DeleteProduct(id string) error {
	product, err := r.FindProductById(id)

	if err != nil {
		return err
	}

	err = r.Db.Delete(&product).Error

	if err != nil {
		return err
	}

	return nil
}
