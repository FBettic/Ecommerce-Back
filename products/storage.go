package products

import (
	"github.com/fbettic/ecommerce-back/internal/database"
	"github.com/fbettic/ecommerce-back/internal/logs"
)

type ProductStorageInterface interface {
	FindAll() (*ProductsList, error)
	FindByID(int64) (*Product, error)
	Add(*Product) error
	Update(*Product, int64) error
	DeleteById(int64) error
}

type ProductStorage struct {
	*database.MySQL
}

func (s *ProductStorage) FindAll() (*ProductsList, error) {
	tx, err := s.MySQL.Begin()

	if err != nil {
		logs.Log().Errorf("cannot create transaction: %s", err.Error())
		return nil, err
	}

	rows, err := tx.Query("SELECT * FROM products")
	if err != nil {
		logs.Log().Errorf("cannot obtain the product table: %s", err.Error())
		return nil, err
	}
	defer rows.Close()

	var p Product
	products := ProductsList{}

	for rows.Next() {
		err = rows.Scan(&p.Id, &p.Title, &p.Description, &p.Price)
		if err != nil {
			logs.Log().Errorf("cannot scan the product row: %s", err.Error())
			return nil, err
		}
		defer rows.Close()
		products = append(products, p)
	}

	return &products, nil
}

func (s *ProductStorage) FindByID(id int64) (*Product, error) {
	tx, err := s.MySQL.Begin()

	if err != nil {
		logs.Log().Errorf("cannot create transaction: %s", err.Error())
		return nil, err
	}

	row := tx.QueryRow("SELECT * FROM products WHERE id=? ", id)

	product := Product{}
	err = row.Scan(&product.Id, &product.Title, &product.Description, &product.Price)
	if err != nil {
		logs.Log().Errorf("cannot scan the product row: %s", err.Error())
		return nil, err
	}

	_ = tx.Commit()
	return &product, nil
}

func (s *ProductStorage) Add(p *Product) error {
	tx, err := s.MySQL.Begin()

	if err != nil {
		logs.Log().Errorf("cannot create transaction: %s", err.Error())
		return err
	}

	res, err := tx.Exec("INSERT INTO products (title, description, price) VALUES (?, ?, ?)", p.Title, p.Description, p.Price)

	if err != nil {
		logs.Log().Errorf("cannot execute statement: %s", err.Error())
		_ = tx.Rollback()
		return err
	}

	_, err = res.LastInsertId()
	if err != nil {
		logs.Log().Errorf("cannot fetch the last id: %s", err.Error())
		_ = tx.Rollback()
		return err
	}

	_ = tx.Commit()
	return nil
}

func (s *ProductStorage) Update(p *Product, id int64) error {
	tx, err := s.MySQL.Begin()

	if err != nil {
		logs.Log().Errorf("cannot create transaction: %s", err.Error())
		return err
	}

	res, err := tx.Exec("UPDATE products SET title = ?, description = ?, price = ? WHERE id = ?", p.Title, p.Description, p.Price, id)

	if err != nil {
		logs.Log().Errorf("cannot execute statement: %s", err.Error())
		_ = tx.Rollback()
		return err
	}

	_, err = res.LastInsertId()

	if err != nil {
		logs.Log().Errorf("cannot fetch the last id: %s", err.Error())
		_ = tx.Rollback()
		return err
	}

	_ = tx.Commit()
	return nil
}

func (s *ProductStorage) DeleteById(id int64) error {
	tx, err := s.MySQL.Begin()

	if err != nil {
		logs.Log().Errorf("cannot create transaction: %s", err.Error())
		return err
	}

	_, err = tx.Exec("DELETE FROM products WHERE id=? ", id)

	if err != nil {
		logs.Log().Errorf("cannot execute statement: %s", err.Error())
		_ = tx.Rollback()
		return err
	}

	_ = tx.Commit()
	return nil
}

func GetProductStorage(db *database.MySQL) ProductStorage {
	return ProductStorage{db}
}
