package dto

type UserLoginResponse struct {
	NoTelepon      string `json:"no_telepon"`
	Username       string `json:"username"`
	Name           string `json:"name"`
	Token          string `json:"token"`
	TokenExpiredAt int64  `json:"token_expired_at"`
}
