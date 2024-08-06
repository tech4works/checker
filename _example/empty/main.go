package main

import (
	"fmt"
	"github.com/tech4works/checker"
)

type student struct {
	name string
	age  int
}

func main() {
	var nilPointer *int

	fmt.Println("IsNil results:")
	fmt.Println(checker.IsNil(nil))            // Should return true.
	fmt.Println(checker.IsNil(3))              // Should return false.
	fmt.Println(checker.IsNil(student{}))      // Should return false.
	fmt.Println(checker.IsNil(&map[int]any{})) // Should return false.
	fmt.Println(checker.IsNil(nilPointer))     // Should return true.

	fmt.Println("NonNil results:")
	fmt.Println(checker.NonNil(nil))               // Should return false.
	fmt.Println(checker.NonNil(nilPointer))        // Should return false.
	fmt.Println(checker.NonNil(5))                 // Should return true.
	fmt.Println(checker.NonNil(&map[string]any{})) // Should return true.
	fmt.Println(checker.NonNil(student{}))         // Should return true.
	fmt.Println(checker.NonNil(false))             // Should return true.

	fmt.Println("IsEmpty results:")
	fmt.Println(checker.IsEmpty(" "))                   // Should return true.
	fmt.Println(checker.IsEmpty([]int{}))               // Should return true.
	fmt.Println(checker.IsEmpty(student{}))             // Should return true.
	fmt.Println(checker.IsEmpty(map[string]any{}))      // Should return true.
	fmt.Println(checker.IsEmpty(0))                     // Should return true.
	fmt.Println(checker.IsEmpty(student{name: "John"})) // Should return false.
	fmt.Println(checker.IsEmpty("Go"))                  // Should return false.
	fmt.Println(checker.IsEmpty(1))                     // Should return false.

	fmt.Println("IsNotEmpty results:")
	fmt.Println(checker.IsNotEmpty(" "))                            // Should return false.
	fmt.Println(checker.IsNotEmpty("Go"))                           // Should return true.
	fmt.Println(checker.IsNotEmpty(0))                              // Should return false.
	fmt.Println(checker.IsNotEmpty(1))                              // Should return true.
	fmt.Println(checker.IsNotEmpty(false))                          // Should return false.
	fmt.Println(checker.IsNotEmpty(true))                           // Should return true.
	fmt.Println(checker.IsNotEmpty(map[string]any{}))               // Should return false.
	fmt.Println(checker.IsNotEmpty(map[string]any{"test": "asa"}))  // Should return true.
	fmt.Println(checker.IsNotEmpty(student{}))                      // Should return false.
	fmt.Println(checker.IsNotEmpty(student{name: "John", age: 23})) // Should return true.

	fmt.Println("AllNil results:")
	fmt.Println(checker.AllNil(nil, nil, nilPointer))                         // Should return true.
	fmt.Println(checker.AllNil(nil, 0, "hello", student{}, map[string]any{})) // Should return false.

	fmt.Println("AllNonNil results:")
	fmt.Println(checker.NoneNil(3, 22, student{}, "test", 21.23, true, map[string]any{})) // Should return true.
	fmt.Println(checker.NoneNil(3, nilPointer, nil, "test"))                              // Should return false.

	fmt.Println("AllEmpty results:")
	fmt.Println(checker.AllEmpty(" ", "", 0, []int{}, nil, false, student{}, map[string]any{}, 0.0)) // Should return true.
	fmt.Println(checker.AllEmpty("hello", 10, []int{1, 2, 3}, true, map[string]any{"test": 1}))      // Should return false.

	fmt.Println("AllNotEmpty results:")
	fmt.Println(checker.NoneEmpty("hello", 10, []int{1, 2, 3}, true, map[string]any{"test": 1})) // Should return true.
	fmt.Println(checker.NoneEmpty(" ", "", 0, []int{}, nil, false, student{}, map[string]any{})) // Should return false.
}
