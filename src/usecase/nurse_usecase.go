package usecase

import (
	"context"
	"errors"
	"health-record/model/database"
	"health-record/model/dto"
	"health-record/src/repository"
)

type NurseUsecase struct {
	iNurseRepository repository.NurseRepositoryInterface
}

func NewNurseUsecase(
	iNurseRepository repository.NurseRepositoryInterface) NurseUsecaseInterface {
	return &NurseUsecase{iNurseRepository}
}

func (uc *NurseUsecase) RegisterNurse(request dto.RequestCreateNurse) (string, error) {
	// Check if the nurse NIP already exists in the database
	existingNurse, err := uc.iNurseRepository.GetNurseByNIP(context.TODO(), request.Nip)
	if err == nil && existingNurse.Id != "" {
			return "", errors.New("a nurse with this NIP already exists")
	}

	// If the nurse does not exist, proceed to create a new nurse
	userId, err := uc.iNurseRepository.CreateNurse(context.TODO(), request)
	if err != nil {
			return "", err
	}
	return userId, nil
}

func (uc *NurseUsecase) GetUsers(request dto.RequestGetUser) ([]dto.UserDTO, error) {
	params := dto.RequestGetUser{
		Limit:    validateLimit(request.Limit),
		Offset:   validateOffset(request.Offset),
		UserId: request.UserId,
		Name:     request.Name,
    NIP:      request.NIP,
    Role: request.Role,
		CreatedAt: request.CreatedAt,
	}

	response, err := uc.iNurseRepository.GetUsers(context.TODO(), params)
	
	
	return response, err
}

// UpdateNurse handles the updating of an existing nurse's information.
func (uc *NurseUsecase) UpdateNurse(userId string, nurse dto.RequestUpdateNurse) int {
	// Ensure the nurse exists before attempting to update
	_, err := uc.iNurseRepository.GetNurseByID(context.TODO(), userId)
	if err != nil {
			return 404
	}

	// Proceed with updating the nurse
	return uc.iNurseRepository.UpdateNurse(context.TODO(), userId, nurse)
}

// DeleteNurse handles the deletion of a nurse.
func (uc *NurseUsecase) DeleteNurse(userId string) int {
	return uc.iNurseRepository.DeleteNurse(context.TODO(), userId)
}

func (uc *NurseUsecase) AddAccess(userId string, password dto.RequestAddAccess) int {
  return uc.iNurseRepository.AddAccess(context.TODO(), userId, password)
}

func (u *NurseUsecase) GetNurseByNIP(nip int64) (bool, error) {
	_, err := u.iNurseRepository.GetNurseByNIP(context.TODO(), nip)
	if err != nil {
    return false, err
  }
	return true, nil
}

func (u *NurseUsecase) GetNurseByID(id string) (database.User, error) {
	user, err := u.iNurseRepository.GetNurseByID(context.TODO(), id)
  if err != nil {
    return database.User{}, err
  }
  return user, nil
}

func validateLimit(limit int) int {
	if limit >= 0 {
			return limit
	}
return 5
}

func validateOffset(offset int) int {
	if offset >= 0 {
			return offset
	}
return 0
}