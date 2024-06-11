package main

import (
	"fmt"
	"github.com/play-innovat/go-comparator"
)

type student struct {
	Name string
	Age  int
}

func main() {
	sliceData := []int{1, 2, 3, 4, 5}
	mapData := map[int]string{1: "One", 2: "Two", 3: "Three"}
	stringData := "Hello Go Programming"
	studentData := student{Name: "John", Age: 20}
	studentsData := []student{
		{Name: "John", Age: 20},
		{Name: "Doe", Age: 22},
		{Name: "Randy", Age: 25},
	}

	fmt.Println("Contains results:")
	fmt.Println(comparator.Contains(sliceData, 3))              // Should return true.
	fmt.Println(comparator.Contains(mapData, "Two"))            // Should return true.
	fmt.Println(comparator.Contains(stringData, "Go"))          // Should return true.
	fmt.Println(comparator.Contains(studentsData, studentData)) // Should return true.
	fmt.Println(comparator.Contains(studentData, "John"))       // Should return true.

	fmt.Println("NotContains results:")
	fmt.Println(comparator.NotContains(sliceData, 6))         // Should return true.
	fmt.Println(comparator.NotContains(mapData, "Four"))      // Should return true.
	fmt.Println(comparator.NotContains(stringData, "Python")) // Should return true.
	fmt.Println(comparator.NotContains(studentData, "Peter")) // Should return true.

	fmt.Println("ContainsIgnoreCase results:")
	fmt.Println(comparator.ContainsIgnoreCase(stringData, "hello")) // Should return true.
	fmt.Println(comparator.ContainsIgnoreCase(stringData, "go"))    // Should return true.

	fmt.Println("NotContainsIgnoreCase results:")
	fmt.Println(comparator.NotContainsIgnoreCase(stringData, "bye"))    // Should return true.
	fmt.Println(comparator.NotContainsIgnoreCase(stringData, "python")) // Should return true.

	fmt.Println("ContainsKey results:")
	fmt.Println(comparator.ContainsKey(mapData, 2))          // Should return true.
	fmt.Println(comparator.ContainsKey(studentData, "Name")) // Should return true.

	fmt.Println("NotContainsKey results:")
	fmt.Println(comparator.NotContainsKey(mapData, 4))             // Should return true.
	fmt.Println(comparator.NotContainsKey(studentData, "Address")) // Should return true.

	fmt.Println("ContainsOnSlice results:")
	fmt.Println(comparator.ContainsOnSlice(studentsData, func(index student) bool {
		return index.Name == "John"
	})) //Should return true

	fmt.Println("NotContainsOnSlice results:")
	fmt.Println(comparator.NotContainsOnSlice(studentsData, func(index student) bool {
		return index.Name == "Mike"
	})) //Should return true
}
