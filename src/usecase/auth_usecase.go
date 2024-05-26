package usecase

import (
	"context"
	"errors"
	"fmt"
	"health-record/helpers"
	"health-record/model/database"
	"health-record/model/dto"
	"health-record/src/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	iUserRepository repository.UserRepositoryInterface
	helper          helpers.HelperInterface
}

func NewAuthUsecase(
	iUserRepository repository.UserRepositoryInterface,
	helper helpers.HelperInterface) AuthUsecaseInterface {
	return &AuthUsecase{iUserRepository, helper}
}

func (u *AuthUsecase) Register(request dto.RequestCreateUser) (token string, userId string, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	if err != nil {
		return "", "", err
	}

	data := database.User{
		Nip:     request.Nip,
		Password:  string(hash),
		Name:      request.Name,
		CreatedAt: time.Now(),
	}

	u.iUserRepository.CreateUser(context.TODO(), data)

	fmt.Println(err)

	userData, err := u.iUserRepository.GetUserByNIP(context.TODO(), request.Nip)

	fmt.Println(userData)

	token, _ = u.helper.GenerateToken(userData.Id, userData.Role)

	return token, userData.Id, err
}

func (u *AuthUsecase) Login(request dto.RequestAuth) (token string, user database.User, err error) {
	// check creds on database
	userData, err := u.iUserRepository.GetUserByNIP(context.TODO(), request.Nip)
	if err != nil {
		return "", database.User{}, errors.New("user not found")
	}

	fmt.Println(userData)

	// check the password
	isValid := u.verifyPassword(request.Password, userData.Password)
	if !isValid {
		return "", database.User{}, errors.New("wrong password")
	}

	token, _ = u.helper.GenerateToken(userData.Id, userData.Role)

	return token, userData, nil
}

func (u *AuthUsecase) LoginNurse(request dto.RequestAuth) (token string, user database.User, err error) {
	userData, err := u.iUserRepository.GetUserByNIP(context.TODO(), request.Nip)
	if err != nil {
		return "", database.User{}, errors.New("user not found")
	}

	fmt.Println(userData)

	// check the password
	isValid := u.verifyPassword(request.Password, userData.Password)
	if !isValid {
		return "", database.User{}, errors.New("wrong password")
	}

	token, _ = u.helper.GenerateToken(userData.Id, userData.Role)

	return token, userData, nil
}

func (u *AuthUsecase) verifyPassword(password, passwordHash string) bool {
	byteHash := []byte(passwordHash)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(password))

	return err == nil
}

func (u *AuthUsecase) GetUserByNIP(nip int64) (bool, error) {
	_, err := u.iUserRepository.GetUserByNIP(context.TODO(), nip)
	if err != nil {
    return false, err
  }
	return true, nil
}