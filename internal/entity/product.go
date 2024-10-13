package entity

import "time"

type Product struct {
	Id                string
	StoreId           string
	PhotoProduct      string
	Name              string
	Category          string
	Description       string
	DangeriousProduct string
	Price             int
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
