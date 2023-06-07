package controllers

import (
	"io"
	"log"
	"net/http"

	"github.com/ak-karimzai/cp-db/internal/logger"
	"github.com/ak-karimzai/cp-db/internal/models"
	"github.com/ak-karimzai/cp-db/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

type UsersController struct {
	usersServices *services.UsersServices
}

func NewUsersController(
	usersServices *services.UsersServices) *UsersController {
	return &UsersController{
		usersServices: usersServices,
	}
}

func (uc *UsersController) GetUser(ctx *gin.Context) {
	id := ctx.Param("user_id")
	user, err := uc.usersServices.GetUser(id)
	if err != nil {
		logger.GetLogger().
			Error("Error while reading data from db: ", err)
		ctx.JSON(http.StatusBadRequest,
			jsonResponse(err, nil))
		return
	}

	ctx.JSON(http.StatusOK,
		jsonResponse(nil, user))
}

func (uc *UsersController) GetAll(ctx *gin.Context) {
	users, err := uc.usersServices.GetAll()
	if err != nil {
		logger.GetLogger().
			Error("Error while reading data from db: ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK,
		jsonResponse(nil, users))
}

func (uc UsersController) Login(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		logger.GetLogger().Error("Error while parsing login " +
			"credintials")
		ctx.AbortWithError(http.StatusBadRequest,
			err)
		return
	}

	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		logger.GetLogger().Error("Error while unmarshaling login " +
			"credintials")
		return
	}
	userData, token, responseErr := uc.usersServices.Login(user.UserName,
		user.Password)
	if responseErr != nil {
		ctx.JSON(http.StatusBadRequest,
			jsonResponse(responseErr, nil))
		return
	}

	ctx.SetCookie("Authorization", token,
		3600, "", "", false, true)
	ctx.JSON(http.StatusCreated,
		jsonResponse(nil, gin.H{"token": token, "user": userData}))
}

func (uc UsersController) Logout(ctx *gin.Context) {
	ctx.SetCookie("Authorization", "",
		3600, "", "", false, true)
	ctx.Status(http.StatusOK)
}

func (uc UsersController) Signup(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while parsing login " +
			"credintials")
		ctx.AbortWithError(http.StatusBadRequest,
			err)
		return
	}

	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		logger.GetLogger().Error("Error while unmarshaling login " +
			"credintials")
		return
	}
	response, responseErr := uc.usersServices.SignUp(&user)
	if responseErr != nil {
		ctx.JSON(http.StatusBadRequest,
			jsonResponse(responseErr, nil))
		return
	}

	ctx.JSON(http.StatusCreated,
		jsonResponse(nil, response))
}
