package service

import (
	"avito_internship/pkg/models"
	"avito_internship/pkg/repository"
	"time"
)

type OrderService struct {
	repo repository.Order
}

func NewOrderService(repo repository.Order) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) AddRecord(record models.AddRecordRequest) (models.AddRecordResponse, error) {
	record.TimeCreated = time.Now()
	record.StatusOrder = "waited"

	return s.repo.AddRecordOrder(record)
}
