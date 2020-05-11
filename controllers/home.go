package controllers

import "github.com/gin-gonic/gin"

func Welcome(c *gin.Context) {
	c.JSON(200, map[string]string{"message": "Welcome to code challenger hostgator"})
}

func Ping(c *gin.Context) {
	c.String(200, "pong")
}
