package database

import "time"

type Item struct {
	ID              int
	Name            string
	ProductCategory string
	Price           int
	ImageUrl        string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
