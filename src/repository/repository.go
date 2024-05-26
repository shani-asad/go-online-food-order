package repository

import (
	"context"
	"online-food/model/database"
	"online-food/model/dto"
)

type UserRepositoryInterface interface {
	GetUserByUsername(ctx context.Context, username string) (response database.User, err error)
	CreateUser(ctx context.Context, data database.User) (err error)
	GetExistingUserInTheRoleByEmail(ctx context.Context, email, role string) (response database.User, err error)
}

type MerchantRepositoryInterface interface {
	CreateMerchant(ctx context.Context, data database.Merchant) (err error)
	GetMerchants(ctx context.Context, filter dto.RequestGetMerchant) ([]database.Merchant, error)
}
