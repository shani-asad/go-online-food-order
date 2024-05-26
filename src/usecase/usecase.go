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
}
