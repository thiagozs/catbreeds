package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thiagozs/hostgator-challenge/api/models"
	"github.com/thiagozs/hostgator-challenge/libs"
)

// @Summary Get information about breed cats
// @Description Get a JSON with search by name
// @Accept  json
// @Produce  json
// @Param cat path string true "Cat Name"
// @Success 200 {object} models.CatAPI
// @Router /breeds/{cat} [get]
func (ctl *CtlRepo) Breeds(c *gin.Context) {
	key := c.Param("cat")

	var model models.CatAPI
	var results []models.CatAPI
	gdb := ctl.DB.GetDB()
	err := gdb.Table(model.TableName()).
		Where("name LIKE ?", "%"+key+"%").
		Find(&results).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(results) == 0 {
		data, err := libs.WriteData(ctl.DB, key)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		results = append(results, data...)
	}

	c.JSON(http.StatusOK, gin.H{"data": results})
}
