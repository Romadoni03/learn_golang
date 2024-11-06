package dto

type UserUpdateRequest struct {
	Username     string `json:"username"`
	Name         string `json:"name"`
	PhotoProfile string `json:"photo_profile"`
	NameStore    string `json:"name_store"`
	Gender       string `json:"gender"`
	BirthDate    string `json:"birth_date"`
}
