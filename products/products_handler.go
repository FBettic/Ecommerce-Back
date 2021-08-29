package products

import (
	"ecommerce-back/internal/database"
	"ecommerce-back/internal/logs"
)

type ProductHandlerInterface interface {
	GetProducts() (*ProductsList, error)
	GetProductsById(int64) (*Product, error)
	CreateProduct(*Product) error
	UpdateProduct(*Product, error) error
	DeleteProduct(int64) error
}

type ProductHandler struct {
	ProductStorage
}

func (ps *ProductHandler) GetProducts() (*ProductsList, error) {

	products, err := ps.ProductStorage.FindAll()

	if err != nil {
		logs.Log().Errorf("cannot obtain products list: %s", err.Error())
		return nil, err
	}

	return products, nil
}

func (ps *ProductHandler) GetProductById(id int64) (*Product, error) {

	product, err := ps.ProductStorage.FindByID(id)

	if err != nil {
		logs.Log().Errorf("cannot obtain product: %s", err.Error())
		return nil, err
	}

	return product, nil
}

func (ps *ProductHandler) CreateProduct(product *Product) error {
	err := ps.ProductStorage.Add(product)
	if err != nil {
		logs.Log().Errorf("cannot create product: %s", err.Error())
		return err
	}
	return nil
}

func (ps *ProductHandler) UpdateProduct(product *Product, id int64) error {

	err := ps.ProductStorage.Update(product, id)

	if err != nil {
		logs.Log().Errorf("cannot update product: %s", err.Error())
		return err
	}

	return nil
}

func (ps *ProductHandler) DeleteProductById(id int64) error {
	err := ps.ProductStorage.DeleteById(id)
	if err != nil {
		logs.Log().Errorf("cannot delete product: %s", err.Error())
		return err
	}
	return nil
}

func GetProductHandler(db *database.MySQL) ProductHandler {
	return ProductHandler{GetProductStorage(db)}
}
