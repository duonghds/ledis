package ledis_global

import (
	"errors"
	"time"
)

type service struct {
	backup map[string]PayloadValue
	keys   map[string]PayloadValue
}

type GlobalService interface {
	AddKey(key string, value PayloadValue)
	Keys() []string
	Del(key string) error
	GetKeys() map[string]PayloadValue
	Save() error
	Restore() error
	FlushDB() error
	Expire(key string, second int) int
	TTL(key string) int
}

func InitGlobal() GlobalService {
	keys := make(map[string]PayloadValue)
	backup := make(map[string]PayloadValue)
	return &service{
		keys:   keys,
		backup: backup,
	}
}

func (s *service) AddKey(key string, value PayloadValue) {
	s.keys[key] = value
}

func (s *service) Keys() []string {
	listKeys := make([]string, 0, len(s.keys))
	for k := range s.keys {
		listKeys = append(listKeys, k)
	}
	return listKeys
}

func (s *service) Del(key string) error {
	if _, ok := s.keys[key]; !ok {
		return errors.New("key not exists")
	}
	delete(s.keys, key)
	return nil
}

func (s *service) GetKeys() map[string]PayloadValue {
	//check and delete key expire
	for k, v := range s.keys {
		if v.ExpireTime.After(time.Now()) {
			delete(s.keys, k)
		}
	}
	return s.keys
}

func (s *service) Save() error {
	s.backup = make(map[string]PayloadValue)
	for k, v := range s.keys {
		s.backup[k] = v
	}
	return nil
}

func (s *service) Restore() error {
	if len(s.backup) == 0 {
		return errors.New("not have snapshot before")
	}
	s.keys = make(map[string]PayloadValue)
	for k, v := range s.backup {
		s.keys[k] = v
	}
	return nil
}

func (s *service) FlushDB() error {
	s.keys = map[string]PayloadValue{}
	return nil
}

func (s *service) Expire(key string, second int) int {
	if _, ok := s.keys[key]; !ok {
		return -1
	}
	payloadValue, _ := s.keys[key]
	payloadValue.ExpireTime = time.Now().Add(time.Second * time.Duration(second))
	s.keys[key] = payloadValue
	return second
}
func (s *service) TTL(key string) int {
	if _, ok := s.keys[key]; !ok {
		return -1
	}
	payloadValue, _ := s.keys[key]
	return payloadValue.ExpireTime.Second() - time.Now().Second()
}
