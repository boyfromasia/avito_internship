package repository

import (
	"avito_internship/pkg/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type HistoryUserPostgres struct {
	db *sqlx.DB
}

func NewHistoryUserPostgres(db *sqlx.DB) *HistoryUserPostgres {
	return &HistoryUserPostgres{db: db}
}

func (r *HistoryUserPostgres) AddRecordHistory(record models.AddHistoryRequest) (models.AddHistoryResponse, error) {
	var response models.AddHistoryResponse
	var id int

	query := fmt.Sprintf("INSERT INTO %s (userid, cost, comment, timecreated) VALUES ($1, $2, $3, $4) RETURNING historyid", historyUserTable)

	row := r.db.QueryRow(query, record.UserId, record.Cost, record.Comment, record.TimeCreated)
	if err := row.Scan(&id); err != nil {
		response.Status = "Error"
		return response, err

	}

	response.Status = "Ok"

	return response, nil
}
