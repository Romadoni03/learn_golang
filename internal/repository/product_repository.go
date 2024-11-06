package repository

import (
	"context"
	"database/sql"
	entity "ecommerce-cloning-app/entities"
	"ecommerce-cloning-app/internal/exception"
	"ecommerce-cloning-app/internal/logger"
	"errors"
)

type ProductRepository struct {
}

func (repository *ProductRepository) Insert(ctx context.Context, tx *sql.Tx, product entity.Product) error {
	SQL := "INSERT INTO products(id,store_id,photo_product,name,category,description,dangerious_product,price,stock,wholesaler,shipping_cost,shipping_insurance,conditions,pre_order,status,created_at,last_updated_at) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

	_, err := tx.ExecContext(ctx, SQL, product.Id, product.StoreId, product.PhotoProduct, product.Name, product.Category, product.Description, product.DangeriousProduct, product.Price, product.Stock, product.Wholesaler, product.ShippingCost, product.ShippingInsurance, product.Conditions, product.PreOrder, product.Status, product.CreatedAt, product.LastUpdatedAt)

	if err != nil {
		logger.Logging().Error(err)
		return err
	}

	logger.Logging().Info("Success create Product : " + product.Name + " id : " + product.Id)
	return nil
}

func (repository *ProductRepository) FindAll(ctx context.Context, tx *sql.Tx, store entity.Store) []entity.Product {
	SQL := "select id, name, wholesaler, price, stock from products where products.store_id = ?"

	rows, err := tx.QueryContext(ctx, SQL, store.StoreId)
	if err != nil {
		logger.Logging().Error(err)
		return []entity.Product{}
	}
	defer rows.Close()

	var products []entity.Product
	for rows.Next() {
		product := entity.Product{}
		err := rows.Scan(&product.Id, &product.Name, &product.Wholesaler, &product.Price, &product.Stock)
		if err != nil {
			logger.Logging().Error(err)
			return []entity.Product{}
		}
		products = append(products, product)
	}
	return products
}

func (repository *ProductRepository) FindById(ctx context.Context, tx *sql.Tx, productId string) (entity.Product, error) {
	SQL := "select id,store_id,photo_product,name,category,description,dangerious_product,price,stock,wholesaler,shipping_cost,shipping_insurance,conditions,pre_order,status,created_at,last_updated_at from products where id = ?"

	rows, err := tx.QueryContext(ctx, SQL, productId)
	if err != nil {
		logger.Logging().Error(err)
		return entity.Product{}, err
	}
	defer rows.Close()

	product := entity.Product{}
	if rows.Next() {
		errScan := rows.Scan(&product.Id, &product.StoreId, &product.PhotoProduct, &product.Name, &product.Category, &product.Description, &product.DangeriousProduct, &product.Price, &product.Stock, &product.Wholesaler, &product.ShippingCost, &product.ShippingInsurance, &product.Conditions, &product.PreOrder, &product.Status, &product.CreatedAt, &product.LastUpdatedAt)
		if errScan != nil {
			logger.Logging().Error(errScan)
			return entity.Product{}, errScan
		}
		return product, nil
	} else {
		return product, errors.New("product is not found")
	}
}

func (repository *ProductRepository) Update(ctx context.Context, tx *sql.Tx, product entity.Product, store entity.Store) entity.Product {
	SQL := "UPDATE products SET photo_product = ?,name = ?,category = ?,description = ?,dangerious_product = ?,price = ?,stock = ?,wholesaler = ?,shipping_cost = ?,shipping_insurance = ?,conditions = ?,pre_order = ?,status = ?,last_updated_at = ? WHERE products.id = ? AND products.store_id = ?"

	_, err := tx.ExecContext(ctx, SQL, product.PhotoProduct, product.Name, product.Category, product.Description, product.DangeriousProduct, product.Price, product.Stock, product.Wholesaler, product.ShippingCost, product.ShippingInsurance, product.Conditions, product.PreOrder, product.Status, product.LastUpdatedAt, product.Id, store.StoreId)

	if err != nil {
		logger.Logging().Error(err)
		panic(exception.NewInternalServerError(err.Error()))
	}
	return product
}

func (repository *ProductRepository) Delete(ctx context.Context, tx *sql.Tx, product entity.Product, store entity.Store) error {
	SQL := "delete from products where id = ? AND store_id = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Id, store.StoreId)
	if err != nil {
		return err
	} else {
		return nil
	}
}
