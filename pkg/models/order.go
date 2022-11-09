package models

import (
	"time"
)

type AddRecordRequest struct {
	OrderId     int       `json:"order_id" binding:"required"`
	UserId      int       `json:"user_id" binding:"required"`
	PurchaseId  int       `json:"purchase_id" binding:"required"`
	Price       float64   `json:"price" binding:"required"`
	TimeCreated time.Time `json:"time_created"`
	StatusOrder string    `json:"status_order"`
}

type AddRecordResponse struct {
	OrderId int `json:"order_id"`
}
