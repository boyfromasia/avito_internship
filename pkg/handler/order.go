package handler

import (
	"avito_internship/pkg/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) DecisionOrder(c *gin.Context) {
	var requestOrder models.DecisionReserveRequest
	var responseOrder models.DecisionReserveResponse

	if err := c.BindJSON(&requestOrder); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// approve or cancel order table
	responseOrder, errOrder := h.services.Order.DecisionReserveUser(requestOrder)
	if errOrder != nil {
		if responseOrder.Status == "Error" {
			newErrorResponse(c, http.StatusBadRequest, errOrder.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, errOrder.Error())
		}
		return
	}

	// change balance of user
	if requestOrder.Decision == "approved" {
		// approve order
		_, errUser := h.services.User.ApproveReserveUser(models.UserDecisionRequest{
			UserId: requestOrder.UserId,
			Cost:   requestOrder.Price,
		})
		if errUser != nil {
			newErrorResponse(c, http.StatusInternalServerError, errOrder.Error())
			return
		}

		// get purchase name
		purchaseName, errPurchase := h.services.Purchase.GetPurchase(models.GetPurchaseRequest{
			PurchaseId: requestOrder.PurchaseId,
		})

		if errPurchase != nil {
			newErrorResponse(c, http.StatusInternalServerError, errPurchase.Error())
			return
		}

		// add record to history table
		_, errHist := h.services.HistoryUser.AddRecordHistory(models.AddHistoryRequest{
			UserId:  requestOrder.UserId,
			Comment: fmt.Sprintf("approve purchase %s", purchaseName.Name),
			Cost:    requestOrder.Price * -1,
		})

		if errHist != nil {
			newErrorResponse(c, http.StatusInternalServerError, errHist.Error())
			return
		}
	} else if requestOrder.Decision == "cancelled" {
		// cancel order
		_, errUser := h.services.User.RejectReserveUser(models.UserDecisionRequest{
			UserId: requestOrder.UserId,
			Cost:   requestOrder.Price,
		})
		if errUser != nil {
			newErrorResponse(c, http.StatusInternalServerError, errOrder.Error())
			return
		}
	}

	c.JSON(http.StatusOK, responseOrder)
}

func (h *Handler) CreateOrder(c *gin.Context) {
	var requestOrder models.AddRecordRequest
	var responseOrder models.AddRecordResponse
	var responseUser models.StatusResponse

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
