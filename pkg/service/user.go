package service

import (
	"avito_internship/pkg/models"
	"avito_internship/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetBalanceUser(user models.UserGetBalanceRequest) (models.UserGetBalanceResponse, error) {
	return s.repo.GetBalanceUser(user)
}

func (s *UserService) AddBalanceUser(user models.UserAddBalanceRequest) (models.StatusResponse, error) {
	var response models.StatusResponse

	if err := user.Validate(); err != nil {
		response.Status = "Error"
		return response, err
	}

	return s.repo.AddBalanceUser(user)
}

func (s *UserService) ReserveMoneyUser(user models.UserReserveMoneyRequest) (models.StatusResponse, error) {
	return s.repo.ReserveMoneyUser(user)
}

func (s *UserService) ApproveReserveUser(user models.UserDecisionRequest) (models.StatusResponse, error) {
	return s.repo.ApproveReserveUser(user)
}

func (s *UserService) RejectReserveUser(user models.UserDecisionRequest) (models.StatusResponse, error) {
	return s.repo.RejectReserveUser(user)
}
