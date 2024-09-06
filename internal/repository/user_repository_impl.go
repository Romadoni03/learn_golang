package repository

import (
	"context"
	"database/sql"
	"ecommerce-cloning-app/internal/entity"
	"ecommerce-cloning-app/internal/helper"
	"errors"
	"fmt"
)

type UserRepositoryImpl struct {
}

func (repository *UserRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, user entity.User) error {
	SQL := "insert into users( no_telepon, password, username, last_updated_username, name, email, photo_profile, bio, gender, status_member, birth_date, created_at, token, token_expired_at) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

	_, err := tx.ExecContext(
		ctx,
		SQL,
		user.NoTelepon,
		user.Password,
		user.Username,
		user.LastUpdatedUsername,
		user.Name,
		user.Email,
		user.PhotoProfile,
		user.Bio, user.Gender,
		user.StatusMember,
		user.BirthDate,
		user.CreatedAt,
		user.Token,
		user.TokenExpiredAt)

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("success add to database")
	return nil
}

func (repository *UserRepositoryImpl) FindByPhone(ctx context.Context, tx *sql.Tx, userPhone string) (entity.User, error) {
	SQL := "select no_telepon, password, username from users where no_telepon = ?"
	rows, err := tx.QueryContext(ctx, SQL, userPhone)
	helper.IfPanicError(err)
	defer rows.Close()

	user := entity.User{}
	if rows.Next() {
		err := rows.Scan(&user.NoTelepon, &user.Password, &user.Username)
		helper.IfPanicError(err)
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}
