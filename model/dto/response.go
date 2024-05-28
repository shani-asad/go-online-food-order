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
