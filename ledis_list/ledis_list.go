package ledis_list

import (
	"errors"
	"fmt"
	"github.com/duonghds/ledis/common"
	"github.com/duonghds/ledis/ledis_global"
)

const (
	Type = "List"
)

type service struct {
	globalService ledis_global.GlobalService
}

type ListService interface {
	LLen(key string) int
	RPush(key string, value []string) (int, error)
	LPop(key string) string
	RPop(key string) string
	LRange(key string, start int, stop int) ([]string, error)
}

func NewService(globalService ledis_global.GlobalService) ListService {
	return &service{
		globalService: globalService,
	}
}

func (s *service) LLen(key string) int {
	return 0
}

func (s *service) RPush(key string, addingValue []string) (int, error) {
	if key == "" {
		return 0, errors.New(common.ErrKeyEmpty)
	}
	keys := s.globalService.GetKeys()
	var payload ledis_global.PayloadValue
	if _, ok := keys[key]; !ok { //Create new key if not exists
		payload = ledis_global.PayloadValue{
			Type:  Type,
			Value: addingValue,
		}
	} else { //Append new list value to existing key
		payloadValue := keys[key]
		value := payloadValue.Value.([]string)
		value = append(value, addingValue...)
		payload = ledis_global.PayloadValue{
			Type:  Type,
			Value: value,
		}
	}
	s.globalService.AddKey(key, payload)
	keys = s.globalService.GetKeys()
	fmt.Println(keys)
	return len(addingValue), nil
}

func (s *service) LPop(key string) string {
	return ""
}

func (s *service) RPop(key string) string {
	return ""
}

func (s *service) LRange(key string, start int, stop int) ([]string, error) {
	if key == "" {
		return []string{}, errors.New(common.ErrKeyEmpty)
	}
	if start > stop {
		return []string{}, errors.New(common.ErrStopLessThanStart)
	}
	keys := s.globalService.GetKeys()
	if _, ok := keys[key]; !ok {
		return []string{}, errors.New(common.ErrKeyNotFound)
	}
	payloadValue := keys[key]
	value := payloadValue.Value.([]string)

	if stop > len(value)-1 {
		stop = len(value) - 1
	}
	return value[start : stop+1], nil
}
