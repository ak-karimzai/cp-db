package server

import (
	"database/sql"

	"github.com/ak-karimzai/cp-db/internal/controllers"
	"github.com/ak-karimzai/cp-db/internal/logger"
	"github.com/ak-karimzai/cp-db/internal/repositories"
	"github.com/ak-karimzai/cp-db/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config               *viper.Viper
	router               *gin.Engine
	usersController      *controllers.UsersController
	billsController      *controllers.BillsController
	servicesController   *controllers.ServicesController
	apartmentsController *controllers.ApartmentsController
	paymentsController   *controllers.PaymentsController
}

func InitHttpServer(config *viper.Viper,
	dbHandler *sql.DB) HttpServer {
	logger.GetLogger().Info("initializing repositories")
	usersRepository := repositories.
		NewUsersRepository(dbHandler)
	billsRepository := repositories.
		NewBillsRepository(dbHandler)
	servicesRepository := repositories.
		NewServicesRepository(dbHandler)
	apartmentsRepository := repositories.
		NewApartmentsRepository(dbHandler)
	paymentsRepository := repositories.
		NewPaymentsRepository(dbHandler)

	logger.GetLogger().Info("initializing services")
	usersServices := services.
		NewUsersServices(usersRepository)
	billsServices := services.
		NewBillsServices(billsRepository)
	servicesServices := services.
		NewServicesServices(servicesRepository)
	apartmentsServices := services.
		NewApartmentsService(apartmentsRepository)
	paymentsServices := services.
		NewPaymentsServices(paymentsRepository)

	logger.GetLogger().Info("initializing controllers")
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
	paymentsController := controllers.
		NewPaymentsController(paymentsServices,
			usersServices)

	server := HttpServer{
		config:               config,
		usersController:      userController,
		billsController:      billsController,
		servicesController:   servicesController,
		apartmentsController: apartmentsController,
		paymentsController:   paymentsController,
	}
	server.router = setHandlers(&server)
	return server
}

func (hs HttpServer) Start() {
	var port = hs.config.GetString("http.server_address")
	err := hs.router.Run(port)
	if err != nil {
		logger.GetLogger().Fatalf("Error while starting HTTP server: %v", err)
	}
}
