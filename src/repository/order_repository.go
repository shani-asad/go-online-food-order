package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"online-food/model/dto"
	"time"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) OrderRepositoryInterface {
	return &OrderRepository{db}
}

func (r *OrderRepository) CreateEstimation(ctx context.Context, data dto.RequestEstimate, userId string) (int, error) {
	query := `
	INSERT INTO orders (
		user_id,
		already_placed,
		created_at)
	VALUES ($1, $2, $3)
	RETURNING id
	`
	var orderId string
	err := r.db.QueryRowContext(
		ctx,
		query,
		userId,
		"false",
		time.Now(),
	).Scan(&orderId)
	if err != nil {
		log.Println("Error insert into orders", err)
		return 0, err
	}

	for _, o := range data.Orders {
		for _, i := range o.Items {
			query = `
			INSERT INTO item_orders (
				item_id,
				order_id,
				quantity,
				created_at)
			VALUES ($1, $2, $3, $4)
			`

			_, err := r.db.ExecContext(
				ctx,
				query,
				i.ItemId,
				orderId,
				i.Quantity,
				time.Now(),
			)
			if err != nil {
				log.Println("Error insert into item_orders", err)
				return 0, err
			}
		}
	}

	query = `
	INSERT INTO estimations (
		order_id,
		created_at)
	VALUES ($1, $2)
	RETURNING id
	`
	var estimationId int
	err = r.db.QueryRowContext(
		ctx,
		query,
		orderId,
		time.Now(),
	).Scan(&estimationId)

	return estimationId, err
}

func (r *OrderRepository) CreateOrder(ctx context.Context, estimationId string) (res string, err error) {
	query := `
	SELECT order_id FROM estimations
	where id = $1
	`
	err = r.db.QueryRowContext(
		ctx,
		query,
		estimationId,
	).Scan(&res)

	if err != nil {
		log.Println("Error create order 1st step", err)
		return "", err
	}

	if res != "" {
		query := `
		UPDATE orders SET already_placed = true
		where id = $1
		`
		_, err := r.db.ExecContext(
			ctx,
			query,
			res,
		)

		if err != nil {
			log.Println("Error create order 2nd step", err)
			return "", err
		}
	}
	return res, err
}

func (r *OrderRepository) GetOrders(ctx context.Context, filter dto.RequestGetOrders, userId string) (res []dto.ResponseGetOrders, err error) {
	query := fmt.Sprintf(`
	SELECT 
		o.id,
		m.id,
		m.name,
		m.merchant_category,
		m.image_url,
		m.location_lat,
		m.location_long,
		m.created_at,
		i.id,
		i.name,
		i.product_category,
		i.price,
		io.quantity,
		i.image_url,
		i.created_at
		FROM
	orders o JOIN item_orders io
	ON o.id = io.order_id
	JOIN items i
	ON io.item_id = i.id
	JOIN merchants m
	ON m.id = i.merchant_id
	WHERE o.user_id = %v
	AND o.already_placed = true
	`, userId)

	if filter.Name != nil {
		query += fmt.Sprintf(" AND (m.name ILIKE '%%%v%%' OR i.name ILIKE '%%%v%%')", *filter.Name, *filter.Name)
	}

	if filter.MerchantId != nil {
		query += fmt.Sprintf(" AND id = %v", *filter.MerchantId)
	}

	if filter.MerchantCategory != nil {
		query += fmt.Sprintf(" AND merchant_category = %v", *filter.MerchantCategory)
	}

	query += fmt.Sprintf(" OFFSET %v", *filter.Offset)
	query += fmt.Sprintf(" LIMIT %v", *filter.Limit)

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		log.Println("Error get order list :>>", err)
		return []dto.ResponseGetOrders{}, err
	}
	defer rows.Close()

	dbRes := dto.ResponseGetOrdersDB{}
	dbResSlice := []dto.ResponseGetOrdersDB{}
	for rows.Next() {
		if err := rows.Scan(
			&dbRes.OrderId,
			&dbRes.MerchantId,
			&dbRes.MerchantName,
			&dbRes.MerchantCategory,
			&dbRes.MercahntImageUrl,
			&dbRes.MerchantLocationLat,
			&dbRes.MerchantLocationLon,
			&dbRes.MerchantCreatedAt,
			&dbRes.ItemId,
			&dbRes.ItemName,
			&dbRes.ItemProductCategory,
			&dbRes.ItemPrice,
			&dbRes.ItemQuantity,
			&dbRes.ItemImageUrl,
			&dbRes.ItemCreatedAt,
		); err != nil {
			return res, err
		}
		dbResSlice = append(dbResSlice, dbRes)
	}

	log.Printf("====================\ndbResSlice %+v\n",dbResSlice)

	res = mapDBToResponseGetOrders(dbResSlice)

	return res, err
}


func mapDBToResponseGetOrders(dbRecords []dto.ResponseGetOrdersDB) []dto.ResponseGetOrders {
	orderMap := make(map[string]*dto.ResponseGetOrders)
	merchantMap := make(map[string]map[int]*dto.OrderDetail)

	for _, record := range dbRecords {
		if _, exists := orderMap[record.OrderId]; !exists {
			orderMap[record.OrderId] = &dto.ResponseGetOrders{
				OrderId: record.OrderId,
				Orders:  []dto.OrderDetail{},
			}
			merchantMap[record.OrderId] = make(map[int]*dto.OrderDetail)
		}

		if _, exists := merchantMap[record.OrderId][record.MerchantId]; !exists {
			merchantDetail := dto.MerchantDetail{
				MerchantId:       record.MerchantId,
				Name:             record.MerchantName,
				MerchantCategory: record.MerchantCategory,
				ImageUrl:         record.MercahntImageUrl,
				Location: dto.Location{
					Lat:  record.MerchantLocationLat,
					Long: record.MerchantLocationLon,
				},
				CreatedAt: record.MerchantCreatedAt,
			}
			orderDetail := dto.OrderDetail{
				Merchant: merchantDetail,
				Items:    []dto.OrderItemDetail{},
			}
			orderMap[record.OrderId].Orders = append(orderMap[record.OrderId].Orders, orderDetail)
			merchantMap[record.OrderId][record.MerchantId] = &orderMap[record.OrderId].Orders[len(orderMap[record.OrderId].Orders)-1]
		}

		item := dto.OrderItemDetail{
			Item: dto.Item{
				ItemId:          fmt.Sprint(record.ItemId), // Converting int ItemId to string
				Name:            record.ItemName,
				ProductCategory: record.ItemProductCategory,
				Price:           record.ItemPrice,
				ImageUrl:        record.ItemImageUrl,
				CreatedAt:       record.ItemCreatedAt,
			},
			Quantity: record.ItemQuantity,
		}

		// Add item to the corresponding order detail
		merchantMap[record.OrderId][record.MerchantId].Items = append(merchantMap[record.OrderId][record.MerchantId].Items, item)
	}

	// Convert map to slice
	response := make([]dto.ResponseGetOrders, 0, len(orderMap))
	for _, order := range orderMap {
		response = append(response, *order)
	}

	return response
}