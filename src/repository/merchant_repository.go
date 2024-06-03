package repository

import (
	"context"
	"database/sql"
	"fmt"
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
