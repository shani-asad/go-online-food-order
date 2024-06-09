package usecase

import (
	"context"
	"errors"
	"log"
	"online-food/helpers"
	"online-food/model/dto"
	"online-food/src/repository"
	"strconv"
	"strings"
)

type PurchaseUsecase struct {
	merchantRepository repository.MerchantRepositoryInterface
	orderRepository    repository.OrderRepositoryInterface
}

func NewPurchaseUsecase(
	merchantRepository repository.MerchantRepositoryInterface,
	orderRepository repository.OrderRepositoryInterface) PurchaseUsecaseInterface {
	return &PurchaseUsecase{merchantRepository, orderRepository}
}

func (u *PurchaseUsecase) GetNearbyMerchants(long float64, lat float64, request dto.RequestNearbyMerchants) (res dto.ResponseNearbyMerchants, err error) {
	res, err = u.merchantRepository.GetNearbyMerchants(context.TODO(), long, lat, request)
	return res, err
}

func (u *PurchaseUsecase) CreateEstimation(request dto.RequestEstimate, userId string) (res dto.ResponseOrder, err error) {
	id, err := u.orderRepository.CreateEstimation(context.TODO(), request, userId)

	if err != nil {
		log.Println("Error create estimation", err.Error())
		return dto.ResponseOrder{}, err
	}

	totalPrice, err := getTotalPrice(request.Orders, u)

	if err != nil {
		return dto.ResponseOrder{}, err
	}

	var startLat, startLon float64
	var destinations [][2]float64
	userLat := request.UserLocation.Lat
	userLong := request.UserLocation.Long

	destinations = append(destinations, [2]float64{
		userLat,
		userLong,
	})

	locationMap, err := getMerchantLocations(request.Orders, u)
	if err != nil {
		return dto.ResponseOrder{}, err
	}

	log.Printf("======%+v\n", locationMap)
	for _, v := range request.Orders {
		merchantLat := locationMap[v.MerchantId].Lat
		merchantLon := locationMap[v.MerchantId].Long

		distance := helpers.NewDistanceHelper().GetHaversineDistance(userLat, userLong, merchantLat, merchantLon)
		if distance > 3 {
			log.Println("Distance: ", distance)
			return dto.ResponseOrder{}, errors.New("merchant's distance is too far from user")
		}

		if v.IsStartingPoint {
			startLat = merchantLat
			startLon = merchantLon
		} else {
			destinations = append(destinations, [2]float64{
				merchantLat,
				merchantLon,
			})
		}
	}

	estimatedTime := helpers.CalculateShortestTime(
		// userLat,
		// userLong,
		startLat,
		startLon,
		destinations,
	)

	res = dto.ResponseOrder{
		TotalPrice:                     totalPrice,
		EstimatedDeliveryTimeInMinutes: estimatedTime,
		CalculatedEstimateId:           strconv.Itoa(id),
	}
	return res, err
}

func getMerchantLocations(orders []dto.Order, u *PurchaseUsecase) (map[string]dto.Location, error) {
	merchantIds := []string{}
	for _, o := range orders {
		merchantIds = append(merchantIds, o.MerchantId)
	}

	merchantIdsString := strings.Join(merchantIds, ", ")

	locations, err := u.merchantRepository.GetMerchantLocations(context.TODO(), merchantIdsString)
	return locations, err
}

func getTotalPrice(orders []dto.Order, u *PurchaseUsecase) (int, error) {
	itemIds := []string{}
	for _, o := range orders {
		for _, i := range o.Items {
			itemIds = append(itemIds, i.ItemId)
		}
	}
	itemIdsString := strings.Join(itemIds, ", ")

	price, err := u.merchantRepository.GetTotalPriceOfItems(context.TODO(), itemIdsString)

	if err != nil {
		return 0, err
	}
	return price, nil
}

func (u *PurchaseUsecase) CreateOrder(estimateId string) (res string, err error){
	res, err = u.orderRepository.CreateOrder(context.TODO(), estimateId)
	return res, err
}