package usecase

import (
	"context"
	"online-food/model/database"
	"online-food/model/dto"
	"online-food/src/repository"
	"strconv"
	"time"
)

type MerchantUsecase struct {
	merchantRepository repository.MerchantRepositoryInterface
}

func NewMerchantUsecase(
	merchantRepository repository.MerchantRepositoryInterface) MerchantUsecaseInterface {
	return &MerchantUsecase{merchantRepository}
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

func (u *MerchantUsecase) GetMerchants(request dto.RequestGetMerchant) (res dto.ResponseGetMerchants, err error) {
	data, err := u.merchantRepository.GetMerchants(context.TODO(), request)
	if err != nil {
		return res, err
	}

	for _, v := range data {
		merchant := dto.ResponseMerchant{
			MerchantId:       v.ID,
			Name:             v.Name,
			MerchantCategory: v.MerchantCategory,
			ImageUrl:         v.ImageUrl,
			Location: dto.Location{
				Lat:  v.LocationLat,
				Long: v.LocationLong,
			},
			CreatedAt: v.CreatedAt,
		}

		res.Data = append(res.Data, merchant)
	}
	res.Meta.Limit = *request.Limit
	res.Meta.Offset = *request.Offset
	res.Meta.Total = len(res.Data)

	return res, err
}

func (u *MerchantUsecase) CreateMerchantItem(request dto.RequestCreateMerchantItem) (res dto.ResponseCreateMerchantItem, err error) {
	return res, err
}

func (u *MerchantUsecase) GetMerchantItems(request dto.RequestGetMerchantItems) (res dto.ResponseGetMerchantItems, err error) {
	return res, err
}


