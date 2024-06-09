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
	GetMerchantCountByIds(ids string) int
	GetItemCountByIds(ids string) int
}

type PurchaseUsecaseInterface interface {
	GetNearbyMerchants(long float64, lat float64, param dto.RequestNearbyMerchants) (dto.ResponseNearbyMerchants, error)
	CreateEstimation(request dto.RequestEstimate, userId string) (res dto.ResponseOrder, err error)
	CreateOrder(id string) (orderId string, err error)
	GetOrders(filter dto.RequestGetOrders, userId string) (res []dto.ResponseGetOrders, err error)
}
