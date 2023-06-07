package server

import (
	"github.com/ak-karimzai/cp-db/cmd/middlewares"
	"github.com/ak-karimzai/cp-db/internal/logger"
	"github.com/gin-gonic/gin"
)

func setHandlers(server *HttpServer) *gin.Engine {
	gin.DefaultWriter = logger.GetLogger().Writer()
	gin.DefaultErrorWriter = logger.GetLogger().Writer()
	router := gin.Default()

	router.Use(middlewares.CORSMiddleware())

	router.POST("/users/login", server.usersController.Login)

	router.POST("/users/logout", server.usersController.Logout)
	router.POST("/users/signup", server.usersController.Signup)
	router.GET("/users", server.usersController.GetAll)
	router.GET("/users/:user_id", server.usersController.GetUser)

	router.GET("/services", server.servicesController.GetAllServices)
	router.GET("/services/:id", server.servicesController.GetService)
	router.POST("/services", server.servicesController.CreateService)

	router.GET("/apartments", server.apartmentsController.GetAll)
	router.GET("/apartments/:id", server.apartmentsController.GetApartment)
	router.POST("/apartments", server.apartmentsController.CreateApartment)

	router.Group("/bills")
	{
		router.POST("/bills", server.billsController.CreateBill)
		router.GET("/bills", server.billsController.GetAll)
		router.GET("/bills/:id", server.billsController.GetBill)
		router.GET("/users/:user_id/bills", server.billsController.GetUserBills)
	}

	router.GET("/payments/:id", server.paymentsController.GetPayment)
	router.GET("/payments", server.paymentsController.GetAllPayments)
	router.POST("/payments", server.paymentsController.CreatePayment)

	return router
}
