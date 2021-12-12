package handler

import (
	"fmt"
	"github.com/duonghds/ledis/ledis_set"
	"strconv"
)

func handleSAddCommand(setService ledis_set.SetService, splitCommand []string) string {
	msg, err := setService.SAdd(splitCommand[1], splitCommand[2:])
	if err != nil {
		return err.Error()
	}
	return msg
}

func handleSCardCommand(setService ledis_set.SetService, key string) string {
	result := setService.SCard(key)
	return strconv.Itoa(result)
}

func handleSMembersCommand(setService ledis_set.SetService, key string) string {
	result := setService.SMembers(key)
	return fmt.Sprintf("%+v", result)
}

func handleSRemCommand(setService ledis_set.SetService, splitCommand []string) string {
	msg := setService.SRem(splitCommand[1], splitCommand[2:])

	return msg
}

func handleSInterCommand(setService ledis_set.SetService, splitCommand []string) string {
	result := setService.SInter(splitCommand[1:])

	return fmt.Sprintf("%+v", result)
}
