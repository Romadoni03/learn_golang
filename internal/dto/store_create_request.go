package dto

type StoreCreateRequest struct {
	Name string `validate:"required" json:"name"`
}
