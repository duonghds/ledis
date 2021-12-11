package handler

import (
	"fmt"
	"github.com/duonghds/ledis/common"
	"github.com/duonghds/ledis/ledis_global"
	"strconv"
)

func handleKeysCommand(globalService ledis_global.GlobalService) string {
	keys := globalService.Keys()
	return fmt.Sprintf("%+v", keys)
}

func handleDelCommand(globalService ledis_global.GlobalService, key string) string {
	err := globalService.Del(key)
	if err != nil {
		return err.Error()
	}
	return "OK"
}

func handleFlushDBCommand(globalService ledis_global.GlobalService) string {
	err := globalService.FlushDB()
	if err != nil {
		return err.Error()
	}
	return "OK"
}

func handleExpireCommand(globalService ledis_global.GlobalService, key string, secondStr string) string {
	second, err := strconv.Atoi(secondStr)
	if err != nil {
		return common.ErrWrongType
	}
	sec := globalService.Expire(key, second)
	return strconv.Itoa(sec)
}

func handleTTLCommand(globalService ledis_global.GlobalService, key string) string {
	sec := globalService.TTL(key)
	return strconv.Itoa(sec)
}
