package controller

import (
	"assignment2/payload"
	"assignment2/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var orderService service.Service

func NewController(service service.Service) {
	orderService = service
}

func GetOrders(c *gin.Context) {
	getOrders, err := orderService.GetOrders()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, getOrders)
}

func CreateOrder(c *gin.Context) {
	var orderPayload payload.OrderPayload
	err := c.ShouldBind(&orderPayload)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	order, err := orderService.CreateOrder(orderPayload)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	message := gin.H{"order": order, "message": "success creating order"}
	c.JSON(http.StatusOK, message)
}

func UpdateOrders(c *gin.Context) {
	var ordersPayload payload.OrderPayload
	orderID := c.Param("orderId")
	err := c.ShouldBind(&ordersPayload)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	updatedOrder, err := orderService.UpdateOrder(orderID, ordersPayload)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	message := gin.H{"updated order": updatedOrder, "message": "success updating order"}
	c.JSON(http.StatusOK, message)
}

func DeleteOrders(c *gin.Context) {
	orderID := c.Param("orderId")
	err := orderService.DeleteOrder(orderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	responseMessage, err := fmt.Printf("success deleting order with id %s", orderID)
	message := gin.H{"message": responseMessage}
	c.JSON(http.StatusOK, message)
}
