//	MIT License
//
//	Copyright (c) 2024 Gabriel Henrique Cataldo Moskorz
//
//	Permission is hereby granted, free of charge, to any person obtaining a copy
//	of this software and associated documentation files (the "Software"), to deal
//	in the Software without restriction, including without limitation the rights
//	to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//	copies of the Software, and to permit persons to whom the Software is
//	furnished to do so, subject to the following conditions:
//
//	The above copyright notice and this permission notice shall be included in all
//	copies or substantial portions of the Software.
//
//	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//	IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//	AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//	LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//	OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//	SOFTWARE.

package checker

import (
	"fmt"
	"reflect"
	"strings"
)

// Equals checks whether two parameters a and b are profoundly equal.
// This function internally uses the DeepEqual function from the reflect package for comparison,
// which can correctly compare different data types, including strings, ints, floats, bools, slices, maps, structs, etc.
//
// Example usage:
//
//	fmt.Println(Equals("test", "test"))  // Outputs: true
//
//	m1 := map[string]int{"one": 1}
//	m2 := map[string]int{"one": 1}
//	fmt.Println(Equals(m1, m2))   // Outputs: true
//
//	s1 := struct{Name string}{Name: "test"}
//	s2 := struct{Name string}{Name: "test"}
//	fmt.Println(Equals(s1, s2))   // Outputs: true
//
// Returns true if a and b are deeply equal, false otherwise.
func Equals(a, b any) bool {
	reflectValueA := reflect.ValueOf(a)
	if (reflectValueA.Kind() == reflect.Ptr || reflectValueA.Kind() == reflect.Interface) && !reflectValueA.IsNil() {
		return Equals(reflectValueA.Elem().Interface(), b)
	}

	reflectValueB := reflect.ValueOf(b)
	if (reflectValueB.Kind() == reflect.Ptr || reflectValueB.Kind() == reflect.Interface) && !reflectValueB.IsNil() {
		return Equals(a, reflectValueB.Elem().Interface())
	}

	return reflect.DeepEqual(a, b)
}

// EqualsIgnoreCase compares two values and returns true if they are equal, ignoring case.
//
// This function takes two parameters, `a` and `b`, of any type. It first validates the input
// by calling the `validateEqualsIgnoreCaseParams` function. Then, it uses reflection to determine
// the kind of the values `a` and `b`. If either `a` or `b` is a pointer or interface, the function
// recursively invokes itself with the dereferenced values. Finally, if both `a` and `b` are strings,
// the function compares their lowercase versions and returns true if they are equal.
//
// Example Usage:
//
//	fmt.Println(EqualsIgnoreCase("GoLang", "golang")) // true
//	fmt.Println(EqualsIgnoreCase("Hello", "hello"))   // true
//	fmt.Println(EqualsIgnoreCase("GoLang", "Java"))   // false
func EqualsIgnoreCase(a, b any) bool {
	validateEqualsIgnoreCaseParams(a)

	reflectValueA := reflect.ValueOf(a)
	if reflectValueA.Kind() == reflect.Ptr || reflectValueA.Kind() == reflect.Interface {
		return EqualsIgnoreCase(reflectValueA.Elem().Interface(), b)
	}

	reflectValueB := reflect.ValueOf(b)
	if reflectValueB.Kind() == reflect.Ptr || reflectValueB.Kind() == reflect.Interface {
		return EqualsIgnoreCase(a, reflectValueB.Elem().Interface())
	}

	return reflectValueB.Kind() == reflect.String &&
		strings.ToLower(reflectValueA.String()) == strings.ToLower(reflectValueB.String())
}

// AllEquals checks whether all parameters in the variadic arguments b are
// profoundly equal to the parameter a. This function internally calls
// the IsNotEqualTo function for comparison, which returns true if the parameters
// are not deeply equal, and false otherwise.
//
// Example usage:
//
//	a := 10
//	b := 10
//	c := 10
//	fmt.Println(AllEquals(a, b, c))  // Outputs: true
//
//	x := "Hello"
//	y := "Hello"
//	z := "World"
//	fmt.Println(AllEquals(x, y, z))  // Outputs: false
//
// Returns true if all parameters are deeply equal to a, false otherwise.
func AllEquals(a, b any, c ...any) bool {
	c = append([]any{a, b}, c...)
	for _, v1 := range c {
		for _, v2 := range c {
			if !Equals(v1, v2) {
				return false
			}
		}
	}
	return true
}

// validateEqualsIgnoreCaseParams validates the input value to ensure that it is not nil and is either a string or a pointer.
// If the value is nil, it panics with an error message "A is nil".
// If the value is not a string or a pointer, it panics with an error message "Unsupported type: {type}".
// The function uses reflection to determine the kind of the value and perform the necessary checks.
func validateEqualsIgnoreCaseParams(a any) {
	reflectValueA := reflect.ValueOf(a)

	if IsNil(a) {
		panic("A is nil")
	} else if reflectValueA.Kind() != reflect.String && reflectValueA.Kind() != reflect.Ptr {
		panic(fmt.Sprintf("Unsupported type: %s", reflectValueA.Kind().String()))
	}
}
