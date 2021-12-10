package ledis_string

import (
	"github.com/duonghds/ledis/common"
	"github.com/duonghds/ledis/ledis_global"
	"testing"
)

func setupTest() StringService {
	globalService := ledis_global.InitGlobal()

	stringService := NewService(globalService)
	return stringService
}

func TestSetAndGet(t *testing.T) {
	stringService := setupTest()
	stringService.Set("key1", "value1")
	stringService.Set("key2", "value2")
	stringService.Set("key2", "value3")
	stringService.Set("key3", "value4 and 5")

	value1, err := stringService.Get("key1")
	if value1 != "value1" || err != nil {
		t.Errorf("Test Get and Set string failed, got: %s, want: %s", value1, "value1")
	}
	value2, err := stringService.Get("key2")
	if value2 != "value3" || err != nil {
		t.Errorf("Test Get and Set string failed, got: %s, want: %s", value2, "value3")
	}
	value3, err := stringService.Get("key3")
	if value3 != "value4 and 5" || err != nil {
		t.Errorf("Test Get and Set string failed, got: %s, want: %s", value3, "value4 and 5")
	}
	_, err = stringService.Get("key4")
	if err == nil || err.Error() != common.KeyNotFound {
		t.Errorf("Test Get and Set string failed, got: %s, want: %s", err.Error(), common.KeyNotFound)
	}
}
