package database

import "time"

type User struct {
	Id			string
	Nip		int64
	Password	string
	Name		string
	Role string
	IdentityCardScanImg string
	CreatedAt 	time.Time
}