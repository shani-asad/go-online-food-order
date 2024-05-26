package database

import "time"

type User struct {
	Username			string
	Email string
	Password	string
	Role string
	CreatedAt 	time.Time
}