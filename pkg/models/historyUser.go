package models

import "time"

type AddHistoryRequest struct {
	UserId      int       `json:"user_id" binding:"required"`
	Cost        string    `json:"cost" binding:"required"`
	Comment     string    `json:"comment" binding:"required"`
	TimeCreated time.Time `json:"time_created"`
}

type AddHistoryResponse struct {
	Status string `json:"status"`
}
