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
	CreatedAt        *string `json:"createdAt" form:"createdAt" validate:"omitempty"`
}

type RequestCreateMerchantItem struct {
	Name            string  `json:"name" validate:"required,min=2,max=30"`
	ProductCategory string  `json:"productCategory" validate:"required,oneof=Beverage Food Snack Condiments Additions"`
	Price           float64 `json:"price" validate:"required,min=1"`
	ImageUrl        string  `json:"imageUrl" validate:"required,completeURL"`
	MerchantID      int     `json:"-"`
}

type RequestGetMerchantItems struct {
	ItemID          *string `json:"itemId" form:"itemId" validate:"omitempty"`
	Limit           *int    `json:"limit" form:"limit" validate:"omitempty,min=1"`
	Offset          *int    `json:"offset" form:"offset" validate:"omitempty,min=0"`
	Name            *string `json:"name" form:"name" validate:"omitempty"`
	MerchantID      *int    `json:"-"`
	ProductCategory *string `json:"productCategory" form:"productCategory" validate:"omitempty,oneof=Beverage Food Snack Condiments Additions"`
	CreatedAt       *string `json:"createdAt" form:"createdAt" validate:"omitempty"`
}

type RequestBindUrlID struct {
	ID int `uri:"id" binding:"required"`
}

type RequestNearbyMerchants struct {
	MerchantId       *string `form:"merchantId" validate:"omitempty"`
	Limit            *int    `form:"limit" validate:"omitempty,min=1"`
	Offset           *int    `form:"offset" validate:"omitempty,min=0"`
	Name             *string `form:"name" validate:"omitempty,max=255"`
	MerchantCategory *string `form:"merchantCategory" validate:"omitempty,oneof=SmallRestaurant MediumRestaurant LargeRestaurant MerchandiseRestaurant BoothKiosk ConvenienceStore"`
}

type OrderItem struct {
	ItemId   string `json:"itemId" validate:"required"`
	Quantity int    `json:"quantity" validate:"required,min=1"`
}

type Order struct {
	MerchantId      string      `json:"merchantId" validate:"required"`
	IsStartingPoint bool        `json:"isStartingPoint"`
	Items           []OrderItem `json:"items" validate:"required,dive"`
}

type RequestEstimate struct {
	UserLocation Location `json:"userLocation" validate:"required"`
	Orders       []Order  `json:"orders" validate:"required,dive"`
}

type RequestOrder struct {
	CalculatedEstimateId string `json:"calculatedEstimateId" validate:"required"`
}

type RequestGetOrders struct {
	MerchantId       *string `form:"merchantId" validate:"omitempty"`
	Limit            *int    `form:"limit" validate:"omitempty,number,min=1"`
	Offset           *int    `form:"offset" validate:"omitempty,number,min=0"`
	Name             *string `form:"name" validate:"omitempty"`
	MerchantCategory *string `form:"merchantCategory" validate:"omitempty,oneof=SmallRestaurant MediumRestaurant LargeRestaurant MerchandiseRestaurant BoothKiosk ConvenienceStore"`
}
