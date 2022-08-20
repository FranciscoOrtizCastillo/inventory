package repository

import (
	context "context"

	entity "github.com/FranciscoOrtizCastillo/inventory/internal/entity"
)

const (
	queryInsertProduct = `
		INSERT INTO products (name, description, price, created_by) 
		VALUES (?,?,?,?);`

	queryGetProducts = `
		SELECT id, name, description, price, created_by FROM products;`

	queryGetProductByID = `
		SELECT id, name, description, price, created_by FROM products WHERE id = ?;`
)

func (r *repo) SaveProduct(ctx context.Context, name, description string, price float32, createdBy int64) error {
	_, err := r.db.ExecContext(ctx, queryInsertProduct, name, description, price, createdBy)
	return err
}

func (r *repo) GetProducts(ctx context.Context) ([]entity.Product, error) {

	//En Goland las variables se nombran con un caracter minuscula, y si es plurar se duplica
	pp := []entity.Product{}

	err := r.db.SelectContext(ctx, &pp, queryGetProducts)

	if err != nil {
		return nil, err
	}

	return pp, err
}

func (r *repo) GetProductByID(ctx context.Context, id int64) (*entity.Product, error) {
	p := &entity.Product{}

	err := r.db.GetContext(ctx, p, queryGetProductByID, id)

	if err != nil {
		return nil, err
	}

	return p, err
}
