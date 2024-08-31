package test

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/portofolio_golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
