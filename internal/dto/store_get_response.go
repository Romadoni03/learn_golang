package dto

type StoreGetResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Logo        string `json:"logo"`
	Description string `json:"description"`
	NoTelepon   string `json:"no_telepon"`
	Email       string `json:"email"`
}
