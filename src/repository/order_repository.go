package repository

import (
	"context"
	"database/sql"
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

	if(res != "") {
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
			log.Println("Error create order", err)
			return "", err
		}
	}

	if err != nil {
		log.Println("Error create order", err)
	}

	return res, err
}