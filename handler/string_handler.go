package handler

import (
	"github.com/duonghds/ledis/ledis_string"
)

func handleSetCommand(stringService ledis_string.StringService, splitCommand []string) string {
	//all is string value after element[0](command) and element[1](key)
	value := ""
	for i := 2; i < len(splitCommand); i++ {
		value = value + splitCommand[i]
		if i < len(splitCommand)-1 {
			value = value + " "
		}
	}
	msg, err := stringService.Set(splitCommand[1], value)
	if err != nil {
		return err.Error()
	}
	return msg
}

func handleGetCommand(stringService ledis_string.StringService, key string) string {
	value, err := stringService.Get(key)
	if err != nil {
		return err.Error()
	}
	return value
}
