package repository

import (
	"context"
	"database/sql"
	"ecommerce-cloning-app/internal/entity"
	"fmt"
)

type UserRepositoryImpl struct {
}

func (repository *UserRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, user entity.User) error {
	SQL := "insert into users(user_id, no_telepon, password, username, last_updated_username, name, email, photo_profile, bio, gender, status_member, birth_date, created_at, token, token_expired_at) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

	_, err := tx.ExecContext(ctx, SQL, user.Id, user.NoTelepon, user.Password, user.Username, user.LastUpdatedUsername, user.Name, user.Email, user.PhotoProfile, user.Bio, user.Gender, user.StatusMember, user.BirthDate, user.CreatedAt, user.Token, user.TokenExpiredAt)

	if err != nil {
		return err
	} else {
		fmt.Println("success add to database")
		return nil
	}
}
