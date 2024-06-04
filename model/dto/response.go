package dto

import "time"

type ResponseStatusAndMessage struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ResponseCreateMerchant struct {
	MerchantID string `json:"merchantId"`
}

type ResponseMeta struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

type ResponseMerchant struct {
	MerchantId       int       `json:"merchantId"`
	Name             string    `json:"name"`
	MerchantCategory string    `json:"merchantCategory"`
	ImageUrl         string    `json:"imageUrl"`
	Location         Location  `json:"location"`
	CreatedAt        time.Time `json:"createdAt"`
}

type ResponseGetMerchants struct {
	Data []ResponseMerchant `json:"data"`
	Meta ResponseMeta       `json:"meta"`
}

type ResponseCreateMerchantItem struct {
	ItemID string `json:"itemId"`
}

type ResponseGetItems struct {
	ItemID          string    `json:"itemId"`
	Name            string    `json:"name"`
	ProductCategory string    `json:"productCategory"`
	Price           int       `json:"price"`
	ImageUrl        string    `json:"imageUrl"`
	CreatedAt       time.Time `json:"craetedAt"`
}

type ResponseGetMerchantItems struct {
	Data []ResponseGetItems `json:"data"`
	Meta ResponseMeta       `json:"meta"`
}
type Item struct {
	ItemId			string	`json:"itemId"`
	Name			string	`json:"name"`
	ProductCategory	string	`json:"productCategory"`
	Price			int		`json:"price"`
	ImageUrl		string	`json:"imageUrl"`
	CreatedAt		string	`json:"createdAt"`
}
type NearbyMerchantsDbResponse struct {
	Merchant	ResponseMerchant	`json:"merchant"`
	Items		Item				`json:"items"`
}

type NearbyMerchants struct {
	Merchant	ResponseMerchant	`json:"merchant"`
	Items		[]Item				`json:"items"`
}

type ResponseNearbyMerchants struct {
	Data	[]NearbyMerchants	`json:"data"`
	Meta	ResponseMeta		`json:"meta"`
}
