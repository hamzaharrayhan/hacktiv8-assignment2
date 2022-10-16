package main

import (
	"assignment2/controller"
	"assignment2/database"
	"assignment2/repository"
	"assignment2/service"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.StartDB()
	fmt.Println("Starting DB")

	ConfigApp()
	router := startServer()
	router.Run()
}

func startServer() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodHead, http.MethodDelete, http.MethodOptions, http.MethodPut},
		AllowHeaders:     []string{"Content-Type", "Accept", "Origin", "X-Requested-With", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.POST("/orders", controller.CreateOrder)
	router.GET("/orders", controller.GetOrders)
	router.PUT("/orders/:orderId", controller.UpdateOrders)
	router.DELETE("/orders/:orderId", controller.DeleteOrders)
	return router
}

var repo repository.Repository
var orderService service.Service

func ConfigApp() {
	repo = repository.NewRepository(database.GetDB())
	orderService = service.NewService(repo)
	controller.NewController(orderService)
}
