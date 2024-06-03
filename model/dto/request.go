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

type Location struct {
	Lat  float64 `json:"lat" validate:"required"`
	Long float64 `json:"long" validate:"required"`
}

type RequestCreateMerchant struct {
	Name             string   `json:"name" validate:"required,min=2,max=30"`
	MerchantCategory string   `json:"merchantCategory" validate:"required,oneof=SmallRestaurant MediumRestaurant LargeRestaurant MerchandiseRestaurant BoothKiosk ConvenienceStore"`
	ImageUrl         string   `json:"imageUrl" validate:"required,completeURL"`
	Location         Location `json:"location" validate:"required"`
}

type RequestGetMerchant struct {
	MerchantID       *string `json:"merchantId" form:"merchantId"`
	Limit            *int    `json:"limit" form:"limit" validate:"omitempty,min=1"`
	Offset           *int    `json:"offset" form:"offset" validate:"omitempty,min=0"`
	Name             *string `json:"name" form:"name"`
	MerchantCategory *string `json:"merchantCategory" form:"merchantCategory" validate:"omitempty,oneof=SmallRestaurant MediumRestaurant LargeRestaurant MerchandiseRestaurant BoothKiosk ConvenienceStore"`
	CreatedAt        *string `json:"createdAt" form:"createdAt" validate:"omitempty,oneof=asc desc"`
}

type RequestCreateMerchantItem struct {
	Name            string  `json:"name" validate:"required,min=2,max=30"`
	ProductCategory string  `json:"productCategory" validate:"required,oneof=Beverage Food Snack Condiments Additions"`
	Price           float64 `json:"price" validate:"required,min=1"`
	ImageUrl        string  `json:"imageUrl" validate:"required,completeURL"`
}

type RequestGetMerchantItems struct {
	ItemID          *string `json:"itemId" validate:"omitempty"`
	Limit           *int    `json:"limit" validate:"omitempty,min=1"`
	Offset          *int    `json:"offset" validate:"omitempty,min=0"`
	Name            *string `json:"name" validate:"omitempty"`
	ProductCategory *string `json:"productCategory" validate:"omitempty,oneof=Beverage Food Snack Condiments Additions"`
	CreatedAt       *string `json:"createdAt" validate:"omitempty,oneof=asc desc"`
}

type RequestBindUrlID struct {
	ID int `uri:"id" binding:"required"`
}
