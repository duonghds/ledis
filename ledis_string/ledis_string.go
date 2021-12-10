package ledis_string

import (
	"errors"
	"fmt"
	"github.com/duonghds/ledis/ledis_global"
)

const (
	Type = "String"
)

type service struct {
	globalService ledis_global.GlobalService
}

type StringService interface {
	Set(key string, value string) string
	Get(key string) (string, error)
}

func NewService(globalService ledis_global.GlobalService) StringService {
	return &service{
		globalService: globalService,
	}
}

//Set string
func (s *service) Set(key string, value string) string {
	payload := ledis_global.PayloadValue{
		Type:  Type,
		Value: value,
	}
	s.globalService.AddKey(key, payload)

	return "OK"
}

//Get string
func (s *service) Get(key string) (string, error) {
	keys := s.globalService.GetKeys()
	if _, ok := keys[key]; !ok {
		fmt.Println(fmt.Sprintf("[Get]key %s not exists", key))
		return "", errors.New("key not exists")
	}
	payloadValue := keys[key]
	value := payloadValue.Value.(string)
	return value, nil
}
