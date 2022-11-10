package handler

import (
	"avito_internship/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) GetHistory(c *gin.Context) {
	var requestHistory models.GetHistoryRequest

	if err := c.BindJSON(&requestHistory); err != nil {
		logrus.Println(err.Error())
		newErrorResponse(c, http.StatusBadRequest, "Wrong Data")
		return
	}

	responseHistory, err := h.services.HistoryUser.GetHistory(requestHistory)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, responseHistory)
}
