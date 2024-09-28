package repository

import (
	"context"
	"database/sql"
	"ecommerce-cloning-app/internal/entity"
	"ecommerce-cloning-app/internal/logger"
)

type StoreRepositoryImpl struct {
}

func (repository *StoreRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, store entity.Store) error {
	SQL := "insert into stores(store_id, no_telepon, name, last_updated_name, logo, description, status, link_store, total_comment, total_following, total_follower, total_product, conditions, created_at) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

	_, err := tx.ExecContext(
		ctx,
		SQL,
		store.StoreId,
		store.NoTelepon,
		store.Name,
		store.LastUpdatedName,
		store.Logo,
		store.Description,
		store.Status,
		store.LinkStore,
		store.TotalComment,
		store.TotalFollowing,
		store.TotalFollower,
		store.TotalProduct,
		store.Condition,
		store.CreatedAt)

	if err != nil {
		logger.Logging().Error(err)
		return err
	}

	logger.Logging().Info("Success create store : " + store.Name.String)
	return nil
}
