package controllers

import (
	"hostgator-challenge/database"
	"hostgator-challenge/libs"
	"hostgator-challenge/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateLogin controller create login
func CreateLogin(c *gin.Context) {
	db := c.MustGet("db").(*database.GormRepo)

	var input models.Login
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pwd := libs.NewPasswordGen()
	gen, err := pwd.Generate(input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	model := &models.Login{
		UserName: input.UserName,
		Password: gen,
	}

	if err := db.Create(model); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": model})
}

// FindLogin controller find login by ID
func FindLogin(c *gin.Context) {
	db := c.MustGet("db").(*database.GormRepo)

	uit, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var login models.Login
	err = db.FindOne(models.Login{ID: uit}, &login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": login})
}
