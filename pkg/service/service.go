package service

import (
	"avito_internship/pkg/models"
	"avito_internship/pkg/repository"
)

type User interface {
	GetBalanceUser(user models.UserGetBalance) (models.UserGetBalanceResponse, error)
	AddBalanceUser(user models.UserAddBalance) (models.UserAddBalanceResponse, error)
}

type Purchase interface {
}

type Order interface {
}

type Service struct {
	User
	Purchase
	Order
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repos.User),
	}
}
