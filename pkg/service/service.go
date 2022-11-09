package service

import (
	"avito_internship/pkg/models"
	"avito_internship/pkg/repository"
)

type User interface {
	GetBalanceUser(user models.UserGetBalanceRequest) (models.UserGetBalanceResponse, error)
	AddBalanceUser(user models.UserAddBalanceRequest) (models.UserAddBalanceResponse, error)
}

type Purchase interface {
}

type Order interface {
	AddRecord(record models.AddRecordRequest) (models.AddRecordResponse, error)
}

type Service struct {
	User
	Purchase
	Order
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User:  NewUserService(repos.User),
		Order: NewOrderService(repos.Order),
	}
}
