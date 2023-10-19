package usecase

import (
	"desafiogrpc/domain/model"
)

type ProductsUseCase struct {
	ProductsRepository model.ProductsRepositoryInterface
}

func (p *ProductsUseCase) AddProduct(name string, description string, price float64) (*model.Product, error) {
	product, err := model.NewProduct(name, description, price)
	if err != nil {
		return nil, err
	}

	err = p.ProductsRepository.AddProduct(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductsUseCase) FindProductByName(name string) (*model.Product, error) {
	product, err := p.ProductsRepository.FindProductByName(name)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductsUseCase) FindProductById(id string) (*model.Product, error) {
	product, err := p.ProductsRepository.FindProductById(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductsUseCase) FindAllProducts() ([]model.Product, error) {
	products, err := p.ProductsRepository.FindAllProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (usecase *ProductsUseCase) DeleteProduct(id string) error {
	return usecase.ProductsRepository.DeleteProduct(id)

}
