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

func GetUsers(context *gin.Context) {
	context.JSON(http.StatusOK, usersDb)
}

func CreateUser(context *gin.Context) {
	var users model.User

	err := context.ShouldBindBodyWithJSON(&users)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users.ID = len(usersDb) + 1
	usersDb = append(usersDb, users)
	context.JSON(http.StatusCreated, users)
}
