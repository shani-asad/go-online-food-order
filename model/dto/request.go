package dto

type RequestAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RequestCreateUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type RequestLocation struct {
	Lat  float64 `json:"lat" validate:"required"`
	Long float64 `json:"long" validate:"required"`
}

type RequestCreateMerchant struct {
	Name             string          `json:"name" validate:"required,min=2,max=30"`
	MerchantCategory string          `json:"merchantCategory" validate:"required,oneof=SmallRestaurant MediumRestaurant LargeRestaurant MerchandiseRestaurant BoothKiosk ConvenienceStore"`
	ImageUrl         string          `json:"imageUrl" validate:"required,completeURL"`
	Location         RequestLocation `json:"location" validate:"required"`
}

type RequestGetMerchant struct {
	MerchantID       *string `json:"merchantId" query:"merchantId"`
	Limit            *int    `json:"limit" query:"limit" validate:"omitempty,min=1"`
	Offset           *int    `json:"offset" query:"offset" validate:"omitempty,min=0"`
	Name             *string `json:"name" query:"name"`
	MerchantCategory *string `json:"merchantCategory" query:"merchantCategory" validate:"omitempty,oneof=SmallRestaurant MediumRestaurant LargeRestaurant MerchandiseRestaurant BoothKiosk ConvenienceStore"`
	CreatedAt        *string `json:"createdAt" query:"createdAt" validate:"omitempty,oneof=asc desc"`
}
