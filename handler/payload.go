package handler

import (
	"github.com/duonghds/ledis/ledis_global"
	"github.com/duonghds/ledis/ledis_list"
	"github.com/duonghds/ledis/ledis_set"
	"github.com/duonghds/ledis/ledis_string"
)

type RequestData struct {
	Command string `json:"command"`
}

type ResponseData struct {
	Message string `json:"message"`
}

type CommandListService struct {
	GlobalService ledis_global.GlobalService
	StringService ledis_string.StringService
	ListService   ledis_list.ListService
	SetService    ledis_set.SetService
}
