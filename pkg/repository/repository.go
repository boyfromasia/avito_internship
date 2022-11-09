package repository

import (
	"avito_internship/pkg/models"
	"github.com/jmoiron/sqlx"
)

type User interface {
	GetBalanceUser(user models.UserGetBalance) (models.UserGetBalanceResponse, error)
	AddBalanceUser(user models.UserAddBalance) (models.UserAddBalanceResponse, error)
}

type Purchase interface {
}

type Order interface {
}

type Repository struct {
	User
	Purchase
	Order
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserPostgres(db),
	}
}
