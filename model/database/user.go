package database

import "time"

type User struct {
	ID        string
	Username  string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
}
