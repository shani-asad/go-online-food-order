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
	var responseMerchants []dto.ResponseMerchant
	if request.Limit == nil {
		defaultLimit := 5
		request.Limit = &defaultLimit
	}

	if request.Offset == nil {
		defaultOffset := 0
		request.Offset = &defaultOffset
	}

	data, err := u.merchantRepository.GetMerchants(context.TODO(), request)
	if err != nil {
		return res, err
	}

	for _, v := range data {
		merchant := dto.ResponseMerchant{
			MerchantId:       strconv.Itoa(v.ID),
			Name:             v.Name,
			MerchantCategory: v.MerchantCategory,
			ImageUrl:         v.ImageUrl,
			Location: dto.Location{
				Lat:  v.LocationLat,
				Long: v.LocationLong,
			},
			CreatedAt: v.CreatedAt,
		}
		responseMerchants = append(responseMerchants, merchant)
	}

	res.Data = responseMerchants
	res.Meta.Limit = *request.Limit
	res.Meta.Offset = *request.Offset
	res.Meta.Total = len(res.Data)

	return res, err
}

func (u *MerchantUsecase) CreateMerchantItem(request dto.RequestCreateMerchantItem) (res dto.ResponseCreateMerchantItem, err error) {
	data := database.Item{
		Name:            request.Name,
		ProductCategory: request.ProductCategory,
		Price:           int(request.Price),
		ImageUrl:        request.ImageUrl,
		MerchantID:      request.MerchantID,
		CreatedAt:       time.Now(),
	}

	id, err := u.merchantRepository.CreateMerchantItem(context.TODO(), data)
	if err == nil || id != 0 {
		res.ItemID = strconv.Itoa(id)
	}

	return res, err
}

func (u *MerchantUsecase) GetMerchantItems(request dto.RequestGetMerchantItems) (res dto.ResponseGetMerchantItems, err error) {
	responseData := []dto.ResponseGetItems{}
	if request.Limit == nil {
		defaultLimit := 5
		request.Limit = &defaultLimit
	}

	if request.Offset == nil {
		defaultOffset := 0
		request.Offset = &defaultOffset
	}

	data, err := u.merchantRepository.GetMerchantItems(context.TODO(), request)
	if err != nil {
		return res, err
	}

	for _, v := range data {
		merchant := dto.ResponseGetItems{
			ItemID:          strconv.Itoa(v.ID),
			Name:            v.Name,
			ProductCategory: v.ProductCategory,
			Price:           v.Price,
			ImageUrl:        v.ImageUrl,
			CreatedAt:       v.CreatedAt,
		}

		responseData = append(responseData, merchant)
	}

	res.Data = responseData
	res.Meta.Limit = *request.Limit
	res.Meta.Offset = *request.Offset
	res.Meta.Total = len(res.Data)

	return res, err
}

func (u *MerchantUsecase) GetMerchantCountByIds(ids string) (res int) {
	res = u.merchantRepository.GetMerchantCountByIds(context.TODO(), ids)
	return res
}

func (u *MerchantUsecase) GetItemCountByIds(ids string) (res int) {
	res = u.merchantRepository.GetItemCountByIds(context.TODO(), ids)
	return res
}
