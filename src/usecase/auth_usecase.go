package usecase

import (
	"context"
	"fmt"
	"online-food/helpers"
	"online-food/model/database"
	"online-food/model/dto"
	"online-food/src/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	iUserRepository repository.UserRepositoryInterface
	helper          helpers.AuthHelperInterface
}

func NewAuthUsecase(
	iUserRepository repository.UserRepositoryInterface,
	helper helpers.AuthHelperInterface) AuthUsecaseInterface {
	return &AuthUsecase{iUserRepository, helper}
}

func (u *AuthUsecase) Register(request dto.RequestCreateUser, role string) (token string, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	data := database.User{
		Email:     request.Email,
		Password:  string(hash),
		Username:  request.Username,
		Role:      role,
		CreatedAt: time.Now(),
	}

	u.iUserRepository.CreateUser(context.TODO(), data)

	fmt.Println(err)

	userData, err := u.iUserRepository.GetUserByUsername(context.TODO(), request.Username)

	fmt.Println(userData)

	token, _ = u.helper.GenerateToken(userData.ID, userData.Role)

	return token, err
}

func (u *AuthUsecase) Login(request dto.RequestAuth, role string) (token string, status int) {
	// check creds on database
	userData, err := u.iUserRepository.GetUserByUsername(context.TODO(), request.Username)
	if err != nil {
		return "", 404
	}

	if userData.Role != role {
		return "", 400
	}

	fmt.Println(userData)

	// check the password
	isValid := u.verifyPassword(request.Password, userData.Password)
	if !isValid {
		return "", 400
	}

	token, _ = u.helper.GenerateToken(userData.ID, userData.Role)

	return token, 200
}

func (u *AuthUsecase) verifyPassword(password, passwordHash string) bool {
	byteHash := []byte(passwordHash)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(password))

	return err == nil
}

func (u *AuthUsecase) GetUserByUsername(username string) (bool, error) {
	_, err := u.iUserRepository.GetUserByUsername(context.TODO(), username)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u *AuthUsecase) GetExistingUserInTheRoleByEmail(email, role string) (bool, error) {
	_, err := u.iUserRepository.GetExistingUserInTheRoleByEmail(context.TODO(), email, role)
	if err != nil {
		return false, err
	}
	return true, nil
}
