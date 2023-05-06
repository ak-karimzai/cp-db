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

type UsersController struct {
	usersServices *services.UsersServices
}

func NewUsersController(
	usersServices *services.UsersServices) *UsersController {
	return &UsersController{
		usersServices: usersServices,
	}
}

func (uc UsersController) Login(ctx *gin.Context) {
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
		log.Println("Error while unmarshaling login " +
			"credintials")
		return
	}
	token, responseErr := uc.usersServices.Login(user.UserName,
		user.Password)
	if responseErr != nil {
		ctx.JSON(http.StatusBadRequest,
			responseErr)
		return
	}

	ctx.SetCookie("Authorization", token,
		3600, "", "", false, true)
	ctx.JSON(http.StatusCreated, gin.H{"token": token})
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
		log.Println("Error while unmarshaling login " +
			"credintials")
		return
	}
	response, responseErr := uc.usersServices.SignUp(&user)
	if responseErr != nil {
		ctx.JSON(http.StatusBadRequest,
			responseErr)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}
