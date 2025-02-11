package handlers

import (
    "net/http"
    "tax-travel-tracker/src/db/models"
	"tax-travel-tracker/src/db/repository"

    "github.com/gin-gonic/gin"
)

// CreateUserHandler godoc
// @Summary Create a new user
// @Description Add a new user record to the users table
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.User true "User"
// @Success 201 {object} models.User
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users [post]
func CreateUserHandler(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // CreateUserRecord is assumed to be a function that saves the user to the DB
    if err := repository.CreateUserRecord(&user); err != nil {
        // You could handle "email already in use" or other DB errors differently
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Return the created user with 201 status
    c.JSON(http.StatusCreated, user)
}
