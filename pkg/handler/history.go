package handler

import (
	"avito_internship/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetHistory(c *gin.Context) {
	var requestHistory models.GetHistoryRequest

	if err := c.BindJSON(&requestHistory); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	responseHistory, err := h.services.HistoryUser.GetHistory(requestHistory)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, responseHistory)
}
