package models

type GetPurchaseRequest struct {
	PurchaseId int `json:"purchase_id" binding:"required"`
}

type GetPurchaseResponse struct {
	Name string `json:"name"`
}
