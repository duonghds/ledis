package ledis_set

import (
	"errors"
	"github.com/duonghds/ledis/common"
	"github.com/duonghds/ledis/ledis_global"
)

const (
	Type = "Set"
)

type service struct {
	globalService ledis_global.GlobalService
}

type SetService interface {
	SAdd(key string, values []string) (string, error)
	SCard(key string) int
	SMembers(key string) []string
	SRem(key string, values []string) string
	SInter(inputKeys []string) []string
}

func NewService(globalService ledis_global.GlobalService) SetService {
	return &service{
		globalService: globalService,
	}
}

func (s *service) SAdd(key string, values []string) (string, error) {
	if key == "" {
		return "", errors.New(common.ErrKeyEmpty)
	}
	keys := s.globalService.GetKeys()
	var payload ledis_global.PayloadValue
	if _, ok := keys[key]; !ok { //Create new key if not exists
		setValues := make(map[string]bool)
		for i := 0; i < len(values); i++ {
			setValues[values[i]] = true
		}
		payload = ledis_global.PayloadValue{
			Type:  Type,
			Value: setValues,
		}
	} else { //Append new list value to existing key
		payloadValue := keys[key]
		setValues := payloadValue.Value.(map[string]bool)
		for i := 0; i < len(values); i++ {
			setValues[values[i]] = true
		}
		payload = ledis_global.PayloadValue{
			Type:  Type,
			Value: setValues,
		}
	}
	s.globalService.AddKey(key, payload)
	return "OK", nil
}

func (s *service) SCard(key string) int {
	if key == "" {
		return 0
	}
	keys := s.globalService.GetKeys()
	if _, ok := keys[key]; !ok {
		return 0
	}
	payloadValue, _ := keys[key]
	setValues := payloadValue.Value.(map[string]bool)
	return len(setValues)
}

func (s *service) SMembers(key string) []string {
	if key == "" {
		return []string{}
	}
	keys := s.globalService.GetKeys()
	if _, ok := keys[key]; !ok {
		return []string{}
	}
	payloadValue, _ := keys[key]
	setValues := payloadValue.Value.(map[string]bool)
	var result []string
	for k, _ := range setValues {
		result = append(result, k)
	}
	return result
}

func (s *service) SRem(key string, values []string) string {
	if key == "" {
		return common.ErrKeyEmpty
	}
	keys := s.globalService.GetKeys()
	if _, ok := keys[key]; !ok {
		return common.ErrKeyNotFound
	}
	payloadValue := keys[key]
	setValues := payloadValue.Value.(map[string]bool)
	for i := 0; i < len(values); i++ {
		delete(setValues, values[i])
	}
	payloadValue = ledis_global.PayloadValue{
		Type:  Type,
		Value: setValues,
	}
	s.globalService.AddKey(key, payloadValue)
	return "OK"
}

func (s *service) SInter(inputKeys []string) []string {
	if len(inputKeys) == 0 {
		return []string{}
	}
	var result []string
	interSet := make(map[string]bool)
	keys := s.globalService.GetKeys()
	for i := 0; i < len(inputKeys); i++ {
		if _, ok := keys[inputKeys[i]]; !ok {
			continue
		}
		payloadValue := keys[inputKeys[i]].Value
		set := payloadValue.(map[string]bool)
		for k, _ := range set {
			if _, ok := interSet[k]; !ok {
				interSet[k] = true
				result = append(result, k)
			}
		}
	}
	payload := ledis_global.PayloadValue{
		Type:  Type,
		Value: interSet,
	}
	s.globalService.AddKey(inputKeys[0], payload)
	return result
}
