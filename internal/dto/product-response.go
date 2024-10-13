package dto

import "github.com/shopspring/decimal"

type ProductRespone struct {
	Id         string          `json:"id"`
	Name       string          `json:"name"`
	Wholesaler string          `json:"wholesaler"`
	Price      decimal.Decimal `json:"price"`
	Stock      int             `json:"stock"`
}
