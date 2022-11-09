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
	var response models.AddRecordResponse

	if err := record.Validate(); err != nil {
		response.Status = "Error"
		return response, err
	}

	record.TimeCreated = time.Now()
	record.Status = "waited"

	return s.repo.AddRecordOrder(record)
}

func (s *OrderService) DecisionReserveUser(decision models.DecisionReserveRequest) (models.DecisionReserveResponse, error) {
	var response models.DecisionReserveResponse

	if err := decision.Validate(); err != nil {
		response.Status = "Error"
		return response, err
	}

	return s.repo.DecisionReserveUser(decision)
}
