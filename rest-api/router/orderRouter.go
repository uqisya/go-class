package router

import (
	"rest-api/controller"

	"github.com/gin-gonic/gin"
)

func StartOrderServer() *gin.Engine {
	// return router
	router := gin.Default()

	router.POST("/orders", controller.CreateOrder)
	router.GET("/orders", controller.GetAllOrders)
	router.GET("/orders/:id", controller.GetOrderDetailByID)
	router.PUT("/orders/:id", controller.UpdateOrderByID)
	router.DELETE("/orders/:id", controller.DeleteOrderByID)

	return router
}
