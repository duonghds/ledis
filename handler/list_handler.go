package handler

import (
	"fmt"
	"github.com/duonghds/ledis/common"
	"github.com/duonghds/ledis/ledis_list"
	"strconv"
)

func handleRPushCommand(listService ledis_list.ListService, splitCommand []string) string {
	if len(splitCommand) < 3 {
		return common.ErrNotEnoughParams
	}
	countAppend, err := listService.RPush(splitCommand[1], splitCommand[2:])
	if err != nil {
		return err.Error()
	}
	return strconv.Itoa(countAppend)
}

func handleLRangeCommand(listService ledis_list.ListService, splitCommand []string) string {
	if len(splitCommand) < 3 {
		return common.ErrNotEnoughParams
	}
	start, err := strconv.Atoi(splitCommand[2])
	if err != nil {
		return common.ErrWrongType
	}
	stop := start
	if len(splitCommand) > 3 {
		stop, err = strconv.Atoi(splitCommand[3])
		if err != nil {
			return common.ErrWrongType
		}
	}
	rangeList, err := listService.LRange(splitCommand[1], start, stop)
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("%+v", rangeList)
}
