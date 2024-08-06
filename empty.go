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
	"reflect"
	"strings"
)

// IsNil determines whether a given value is nil using reflection.
//
// This function takes a {} interface parameter `a` and returns a boolean value
// indicating whether the value is nil. It uses reflection to determine the type
// of the value and checks if it is nil based on its type.
//
// The following types of values can be considered nil:
//   - Pointers
//   - Maps
//   - Matrices
//   - Channels
//   - Slices
//   - Functions
//   - Interfaces
//
// Parameters:
//   - a: Any interface value to be checked for nil.
//
// Returns:
//   - bool: A boolean value indicating whether the value is nil.
//
// Example:
//
//	var x *int
//	y := 10
//	fmt.Println(IsNil(x)) // true
//	fmt.Println(IsNil(y)) // false
func IsNil(a any) bool {
	rv := reflect.ValueOf(a)

	switch rv.Kind() {
	case reflect.Invalid:
		return true
	case reflect.Interface, reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice, reflect.Func:
		return rv.IsNil()
	default:
		return false
	}
}

// NonNil determines whether a given value is not nil. It uses the IsNil function
// to check if the value is nil and returns the negation of its result.
//
// Parameters:
//   - a: Any interface value to be checked for nil.
//
// Returns:
//   - bool: A boolean value indicating whether the value is not nil.
//
// Example:
//
//	var x *int
//	y := 10
//	fmt.Println(NonNil(x)) // false
//	fmt.Println(NonNil(y)) // true
func NonNil(a any) bool {
	return !IsNil(a)
}

// AllNil determines whether multiple values are all nil by iterating over them and checking if any of them is not nil.
//
// Parameters:
//   - a: Any interface value to be checked for nil.
//   - b: Variadic parameter for additional values to be checked for nil.
//
// Returns:
//   - bool: A boolean value indicating whether all the values are nil.
//
// Example:
//
//	fmt.Println(AllNil(nil, nil))          // true
//	fmt.Println(AllNil(nil, 0, "hello"))   // false
//	fmt.Println(AllNil(nil, nil, nil, nil))// true
//	fmt.Println(AllNil(nil))               // true
func AllNil(a, b any, c ...any) bool {
	for _, i := range append([]any{a, b}, c...) {
		if NonNil(i) {
			return false
		}
	}
	return true
}

// NoneNil determines whether all given values are not nil.
// It uses the IsNil function to check if each value is nil and returns false if any value is nil.
//
// Parameters:
//   - a: The first value to be checked for nil.
//   - b: The second value to be checked for nil.
//   - c: A variadic slice containing other values to be checked for nil.
//
// Returns:
//   - bool: A boolean value indicating whether all values are not nil.
//
// Example:
//
//	var x *int
//	y := 10
//	z := "Hello"
//	fmt.Println(NoneNil(x, y)) // false
//	fmt.Println(NoneNil(y, z)) // true
//
// NoneNil can also be used with a varying number of values:
//
//	var i *int
//	j := "Go"
//	k := 15
//	fmt.Println(NoneNil(j, k, "example", 20)) // true
//	fmt.Println(NoneNil(i, j, k)) // false
func NoneNil(a, b any, c ...any) bool {
	c = append([]any{a, b}, c...)
	for _, i := range c {
		if IsNil(i) {
			return false
		}
	}
	return true
}

// IsEmpty checks if a given value is empty based on its type.
//
// The function first calls the IsNil function to check if the value is nil. If it is nil,
// the function immediately returns true.
//
// If the value is not nil, it then uses reflection to determine the type of the value and
// checks if it is empty based on its type. The following checks are performed:
//   - For strings, it trims whitespace from the string and checks if the resulting string has zero length.
//   - For slices and arrays, it checks if the length is zero.
//   - For maps, it checks if the number of keys is zero.
//   - For all other types, it uses the IsZero method of the reflect.Value to check if the value is zero.
//
// If the value is nil or empty, the function returns true; otherwise, it returns false.
//
// Examples:
//
//	var strA string
//	fmt.Println(IsEmpty(strA))  // true
//
//	strB := "    "
//	fmt.Println(IsEmpty(strB))  // true
//
//	list := []int{}
//	fmt.Println(IsEmpty(list))  // true
//
//	m := make(map[string]int)
//	fmt.Println(IsEmpty(m))  // true
//
//	var ptr *int
//	fmt.Println(IsEmpty(ptr))  // true
func IsEmpty(a any) bool {
	if IsNil(a) {
		return true
	}

	reflectValue := reflect.ValueOf(a)
	if reflectValue.Kind() == reflect.Pointer || reflectValue.Kind() == reflect.Interface {
		reflectValue = reflectValue.Elem()
	}

	switch reflectValue.Kind() {
	case reflect.String:
		return len(strings.TrimSpace(reflectValue.String())) == 0
	case reflect.Slice, reflect.Array:
		return reflectValue.Len() == 0
	case reflect.Map:
		return len(reflectValue.MapKeys()) == 0
	default:
		return reflectValue.IsZero()
	}
}

// IsNotEmpty checks if a given value is not empty based on its type by calling the IsEmpty function and negating its result.
// This function takes a value `a` of any type and returns a boolean value indicating whether the value is not empty.
// It first calls the IsEmpty function with `a`. If the result of IsEmpty is true, it returns false. Otherwise, it returns true.
//
// Parameters:
//   - a: Any interface value to be checked for not being empty.
//
// Returns:
//   - bool: A boolean value indicating whether the value is not empty.
//
// Examples:
//
//	var strA string
//	fmt.Println(IsNotEmpty(strA))  // false
//
//	strB := "Hello, World!"
//	fmt.Println(IsNotEmpty(strB))  // true
//
//	list := []int{1, 2, 3}
//	fmt.Println(IsNotEmpty(list))  // true
//
//	m := make(map[string]int)
//	m["one"] = 1
//	fmt.Println(IsNotEmpty(m))  // true
//
//	var ptr *int
//	fmt.Println(IsNotEmpty(ptr))  // false
func IsNotEmpty(a any) bool {
	return !IsEmpty(a)
}

// AllEmpty checks if all given values are empty. Empty means either a value
// is nil or in case of strings and slices, their length is zero.
//
// Parameters:
//   - a: The first value to be checked.
//   - b: The second value to be checked.
//   - c: A variadic parameter that represents a list of additional values to be checked.
//
// Returns:
//   - bool: A boolean indicating whether all passed values are empty. It returns true when all values are empty and false otherwise.
//
// Example:
//
//	x, y := "", 0
//	z, w := []int{}, nil
//	fmt.Println(AllEmpty(x, y, z, w)) // true
//
//	a, b := "hello", 10
//	c, d := []int{1, 2, 3}, "world"
//	fmt.Println(AllEmpty(a, b, c, d)) // false
func AllEmpty(a, b any, c ...any) bool {
	c = append([]any{a, b}, c...)
	for _, i := range c {
		if IsNotEmpty(i) {
			return false
		}
	}
	return true
}

// NoneEmpty determines whether all given values are not empty using the IsEmpty function.
//
// This function first calls the IsEmpty function on each given value to check if any value is empty.
// If any value is empty, it immediately returns false.
// If all values are not empty, the function returns true.
// The function takes any number of arguments of any types as it variadically accepts argument of type `any`.
//
// Parameters:
//   - a: The first value to be checked if it is not empty.
//   - b: The second value to be checked if it is not empty.
//   - c: The rest of the values to be checked if they are not empty.
//
// Returns:
//   - bool: A boolean value indicating whether all given values are not empty.
//
// Example:
//
//	strA := "Hello"
//	strB := "Go"
//	fmt.Println(NoneEmpty(strA, strB)) // true
//
//	strC := "    "
//	fmt.Println(NoneEmpty(strA, strC)) // false
//
//	list1 := []int{1, 2, 3}
//	list2 := []int{}
//	fmt.Println(NoneEmpty(list1, list2)) // false
//
//	m1 := make(map[string]int{"key": 1})
//	m2 := make(map[string]int)
//	fmt.Println(NoneEmpty(m1, m2))  // false
func NoneEmpty(a, b any, c ...any) bool {
	c = append([]any{a, b}, c...)
	for _, i := range c {
		if IsEmpty(i) {
			return false
		}
	}
	return true
}

// IsNilOrEmpty checks whether a given value is nil or empty.
//
// Parameters:
//   - a: Any interface value to be checked for nil or empty.
//
// Returns:
//   - bool: A Boolean value indicating whether the value is either nil or empty.
//
// Example:
//
//	var str string
//	fmt.Println(IsNilOrEmpty(str)) // true
//
//	list := []int{}
//	fmt.Println(IsNilOrEmpty(list))  // true
//
//	m := make(map[string]int)
//	fmt.Println(IsNilOrEmpty(m))  // true
//
//	var x *int
//	fmt.Println(IsNilOrEmpty(x)) // true
//
//	y := 10
//	fmt.Println(IsNilOrEmpty(y)) // false
func IsNilOrEmpty(a any) bool {
	return IsNil(a) || IsEmpty(a)
}

// IsNotNilOrEmpty checks whether a given value is not nil or empty. It uses the IsNilOrEmpty function
// to check if the value is nil or empty and returns the negation of its result.
//
// Parameters:
//   - a: Any interface value to be checked for nil or empty.
//
// Returns:
//   - bool: A boolean value indicating whether the value is not nil or empty.
//
// Example:
//
//	var str string
//	fmt.Println(IsNotNilOrEmpty(str)) // false
//
//	list := []int{}
//	fmt.Println(IsNotNilOrEmpty(list))  // false
//
//	m := make(map[string]int)
//	fmt.Println(IsNotNilOrEmpty(m))  // false
//
//	var x *int
//	fmt.Println(IsNotNilOrEmpty(x)) // false
//
//	y := 10
//	fmt.Println(IsNotNilOrEmpty(y)) // true
func IsNotNilOrEmpty(a any) bool {
	return !IsNilOrEmpty(a)
}

// AllNilOrEmpty checks whether all supplied values are nil or empty.
// Multiple values can be passed through variadic arguments, which are then
// checked one by one using the IsNotNilOrEmpty function. If any of the values
// are not nil or empty, it returns false.
//
// Parameters:
//   - a: The initial value to check for being nil or empty.
//   - b: Additional values that should be checked for being nil or empty.
//
// Returns:
//   - bool: A boolean value indicating whether all values are nil or empty.
//
// Example:
//
//	var str string
//	x := 10
//	var m map[string]int
//
//	fmt.Println(AllNilOrEmpty(str)) // true
//	fmt.Println(AllNilOrEmpty(x)) // false
//	fmt.Println(AllNilOrEmpty(m, x, str)) // false
func AllNilOrEmpty(a, b any, c ...any) bool {
	c = append([]any{a, b}, c...)
	for _, v := range c {
		if IsNotNilOrEmpty(v) {
			return false
		}
	}
	return true
}

// NoneNilOrEmpty checks whether all the given values are neither nil nor empty.
//
// Parameters:
//   - a: The first interface value to be checked for nil or empty.
//   - b: The second interface value to be checked for nil or empty.
//   - c: Additional interfaces (variadic) to be checked for nil or empty.
//
// The function appends all parameters into a single slice and then iterates over this slice, for each value checking
// if it is nil or empty by invoking the IsNilOrEmpty() function.
//
// Returns:
//   - bool: A boolean value indicating whether all values are neither nil nor empty.
//     If IsNilOrEmpty() ever returns true for any value, NoneNilOrEmpty() immediately returns false.
//     If IsNilOrEmpty() returns false for all values, NoneNilOrEmpty() will return true, indicating that no values
//     were nil or empty.
//
// Example:
//
//	x := 10
//	y := "NotEmpty"
//	z := []int{1, 2, 3}
//	w := map[string]int{"one": 1}
//	fmt.Println(NoneNilOrEmpty(x, y, z, w)) // true
//
//	x := (*int)(nil)
//	y := ""
//	z := []int{}
//	w := map[string]int{}
//	fmt.Println(NoneNilOrEmpty(x, y, z, w)) // false
func NoneNilOrEmpty(a, b any, c ...any) bool {
	c = append([]any{a, b}, c...)
	for _, v := range c {
		if IsNilOrEmpty(v) {
			return false
		}
	}
	return true
}

// IfNilReturns checks if the given value is nil and if so, returns the other
// specified value. It uses the IsNil function to perform the nil check.
//
// Parameters:
//   - a: The pointer to an arbitrary value that will be checked for nil.
//   - b: The arbitrary value that will be returned if 'a' is nil.
//
// Returns:
//   - T: The same kind of value as 'a' and 'b'. If 'a' is not nil, the referenced value of 'a' will
//     be returned. If 'a' is nil, 'b' will be returned instead.
//
// Panic:
//   - This function might panic if 'a' is nil and the dereference operation is performed on it.
//
// Example:
//
//	var x *int
//	var y int = 10
//	fmt.Println(IfNilReturns(x, y)) // 10
//	x = new(int)
//	*x = 5
//	fmt.Println(IfNilReturns(x, y)) // 5
func IfNilReturns[T any](a *T, b T) T {
	if IsNil(a) {
		return b
	}
	return *a
}

// IfEmptyReturns checks if the first parameter 'a' value is empty, and returns the second parameter 'b' if 'a' is empty.
// The empty check is done using the IsEmpty function. If 'a' is not empty, it simply returns 'a'
//
// Parameters:
//   - a: The primary value to be checked for emptiness.
//   - b: The value to be returned if 'a' is empty.
//
// Returns:
//   - T: The original value 'a' if it is not empty, or the backup value 'b' if 'a' is empty.
//
// Example:
//
//	strA := ""
//	strB := "backup"
//	fmt.Println(ifEmptyReturns(strA, strB))  // output: "backup"
//
//	numA := 0
//	numB := 10
//	fmt.Println(ifEmptyReturns(numA, numB))  // output: 10
//
//	listA := []int{}
//	listB := []int{1, 2, 3}
//	fmt.Println(ifEmptyReturns(listA, listB))  // output: [1 2 3]
func IfEmptyReturns[T any](a T, b T) T {
	if IsEmpty(a) {
		return b
	}
	return a
}
