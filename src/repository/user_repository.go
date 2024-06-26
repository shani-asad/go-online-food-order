package repository

import (
	"context"
	"database/sql"
	"fmt"
	"online-food/model/database"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepositoryInterface {
	return &UserRepository{db}
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (response database.User, err error) {
	err = r.db.QueryRowContext(ctx, "SELECT id, username, email, password, role FROM users WHERE username = $1", username).Scan(&response.ID, &response.Username, &response.Email, &response.Password, &response.Role)
	if err != nil {
		return
	}
	return
}

func (r *UserRepository) CreateUser(ctx context.Context, data database.User) (id int, err error) {
	query := `
	INSERT INTO users (username, email, password, role, created_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id`
	
	err = r.db.QueryRowContext(
		ctx,
		query,
		data.Username,
		data.Email,
		data.Password,
		data.Role,
		data.CreatedAt,
	).Scan(&id)

	fmt.Println(err)

	return id, err
}

func (r *UserRepository) GetExistingUserInTheRoleByEmail(ctx context.Context, email, role string) (response database.User, err error) {
	err = r.db.QueryRowContext(ctx, "SELECT username, email, password, role FROM users WHERE email = $1 and role = $2", email, role).Scan(&response.Username, &response.Email, &response.Password, &response.Role)
	if err != nil {
		return
	}
	return
}
