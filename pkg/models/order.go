package models

type Order struct {
	OrderId   int `json:"order_id"`
	UserId    int `json:"user_id"`
	ServiceId int `json:"service_id"`
	Price     int `json:"price"`
}
