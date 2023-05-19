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
	response, responseErr := sc.servicesServices.
		CreateService(&service)
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
	userInfo := ctx.MustGet("user").(*models.User)
	if userInfo == nil ||
		userInfo.ID == "" ||
		userInfo.Role == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	response, responseErr := sc.servicesServices.
		GetAllServices()
	if responseErr != nil {
		ctx.JSON(responseErr.Status,
			jsonResponse(responseErr, nil))
		return
	}

	ctx.JSON(http.StatusOK,
		jsonResponse(nil, response))
}
