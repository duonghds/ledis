package ledis_list

import (
	"errors"
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
	RPush(key string, values []string) (int, error)
	LPop(key string) (string, error)
	RPop(key string) (string, error)
	LRange(key string, start int, stop int) ([]string, error)
}

func NewService(globalService ledis_global.GlobalService) ListService {
	return &service{
		globalService: globalService,
	}
}

func (s *service) LLen(key string) int {
	if key == "" {
		return 0
	}
	keys := s.globalService.GetKeys()
	if _, ok := keys[key]; !ok {
		return 0
	}
	payloadValue := keys[key]
	value := payloadValue.Value.([]string)
	return len(value)
}

func (s *service) RPush(key string, addingValues []string) (int, error) {
	if key == "" {
		return 0, errors.New(common.ErrKeyEmpty)
	}
	keys := s.globalService.GetKeys()
	var payload ledis_global.PayloadValue
	if _, ok := keys[key]; !ok { //Create new key if not exists
		payload = ledis_global.PayloadValue{
			Type:  Type,
			Value: addingValues,
		}
	} else { //Append new list value to existing key
		payloadValue := keys[key]
		value := payloadValue.Value.([]string)
		value = append(value, addingValues...)
		payload = ledis_global.PayloadValue{
			Type:  Type,
			Value: value,
		}
	}
	s.globalService.AddKey(key, payload)
	keys = s.globalService.GetKeys()
	return len(addingValues), nil
}

func (s *service) LPop(key string) (string, error) {
	if key == "" {
		return "", errors.New(common.ErrKeyEmpty)
	}
	keys := s.globalService.GetKeys()
	if _, ok := keys[key]; !ok {
		return "", errors.New(common.ErrKeyNotFound)
	}
	payloadValue := keys[key]
	value := payloadValue.Value.([]string)
	popValue := value[0]
	payload := ledis_global.PayloadValue{
		Type:  Type,
		Value: value[1:], //Remove first element
	}
	s.globalService.AddKey(key, payload)
	return popValue, nil
}

func (s *service) RPop(key string) (string, error) {
	if key == "" {
		return "", errors.New(common.ErrKeyEmpty)
	}
	keys := s.globalService.GetKeys()
	if _, ok := keys[key]; !ok {
		return "", errors.New(common.ErrKeyNotFound)
	}
	payloadValue := keys[key]
	value := payloadValue.Value.([]string)
	popValue := value[len(value)-1]
	payload := ledis_global.PayloadValue{
		Type:  Type,
		Value: value[:len(value)-1], //Remove last element
	}
	s.globalService.AddKey(key, payload)
	return popValue, nil
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
