package main

import (
	"fmt"
	"github.com/tech4works/checker"
)

func main() {
	fmt.Println("IsGreaterThan results:")
	fmt.Println(checker.IsGreaterThan(5, 4))   // Outputs: true
	fmt.Println(checker.IsGreaterThan(-2, -3)) // Outputs: true
	fmt.Println(checker.IsGreaterThan(0, 4))   // Outputs: false

	fmt.Println("IsGreaterThanOrEqual results:")
	fmt.Println(checker.IsGreaterThanOrEqual(5, 5)) // Outputs: true
	fmt.Println(checker.IsGreaterThanOrEqual(5, 4)) // Outputs: true
	fmt.Println(checker.IsGreaterThanOrEqual(0, 4)) // Outputs: false

	fmt.Println("IsLessThan results:")
	fmt.Println(checker.IsLessThan(4, 5)) // Outputs: true
	fmt.Println(checker.IsLessThan(0, 4)) // Outputs: true

	fmt.Println("IsLessThanOrEqual results:")
	fmt.Println(checker.IsLessThanOrEqual(4, 4)) // Outputs: true
	fmt.Println(checker.IsLessThanOrEqual(4, 5)) // Outputs: true
	fmt.Println(checker.IsLessThanOrEqual(0, 4)) // Outputs: true

	fmt.Println("IsLengthEquals results:")
	fmt.Println(checker.IsLengthEquals("test", "test"))                         // Outputs: true
	fmt.Println(checker.IsLengthEquals(struct{ Name string }{Name: "test"}, 1)) // Outputs: true
	fmt.Println(checker.IsLengthEquals([]int{1, 2, 3, 4}, 4))                   // Outputs: true
	fmt.Println(checker.IsLengthEquals("test", 4))                              // Outputs: false

	fmt.Println("IsLengthGreaterThan results:")
	fmt.Println(checker.IsLengthGreaterThan("test", 2))           // Outputs: true
	fmt.Println(checker.IsLengthGreaterThan([]int{1, 2}, "test")) // Outputs: false

	fmt.Println("IsLengthGreaterThanOrEqual results:")
	fmt.Println(checker.IsLengthGreaterThanOrEqual("test", 2))                 // Outputs: true
	fmt.Println(checker.IsLengthGreaterThanOrEqual("test", []int{1, 2, 3, 4})) // Outputs: true

	fmt.Println("IsLengthLessThan results:")
	fmt.Println(checker.IsLengthLessThan("test", "testexample")) // Outputs: true
	fmt.Println(checker.IsLengthLessThan("test", 12))            // Outputs: true

	fmt.Println("IsLengthLessThanOrEqual results:")
	fmt.Println(checker.IsLengthLessThanOrEqual("test", "testexample")) // Outputs: true
	fmt.Println(checker.IsLengthLessThanOrEqual("test", 12))            // Outputs: true
}
