package helper

import (
	"regexp"
	"strings"
)

func RemoveNonAlphaNumSpace(s string) string {
	s = strings.TrimSpace(s)
	reg, _ := regexp.Compile("[^a-zA-Z0-9 ]+")
	return reg.ReplaceAllString(s, "")
}
