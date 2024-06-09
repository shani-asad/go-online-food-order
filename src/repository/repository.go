package repository

import (
	"context"
	"online-food/model/database"
	"online-food/model/dto"
)

type UserRepositoryInterface interface {
	GetUserByUsername(ctx context.Context, username string) (response database.User, err error)
	CreateUser(ctx context.Context, data database.User) (id int, err error)
	GetExistingUserInTheRoleByEmail(ctx context.Context, email, role string) (response database.User, err error)
}

type MerchantRepositoryInterface interface {
	CreateMerchant(ctx context.Context, data database.Merchant) (id int, err error)
	GetMerchants(ctx context.Context, filter dto.RequestGetMerchant) ([]database.Merchant, error)
	CreateMerchantItem(ctx context.Context, data database.Item) (id int, err error)
	GetMerchantItems(ctx context.Context, filter dto.RequestGetMerchantItems) ([]database.Item, error)
	GetMerchantLocations(ctx context.Context, merchantIdsString string) (locationMap map[string]dto.Location, err error)
	GetTotalPriceOfItems(ctx context.Context, idsString string) (int, error)
	GetNearbyMerchants(ctx context.Context, long float64, lat float64, filter dto.RequestNearbyMerchants) (response dto.ResponseNearbyMerchants, err error)
	GetMerchantCountByIds(ctx context.Context, ids string) int
	GetItemCountByIds(ctx context.Context, ids string) int
}

type OrderRepositoryInterface interface {
	CreateEstimation(ctx context.Context, filter dto.RequestEstimate, userId string) (int, error)
	CreateOrder(ctx context.Context, estimationId string) (res string, err error)
	GetOrders(ctx context.Context, filter dto.RequestGetOrders, userId string) (res []dto.ResponseGetOrders, err error)
}
