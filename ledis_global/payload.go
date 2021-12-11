package ledis_global

import "time"

type PayloadValue struct {
	Type       string
	Value      interface{}
	ExpireTime time.Time
}
