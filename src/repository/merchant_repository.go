package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"online-food/model/database"
	"online-food/model/dto"
)

type MerchantRepository struct {
	db *sql.DB
}

func NewMerchantRepository(db *sql.DB) MerchantRepositoryInterface {
	return &MerchantRepository{db}
}

func (r *MerchantRepository) CreateMerchant(ctx context.Context, data database.Merchant) (id int, err error) {
	query := `
	INSERT INTO merchants (
		name,
		merchant_category,
		image_url,
		location_lat,
		location_long,
		created_at,
		updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id`

	err = r.db.QueryRowContext(
		ctx,
		query,
		data.Name,
		data.MerchantCategory,
		data.ImageUrl,
		data.LocationLat,
		data.LocationLong,
		data.CreatedAt,
		data.UpdatedAt,
	).Scan(&id)

	return id, err
}

func (r *MerchantRepository) GetMerchants(ctx context.Context, filter dto.RequestGetMerchant) (response []database.Merchant, err error) {
	query := `SELECT id, name, merchant_category, image_url, location_lat, location_long, created_at, updated_at FROM merchants WHERE 1=1`
	args := []interface{}{}

	if filter.MerchantID != nil {
		query += fmt.Sprintf(" AND id = %v", *filter.MerchantID)
	}
	if filter.Name != nil {
		query += fmt.Sprintf(" AND name LIKE '%s'", *filter.Name)
	}
	if filter.MerchantCategory != nil {
		query += fmt.Sprintf(" AND merchant_category = '%s'", *filter.MerchantCategory)
	}
	if filter.CreatedAt != nil {
		if *filter.CreatedAt == "asc" {
			query += " ORDER BY created_at ASC"
		} else if *filter.CreatedAt == "desc" {
			query += " ORDER BY created_at DESC"
		}
	}
	if filter.Limit != nil {
		query += fmt.Sprintf(" LIMIT %d", *filter.Limit)
	}
	if filter.Offset != nil {
		query += fmt.Sprintf(" OFFSET %d", *filter.Offset)
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var merchants []database.Merchant
	for rows.Next() {
		var merchant database.Merchant
		if err := rows.Scan(
			&merchant.ID,
			&merchant.Name,
			&merchant.MerchantCategory,
			&merchant.ImageUrl,
			&merchant.LocationLat,
			&merchant.LocationLong,
			&merchant.CreatedAt,
			&merchant.UpdatedAt,
		); err != nil {
			return nil, err
		}
		merchants = append(merchants, merchant)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return merchants, nil
}

func (r *MerchantRepository) CreateMerchantItem(ctx context.Context, data database.Item) (id int, err error) {
	query := `
	INSERT INTO items (
		name,
		product_category,
		price,
		image_url,
		merchant_id,
		created_at,
		updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id`

	err = r.db.QueryRowContext(
		ctx,
		query,
		data.Name,
		data.ProductCategory,
		data.Price,
		data.ImageUrl,
		data.MerchantID,
		data.CreatedAt,
		data.UpdatedAt,
	).Scan(&id)

	return id, err
}

func (r *MerchantRepository) GetMerchantItems(ctx context.Context, filter dto.RequestGetMerchantItems) ([]database.Item, error) {
	query := `SELECT id, name, product_category, price, image_url, created_at, updated_at FROM items WHERE 1=1`
	args := []interface{}{}

	if filter.ItemID != nil {
		query += fmt.Sprintf(" AND id = %v", *filter.ItemID)
	}
	if filter.Name != nil {
		query += fmt.Sprintf(" AND name LIKE '%s'", *filter.Name)
	}
	if filter.ProductCategory != nil {
		query += fmt.Sprintf(" AND merchant_category = '%s'", *filter.ProductCategory)
	}
	if filter.CreatedAt != nil {
		if *filter.CreatedAt == "asc" {
			query += " ORDER BY created_at ASC"
		} else if *filter.CreatedAt == "desc" {
			query += " ORDER BY created_at DESC"
		}
	}
	if filter.Limit != nil {
		query += fmt.Sprintf(" LIMIT %d", *filter.Limit)
	}
	if filter.Offset != nil {
		query += fmt.Sprintf(" OFFSET %d", *filter.Offset)
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []database.Item
	for rows.Next() {
		var item database.Item
		if err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.ProductCategory,
			&item.Price,
			&item.ImageUrl,
			&item.CreatedAt,
			&item.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *MerchantRepository) GetNearbyMerchants(ctx context.Context, long float64, lat float64, filter dto.RequestNearbyMerchants) (response dto.ResponseNearbyMerchants, err error) {

	query := fmt.Sprintf(`
	WITH limited_merchants AS (
		SELECT distinct m.*, earth_distance(ll_to_earth(%v, %v), earth_location)
		FROM merchants m
		JOIN items i ON m.id = i.merchant_id
		WHERE 1 = 1
	`, lat, long)
	if filter.Name != nil {
		query += fmt.Sprintf(" AND (m.name ILIKE '%%%v%%' OR i.name ILIKE '%%%v%%')", *filter.Name, *filter.Name)
	}

	if filter.MerchantId != nil {
		query += fmt.Sprintf(" AND id = %v", *filter.MerchantId)
	}

	if filter.MerchantCategory != nil {
		query += fmt.Sprintf(" AND merchant_category = %v", *filter.MerchantCategory)
	}

	query += fmt.Sprintf("	ORDER BY earth_distance(ll_to_earth(%v, %v), earth_location)", lat, long)
	query += fmt.Sprintf(" LIMIT %v )", *filter.Limit)

	query += `
	SELECT
		m.id, m.name, m.merchant_category, m.image_url, m.location_lat, m.location_long, m.created_at,
		i.id, i.name, i.product_category, i.price, i.image_url, i.created_at
	FROM limited_merchants m
	JOIN items i ON m.id = i.merchant_id
	ORDER BY m.id
	`

	query += fmt.Sprintf(" OFFSET %v", *filter.Offset)

	fmt.Println("filter", filter)
	fmt.Println("query>>", query)

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return dto.ResponseNearbyMerchants{}, err
	}
	defer rows.Close()

	var nearbyMerchantsDbResponse []dto.NearbyMerchantsDbResponse
	for rows.Next() {
		var m dto.NearbyMerchantsDbResponse
		if err := rows.Scan(
			&m.Merchant.MerchantId,
			&m.Merchant.Name,
			&m.Merchant.MerchantCategory,
			&m.Merchant.ImageUrl,
			&m.Merchant.Location.Lat,
			&m.Merchant.Location.Long,
			&m.Merchant.CreatedAt,
			&m.Items.ItemId,
			&m.Items.Name,
			&m.Items.ProductCategory,
			&m.Items.Price,
			&m.Items.ImageUrl,
			&m.Items.CreatedAt,
		); err != nil {
			return dto.ResponseNearbyMerchants{}, err
		}
		nearbyMerchantsDbResponse = append(nearbyMerchantsDbResponse, m)
	}

	if len(nearbyMerchantsDbResponse) == 0 {
		return dto.ResponseNearbyMerchants{
			Data: []dto.NearbyMerchants{},
			Meta: dto.ResponseMeta{
				Limit:  *filter.Limit,
				Offset: *filter.Offset,
				Total:  0,
			},
		}, err
	}

	var nearbyMerchant dto.NearbyMerchants
	var nearbyMerchants []dto.NearbyMerchants
	merchantId := nearbyMerchantsDbResponse[0].Merchant.MerchantId
	var items []dto.Item

	log.Println("len(nearbyMerchantsDbResponse)", len(nearbyMerchantsDbResponse))
	for idx, v := range nearbyMerchantsDbResponse {
		if v.Merchant.MerchantId == merchantId {
			i := v.Items
			items = append(items, dto.Item{
				ItemId:          i.ItemId,
				Name:            i.Name,
				ProductCategory: i.ProductCategory,
				Price:           i.Price,
				ImageUrl:        i.ImageUrl,
				CreatedAt:       i.CreatedAt,
			})
		}
		if idx == len(nearbyMerchantsDbResponse)-1 || nearbyMerchantsDbResponse[idx+1].Merchant.MerchantId != merchantId {
			nearbyMerchant.Items = items
			nearbyMerchant.Merchant.MerchantId = merchantId
			nearbyMerchant.Merchant.Name = v.Merchant.Name
			nearbyMerchant.Merchant.MerchantCategory = v.Merchant.MerchantCategory
			nearbyMerchant.Merchant.ImageUrl = v.Merchant.ImageUrl
			nearbyMerchant.Merchant.Location = v.Merchant.Location
			nearbyMerchant.Merchant.CreatedAt = v.Merchant.CreatedAt
			items = []dto.Item{}

			nearbyMerchants = append(nearbyMerchants, nearbyMerchant)

			if idx != len(nearbyMerchantsDbResponse)-1 {
				merchantId = nearbyMerchantsDbResponse[idx+1].Merchant.MerchantId
			}
		}

	}
	response.Data = nearbyMerchants
	response.Meta = dto.ResponseMeta{
		Limit:  *filter.Limit,
		Offset: *filter.Offset,
		Total:  len(nearbyMerchants),
	}

	if err := rows.Err(); err != nil {
		return dto.ResponseNearbyMerchants{}, err
	}

	return response, nil
}
