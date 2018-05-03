package main

import (
	"log"
	"net/http"

	"code.czg666.cf/wechat_bot/client"
	"github.com/gin-gonic/gin"
)

//Ping ...
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

//Login wechat
func Login(c *gin.Context) {
	uuid, err := client.GetUUID()
	if err != nil {
		log.Println(err)
		c.Abort()
		return
	}
	str, err := client.GetQRcode(uuid)
	if err != nil {
		log.Println(err)
		c.Abort()
		return
	}
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"Image": str,
	})
}
