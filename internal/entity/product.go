package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type Product struct {
	Id                string
	StoreId           string
	PhotoProduct      string
	Name              string
	Category          string
	Description       string
	DangeriousProduct string
	Price             decimal.Decimal
	Stock             int
	Wholesaler        string
	ShippingCost      int
	ShippingInsurance int
	Conditions        string
	PreOrder          string
	Status            string
	CreatedAt         time.Time
	LastUpdatedAt     time.Time
}
