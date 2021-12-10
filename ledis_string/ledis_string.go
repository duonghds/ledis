package ledis_string

import (
	"errors"
	"github.com/duonghds/ledis/common"
	"github.com/duonghds/ledis/ledis_global"
)

const (
	Type = "String"
)

type service struct {
	globalService ledis_global.GlobalService
}

type StringService interface {
	Set(key string, value string) (string, error)
	Get(key string) (string, error)
}

func NewService(globalService ledis_global.GlobalService) StringService {
	return &service{
		globalService: globalService,
	}
}

//Set string
func (s *service) Set(key string, value string) (string, error) {
	if key == "" {
		return "", errors.New(common.ErrKeyEmpty)
	}
	payload := ledis_global.PayloadValue{
		Type:  Type,
		Value: value,
	}
	s.globalService.AddKey(key, payload)

	return "OK", nil
}

//Get string
func (s *service) Get(key string) (string, error) {
	if key == "" {
		return "", errors.New(common.ErrKeyEmpty)
	}
	keys := s.globalService.GetKeys()
	if _, ok := keys[key]; !ok {
		return "", errors.New(common.ErrKeyNotFound)
	}
	payloadValue := keys[key]
	value := payloadValue.Value.(string)
	return value, nil
}
