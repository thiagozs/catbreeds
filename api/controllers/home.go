package controllers

import "github.com/gin-gonic/gin"

// @Summary Welcome API
// @Description Home for welcome API
// @Accept  json
// @Produce  json
// @Router / [get]
func (ctl *CtlRepo) Welcome(c *gin.Context) {
	c.JSON(200, map[string]string{"message": "Welcome to code github.com/thiagozs/hostgator-challenger hostgator"})
}

// @Summary Ping service
// @Description Test if API are alive
// @Accept  json
// @Produce  json
// @Success 200 {string} string "pong"
// @Router /ping [get]
func (ctl *CtlRepo) Ping(c *gin.Context) {
	c.String(200, "pong")
}
