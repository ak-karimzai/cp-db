package server

import (
	"database/sql"
	"log"

	"github.com/ak-karimzai/cp-db/controllers"
	"github.com/ak-karimzai/cp-db/middlewares"
	"github.com/ak-karimzai/cp-db/repositories"
	"github.com/ak-karimzai/cp-db/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config              *viper.Viper
	router              *gin.Engine
	usersController     *controllers.UsersController
	billsController     *controllers.BillsController
	serviceController   *controllers.ServicesController
	apartmentController *controllers.ApartmentsController
}

func InitHttpServer(config *viper.Viper,
	dbHandler *sql.DB) HttpServer {
	log.Println("initializing repositories")
	usersRepository := repositories.
		NewUsersRepository(dbHandler)
	billsRepository := repositories.
		NewBillsRepository(dbHandler)
	servicesRepository := repositories.
		NewServicesRepository(dbHandler)
	apartmentsRepository := repositories.
		NewApartmentsRepository(dbHandler)

	log.Println("initializing services")
	usersServices := services.
		NewUsersServices(usersRepository)
	billsServices := services.
		NewBillsServices(billsRepository)
	servicesServices := services.
		NewServicesServices(servicesRepository)
	apartmentsServices := services.
		NewApartmentsService(apartmentsRepository)

	log.Println("initializing controllers")
	userController := controllers.
		NewUsersController(usersServices)
	billsController := controllers.
		NewBillsController(billsServices)
	servicesController := controllers.
		NewServiceController(servicesServices,
			usersServices)
	apartmentsController := controllers.
		NewApartmentsController(apartmentsServices,
			usersServices)

	router := gin.Default()

	router.POST("/login", userController.Login)

	router.Use(middlewares.BindUserInfo())
	router.POST("/logout", userController.Logout)
	router.POST("/signup", userController.Signup)

	router.Group("/auth", middlewares.BindUserInfo())
	{
		router.GET("/services", servicesController.GetAllServices)
		router.GET("/services/:id", servicesController.GetService)
		router.POST("/services", servicesController.CreateService)
		router.PUT("/services", servicesController.UpdateService)
	}

	router.Group("/auth", middlewares.BindUserInfo())
	{
		router.POST("/apartments", apartmentsController.CreateApartment)
		router.GET("/apartments/:id", apartmentsController.GetApartment)
	}

	router.Group("/auth", middlewares.BindUserInfo())
	{
		router.POST("/bills", billsController.CreateBill)
		router.GET("/bills", billsController.GetAll)
		router.GET("/bills/:id", billsController.GetBill)
	}

	return HttpServer{
		config:              config,
		router:              router,
		usersController:     userController,
		billsController:     billsController,
		serviceController:   servicesController,
		apartmentController: apartmentsController,
	}
}

func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err != nil {
		log.Fatalf("Error while starting HTTP server: %v", err)
	}
}
