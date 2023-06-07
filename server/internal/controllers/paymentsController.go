package controllers

import (
	"io"
	"log"
	"net/http"

	"github.com/ak-karimzai/cp-db/internal/models"
	"github.com/ak-karimzai/cp-db/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

type PaymentsController struct {
	paymentsServices *services.PaymentsServices
	usersServices    *services.UsersServices
}

func NewPaymentsController(paymentsServices *services.PaymentsServices,
	usersServices *services.UsersServices) *PaymentsController {
	return &PaymentsController{
		paymentsServices: paymentsServices,
		usersServices:    usersServices,
	}
}

func (pc *PaymentsController) CreatePayment(
	ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create payment " +
			"request body")
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var payment models.Payment
	err = json.Unmarshal(body, &payment)
	if err != nil {
		log.Println("Error while unmarshaling create "+
			"payment request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response, responseErr := pc.paymentsServices.
		CreatePayment(&payment)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status,
			jsonResponse(responseErr, nil))
		return
	}

	ctx.JSON(http.StatusOK,
		jsonResponse(nil, &response))
}

func (pc *PaymentsController) GetAllPayments(
	ctx *gin.Context) {
	var usrId = ctx.Query("userId")

	var response []*models.Payment
	var responseErr *models.ResponseError

	if usrId == "" {
		response, responseErr = pc.paymentsServices.
			GetAllPayments()
	} else {
		response, responseErr = pc.paymentsServices.
			GetUserPayments(usrId)
	}
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status,
			jsonResponse(responseErr, nil))
		return
	}
	ctx.JSON(http.StatusOK,
		jsonResponse(nil, &response))
}

func (pc *PaymentsController) GetPayment(
	ctx *gin.Context) {
	id := ctx.Param("id")
	response, responseErr := pc.paymentsServices.
		GetPayment(id)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status,
			jsonResponse(responseErr, nil))
		return
	}
	ctx.JSON(http.StatusOK,
		jsonResponse(nil, &response))
}

func (pc *PaymentsController) GetUserPayments(
	ctx *gin.Context) {
	userId := ctx.Param("user_id")
	response, responseErr := pc.paymentsServices.
		GetUserPayments(userId)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status,
			jsonResponse(responseErr, nil))
		return
	}
	ctx.JSON(http.StatusOK,
		jsonResponse(nil, &response))
}
