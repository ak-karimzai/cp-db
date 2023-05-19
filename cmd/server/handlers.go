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

	router.POST("/users/login", server.usersController.Login)

	router.Use(middlewares.BindUserInfo())
	router.POST("/users/logout", server.usersController.Logout)
	router.POST("/users/signup", server.usersController.Signup)

	router.Group("/services", middlewares.BindUserInfo())
	{
		router.GET("/services", server.servicesController.GetAllServices)
		router.GET("/services/:id", server.servicesController.GetService)
		router.POST("/services", server.servicesController.CreateService)
		router.PUT("/services", server.servicesController.UpdateService)
	}

	router.Group("/apartments", middlewares.BindUserInfo())
	{
		router.POST("/apartments", server.apartmentsController.CreateApartment)
		router.GET("/apartments/:id", server.apartmentsController.GetApartment)
	}

	router.Group("/bills", middlewares.BindUserInfo())
	{
		router.POST("/bills", server.billsController.CreateBill)
		router.GET("/bills", server.billsController.GetAll)
		router.GET("/bills/:id", server.billsController.GetBill)
		router.GET("/users/:user_id/bills", server.billsController.GetUserBills)
	}

	router.Group("/payments", middlewares.BindUserInfo())
	{
		router.POST("/users/:user_id/payments", server.paymentsController.CreatePayment)
		router.GET("/users/:user_id/payments", server.paymentsController.GetUserPayments)
		router.GET("/payments/:id", server.paymentsController.GetPayment)
		router.GET("/payments", server.paymentsController.GetAllPayments)
	}

	router.Use(middlewares.PageNotFound())

	return router
}
