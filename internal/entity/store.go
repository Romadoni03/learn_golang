package entity

import "time"

type Store struct {
	StoreId         string
	NoTelepon       string
	Name            string
	LastUpdatedName time.Time
	Logo            string
	Description     string
	Status          string
	LinkStore       string
	TotalComment    int
	TotalFollowing  int
	TotalFollower   int
	TotalProduct    int
	Condition       string
	CreatedAt       time.Time
}
