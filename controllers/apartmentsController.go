package controllers

import (
	"io"
	"log"
	"net/http"

	"github.com/ak-karimzai/cp-db/models"
	"github.com/ak-karimzai/cp-db/services"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

type ApartmentsController struct {
	apartmentsServices *services.ApartmentsServices
	usersServices      *services.UsersServices
}

func NewApartmentsController(apartmentsServices *services.ApartmentsServices,
	usersServices *services.UsersServices) *ApartmentsController {
	return &ApartmentsController{
		apartmentsServices: apartmentsServices,
		usersServices:      usersServices,
	}
}

func (rh ApartmentsController) CreateApartment(
	ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading "+
			"create apartment request body", err)
		ctx.AbortWithError(http.
			StatusInternalServerError, err)
		return
	}

	var apartment models.Apartment
	err = json.Unmarshal(body, &apartment)
	if err != nil {
		log.Println("Error while unmarshaling result " +
			"create aparment request body")
	}

	response, responseErr := rh.apartmentsServices.
		CreateApartment(&apartment)
	if responseErr != nil {
		ctx.JSON(responseErr.Status,
			responseErr)
		return
	}
	ctx.JSON(http.StatusCreated, response)
}

func (rh ApartmentsController) GetApartment(
	ctx *gin.Context) {
	apartmentId := ctx.Param("id")
	aparment, responseErr := rh.apartmentsServices.
		GetApartment(apartmentId)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, aparment)
}

func (rh ApartmentsController) GetAll(
	ctx *gin.Context) {
	userId := ctx.Query("userId")

	var response []*models.Apartment
	var responseErr *models.ResponseError
	if userId == "" {
		response, responseErr = rh.
			apartmentsServices.GetAll()
	} else {
		response, responseErr = rh.
			apartmentsServices.GetUserAparments(userId)
	}

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}
