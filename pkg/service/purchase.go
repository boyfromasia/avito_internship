package service

import (
	"avito_internship/pkg/models"
	"avito_internship/pkg/repository"
)

type PurchaseService struct {
	repo repository.Purchase
}

func NewPurchaseService(repo repository.Purchase) *PurchaseService {
	return &PurchaseService{repo: repo}
}

func (s *PurchaseService) GetPurchase(purchase models.GetPurchaseRequest) (models.GetPurchaseResponse, error) {
	return s.repo.GetPurchase(purchase)
}
