package twitter_old

import (
	_ "os"
	"regexp"
	"unicode/utf8"
)

var regexNoIllegalPattern = regexp.MustCompile("(?i)twitter")
var onlyContainsLegalRunesRegex = regexp.MustCompile("^[A-Za-z0-9]+$")

func IsShortEnough(s string) bool {
	//runes := []rune(s)
	//return len(runes) <= 15
	return utf8.RuneCountInString(s) <= 15
}

func IsLongEnough(s string) bool {
	//runes := []rune(s)
	//return len(runes) >= 1
	return utf8.RuneCountInString(s) >= 1
}

func ContainsNoIllegalPattern(s string) (result bool) { // no (?i)twitter, using regex
	result = regexNoIllegalPattern.MatchString(s)
	return
}

func OnlyContainsLegalRunes(s string) (result bool) { // ^[0-9A-Z]
	result = onlyContainsLegalRunesRegex.MatchString(s)
	return
}

func Validate(s string) []string {
	result := make([]string, 0)
	if !IsLongEnough(s) {
		result = append(result, "Is not long enough")
	}
	if !IsShortEnough(s) {
		result = append(result, "Is not short enough")
	}
	if ContainsNoIllegalPattern(s) {
		result = append(result, "ContainsNoIllegalPattern")
	}
	if !OnlyContainsLegalRunes(s) {
		result = append(result, "Not OnlyContainsLegalRunes")
	}
	return result
}
