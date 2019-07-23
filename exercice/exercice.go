package exercice

import (
	"fmt"
	_ "os"
	// "unicode/utf8"
)

type Person struct {
	name string
}

const abc int = 0
const cde int = 0

const (
	a int = 32 << (^uint(0) >> 63)
)

type Celsius int

func main() {
	var (
		hello     = "aasaaaa"
		noAddress = "asdfsdf"
		boolean   = true
	)
	hello = "tototo"

	fmt.Printf("%d, %d, %d\n\n", a)

	// this code doesnt work since the pointer pints to a empty adresse, it cannot be modified as a value
	// var intPtr *int
	// // *intPtr = 123
	// fmt.Printf("%v %T\n", intPtr, *intPtr)

	// commentaire
	fmt.Println(&hello)
	fmt.Println(&noAddress)
	fmt.Println(boolean)
	fmt.Println(hello)

	var s string = "测"
	// fmt.Println(utf8.RuneLen(s))
	fmt.Printf("%s\n", s[1])

	var s2 = "loyoyohi"
	s3 := s2[0:2]
	fmt.Println(s3)

	// _, yoyo := foo()

	ss := "阿斯顿😋飞"
	// ss := "helloworld"
	for i, c := range ss {
		fmt.Printf("index: %d, %c\n", i, c)
	}

	switch switchTest() {
	default:
		fmt.Println("1")
	case 0:
		fmt.Println("0")
	}

	fmt.Println(signumSwitch(1, 2))
	f := increment()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func foo() (int, int) {
	return 0, 0
}

func switchTest() int {
	return 0
}

func signum(a int, b int) int {
	if a > b {
		return -1
	}
	if a < b {
		return 1
	}
	return 0
}

func signumSwitch(a int, b int) (signum int) {
	switch {
	case a > b:
		signum = -1
	case a < b:
		signum = 1
	default:
		signum = 0
	}
	return
}

func functionExample() func(s string) int {
	f := func(s string) int {
		return len(s)
	}
	return f
}

func isEven(a uint) bool {
	return a != 0 && !isOdd(a-1)
}

func isOdd(a uint) bool {
	return a == 0 || isEven(a-1)
}

func increment() func() int {
	var toto int = 0

	f := func() int {
		toto++
		return toto
	}
	return f
}
