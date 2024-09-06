package repository_test

import (
	"context"
	"database/sql"
	"ecommerce-cloning-app/internal/entity"
	"ecommerce-cloning-app/internal/helper"
	"ecommerce-cloning-app/internal/repository"
	"fmt"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func setUpDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/portofolio_golang")
	helper.IfPanicError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func truncateUser(db *sql.DB) {
	db.Exec("TRUNCATE users")
}

func TestInsert(t *testing.T) {
	db := setUpDB()
	truncateUser(db)

	user := entity.User{
		NoTelepon:           "123567898765",
		Password:            helper.HashingPassword("rahasia"),
		Username:            helper.GeneratedUsername(),
		LastUpdatedUsername: helper.GeneratedTimeNow(),
		Name:                "",
		Email:               "",
		PhotoProfile:        helper.EncodeImageName("account_profile.png"),
		Bio:                 "",
		Gender:              "",
		StatusMember:        "",
		BirthDate:           "",
		CreatedAt:           helper.GeneratedTimeNow(),
		Token:               "",
		TokenExpiredAt:      0,
	}

	repository := repository.UserRepositoryImpl{}
	tx, err := db.Begin()
	helper.IfPanicError(err)
	defer helper.CommitOrRollback(tx)
	context := context.Background()
	defer context.Done()

	errRequest := repository.Insert(context, tx, user)

	if errRequest != nil {
		fmt.Println("repository bermasalah")
	} else {
		fmt.Println("repository aman")
	}
}

// func TestFindByTelp(t *testing.T) {
// 	db := setUpDB()
// 	truncateUser(db)

// 	repository := repository.UserRepositoryImpl{}
// 	tx, err := db.Begin()
// 	helper.IfPanicError(err)
// 	defer helper.CommitOrRollback(tx)
// 	context := context.Background()
// 	defer context.Done()

// 	user, errRequest := repository.FindByTelepon(context, tx, "083156490686")

// 	if errRequest != nil {
// 		panic(errRequest)
// 	} else {
// 		fmt.Println("sukses")
// 	}

// 	fmt.Println(user.Id)
// 	fmt.Println(user.NoTelepon)
// 	fmt.Println(user.Password)
// 	fmt.Println(user.Username)

// }
