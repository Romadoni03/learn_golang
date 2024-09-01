package config

import (
	"database/sql"
	"time"
)

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/portofolio_golang")

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	if err != nil {
		return nil, err
	}

	return db, nil
}
