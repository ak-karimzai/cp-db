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

type BillsController struct {
	billsServices *services.BillsServices
}

func NewBillsController(
	billsServices *services.BillsServices) *BillsController {
	return &BillsController{
		billsServices: billsServices,
	}
}

func (bc *BillsController) CreateBill(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		logger.GetLogger().Error("Error while reading "+
			"create bill request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var bill models.Bill
	err = json.Unmarshal(body, &bill)
	if err != nil {
		logger.GetLogger().Error("Error while unmarshaling bill " +
			"create bill request body")
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := bc.billsServices.
		CreateBill(&bill)
	if responseErr != nil {
		ctx.JSON(responseErr.Status,
			jsonResponse(responseErr, nil))
		return
	}
	ctx.JSON(http.StatusCreated,
		jsonResponse(nil, &response))
}

func (bc *BillsController) GetBill(ctx *gin.Context) {
	billId := ctx.Param("id")
	response, responseErr := bc.billsServices.
		GetBill(billId)
	if responseErr != nil {
		ctx.JSON(responseErr.Status,
			jsonResponse(responseErr, nil))
		return
	}

	ctx.JSON(http.StatusOK,
		jsonResponse(nil, &response))
}

func (bc *BillsController) GetUserBills(ctx *gin.Context) {
	userId := ctx.Param("user_id")

	response, responseErr := bc.billsServices.
		GetUserBills(userId)
	if responseErr != nil {
		ctx.JSON(responseErr.Status,
			jsonResponse(responseErr, nil))
		return
	}

	ctx.JSON(http.StatusOK,
		jsonResponse(nil, &response))
}

func (bc *BillsController) GetAll(ctx *gin.Context) {
	params := ctx.Request.URL.Query()
	paid := params.Get("paid")

	var response []*models.Bill
	var responseErr *models.ResponseError

	switch paid {
	case "true":
		response, responseErr = bc.billsServices.
			GetAllPaid()
	case "false":
		response, responseErr = bc.billsServices.
			GetAllPaid()
	default:
		response, responseErr = bc.billsServices.
			GetAll()
	}
	if responseErr != nil {
		ctx.JSON(responseErr.Status,
			jsonResponse(responseErr, nil))
		return
	}
	ctx.JSON(http.StatusOK,
		jsonResponse(nil, response))
}
