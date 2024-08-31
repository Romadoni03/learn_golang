package repository

import (
	"context"
	"database/sql"
	"ecommerce-cloning-app/internal/entity"
)

type UserRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, user entity.User) entity.User
}
