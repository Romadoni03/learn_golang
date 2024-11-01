package config_test

import (
	"ecommerce-cloning-app/internal/config"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestOpenConnection(t *testing.T) {

	db, err := config.NewDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
