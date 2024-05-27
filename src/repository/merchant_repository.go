package repository

import (
	"context"
	"database/sql"
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
	return response, err
}
