package handler

import (
	"avito_internship/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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
	var _ models.AddRecordResponse

	if err := c.BindJSON(&requestUser); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	responseUser, errUser := h.services.User.AddBalanceUser(requestUser)
	if errUser != nil {
		newErrorResponse(c, http.StatusInternalServerError, errUser.Error())
		return
	}

	_, errOrder := h.services.Order.AddRecord(models.AddRecordRequest{
		UserId:      requestUser.UserId,
		PurchaseId:  -1,
		Price:       requestUser.Balance,
		Comment:     "Add to balance",
		TimeCreated: time.Now(),
		StatusOrder: "approved",
	})

	if errOrder != nil {
		newErrorResponse(c, http.StatusInternalServerError, errOrder.Error())
		return
	}

	c.JSON(http.StatusOK, responseUser)
}
