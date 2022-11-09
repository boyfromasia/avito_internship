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

func (s *UserService) AddBalanceUser(user models.UserAddBalanceRequest) (models.UserAddBalanceResponse, error) {
	return s.repo.AddBalanceUser(user)
}

func (s *UserService) ReserveMoneyUser(user models.UserReserveMoneyRequest) (models.UserReserveMoneyResponse, error) {
	return s.repo.ReserveMoneyUser(user)
}
