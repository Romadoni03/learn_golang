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

func (repository *UserRepositoryImpl) UpdateToken(ctx context.Context, tx *sql.Tx, user entity.User) error {
	SQL := "update users set token = ?, token_expired_at = ? where no_telepon = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Token, user.TokenExpiredAt, user.NoTelepon)
	if err != nil {
		return err
	}
	return nil
}

func (repository *UserRepositoryImpl) FindFirstByToken(ctx context.Context, tx *sql.Tx, token string) (entity.User, error) {
	SQL := "select no_telepon, token, token_expired_at from users where token = ?"
	rows, err := tx.QueryContext(ctx, SQL, token)
	helper.IfPanicError(err)
	defer rows.Close()

	user := entity.User{}
	if rows.Next() {
		errNext := rows.Scan(&user.NoTelepon, &user.Token, &user.TokenExpiredAt)
		helper.IfPanicError(errNext)
		return user, nil
	} else {
		return user, errors.New("user by token is not found")
	}

}

func (repository *UserRepositoryImpl) DeleteToken(ctx context.Context, tx *sql.Tx, token string) error {
	SQL := "update users set token = '', token_expired_at = 0 where token = ?"
	_, err := tx.ExecContext(ctx, SQL, token)
	if err != nil {
		return err
	}
	return nil
}

func (repository *UserRepositoryImpl) GetByToken(ctx context.Context, tx *sql.Tx, token string) (entity.User, error) {
	SQL := "select username, last_updated_username, users.name, email, users.no_telepon, photo_profile, stores.name, gender, birth_date from users LEFT JOIN stores ON users.no_telepon = stores.no_telepon where token = ?"
	rows, err := tx.QueryContext(ctx, SQL, token)
	helper.IfPanicError(err)
	defer rows.Close()

	user := entity.User{}
	// store := entity.Store{}
	if rows.Next() {
		errNext := rows.Scan(&user.Username, &user.LastUpdatedUsername, &user.Name, &user.Email, &user.NoTelepon, &user.PhotoProfile, &user.Store.Name, &user.Gender, &user.BirthDate)
		helper.IfPanicError(errNext)
		// user.Store = store
		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user entity.User) error {
	SQL := "UPDATE users LEFT JOIN stores ON users.no_telepon = stores.no_telepon set username = ?, last_updated_username = ?, users.name = ?, photo_profile = ?, stores.name = ?, gender = ?, birth_date = ? where users.no_telepon = ?"

	_, err := tx.ExecContext(ctx, SQL, user.Username, user.LastUpdatedUsername, user.Name, user.PhotoProfile, user.Store.Name, user.Gender, user.BirthDate, user.NoTelepon)
	if err != nil {
		return err
	}
	return nil
}
