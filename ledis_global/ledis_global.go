package ledis_global

type service struct {
	keys map[string]PayloadValue
}

type GlobalService interface {
	AddKey(key string, value PayloadValue)
	Keys() []PayloadValue
	Del(key string) (bool, error)
	GetKeys() map[string]PayloadValue
}

func InitGlobal() GlobalService {
	keys := make(map[string]PayloadValue)
	return &service{
		keys: keys,
	}
}

func (s *service) AddKey(key string, value PayloadValue) {
	s.keys[key] = value
}

func (s *service) Keys() []PayloadValue {
	return []PayloadValue{}
}

func (s *service) Del(key string) (bool, error) {
	return true, nil
}

func (s *service) GetKeys() map[string]PayloadValue {
	return s.keys
}
