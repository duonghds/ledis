package handler

import (
	"encoding/json"
	"fmt"
	"github.com/duonghds/ledis/common"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

func CommandHandler(listService *ListService) gin.HandlerFunc {
	return func(c *gin.Context) {
		jsonData, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Println("[Command]cannot parse body command")
		}

		var requestData RequestData
		err = json.Unmarshal(jsonData, &requestData)
		if err != nil {
			fmt.Println("[Command]cannot parse json command")
		}
		command := common.RemoveDuplicateAndTrimSpace(requestData.Command)

		splitCommand := strings.Split(command, " ")
		mainCommand := strings.ToLower(splitCommand[0])
		result := handleCommand(mainCommand, splitCommand, listService)

		c.JSON(http.StatusOK, ResponseData{
			Message: result,
		})
	}
}

func handleCommand(mainCommand string, splitCommand []string, listService *ListService) string {
	switch mainCommand {
	case "set":
		return handleSetCommand(listService.StringService, splitCommand)
	case "get":
		return handleGetCommand(listService.StringService, splitCommand[1])
	default:
		return "ERROR: Unknown command"
	}
}
