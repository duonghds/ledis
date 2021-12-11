package ledis_global

import (
	"testing"
)

func setupTest() GlobalService {
	globalService := InitGlobal()

	return globalService
}

func TestSnapshot(t *testing.T) {
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

	globalService.Save()

	globalService.AddKey("key4", PayloadValue{
		Type: "type",
	})
	globalService.AddKey("key5", PayloadValue{
		Type: "type",
	})

	keyBfSnapshot := globalService.Keys()
	if len(keyBfSnapshot) != 5 {
		t.Errorf("Snapshot fail, got: %d, want: %d", len(keyBfSnapshot), 5)
	}

	globalService.Restore()
	keyAtSnapshot := globalService.Keys()
	if len(keyAtSnapshot) != 3 {
		t.Errorf("Snapshot fail, got: %d, want: %d", len(keyAtSnapshot), 3)
	}
}
