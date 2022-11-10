package models

import (
	"errors"
	"time"
)

type AddHistoryRequest struct {
	UserId      int       `json:"user_id" binding:"required"`
	Cost        float64   `json:"cost" binding:"required"`
	Comment     string    `json:"comment" binding:"required"`
	TimeCreated time.Time `json:"time_created"`
}

type AddHistoryResponse struct {
	Status string `json:"status"`
}

type GetHistoryRequest struct {
	UserId   int    `json:"user_id" binding:"required"`
	SortType string `json:"sort_type"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
	SortCol  string `json:"sort_col"`
}

type GetHistoryResponse struct {
	History []GetHistoryResponseItem ` json:"history" binding:"required"`
}

type GetHistoryResponseItem struct {
	Cost        float64   `json:"cost" binding:"required"`
	Comment     string    `json:"comment" binding:"required"`
	TimeCreated time.Time `json:"time_created"`
}

func (i GetHistoryRequest) Validate() error {
	if i.Offset < 0 || i.Limit < 0 {
		return errors.New("Wrong input - negative value")
	}

	if !(i.SortType == "ASC" || i.SortType == "DESC") {
		return errors.New("Wrong input - 'SortType'")
	}

	if !(i.SortCol == "cost" || i.SortCol == "timecreated") {
		return errors.New("Wrong input - 'SortCol'")
	}

	return nil
}
