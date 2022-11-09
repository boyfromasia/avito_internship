package handler

import (
	"avito_internship/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/user", h.GetBalanceUser)
	router.POST("/user", h.AddBalanceUser)
	router.POST("/order/create", h.CreateOrder)
	router.POST("/order/approve", h.ApproveOrder)

	//order := router.Group("/order")
	//{
	//	order.POST("/")
	//}
	//
	//service := router.Group("/service")

	return router
}
