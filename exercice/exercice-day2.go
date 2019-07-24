package exercice

import (
	"fmt"
	_ "os"
)

func main() {
	set := map[string]bool{
		"abc": true,
		"cde": true,
	}

	setWithStruct := map[string]struct{}{
		"abc": {},
		"cde": {},
	}

	checkPresent(set, "abc")
	checkPresentStruct(setWithStruct, "abc")
	checkPresent(set, "yoyoyo")
	set["yoyoyo"] = true
	fmt.Printf("%v\n", set["yoyoyo"])
	checkPresent(set, "yoyoyo")
}

func checkPresent(set map[string]bool, s string) {
	value, present := set[s]
	if present {
		fmt.Printf("%s present, %v\n", s, value)
	} else {
		fmt.Printf("%s not present\n", s)
	}
}

func checkPresentStruct(set map[string]struct{}, s string) {
	value, present := set[s]
	if present {
		fmt.Printf("%s present, %v\n", s, value)
	} else {
		fmt.Printf("%s not present\n", s)
	}
}
