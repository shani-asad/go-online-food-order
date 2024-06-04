package usecase

import (
	"context"
	"online-food/model/dto"
	"online-food/src/repository"
)

type PurchaseUsecase struct {
	merchantRepository repository.MerchantRepositoryInterface
}

func NewPurchaseUsecase(
	merchantRepository repository.MerchantRepositoryInterface) PurchaseUsecaseInterface {
	return &PurchaseUsecase{merchantRepository}
}

func (u *PurchaseUsecase) GetNearbyMerchants(long float64, lat float64, request dto.RequestNearbyMerchants) (res dto.ResponseNearbyMerchants, err error) {

	res, err = u.merchantRepository.GetNearbyMerchants(context.TODO(), long, lat, request)
	return res, err
}
