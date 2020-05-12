package controllers

import (
	"hostgator-challenge/database"
	"hostgator-challenge/libs"
	"hostgator-challenge/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Breeds godoc
// @Summary Show a cat information
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param cat path string true "Cat name"
// @Success 200 {object} models.CatAPI
// @Header 200 {string} Token "jwt"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /breeds/{cat} [get]

func Breeds(c *gin.Context) {
	db := c.MustGet("db").(*database.GormRepo)
	key := c.Param("cat")

	var model models.CatAPI
	var results []models.CatAPI
	gdb := db.GetDB()
	err := gdb.Table(model.TableName()).
		Where("name LIKE ?", "%"+key+"%").
		Find(&results).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(results) == 0 {
		data, err := libs.WriteData(db, key)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		results = append(results, data...)
	}

	c.JSON(http.StatusOK, gin.H{"data": results})
}
