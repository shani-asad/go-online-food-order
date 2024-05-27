package usecase

import (
	"context"
	"online-food/helpers"
	"online-food/model/database"
	"online-food/model/dto"
	"online-food/src/repository"
	"strconv"
	"time"
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
	data := database.Merchant{
		Name:             request.Name,
		MerchantCategory: request.MerchantCategory,
		ImageUrl:         request.ImageUrl,
		LocationLat:      request.Location.Lat,
		LocationLong:     request.Location.Long,
		CreatedAt:        time.Now(),
	}

	id, err := u.merchantRepository.CreateMerchant(context.TODO(), data)
	if err == nil || id != 0 {
		res.MerchantID = strconv.Itoa(id)
	}

	return res, err
}
