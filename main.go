package main

import (
	"fmt"
	"github.com/duonghds/ledis/handler"
	"github.com/duonghds/ledis/ledis_global"
	"github.com/duonghds/ledis/ledis_list"
	"github.com/duonghds/ledis/ledis_set"
	"github.com/duonghds/ledis/ledis_string"
	"github.com/gin-contrib/cors"
	_ "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"os"
)

func main() {
	commandListService := initCommandService()
	router := gin.Default()
	router.Use(cors.Default())
	v1 := router.Group("/api")
	{
		v1.POST("/", handler.CommandHandler(commandListService))
	}
	port := os.Getenv("PORT")
	router.Run(fmt.Sprintf(":%s", port))
}

func initCommandService() *handler.CommandListService {
	globalService := ledis_global.InitGlobal()

	stringService := ledis_string.NewService(globalService)

	listService := ledis_list.NewService(globalService)

	setService := ledis_set.NewService(globalService)

	return &handler.CommandListService{
		GlobalService: globalService,
		StringService: stringService,
		ListService:   listService,
		SetService:    setService,
	}
}
