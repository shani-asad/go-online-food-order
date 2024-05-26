package repository

import (
	"context"
	"database/sql"
	"fmt"
	"health-record/model/database"
	"math/rand"
	"time"

	"github.com/docker/distribution/uuid"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepositoryInterface {
	return &UserRepository{db}
}

func (r *UserRepository) GetUserByNIP(ctx context.Context, nip int64) (response database.User, err error) {
	err = r.db.QueryRowContext(ctx, "SELECT user_id, name, nip, password, role FROM users WHERE nip = $1", nip).Scan(&response.Id, &response.Name, &response.Nip, &response.Password, &response.Role)
	if err != nil {
		return
	}
	return
}

func randomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min+1)
}

func (r *UserRepository) CreateUser(ctx context.Context, data database.User) (err error) {
	query := `
	INSERT INTO users (user_id, nip, name, password, role, created_at)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING user_id`

	_, err = r.db.ExecContext(
		ctx,
		query,
		uuid.Generate().String(),
		data.Nip,
		data.Name,
		data.Password,
		"it",
		data.CreatedAt,
	)

	fmt.Println(err)

	return err
}
