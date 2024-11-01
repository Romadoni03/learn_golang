package dto

type UserLoginResponse struct {
	Message     string `json:"message"`
	NoTelepon   string `json:"no_telepon"`
	Username    string `json:"username"`
	Name        string `json:"name"`
	AccessToken string `json:"access_token"`
}
