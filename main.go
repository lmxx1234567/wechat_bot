package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.StaticFS("/statics", http.Dir("./statics"))
	router.StaticFile("/", "./statics/index.html")
	SetRouters(router)
	router.Run(":9876")
}
