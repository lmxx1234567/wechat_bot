package main

import "github.com/gin-gonic/gin"

//GetPattern is the map of urls with GET method
var GetPattern = map[string]gin.HandlerFunc{
	"/ping": Ping,
	"/":     Login,
}

//PostPattern is the map of urls with POST method
var PostPattern = map[string]gin.HandlerFunc{}

//BothPattern is the map of urls with GET&POST method
var BothPattern = map[string]gin.HandlerFunc{}

//SetRouters set how to handle with request url
func SetRouters(r *gin.Engine) {

	for path, f := range GetPattern {
		r.GET(path, f)
	}

	for path, f := range PostPattern {
		r.POST(path, f)
	}

	for path, f := range BothPattern {
		r.GET(path, f)
		r.POST(path, f)
	}
}
