package repository

import (
	"context"
	"database/sql"
	"ecommerce-cloning-app/internal/entity"
	"ecommerce-cloning-app/internal/logger"
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
