package regex

import (
	"regexp"
	"strings"
)

var (
	re_url = regexp.MustCompile(`(?m)http[s]?:\/\/[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b[-a-zA-Z0-9()@:%_\+.~#?&\/\/=]*`)
)

func RemoveURLFromString(str string) string {
	for _, match := range re_url.FindAllString(str, -1) {
		str = strings.ReplaceAll(str, match, "")
	}
	return strings.TrimSpace(str)
}

func MatchURLFromString(str string) string {
	for _, match := range re_url.FindAllString(str, -1) {
		return strings.TrimSpace(match)
	}
	return ""
}
