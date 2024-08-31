package dto

type UserCreateRequest struct {
	NoTelepon string `validate:"required" json:"no_telepon"`
	Password  string `validate:"required" json:"password"`
}
