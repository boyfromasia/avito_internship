package repository

import (
	"avito_internship/pkg/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) GetBalanceUser(user models.UserGetBalanceRequest) (models.UserGetBalanceResponse, error) {
	var response models.UserGetBalanceResponse
	var balance float64

	query := fmt.Sprintf("SELECT balance from %s WHERE userid=$1", usersTable)
	row := r.db.QueryRow(query, user.UserId)

	if err := row.Scan(&balance); err != nil {
		response.Balance = 0
		return response, err
	}

	response.Balance = balance

	return response, nil
}

func (r *UserPostgres) AddBalanceUser(user models.UserAddBalanceRequest) (models.UserAddBalanceResponse, error) {
	var response models.UserAddBalanceResponse
	var flagExist bool
	var id int

	query := fmt.Sprintf("SELECT EXISTS(SELECT * from %s WHERE userid=$1)", usersTable)
	row := r.db.QueryRow(query, user.UserId)

	if err := row.Scan(&flagExist); err != nil {
		response.Status = "Error"
		return response, err
	}

	if flagExist {
		query := fmt.Sprintf("UPDATE %s SET balance=$1+balance WHERE userid=$2 RETURNING userid", usersTable)
		row := r.db.QueryRow(query, user.Balance, user.UserId)

		if err := row.Scan(&id); err != nil {
			response.Status = "Error with updating user table"
			return response, err
		}
	} else {
		query := fmt.Sprintf("INSERT INTO %s (userid, balance, reserve) VALUES ($1, $2, 0) RETURNING userid", usersTable)
		row := r.db.QueryRow(query, user.UserId, user.Balance)

		if err := row.Scan(&id); err != nil {
			response.Status = "Error with inserting user table"
			return response, err
		}
	}

	response.Status = "OK"
	return response, nil
}
