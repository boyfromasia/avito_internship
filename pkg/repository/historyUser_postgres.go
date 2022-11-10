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

func (r *HistoryUserPostgres) GetHistory(record models.GetHistoryRequest) (models.GetHistoryResponse, error) {
	var history []models.GetHistoryResponseItem
	var response models.GetHistoryResponse

	if record.Limit == 0 {
		query := fmt.Sprintf("SELECT cost, comment, timecreated FROM %s WHERE userid=$1 ORDER BY $2 %s LIMIT ALL OFFSET $3", historyUserTable, record.SortType)
		err := r.db.Select(&history, query, record.UserId, record.SortCol, record.Offset)

		if err != nil {
			return response, err
		}
	} else {
		query := fmt.Sprintf("SELECT cost, comment, timecreated FROM %s WHERE userid=$1 ORDER BY $2 %s LIMIT $3 OFFSET $4", historyUserTable, record.SortType)
		err := r.db.Select(&history, query, record.UserId, record.SortCol, record.Limit, record.Offset)

		if err != nil {
			return response, err
		}
	}

	response.History = history

	return response, nil
}
