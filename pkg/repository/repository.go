package repository

import (
	"avito_internship/pkg/models"
	"github.com/jmoiron/sqlx"
)

type User interface {
	GetBalanceUser(user models.UserGetBalanceRequest) (models.UserGetBalanceResponse, error)
	AddBalanceUser(user models.UserAddBalanceRequest) (models.UserAddBalanceResponse, error)
}

type Purchase interface {
}

type Order interface {
	AddRecord(record models.AddRecordRequest, isAdd bool) (models.AddRecordResponse, error)
}

type Repository struct {
	User
	Purchase
	Order
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:  NewUserPostgres(db),
		Order: NewOrderPostgres(db),
	}
}
