package usecase

import (
	"online-food/helpers"
	"online-food/model/dto"
	"online-food/src/repository"
)

type MerchantUsecase struct {
	merchantRepository repository.MerchantRepositoryInterface
	helper             helpers.HelperInterface
}

func NewMerchantUsecase(
	merchantRepository repository.MerchantRepositoryInterface,
	helper helpers.HelperInterface) MerchantUsecaseInterface {
	return &MerchantUsecase{merchantRepository, helper}
}

func (u *MerchantUsecase) CreateMerchant(request dto.RequestCreateMerchant) (res dto.ResponseCreateMerchant, err error) {
	return res, err
}
