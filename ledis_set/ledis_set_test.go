package ledis_set

import (
	"github.com/duonghds/ledis/ledis_global"
	"testing"
)

func setupTest() SetService {
	globalService := ledis_global.InitGlobal()

	setService := NewService(globalService)
	return setService
}

func TestSet(t *testing.T) {
	setService := setupTest()
	setService.SAdd("s1", []string{"1", "2", "3"})
	card := setService.SCard("s1")
	if card != 3 {
		t.Errorf("Test set card failed got: %d, want: %d", card, 3)
	}
	members := setService.SMembers("s1")
	if len(members) != 3 {
		t.Errorf("Test set card failed got: %d, want: %d", len(members), 3)
	}

	setService.SRem("s1", []string{"2", "3"})
	card = setService.SCard("s1")
	if card != 1 {
		t.Errorf("Test set card failed got: %d, want: %d", card, 1)
	}
	members = setService.SMembers("s1")
	if len(members) != 1 {
		t.Errorf("Test set card failed got: %d, want: %d", len(members), 1)
	}

	setService.SAdd("s1", []string{"1", "1", "5", "6", "7"})
	card = setService.SCard("s1")
	if card != 4 {
		t.Errorf("Test set card failed got: %d, want: %d", card, 4)
	}
	members = setService.SMembers("s1")
	if len(members) != 4 {
		t.Errorf("Test set card failed got: %d, want: %d", len(members), 4)
	}

	setService.SAdd("s2", []string{"6", "7", "9", "0"})
	setService.SAdd("s3", []string{"12", "15", "17"})
	setService.SInter([]string{"s1", "s2", "s3"})
	card = setService.SCard("s1")
	if card != 9 {
		t.Errorf("Test set card failed got: %d, want: %d", card, 9)
	}
}
