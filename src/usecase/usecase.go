package usecase

import (
	"online-food/model/dto"
)

type AuthUsecaseInterface interface {
	Register(request dto.RequestCreateUser, role string) (token string, err error)
	Login(request dto.RequestAuth, role string) (token string, status int)
	GetUserByUsername(username string) (exists bool, err error)
	GetExistingUserInTheRoleByEmail(email, role string) (exists bool, err error)
}

type MerchantUsecaseInterface interface {
	CreateMerchant(request dto.RequestCreateMerchant) (dto.ResponseCreateMerchant, error)
	GetMerchants(request dto.RequestGetMerchant) (dto.ResponseGetMerchants, error)
	CreateMerchantItem(request dto.RequestCreateMerchantItem) (dto.ResponseCreateMerchantItem, error)
	GetMerchantItems(request dto.RequestGetMerchantItems) (dto.ResponseGetMerchantItems, error)
}

type PurchaseUsecaseInterface interface {
	GetNearbyMerchants(long float64, lat float64, param dto.RequestNearbyMerchants) (dto.ResponseNearbyMerchants, error)
}
