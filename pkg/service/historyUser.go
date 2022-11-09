package service

import (
	"avito_internship/pkg/models"
	"avito_internship/pkg/repository"
	"time"
)

type HistoryUserService struct {
	repo repository.HistoryUser
}

func NewHistoryUserService(repo repository.HistoryUser) *HistoryUserService {
	return &HistoryUserService{repo: repo}
}

func (s *HistoryUserService) AddRecordHistory(record models.AddHistoryRequest) (models.AddHistoryResponse, error) {
	record.TimeCreated = time.Now()

	return s.repo.AddRecordHistory(record)
}
