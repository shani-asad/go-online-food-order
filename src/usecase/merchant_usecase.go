package usecase

import (
	"online-food/helpers"
	"online-food/model/dto"
)

type MerchantUsecase struct {
	helper helpers.AuthHelperInterface
}

func NewMerchantUsecase(
	helper helpers.AuthHelperInterface) MerchantUsecaseInterface {
	return &MerchantUsecase{helper}
}

func (u *MerchantUsecase) CreateMerchant(request dto.RequestCreateMerchant) (res dto.ResponseCreateMerchant, err error) {
	return res, err
}
