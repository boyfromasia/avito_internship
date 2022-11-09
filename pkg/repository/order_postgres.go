package repository

import (
	"avito_internship/pkg/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type OrderPostgres struct {
	db *sqlx.DB
}

func NewOrderPostgres(db *sqlx.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}

func (r *OrderPostgres) AddRecord(record models.AddRecordRequest, isAdd bool) (models.AddRecordResponse, error) {
	var response models.AddRecordResponse
	var id int

	query := fmt.Sprintf("INSERT INTO %s (userid, purchaseid, price, comment, timecreated, statusorder)"+
		" VALUES ($1, $2, $3, $4, $5, $6) RETURNING orderid", orderTable)

	if isAdd {
		row := r.db.QueryRow(query, record.UserId, nil, record.Price, record.Comment, record.TimeCreated, record.StatusOrder)
		if err := row.Scan(&id); err != nil {
			response.OrderId = -1
			return response, err
		}

	} else {
		row := r.db.QueryRow(query, record.UserId, record.PurchaseId, record.Price, record.Comment, record.TimeCreated, record.StatusOrder)
		if err := row.Scan(&id); err != nil {
			response.OrderId = -1
			return response, err
		}
	}

	response.OrderId = id

	return response, nil
}
