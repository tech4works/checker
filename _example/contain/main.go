package main

import (
	"fmt"
	"github.com/tech4works/go-checker"
)

type student struct {
	name string
	age  int
}

func main() {
	sliceData := []int{1, 2, 3, 4, 5}
	mapData := map[int]string{1: "One", 2: "Two", 3: "Three"}
	stringData := "Hello Go Programming"
	studentData := student{name: "John", age: 20}
	studentsData := []student{
		{name: "John", age: 20},
		{name: "Doe", age: 22},
		{name: "Randy", age: 25},
	}

	fmt.Println("Contains results:")
	fmt.Println(checker.Contains(sliceData, 3))              // Should return true.
	fmt.Println(checker.Contains(mapData, "Two"))            // Should return true.
	fmt.Println(checker.Contains(stringData, "Go"))          // Should return true.
	fmt.Println(checker.Contains(studentsData, studentData)) // Should return true.
	fmt.Println(checker.Contains(studentData, "John"))       // Should return true.

	fmt.Println("NotContains results:")
	fmt.Println(checker.NotContains(sliceData, 6))         // Should return true.
	fmt.Println(checker.NotContains(mapData, "Four"))      // Should return true.
	fmt.Println(checker.NotContains(stringData, "Python")) // Should return true.
	fmt.Println(checker.NotContains(studentData, "Peter")) // Should return true.

	fmt.Println("ContainsIgnoreCase results:")
	fmt.Println(checker.ContainsIgnoreCase(stringData, "hello")) // Should return true.
	fmt.Println(checker.ContainsIgnoreCase(stringData, "go"))    // Should return true.

	fmt.Println("NotContainsIgnoreCase results:")
	fmt.Println(checker.NotContainsIgnoreCase(stringData, "bye"))    // Should return true.
	fmt.Println(checker.NotContainsIgnoreCase(stringData, "python")) // Should return true.

	fmt.Println("ContainsKey results:")
	fmt.Println(checker.ContainsKey(mapData, 2))          // Should return true.
	fmt.Println(checker.ContainsKey(studentData, "Name")) // Should return true.

	fmt.Println("NotContainsKey results:")
	fmt.Println(checker.NotContainsKey(mapData, 4))             // Should return true.
	fmt.Println(checker.NotContainsKey(studentData, "Address")) // Should return true.

	fmt.Println("ContainsOnSlice results:")
	fmt.Println(checker.ContainsOnSlice(studentsData, func(index student) bool {
		return index.name == "John"
	})) //Should return true

	fmt.Println("NotContainsOnSlice results:")
	fmt.Println(checker.NotContainsOnSlice(studentsData, func(index student) bool {
		return index.name == "Mike"
	})) //Should return true
}
