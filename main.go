package main

import (
	"github.com/duonghds/ledis/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/api")
	{
		v1.POST("/", handler.CommandHandler)
	}
	router.Run(":9900")
}
