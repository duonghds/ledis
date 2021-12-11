package ledis_global

import (
	"testing"
)

func TestExpiration(t *testing.T) {
	globalService := setupTest()
	globalService.AddKey("key1", PayloadValue{
		Type: "type",
	})
	globalService.AddKey("key2", PayloadValue{
		Type: "type",
	})
	globalService.AddKey("key3", PayloadValue{
		Type: "type",
	})
	keys := globalService.Keys()
	if len(keys) != 3 {
		t.Errorf("check keys fail, got: %d, want: %d", len(keys), 3)
	}
	globalService.Del("key1")
	keys = globalService.Keys()
	if len(keys) != 2 {
		t.Errorf("delete keys fail, got: %d, want: %d", len(keys), 2)
	}
	globalService.Expire("key2", 10)
	ttl := globalService.TTL("key2")
	if ttl != 10 {
		t.Errorf("set expire keys fail, got: %d, want: %d", ttl, 10)
	}
	globalService.FlushDB()
	keys = globalService.Keys()
	if len(keys) != 0 {
		t.Errorf("delete keys fail, got: %d, want: %d", len(keys), 0)
	}
}
