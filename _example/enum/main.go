package main

import (
	"fmt"
	"github.com/tech4works/checker"
)

type MyEnum int

const (
	NUMBER0 MyEnum = 0
	NUMBER1 MyEnum = 1
	NUMBER2 MyEnum = 2
)

func (e MyEnum) IsEnumValid() bool {
	switch e {
	case NUMBER0, NUMBER1, NUMBER2:
		return true
	default:
		return false
	}
}

func main() {
	a := NUMBER2
	b := 3

	fmt.Println(checker.IsEnumValid(a)) // Should print true
	fmt.Println(checker.IsEnumValid(b)) // Should print false
}
