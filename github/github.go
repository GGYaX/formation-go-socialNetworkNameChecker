package github

import (
	"fmt"
	"github.com/GGYaX/namecheck/namecheck"
	"net/http"
	_ "os"
	"regexp"
	"unicode/utf8"
)

var regexNoIllegalPattern = regexp.MustCompile("(?i)twitter")
var onlyContainsLegalRunesRegex = regexp.MustCompile("^[A-Za-z0-9]+$")

type Github struct{}

var twitter *Github

func (t *Github) IsShortEnough(s string) bool {
	//runes := []rune(s)
	//return len(runes) <= 15
	return utf8.RuneCountInString(s) <= 38
}

func (t *Github) IsLongEnough(s string) bool {
	//runes := []rune(s)
	//return len(runes) >= 1
	return utf8.RuneCountInString(s) >= 1
}

func (t *Github) ContainsNoIllegalPattern(s string) (result bool) { // no (?i)twitter, using regex
	result = regexNoIllegalPattern.MatchString(s)
	return
}

func (t *Github) OnlyContainsLegalRunes(s string) (result bool) { // ^[0-9A-Z]
	result = onlyContainsLegalRunesRegex.MatchString(s)
	return
}

func (t *Github) Validate(s string) []string {
	result := make([]string, 0)
	if !twitter.IsLongEnough(s) {
		result = append(result, "Is not long enough")
	}
	if !twitter.IsShortEnough(s) {
		result = append(result, "Is not short enough")
	}
	if twitter.ContainsNoIllegalPattern(s) {
		result = append(result, "ContainsNoIllegalPattern")
	}
	if !twitter.OnlyContainsLegalRunes(s) {
		result = append(result, "Not OnlyContainsLegalRunes")
	}
	return result
}

func (t *Github) IsAvailable(username string) (bool, namecheck.IsAvailablerError) {
	// need to validate username
	resp, err := http.Get(fmt.Sprintf("https://www.github.com/%s", username))
	if err != nil {
		return false, namecheck.IsAvailablerError{E: err}
	}
	if resp.StatusCode == http.StatusNotFound {
		return false, namecheck.IsAvailablerError{}
	}
	return true, namecheck.IsAvailablerError{}
}
