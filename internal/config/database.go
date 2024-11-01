package config

import (
	"database/sql"
	"fmt"
	"os"
	"time"
)

func NewDB() (*sql.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dataSource := fmt.Sprintf("%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dataSource)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	if err != nil {
		return nil, err
	}

	return db, nil
}
