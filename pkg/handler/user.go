package handler

import (
	"avito_internship/pkg/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetBalanceUser(c *gin.Context) {
	var input models.UserGetBalanceRequest
	var response models.UserGetBalanceResponse

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.services.User.GetBalanceUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) AddBalanceUser(c *gin.Context) {
	var requestUser models.UserAddBalanceRequest
	var responseUser models.UserAddBalanceResponse
	var _ models.AddHistoryResponse

	if err := c.BindJSON(&requestUser); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// add to balance
	responseUser, errUser := h.services.User.AddBalanceUser(requestUser)
	if errUser != nil {
		newErrorResponse(c, http.StatusInternalServerError, errUser.Error())
		return
	}

	// create record in historyUser
	_, errHist := h.services.HistoryUser.AddRecordHistory(models.AddHistoryRequest{
		UserId:  requestUser.UserId,
		Comment: "Add to balance from persona X",
		Cost:    fmt.Sprintf("+%f", requestUser.Balance),
	})

	if errHist != nil {
		newErrorResponse(c, http.StatusInternalServerError, errHist.Error())
		return
	}

	c.JSON(http.StatusOK, responseUser)
}
