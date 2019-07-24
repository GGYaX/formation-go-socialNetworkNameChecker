package main

import (
	"fmt"
	_ "os"
)

type Tree struct {
	left  *Tree
	v     string
	right *Tree
}

func main() {
	bar := Tree{v: "bar"}
	foo := Tree{v: "foo", left: &bar}

	fmt.Printf("%v\n", foo)
}
