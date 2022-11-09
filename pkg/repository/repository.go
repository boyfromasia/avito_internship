package repository

import (
	"avito_internship/pkg/models"
	"github.com/jmoiron/sqlx"
)

type User interface {
	GetBalanceUser(user models.UserGetBalanceRequest) (models.UserGetBalanceResponse, error)
	AddBalanceUser(user models.UserAddBalanceRequest) (models.StatusResponse, error)
	ReserveMoneyUser(user models.UserReserveMoneyRequest) (models.StatusResponse, error)
	ApproveReserveUser(user models.UserDecisionRequest) (models.StatusResponse, error)
	RejectReserveUser(user models.UserDecisionRequest) (models.StatusResponse, error)
}

type Purchase interface {
	GetPurchase(purchase models.GetPurchaseRequest) (models.GetPurchaseResponse, error)
}

type Order interface {
	AddRecordOrder(record models.AddRecordRequest) (models.AddRecordResponse, error)
	DecisionReserveUser(decision models.DecisionReserveRequest) (models.DecisionReserveResponse, error)
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
		Purchase:    NewPurchasePostgres(db),
	}
}
