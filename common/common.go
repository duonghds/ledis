package common

import (
	"regexp"
	"strings"
)

//RemoveDuplicateAndTrimSpace ...
func RemoveDuplicateAndTrimSpace(str string) string {
	str = strings.TrimSpace(str)
	space := regexp.MustCompile(`\s+`)
	str = space.ReplaceAllString(str, " ")
	return str
}
