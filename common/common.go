package common

import (
	"regexp"
	"strings"
)

const (
	ErrKeyNotFound = "ERROR: Key not found"
	ErrKeyEmpty    = "ERROR: Key empty"

	ErrStopLessThanStart = "ERROR: Stop less than start"

	ErrUnknownCommand  = "ERROR: Unknown command"
	ErrNotEnoughParams = "ERROR: Not enough params"
	ErrWrongType       = "ERROR: Wrong type"
)

//RemoveDuplicateAndTrimSpace ...
func RemoveDuplicateAndTrimSpace(str string) string {
	str = strings.TrimSpace(str)
	space := regexp.MustCompile(`\s+`)
	str = space.ReplaceAllString(str, " ")
	return str
}
