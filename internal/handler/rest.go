package handler

import (
	"go_learn/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

var usersDb = []model.User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
}

// GetUsers godoc
// @Summary Get list of users
// @Tags users
// @Produce json
// @Success 200 {array} model.User
// @Router /users [get]
func GetUsers(context *gin.Context) {
	context.JSON(http.StatusOK, usersDb)
}

// CreateUser godoc
// @Summary Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.User true "User to create"
// @Success 201 {object} model.User
// @Failure 400 {object} model.ErrorResponse
// @Router /users [post]
func CreateUser(context *gin.Context) {
	var users model.User

	err := context.ShouldBindBodyWithJSON(&users)
	if err != nil {
		context.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	users.ID = len(usersDb) + 1
	usersDb = append(usersDb, users)
	context.JSON(http.StatusCreated, users)
}
