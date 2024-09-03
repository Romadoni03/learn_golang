package entity

import "time"

type User struct {
	Id                  string
	NoTelepon           string
	Password            string
	Username            string
	LastUpdatedUsername time.Time
	Name                string
	Email               string
	PhotoProfile        string
	Bio                 string
	Gender              string
	StatusMember        string
	BirthDate           string
	CreatedAt           time.Time
	Token               string
	TokenExpiredAt      int64
}
