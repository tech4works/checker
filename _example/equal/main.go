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
	fmt.Println("Equals results:")
	fmt.Println(checker.Equals("test", "test"))                                     // Should return true.
	fmt.Println(checker.Equals(map[string]int{"one": 1}, map[string]int{"one": 1})) // Should return true.
	fmt.Println(checker.Equals(student{name: "test"}, student{name: "test"}))       // Should return true.
	fmt.Println(checker.Equals(student{name: "test"}, student{name: "test2"}))      // Should return false.
	fmt.Println(checker.Equals("GoLang", "Java"))                                   // Should return false.

	fmt.Println("EqualsIgnoreCase results:")
	fmt.Println(checker.EqualsIgnoreCase("GoLang", "golang")) // Should return true.
	fmt.Println(checker.EqualsIgnoreCase("Hello", "hello"))   // Should return true.

	fmt.Println("AllEquals results:")
	fmt.Println(checker.AllEquals(10, 10, 10))                // Should return true.
	fmt.Println(checker.AllEquals("Hello", "Hello", "World")) // Should return false.
}
