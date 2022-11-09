package service

import (
	"avito_internship/pkg/models"
	"avito_internship/pkg/repository"
)

type OrderService struct {
	repo repository.Order
}

func NewOrderService(repo repository.Order) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) AddRecord(record models.AddRecordRequest) (models.AddRecordResponse, error) {
	var isAdd bool

	if record.PurchaseId == -1 {
		isAdd = true
	} else {
		isAdd = false
	}

	return s.repo.AddRecord(record, isAdd)
}
