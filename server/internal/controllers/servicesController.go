package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ak-karimzai/cp-db/internal/logger"
	"github.com/ak-karimzai/cp-db/internal/models"
	"github.com/ak-karimzai/cp-db/internal/services"
	"github.com/gin-gonic/gin"
)

type ServicesController struct {
	servicesServices *services.ServicesServices
	usersServices    *services.UsersServices
}

func NewServiceController(servicesServices *services.ServicesServices,
	usersServices *services.UsersServices) *ServicesController {
	return &ServicesController{
		servicesServices: servicesServices,
		usersServices:    usersServices,
	}
}

func (sc *ServicesController) CreateService(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		logger.GetLogger().Error("Error while reading create service " +
			"request body")
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var service models.Service
	err = json.Unmarshal(body, &service)
	if err != nil {
		logger.GetLogger().Error("Error while unmarshaling create "+
			"service request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var response *models.Service
	var responseErr *models.ResponseError
	if service.ID == "" {
		response, responseErr = sc.servicesServices.
			CreateService(&service)
	} else {
		responseErr = sc.servicesServices.
			UpdateService(&service)
	}
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status,
			jsonResponse(responseErr, nil))
		return
	}
	ctx.JSON(http.StatusCreated,
		jsonResponse(nil, &response))
}

func (sc *ServicesController) UpdateService(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		logger.GetLogger().Error("Error while reading update service " +
			"request body")
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var service models.Service
	err = json.Unmarshal(body, &service)
	if err != nil {
		logger.GetLogger().Error("Error while unmarshaling update "+
			"service request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	responseErr := sc.servicesServices.
		UpdateService(&service)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status,
			jsonResponse(responseErr, nil))
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (sc *ServicesController) GetService(ctx *gin.Context) {
	id := ctx.Param("id")
	response, responseErr := sc.servicesServices.
		GetService(id)
	if responseErr != nil {
		ctx.JSON(responseErr.Status,
			jsonResponse(responseErr, nil))
		return
	}
	ctx.JSON(http.StatusOK,
		jsonResponse(nil, &response))
}

func (sc *ServicesController) GetAllServices(ctx *gin.Context) {
	aprId := ctx.Query("aprId")
	var services []*models.Service
	var responseErr *models.ResponseError

	if aprId == "" {
		services, responseErr = sc.servicesServices.
			GetAllServices()
	} else {
		services, responseErr = sc.servicesServices.
			GetApartmentServices(aprId)
	}
	if responseErr != nil {
		ctx.JSON(responseErr.Status,
			jsonResponse(responseErr, nil))
		return
	}

	ctx.JSON(http.StatusOK,
		jsonResponse(nil, services))
}

// func (sc *ServicesController) GetApartmentServices(ctx *gin.Context) {
// 	aprId := ctx.Param("aprId")
// 	response, responseErr := sc.servicesServices.
// 		GetApartmentServices(aprId)
// }
