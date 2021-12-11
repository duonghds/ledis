package ledis_list

import (
	"fmt"
	"github.com/duonghds/ledis/ledis_global"
	"testing"
)

func setupTest() ListService {
	globalService := ledis_global.InitGlobal()

	listService := NewService(globalService)
	return listService
}

func TestList(t *testing.T) {
	listService := setupTest()
	key := "key1"
	count, err := listService.RPush(key, []string{"1", "2", "3", "4", "5"})
	if err != nil || count != 5 {
		t.Errorf("Test list rpush failed got: %d, want: %d", count, 5)
	}

	lPopValue, err := listService.LPop(key)
	if err != nil || lPopValue != "1" {
		t.Errorf("Test list lpop failed got: %s, want: %s", lPopValue, "1")
	}

	rPopValue, err := listService.RPop(key)
	if err != nil || rPopValue != "5" {
		t.Errorf("Test list rpop failed got: %s, want: %s", rPopValue, "5")
	}

	llen := listService.LLen(key)
	if llen != 3 {
		t.Errorf("Test list llen failed got: %d, want: %d", llen, 3)
	}

	listValue, err := listService.LRange(key, 0, 999)
	if err != nil || len(listValue) != 3 {
		t.Errorf("Test list lrange failed got: %d, want: %d", len(listValue), 3)
	}
	realListValue := fmt.Sprintf("%+v", listValue)
	expectedListValue := fmt.Sprintf("%+v", []string{"2", "3", "4"})
	if realListValue != expectedListValue {
		t.Errorf("Test list lrange failed got: %s, want: %s", realListValue, expectedListValue)
	}
	count, err = listService.RPush(key, []string{"6"})
	if err != nil || count != 1 {
		t.Errorf("Test list rpush failed got: %d, want: %d", count, 1)
	}

	listValue, err = listService.LRange(key, 0, 999)
	if err != nil || len(listValue) != 4 {
		t.Errorf("Test list lrange failed got: %d, want: %d", len(listValue), 4)
	}
	realListValue = fmt.Sprintf("%+v", listValue)
	expectedListValue = fmt.Sprintf("%+v", []string{"2", "3", "4", "6"})
	if realListValue != expectedListValue {
		t.Errorf("Test list lrange failed got: %s, want: %s", realListValue, expectedListValue)
	}
}
