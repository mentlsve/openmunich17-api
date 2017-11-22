package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	r := gin.New()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	http.Handle("/",r)
}
