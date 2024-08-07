//	MIT License
//
//	Copyright (c) 2024 Tech4Works
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

// Contains checks if the provided value 'b' is contained within the value 'a'.
// It uses reflection to determine the type of 'a' and performs appropriate checks
// for slice, array, map, struct, and string types. For a slice or an array,
// it iterates over each element and uses reflect.DeepEqual to compare with 'b'.
//
// Example usage:
//
//	sliceA := []int{1, 2, 3, 4}
//	fmt.Println(Contains(sliceA, 3))   // true
//
//	mapA := map[string]int{"one": 1, "two": 2, "three": 3}
//	fmt.Println(Contains(mapA, 1)) // true
//
//	type structA struct {
//	    field1 int
//	    field2 string
//	}
//	s := structA{field1: 1, field2: "value"}
//	fmt.Println(Contains(s, "value"))  // true
//
//	strA := "Hello World"
//	fmt.Println(Contains(strA, "World"))  // true
func Contains(a, b any) bool {
	validateContainsParams(a)

	reflectValueA := reflect.ValueOf(a)
	if reflectValueA.Kind() == reflect.Ptr || reflectValueA.Kind() == reflect.Interface {
		return Contains(reflectValueA.Elem().Interface(), b)
	}

	reflectValueB := reflect.ValueOf(b)
	if reflectValueB.Kind() == reflect.Ptr || reflectValueB.Kind() == reflect.Interface {
		return Contains(a, reflectValueB.Elem().Interface())
	}

	if reflectValueA.Kind() == reflect.Slice || reflectValueA.Kind() == reflect.Array {
		return containsValueInSlice(reflectValueA, b)
	} else if reflectValueA.Kind() == reflect.Map {
		return containsValueInMap(reflectValueA, b)
	} else if reflectValueA.Kind() == reflect.Struct {
		return containsValueInStruct(reflectValueA, b)
	}

	return reflectValueB.Kind() == reflect.String &&
		strings.Contains(reflectValueA.String(), reflectValueB.String())
}

// NotContains checks if the value 'b' is not contained within the value 'a'.
// This function uses the 'Contains' function for the check and
// returns the negation of the result.
// If 'a' or 'b' is a pointer or interface, it gets the underlying value before checking.
// If 'a' is a slice or array, it checks if 'b' is included in 'a'.
// If 'a' is a map, it checks if 'b' is a value in 'a'.
// If 'a' is a struct, it checks if 'b' is a value of any field in 'a'.
// If 'a' is a string, it handles 'b' as a string and checks if 'a' contains 'b' as a substring.
//
// Parameters:
//   - a: An interface value that 'b' is checked against. The value should be a slice, array, map, struct, or string.
//   - b: Any interface value to be checked for its existence in 'a'.
//
// Returns:
//   - bool: A boolean value indicating whether 'b' is not contained within 'a'.
//
// Example:
//
//	strA := "Hello World"
//	fmt.Println(NotContains(strA, "Moon"))  // true
//
//	sliceB := []int{1, 2, 3, 4}
//	fmt.Println(NotContains(sliceB, 5))   // true
//
//	mapC := map[string]int{"one":1, "two":2, "three":3}
//	fmt.Println(NotContains(mapC, 4)) // true
//
//	type structD struct {
//	    field1 string
//	    field2 int
//	}
//	s := structD{field1: "John", field2: 30}
//	fmt.Println(NotContains(s, "Jane"))  // true
func NotContains(a, b any) bool {
	return !Contains(a, b)
}

// ContainsIgnoreCase checks if the provided value 'b' is contained within the value 'a',
// ignoring case sensitivity. It uses reflection to determine the type of 'a' and performs
// appropriate checks for string types.
//
// Example usage:
//
//	strA := "Hello World"
//	fmt.Println(ContainsIgnoreCase(strA, "WORLD"))   // true
//	fmt.Println(ContainsIgnoreCase(strA, "goodbye")) // false
func ContainsIgnoreCase(a, b any) bool {
	validateContainsIgnoreCaseParams(a)

	reflectValueA := reflect.ValueOf(a)
	if reflectValueA.Kind() == reflect.Ptr || reflectValueA.Kind() == reflect.Interface {
		return ContainsIgnoreCase(reflectValueA.Elem().Interface(), b)
	}

	reflectValueB := reflect.ValueOf(b)
	if reflectValueB.Kind() == reflect.Ptr || reflectValueB.Kind() == reflect.Interface {
		return ContainsIgnoreCase(a, reflectValueB.Elem().Interface())
	}

	return reflectValueB.Kind() == reflect.String &&
		strings.Contains(strings.ToLower(reflectValueA.String()), strings.ToLower(reflectValueB.String()))
}

// NotContainsIgnoreCase determines whether the provided value 'b' is not contained within
// the value 'a', ignoring case sensitivity.
// This function is the complement of the ContainsIgnoreCase function, returning its negated result.
//
// Parameters:
//   - a: The value that will be checked to see if it contains 'b'.
//   - b: The value that we are looking for within 'a'.
//
// Returns:
//   - bool: A boolean value indicating whether 'b' is not found within 'a', disregarding case.
//
// Example:
//
//	strA := "Hello World"
//	fmt.Println(NotContainsIgnoreCase(strA, "WORLD")) // false
//	fmt.Println(NotContainsIgnoreCase(strA, "goodbye")) // true
//
// Please note:
// Depending on the types of 'a' and 'b', the ContainsIgnoreCase function invoked by this
// function may panic with an error stating "A is nil", or stating "Unsupported type",
// if 'a' is not a string or cannot be converted to one.
func NotContainsIgnoreCase(a, b any) bool {
	return !ContainsIgnoreCase(a, b)
}

// ContainsKey checks if the provided key 'key' is present in the value 'a'.
// It uses reflection to determine the type of 'a' and performs appropriate checks
// for struct and map types.
//
// Example Usage:
//
//	mapA := map[string]int{"one": 1, "two": 2, "three": 3}
//	fmt.Println(ContainsKey(mapA, "one"))  // true
//	fmt.Println(ContainsKey(mapA, "four"))  // false
//
//	type structA struct {
//	    field1 int
//	    field2 string
//	}
//	s := structA{field1: 1, field2: "value"}
//	fmt.Println(ContainsKey(s, "field1"))  // true
//	fmt.Println(ContainsKey(s, "field3"))  // false
func ContainsKey(a, key any) bool {
	validateContainsKeyParams(a)

	reflectValue := reflect.ValueOf(a)
	if reflectValue.Kind() == reflect.Ptr || reflectValue.Kind() == reflect.Interface {
		return ContainsKey(reflectValue.Elem().Interface(), key)
	}

	reflectKey := reflect.ValueOf(key)
	if reflectKey.Kind() == reflect.Ptr || reflectKey.Kind() == reflect.Interface {
		return ContainsKey(a, reflectKey.Elem().Interface())
	}

	if reflectValue.Kind() == reflect.Struct && reflectKey.Kind() == reflect.String {
		_, found := reflectValue.Type().FieldByName(reflectKey.String())
		return found
	}

	return reflectValue.Kind() == reflect.Map && reflectValue.MapIndex(reflectKey).IsValid()
}

// NotContainsKey determines whether a specific 'key' is not present in the given value 'a'.
// It uses the ContainsKey function to check if the key is present in the value, and returns the negation of its result.
//
// Parameters:
//   - a: Any interface value which is either a map or a struct type to be checked
//   - key: The key which presence is being checked
//
// Returns:
//   - bool: A boolean value indicating whether the provided key is not present in the value.
//
// Panic:
//   - If 'a' is nil, it panics with the message "A is nil"
//   - If 'a' is neither a map nor a struct, it panics with a formatted string indicating the unsupported type
//
// Example:
//
//	mapA := map[string]int{"one": 1, "two": 2, "three": 3}
//	fmt.Println(NotContainsKey(mapA, "one"))  // false
//	fmt.Println(NotContainsKey(mapA, "four"))  // true
//
//	structA := struct {
//	    field1 int
//	    field2 string
//	}
//	s := structA{field1: 1, field2: "value"}
//	fmt.Println(NotContainsKey(s, "field1")) // false
//	fmt.Println(NotContainsKey(s, "field3"))  // true
func NotContainsKey(a, key any) bool {
	return !ContainsKey(a, key)
}

// ContainsOnSlice checks if the provided value 'b' is found by the 'found' function when applied to the elements in the slice 'a'.
// It iterates over each element in 'a' and calls the 'found' function with the index and element as arguments.
// If 'found' returns true for any element, the function returns true.
// If 'found' returns false for all elements, the function returns false.
//
// Example Usage:
//
//	elements := []int{1, 2, 3, 4, 5}
//	fmt.Println(ContainsOnSlice(elements, func(index int, element int) bool {
//	    return element > 3
//	}))  // true
//
//	fmt.Println(ContainsOnSlice(elements, func(index int, element int) bool {
//	    return element == 10
//	}))  // false
func ContainsOnSlice[T any](a []T, found func(index int, element T) bool) bool {
	for index, element := range a {
		if found(index, element) {
			return true
		}
	}
	return false
}

// NotContainsOnSlice checks if the provided value 'b' is not present in the slice 'a'.
// It utilizes the ContainsOnSlice function to seek for the value and negates its result.
//
// Parameters:
//   - a: A slice of any type 'T' where we want to check the lack of 'b'.
//   - found: A higher order function that takes an index and an element of type 'T'
//     from the array and returns a boolean result.
//
// Returns:
//   - bool: A boolean value indicating whether the value is not present in the slice.
//
// Example:
//
//	elements := []int{1, 2, 3, 4, 5}
//	fmt.Println(NotContainsOnSlice(elements, func(index int, element int) bool {
//	    return element > 3
//	}))  // false
//
//	fmt.Println(NotContainsOnSlice(elements, func(index int, element int) bool {
//	    return element == 10
//	}))  // true
func NotContainsOnSlice[T any](a []T, found func(index int, element T) bool) bool {
	return !ContainsOnSlice(a, found)
}

// validateContainsParams validates the value 'a' to ensure it is a supported type for
// the Contains function. If 'a' is nil, it panics with the message "A is nil".
// If 'a' is not one of the supported types (slice, array, map, struct, string),
// it panics with a formatted string message indicating the unsupported type.
func validateContainsParams(a any) {
	reflectValueA := reflect.ValueOf(a)

	if IsNil(a) {
		panic("A is nil")
	} else if reflectValueA.Kind() != reflect.Slice && reflectValueA.Kind() != reflect.Array &&
		reflectValueA.Kind() != reflect.Map && reflectValueA.Kind() != reflect.Struct &&
		reflectValueA.Kind() != reflect.String && reflectValueA.Kind() != reflect.Ptr {
		panic(fmt.Sprintf("Unsupported type: %s", reflectValueA.Kind().String()))
	}
}

// validateContainsIgnoreCaseParams validates the value 'a' to ensure it is not nil and of type string.
// If 'a' is nil, it panics with the message "A is nil".
// If 'a' is not of type string, it panics with a formatted message indicating the unsupported type.
func validateContainsIgnoreCaseParams(a any) {
	reflectValueA := reflect.ValueOf(a)

	if IsNil(a) {
		panic("A is nil")
	} else if reflectValueA.Kind() != reflect.String && reflectValueA.Kind() != reflect.Ptr {
		panic(fmt.Sprintf("Unsupported type: %s", reflectValueA.Kind().String()))
	}
}

// validateContainsKeyParams validates the value 'a' to ensure it is a map or struct.
// If 'a' is nil, it panics with the message "A is nil".
// If 'a' is neither a map nor a struct, it panics with a formatted string
// indicating the unsupported type.
// This function uses reflection to determine the type of 'a'.
// It is used by the ContainsKey function to validateStringParams the input value before
// checking if the key is present.
func validateContainsKeyParams(a any) {
	reflectValueA := reflect.ValueOf(a)

	if IsNil(a) {
		panic("A is nil")
	} else if reflectValueA.Kind() != reflect.Map && reflectValueA.Kind() != reflect.Struct &&
		reflectValueA.Kind() != reflect.Ptr && reflectValueA.Kind() != reflect.Interface {
		panic(fmt.Sprintf("Unsupported type: %s", reflectValueA.Kind().String()))
	}
}

// containsValueInSlice checks if the provided value 'value' is contained within the slice 'reflectValueSlice'.
// It iterates over the elements of the slice and uses reflect.DeepEqual to compare each element with the value.
// If a match is found, it returns true. Otherwise, it returns false.
func containsValueInSlice(reflectValueSlice reflect.Value, value any) bool {
	for i := 0; i < reflectValueSlice.Len(); i++ {
		index := reflectValueSlice.Index(i)
		if reflect.DeepEqual(index.Interface(), value) {
			return true
		}
	}
	return false
}

// containsValueInMap checks if the provided value 'value' is contained within the map 'reflectValueMap'.
// It iterates over the keys of the map and uses reflect.DeepEqual to compare the value at each key with the value.
// If a match is found, it returns true. Otherwise, it returns false.
func containsValueInMap(reflectValueMap reflect.Value, value any) bool {
	mapKeys := reflectValueMap.MapKeys()
	for _, key := range mapKeys {
		mapValue := reflectValueMap.MapIndex(key)
		if reflect.DeepEqual(mapValue.Interface(), value) {
			return true
		}
	}
	return false
}

// containsValueInStruct iterates over the fields of the struct 'reflectValueStruct'
// and uses reflect.DeepEqual to compare each field with the value.
// If a match is found, it returns true. Otherwise, it returns false.
func containsValueInStruct(reflectValueStruct reflect.Value, value any) bool {
	for i := 0; i < reflectValueStruct.NumField(); i++ {
		if reflect.DeepEqual(reflectValueStruct.Field(i).Interface(), value) {
			return true
		}
	}
	return false
}
