package repository

import (
	"context"
	"database/sql"
	entity "ecommerce-cloning-app/entities"
	"ecommerce-cloning-app/internal/helper"
	"ecommerce-cloning-app/internal/logger"
	"errors"
)

type StoreRepository struct {
}

func (repository *StoreRepository) Insert(ctx context.Context, tx *sql.Tx, store entity.Store) error {
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

func (repository *StoreRepository) Delete(ctx context.Context, tx *sql.Tx, store entity.Store) error {
	SQL := "delete from stores where store_id = ?"
	_, err := tx.ExecContext(ctx, SQL, store.StoreId)
	if err != nil {
		logger.Logging().Error(err)
		return err
	}
	return nil
}

func (repository *StoreRepository) FindByUser(ctx context.Context, tx *sql.Tx, user entity.User) (entity.Store, error) {
	SQL := "select store_id, no_telepon, name, last_updated_name, logo, description, status, link_store, total_comment, total_following, total_follower, total_product, conditions, created_at from stores where no_telepon = ?"

	rows, err := tx.QueryContext(ctx, SQL, user.NoTelepon)
	helper.IfPanicError(err)
	defer rows.Close()

	store := entity.Store{}

	if rows.Next() {
		errNext := rows.Scan(&store.StoreId, &store.NoTelepon, &store.Name, &store.LastUpdatedName, &store.Logo, &store.Description, &store.Status, &store.LinkStore, &store.TotalComment, &store.TotalFollowing, &store.TotalFollower, &store.TotalProduct, &store.Condition, &store.CreatedAt)
		helper.IfPanicError(errNext)

		return store, nil
	} else {
		return store, errors.New("store not found")
	}
}
