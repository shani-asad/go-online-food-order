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

func (r *MerchantRepository) CreateMerchant(ctx context.Context, data database.Merchant) (err error) {
	return err
}

func (r *MerchantRepository) GetMerchants(ctx context.Context, filter dto.RequestGetMerchant) (response []database.Merchant, err error) {
	return response, err
}
