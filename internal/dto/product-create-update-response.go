package dto

import "github.com/shopspring/decimal"

type ProductCreateUpdateResponse struct {
	Id                string          `validate:"required" json:"id"`
	PhotoProduct      string          `validate:"required" json:"photo"`
	Name              string          `validate:"required" json:"name"`
	Category          string          `validate:"required" json:"category"`
	Description       string          `validate:"required" json:"description"`
	DangeriousProduct string          `validate:"required" json:"dangerious_product"`
	Price             decimal.Decimal `validate:"required" json:"price"`
	Stock             int             `validate:"required" json:"stock"`
	Wholesaler        string          `validate:"required" json:"wholesaler"`
	ShippingCost      decimal.Decimal `validate:"required" json:"shipping_cost"`
	ShippingInsurance int             `validate:"required" json:"shipping_insurance"`
	Conditions        string          `validate:"required" json:"condition"`
	PreOrder          string          `validate:"required" json:"pre_order"`
	Status            string          `validate:"required" json:"status"`
}
