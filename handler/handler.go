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
		if len(splitCommand) < 3 {
			return common.ErrNotEnoughParams
		}
		return handleSetCommand(commandListService.StringService, splitCommand)
	case "get":
		if len(splitCommand) < 2 {
			return common.ErrNotEnoughParams
		}
		return handleGetCommand(commandListService.StringService, splitCommand[1])

	case "llen":
		if len(splitCommand) < 2 {
			return common.ErrNotEnoughParams
		}
		return handleLLenCommand(commandListService.ListService, splitCommand[1])
	case "rpush":
		if len(splitCommand) < 3 {
			return common.ErrNotEnoughParams
		}
		return handleRPushCommand(commandListService.ListService, splitCommand)
	case "lpop":
		if len(splitCommand) < 2 {
			return common.ErrNotEnoughParams
		}
		return handleLPopCommand(commandListService.ListService, splitCommand[1])
	case "rpop":
		if len(splitCommand) < 2 {
			return common.ErrNotEnoughParams
		}
		return handleRPopCommand(commandListService.ListService, splitCommand[1])
	case "lrange":
		if len(splitCommand) < 3 {
			return common.ErrNotEnoughParams
		}
		return handleLRangeCommand(commandListService.ListService, splitCommand)

	case "keys":
		return handleKeysCommand(commandListService.GlobalService)
	case "del":
		if len(splitCommand) < 2 {
			return common.ErrNotEnoughParams
		}
		return handleDelCommand(commandListService.GlobalService, splitCommand[1])
	case "flushdb":
		return handleFlushDBCommand(commandListService.GlobalService)
	case "expire":
		if len(splitCommand) < 3 {
			return common.ErrNotEnoughParams
		}
		return handleExpireCommand(commandListService.GlobalService, splitCommand[1], splitCommand[2])
	case "ttl":
		if len(splitCommand) < 2 {
			return common.ErrNotEnoughParams
		}
		return handleTTLCommand(commandListService.GlobalService, splitCommand[1])

	case "save":
		return handleSaveCommand(commandListService.GlobalService)
	case "restore":
		return handleRestoreCommand(commandListService.GlobalService)
	default:
		return common.ErrUnknownCommand
	}
}
