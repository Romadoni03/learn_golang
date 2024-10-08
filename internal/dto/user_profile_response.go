package dto

type UserProfileResponse struct {
	Username     string `json:"username"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	NoTelepon    string `json:"no_telepon"`
	PhotoProfile string `json:"photo_profile"`
	NameStore    string `json:"name_store"`
	Gender       string `json:"gender"`
	BirthDate    string `json:"birt_date"`
}
