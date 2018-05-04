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
	c.File("./static/index.html")
}

//GetQRcode ...
func GetQRcode(c *gin.Context) {
	uuid, err := client.GetUUID()
	if err != nil {
		log.Println(err)
		c.Abort()
		return
	}
	str := client.GetQRcodeLink(uuid)

	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"code": str,
	})
}
