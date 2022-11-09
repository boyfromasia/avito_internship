package models

import "errors"

type UserGetBalanceRequest struct {
	UserId int `json:"user_id" binding:"required"`
}

type UserGetBalanceResponse struct {
	Balance float64 `json:"balance"`
}

type UserAddBalanceRequest struct {
	UserId  int     `json:"user_id" binding:"required"`
	Balance float64 `json:"balance" binding:"required"`
}

type StatusResponse struct {
	Status string `json:"status"`
}

type UserReserveMoneyRequest struct {
	UserId int     `json:"user_id" binding:"required"`
	Price  float64 `json:"price" binding:"required"`
}

type UserDecisionRequest struct {
	UserId int     `json:"user_id" binding:"required"`
	Cost   float64 `json:"cost" binding:"required"`
}

func (i UserAddBalanceRequest) Validate() error {
	if i.Balance < 0 {
		return errors.New("Wrong input - 'balance'")
	}

	return nil
}
