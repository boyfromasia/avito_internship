package models

import (
	"time"
)

type AddRecordRequest struct {
	UserId      int       `json:"user_id" binding:"required"`
	PurchaseId  int       `json:"purchase_id" binding:"required"`
	Price       float64   `json:"price" binding:"required"`
	Comment     string    `json:"comment" binding:"required"`
	TimeCreated time.Time `json:"time_created" binding:"required"`
	StatusOrder string    `json:"status_order" binding:"required"`
}

type AddRecordResponse struct {
	OrderId int `json:"order_id"`
}
