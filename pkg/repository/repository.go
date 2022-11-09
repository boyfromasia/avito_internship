package repository

import (
	"avito_internship/pkg/models"
	"github.com/jmoiron/sqlx"
)

type User interface {
	GetBalanceUser(user models.UserGetBalanceRequest) (models.UserGetBalanceResponse, error)
	AddBalanceUser(user models.UserAddBalanceRequest) (models.UserAddBalanceResponse, error)
	ReserveMoneyUser(user models.UserReserveMoneyRequest) (models.UserReserveMoneyResponse, error)
}

type Purchase interface {
}

type Order interface {
	AddRecordOrder(record models.AddRecordRequest) (models.AddRecordResponse, error)
}

type HistoryUser interface {
	AddRecordHistory(record models.AddHistoryRequest) (models.AddHistoryResponse, error)
}

type Repository struct {
	User
	Purchase
	Order
	HistoryUser
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:        NewUserPostgres(db),
		Order:       NewOrderPostgres(db),
		HistoryUser: NewHistoryUserPostgres(db),
	}
}
