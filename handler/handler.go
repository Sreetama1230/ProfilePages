package handler

import (
	"ProfilePages/db"
	"ProfilePages/models"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

type profileRequest struct {
	Bio     string `json:"bio"`
	Twitter string `json:"twitter"`
}
type request struct {
	Name    string         `json:"name"`
	Email   string         `json:"email"`
	Profile profileRequest `json:"profile"`
}

func CreateUser(c *gin.Context) {
	var input request

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "request payload is not valid"})
		return
	}
	u := models.User{
		Name:  input.Name,
		Email: input.Email,
		Profile: models.Profile{
			Bio:     input.Profile.Bio,
			Twitter: input.Profile.Twitter,
		},
	}
	if err := db.DB.Create(&u).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to save user details"})
		return
	}
	c.JSON(http.StatusCreated, u)
}

func GetUser(c *gin.Context) {
	providedId := c.Param("id")
	id, err := strconv.Atoi(providedId)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "provided id is invalid"})
		return
	}
	var user models.User
	if err := db.DB.Model(&models.User{}).Preload("Profile").First(&user, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch the user."})
		return
	}
	c.JSON(http.StatusOK, user)

}
