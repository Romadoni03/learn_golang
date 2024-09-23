package repository

import (
	"context"
	"database/sql"
	"ecommerce-cloning-app/internal/entity"
)

type StoreRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, store entity.Store) entity.Store
}
