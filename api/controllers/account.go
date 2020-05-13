package controllers

import (
	"net/http"
	"strconv"

	"hostgator-challenge/api/libs"
	"hostgator-challenge/api/models"

	"github.com/gin-gonic/gin"
)

// CreateAccount controller create Account
// @Summary Create a new Account
// @Description This method you will create a new Login
// @Accept json
// @Produce json
// @Success 201 {object} models.Login
// @Router /account [post]
// @Param login body models.ReqLogin true "Login"
func (ctl *CtlRepo) CreateAccount(c *gin.Context) {
	var input models.ReqLogin
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

	if err := ctl.DB.Create(model); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": model})
}

// FindAccount controller find login by ID
// @Summary Get information about accounts
// @Description Get a JSON with search by ID
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} models.CatAPI
// @Router /account/{id} [get]
func (ctl *CtlRepo) FindAccount(c *gin.Context) {

	uit, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var login models.Login
	if err := ctl.DB.FindOne(models.Login{ID: uit}, &login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, login)
}
