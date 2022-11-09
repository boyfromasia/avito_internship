package models

type GetPurchaseRequest struct {
	PurchaseId int    `json:"purchase_id"`
	Name       string `json:"name"`
}
