package models

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

type UserAddBalanceResponse struct {
	Status string `json:"status"`
}

//type UserReserve
