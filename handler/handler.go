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

func CommandHandler(listService *CommandListService) gin.HandlerFunc {
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

func handleCommand(mainCommand string, splitCommand []string, commandListService *CommandListService) string {
	switch mainCommand {
	case "set":
		return handleSetCommand(commandListService.StringService, splitCommand)
	case "get":
		return handleGetCommand(commandListService.StringService, splitCommand[1])

	case "llen":
		return handleLLenCommand(commandListService.ListService, splitCommand[1])
	case "rpush":
		return handleRPushCommand(commandListService.ListService, splitCommand)
	case "lpop":
		return handleLPopCommand(commandListService.ListService, splitCommand[1])
	case "rpop":
		return handleRPopCommand(commandListService.ListService, splitCommand[1])
	case "lrange":
		return handleLRangeCommand(commandListService.ListService, splitCommand)
	//case "keys":
	//case "del":
	//case "flushdb":
	//case "expire":
	//case "ttl":

	case "save":
		return handleSaveCommand(commandListService.GlobalService)
	case "restore":
		return handleRestoreCommand(commandListService.GlobalService)
	default:
		return common.ErrUnknownCommand
	}
}
