package main

import (
	"fmt"
	"github.com/GGYaX/namecheck/twitter"
	_ "os"
)

func main() {
	containsNoIllegalPatternInput := "asdfasdfadsf"
	containsNoIllegalPatternResult := twitter.ContainsNoIllegalPattern(containsNoIllegalPatternInput)
	shortEnoughInput := "阿迪舒服撒地方阿斯顿飞"
	shortEnough := twitter.IsShortEnough(shortEnoughInput)
	longEnoughInput := "twitter"
	longEnough := twitter.IsLongEnough(longEnoughInput)
	containsLegalRunesInput := "^&*"
	containsLegalRunes := twitter.OnlyContainsLegalRunes(containsLegalRunesInput)

	fmt.Printf("result containsNoIllegalPatternResult of input %s is %v\n", containsNoIllegalPatternInput, containsNoIllegalPatternResult)
	fmt.Printf("result shortEnough of input %s is %v\n", shortEnoughInput, shortEnough)
	fmt.Printf("result longEnough of input %s is %v\n", longEnoughInput, longEnough)
	fmt.Printf("result containsLegalRunes of input %s is %v\n", containsLegalRunesInput, containsLegalRunes)
}
