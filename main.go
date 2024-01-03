package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ctx = context.TODO()
var apiTestDB *pgxpool.Pool

func init() {
	initDogBreeds()
	err := initDB(ctx)
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}
}

func main() {

	defer apiTestDB.Close()

	router := gin.Default()
	router.GET("/getusers", getUsers)
	router.GET("/getusersbyid/:id")
	router.POST("/adduser", addUsers)
	router.PATCH("updateemail/:id", updateEmail)
	//dogapi
	router.GET("/dog/getbreeds", getDogBreeds)
	router.Run("localhost:8080")
}

func getUsers(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, users)
}

func addUsers(context *gin.Context) {
	var newUser user

	err := context.BindJSON(&newUser)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	err = newUser.ValidateNewUser()
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, err)
	}
	users = append(users, newUser)

	context.IndentedJSON(http.StatusCreated, newUser)
}

func getUser(context *gin.Context) {
	id := context.Param("id")
	u, err := getUserByID(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, id)
		return
	}
	context.IndentedJSON(http.StatusOK, u)
}

func getUserByID(id string) (*user, error) {
	for i, u := range users {
		if u.ID == id {
			return &users[i], nil
		}
	}
	return nil, errors.New("ID not found")
}

func updateEmail(context *gin.Context) {
	id := context.Param("id")

	var newEmail email

	err := context.BindJSON(&newEmail)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	err = updateUserEmail(id, newEmail.Email)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, err)
	}
}

func updateUserEmail(id string, email string) error {
	for i, u := range users {
		if u.ID == id {
			users[i].Email = email
			return nil
		}
	}
	return errors.New("ID not found")
}
