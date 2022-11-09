package handler

import (
	"avito_internship/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) ApproveOrder(c *gin.Context) {

}

func (h *Handler) CreateOrder(c *gin.Context) {
	var requestOrder models.AddRecordRequest
	var responseOrder models.AddRecordResponse
	var responseUser models.UserReserveMoneyResponse

	if err := c.BindJSON(&requestOrder); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	_, errPurchase := h.services.Purchase.GetPurchase(models.GetPurchaseRequest{
		PurchaseId: requestOrder.PurchaseId,
	})
	if errPurchase != nil {
		newErrorResponse(c, http.StatusBadRequest, errPurchase.Error())
		return
	}

	// reserve user
	responseUser, errUser := h.services.User.ReserveMoneyUser(models.UserReserveMoneyRequest{
		UserId: requestOrder.UserId,
		Price:  requestOrder.Price,
	})

	if errUser != nil {
		if responseUser.Status == "Error" {
			newErrorResponse(c, http.StatusBadRequest, errUser.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, errUser.Error())
		}
		return
	}

	// create order
	responseOrder, errOrder := h.services.Order.AddRecord(requestOrder)
	if errOrder != nil {
		newErrorResponse(c, http.StatusInternalServerError, errOrder.Error())
		return
	}

	c.JSON(http.StatusOK, responseOrder)
}
