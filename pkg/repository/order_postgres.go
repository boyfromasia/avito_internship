package repository

import (
	"avito_internship/pkg/models"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type OrderPostgres struct {
	db *sqlx.DB
}

func NewOrderPostgres(db *sqlx.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}

func (r *OrderPostgres) AddRecordOrder(record models.AddRecordRequest) (models.AddRecordResponse, error) {
	var response models.AddRecordResponse
	var id int

	query := fmt.Sprintf("INSERT INTO %s (orderid, userid, purchaseid, price, timecreated, statusorder)"+
		" VALUES ($1, $2, $3, $4, $5, $6) RETURNING orderid", orderTable)

	row := r.db.QueryRow(query, record.OrderId, record.UserId, record.PurchaseId, record.Price, record.TimeCreated, record.StatusOrder)
	if err := row.Scan(&id); err != nil {
		response.OrderId = -1
		return response, err
	}

	response.OrderId = id

	return response, nil
}

func (r *OrderPostgres) DecisionReserveUser(decision models.DecisionReserveRequest) (models.DecisionReserveResponse, error) {
	var response models.DecisionReserveResponse
	var flagExist bool
	var status string
	var id int

	query := fmt.Sprintf("SELECT EXISTS(SELECT * from %s WHERE orderid=$1 AND userid=$2 AND purchaseid=$3 AND price=$4)", orderTable)
	row := r.db.QueryRow(query, decision.OrderId, decision.UserId, decision.PurchaseId, decision.Price)

	if err := row.Scan(&flagExist); err != nil {
		response.Status = "Error"
		return response, err
	}

	if !flagExist {
		response.Status = "Error"
		return response, errors.New("Error, this order doesn't exist")
	} else {
		query := fmt.Sprintf("SELECT statusorder FROM %s WHERE orderid=$1 AND userid=$2 AND purchaseid=$3 AND price=$4", orderTable)
		row := r.db.QueryRow(query, decision.OrderId, decision.UserId, decision.PurchaseId, decision.Price)
		if err := row.Scan(&status); err != nil {
			response.Status = "Error"
			return response, err
		}

		if status == "waited" {
			query := fmt.Sprintf("UPDATE %s SET statusorder=$1 WHERE orderid=$2 AND userid=$3 AND purchaseid=$4 AND price=$5 RETURNING orderid", orderTable)
			row := r.db.QueryRow(query, decision.Decision, decision.OrderId, decision.UserId, decision.PurchaseId, decision.Price)

			if err := row.Scan(&id); err != nil {
				response.Status = "Error"
				return response, err
			}
			response.Status = "OK"
			return response, nil
		} else {
			response.Status = "Error"
			return response, errors.New("This order has already already processed")
		}
	}
}
