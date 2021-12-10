package handler

import "github.com/duonghds/ledis/ledis_string"

type RequestData struct {
	Command string `json:"command"`
}

type ResponseData struct {
	Message string `json:"message"`
}

type ListService struct {
	StringService ledis_string.StringService
}
