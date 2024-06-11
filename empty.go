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
	c = append([]any{a, b}, c...)
	for _, i := range c {
		if NonNil(i) {
			return false
		}
	}
	return true
}

// AllNonNil determines whether multiple values are all non-nil by calling the AllNil function and negating its result.
//
// This function takes at least two interface parameters `a` and `b`, and optionally additional values `c ...any` to be checked for nil.
// It calls the AllNil function with the parameters `a`, `b`, and `c` and returns the negation of the result.
//
// Parameters:
//   - a: Any interface value to be checked for nil.
//   - b: Second interface value to be checked for nil.
//   - c: Variadic parameter for additional values to be checked for nil.
//
// Returns:
//   - bool: A boolean value indicating whether all the values are non-nil.
//
// Example:
//
//	AllNonNil(nil)                  // false
//	AllNonNil(nil, nil)             // false
//	AllNonNil(nil, "hello", 42)     // false
//	AllNonNil("hello", 42)          // true
func AllNonNil(a, b any, c ...any) bool {
	return !AllNil(a, b, c...)
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

// AllNotEmpty checks if all given values are not empty. The values are not empty if they are not nil and,
// in the case of strings and slices, their length is not zero.
//
// Parameters:
//   - a: The first value to be checked for not being empty.
//   - b: The second value to be checked for not being empty.
//   - c: A variadic parameter that represents a list of additional values to be checked for not being empty.
//
// Returns:
//   - bool: A boolean value indicating whether all passed values are not empty. It returns true when all provided
//     values are not empty, false otherwise.
//
// Example:
//
//	strA, strB := "Hello, World!", ""
//	intC := []int{1, 2, 3}
//	n := nil
//	fmt.Println(AllNotEmpty(strA, strB, intC, n)) // false (because strB and n are empty)
//
//	strX, strY := "Hello, World!", "Programming is fun!"
//	intZ := []int{1, 2, 3}
//	m := make(map[string]int)
//	m["Gophers"] = 6
//	fmt.Println(AllNotEmpty(strX, strY, intZ, m)) // true (because all values are not empty)
func AllNotEmpty(a, b any, c ...any) bool {
	return !AllEmpty(a, b, c...)
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

// AllNotNilOrEmpty verifies whether all supplied values are not nil or empty.
// The function accepts an initial pair of values and an optional series of additional values.
// These supplementary arguments are processed using the IsNotNilOrEmpty function. If all
// the values are neither nil nor empty, the function returns true.
//
// Parameters:
//   - a: The first value to check for not being nil or empty.
//   - b: The second value to check for not being nil or empty.
//   - c: Optional additional values to check for not being nil or empty.
//
// Returns:
//   - bool: A boolean value indicating whether all the values are neither nil nor empty.
//
// Example:
//
//	var str string
//	x := 10
//	var m map[string]int
//
//	fmt.Println(AllNotNilOrEmpty(str, "NotEmpty")) // false
//	fmt.Println(AllNotNilOrEmpty(x, 20)) // true
//	fmt.Println(AllNotNilOrEmpty(m, x, "NotEmpty")) // false
func AllNotNilOrEmpty(a, b any, c ...any) bool {
	return !AllNilOrEmpty(a, b, c...)
}
