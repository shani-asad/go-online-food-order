package repository

import (
	"context"
	"online-food/model/database"
)

type UserRepositoryInterface interface {
	GetUserByUsername(ctx context.Context, username string) (response database.User, err error)
	CreateUser(ctx context.Context, data database.User) (err error)
	GetExistingUserInTheRoleByEmail(ctx context.Context, email, role string) (response database.User, err error)
}