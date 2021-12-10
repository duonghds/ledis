package main

import (
	"github.com/duonghds/ledis/handler"
	"github.com/duonghds/ledis/ledis_global"
	"github.com/duonghds/ledis/ledis_string"
	"github.com/gin-gonic/gin"
)

func main() {
	commandListService := initCommandService()
	router := gin.Default()
	v1 := router.Group("/api")
	{
		v1.POST("/", handler.CommandHandler(commandListService))
	}

	router.Run(":9900")
}

func initCommandService() *handler.ListService {
	globalService := ledis_global.InitGlobal()

	stringService := ledis_string.NewService(globalService)

	return &handler.ListService{
		StringService: stringService,
	}
}