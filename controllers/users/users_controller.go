package users

import (
	"net/http"
	"strconv"

	"github.com/Teslenk0/bookstore_users-api/domain/users"
	"github.com/Teslenk0/bookstore_users-api/services"
	"github.com/Teslenk0/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

//We are going to implement in MVC

//Entry points of our application

// All functions implements c *gin.Context Interface

//GetUser - function to handle GET request
func GetUser(c *gin.Context) {

	//Get the id from the GET request
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("Invalid User Id")
		c.JSON(err.Status, err)
		return
	}

	result, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

//CreateUser - function to create a new given user
func CreateUser(c *gin.Context) {
	var user users.User

	//Tries to parse the request to JSON
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid Json Object")
		c.JSON(restErr.Status, restErr)
		return
	}

	//Tries to create the user and persist it
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	//If there is no errors
	c.JSON(http.StatusCreated, result)
}

//SearchUser - function to find a user
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "TODO")

}
