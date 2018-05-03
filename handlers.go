package main

import (
	"code.byted.org/gopkg/logs"
	"github.com/gin-gonic/gin"
)

//Ping ...
func Ping(c *gin.Context) {
	logs.Info("a sample app log")
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
