package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/ak-karimzai/cp-db/models"
	"github.com/ak-karimzai/cp-db/services"
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

func (bc BillsController) CreateBill(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading "+
			"create bill request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var bill models.Bill
	err = json.Unmarshal(body, &bill)
	if err != nil {
		log.Println("Error while unmarshaling bill " +
			"create bill request body")
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := bc.billsServices.
		CreateBill(&bill)
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusCreated, response)
}

func (bc BillsController) GetBill(ctx *gin.Context) {
	billId := ctx.Param("id")
	response, responseErr := bc.billsServices.
		GetBill(billId)
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (bc BillsController) GetAll(ctx *gin.Context) {
	params := ctx.Request.URL.Query()
	paid := params.Get("paid")
	unPaid := params.Get("unpaid")

	var response []*models.Bill
	var responseErr *models.ResponseError
	if paid != "" && unPaid != "" {
		ctx.AbortWithError(http.StatusBadRequest,
			fmt.Errorf("invalid query parameters"))
		return
	} else if paid == "" && unPaid == "" {
		response, responseErr = bc.billsServices.
			GetAll()
	} else if paid != "" {
		response, responseErr = bc.billsServices.
			GetAllPaid()
	} else {
		response, responseErr = bc.billsServices.
			GetAllUnpaid()
	}

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)
}
