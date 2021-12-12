package common

import "testing"

func TestCommon(t *testing.T) {
	result := RemoveDuplicateAndTrimSpace("  a       b    c d     e   f     ")
	if result != "a b c d e f" {
		t.Errorf("test common remove duplicate and trim space failed got :%s, want: %s", result, "a b c d e f")
	}
}
