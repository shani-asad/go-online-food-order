package database

import "time"

type Merchant struct {
	ID               int
	Name             string
	MerchantCategory string
	ImageUrl         string
	LocationLat      float64
	LocationLong     float64
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
