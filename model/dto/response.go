package dto

type ResponseStatusAndMessage struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ResponseCreateMerchant struct {
	MerchantID string `json:"merchantId"`
}
