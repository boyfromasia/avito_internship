package handler

import (
	"avito_internship/pkg/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func (h *Handler) GetBalanceUser(c *gin.Context) {
	var input models.UserGetBalanceRequest
	var response models.UserGetBalanceResponse

	if err := c.BindJSON(&input); err != nil {
		logrus.Println(err.Error())
		newErrorResponse(c, http.StatusBadRequest, "Wrong Data")
		return
	}

	response, err := h.services.User.GetBalanceUser(input)
	if err != nil {
		logrus.Println(err.Error())
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) AddBalanceUser(c *gin.Context) {
	var requestUser models.UserAddBalanceRequest
	var responseUser models.StatusResponse
	var _ models.AddHistoryResponse

	if err := c.BindJSON(&requestUser); err != nil {
		logrus.Println(err.Error())
		newErrorResponse(c, http.StatusBadRequest, "Wrong Data")
		return
	}

	// add to balance
	responseUser, errUser := h.services.User.AddBalanceUser(requestUser)
	if errUser != nil {
		logrus.Println(errUser.Error())
		newErrorResponse(c, http.StatusInternalServerError, errUser.Error())
		return
	}

	// create record in historyUser
	_, errHist := h.services.HistoryUser.AddRecordHistory(models.AddHistoryRequest{
		UserId:  requestUser.UserId,
		Comment: "Add to balance from persona X",
		Cost:    requestUser.Balance,
	})

	if errHist != nil {
		logrus.Println(errUser.Error())
		newErrorResponse(c, http.StatusInternalServerError, errHist.Error())
		return
	}

	c.JSON(http.StatusOK, responseUser)
}

func (h *Handler) TransferMoney(c *gin.Context) {
	var requestUser models.UserTransferRequest
	var responseUser models.StatusResponse

	if err := c.BindJSON(&requestUser); err != nil {
		logrus.Println(err.Error())
		newErrorResponse(c, http.StatusBadRequest, "Wrong Data")
		return
	}

	// transfer money
	responseUser, err := h.services.User.TransferMoneyUsers(requestUser)
	if err != nil {
		logrus.Println(err.Error())
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// add addition to history for user 'to_id'
	_, err = h.services.HistoryUser.AddRecordHistory(models.AddHistoryRequest{
		UserId:      requestUser.ToId,
		Cost:        requestUser.Amount,
		Comment:     fmt.Sprintf("Add to balance from user %d", requestUser.FromId),
		TimeCreated: time.Now(),
	})

	// add subtraction to history for user 'from_id'
	_, err = h.services.HistoryUser.AddRecordHistory(models.AddHistoryRequest{
		UserId:      requestUser.FromId,
		Cost:        requestUser.Amount * -1,
		Comment:     fmt.Sprintf("Send money to user %d", requestUser.ToId),
		TimeCreated: time.Now(),
	})

	if err != nil {
		logrus.Println(err.Error())
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, responseUser)
}
