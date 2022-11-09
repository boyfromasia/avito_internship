package handler

import (
	"avito_internship/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetBalanceUser(c *gin.Context) {
	var input models.UserGetBalance
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
	var input models.UserAddBalance
	var response models.UserAddBalanceResponse

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.services.User.AddBalanceUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}
