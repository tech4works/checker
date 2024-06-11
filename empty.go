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
// For all other types of values, the function returns false.
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
//
// Note: The NonNil function is used to determine if a value is not nil.
// Refer to the NonNil function documentation for more details.
func AllNil(a any, b ...any) bool {
	b = append([]any{a}, b...)
	for _, i := range b {
		if NonNil(i) {
			return false
		}
	}
	return true
}

// AllNonNil checks whether any of the given values are nil.
//
// This function takes a variable number of interface parameters `a` and `b`.
// It appends `a` to the `b` slice and iterates over each value in the slice.
// It uses the `IsNil` function to determine if a value is nil. If any of the
// values are nil, it returns false; otherwise, it returns true.
//
// Example usage:
//
//	AllNonNil(nil)                  // false
//	AllNonNil(nil, nil)             // false
//	AllNonNil(nil, "hello", 42)     // false
//	AllNonNil("hello", 42)          // true
func AllNonNil(a any, b ...any) bool {
	return !AllNil(a, b...)
}

// IsEmpty checks if a given value is empty based on its type.
//
// It first calls the IsNil function to check if the value is nil. If it is nil,
// the function returns true immediately.
//
// If the value is not nil, it then uses reflection to determine the type of the value
// and checks if it is empty based on its type:
//   - For strings, it trims whitespace from the string and checks if the resulting string has zero length.
//   - For slices and arrays, it checks if the length is zero.
//   - For maps, it checks if the number of keys is zero.
//   - For all other types, it uses the IsZero method of the reflect.Value to check if the value is zero.
//
// If the value falls into any of the above cases, the function returns true.
// Otherwise, it returns false.
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

// IsNotEmpty returns a boolean value indicating whether a given value is not empty.
//
// It uses the IsEmpty function to determine if the value is empty. If the value is not empty,
// then it returns true; otherwise, it returns false.
func IsNotEmpty(a any) bool {
	return !IsEmpty(a)
}

// AllEmpty checks if all given values are empty.
//
// It takes a value `a` of any type and variadic values `b` of any type. It appends `a` to `b`
// and iterates over the resulting list. For each value `i` in the list, it calls the `IsNotEmpty` function to
// determine if the value is not empty. If any value is not empty, the function returns false.
// Otherwise, it returns true.
func AllEmpty(a any, b ...any) bool {
	b = append([]any{a}, b...)
	for _, i := range b {
		if IsNotEmpty(i) {
			return false
		}
	}
	return true
}

// AllNotEmpty checks if all given values are not empty.
//
// This function takes a variable number of parameters `a` and `b` and checks if all of them are not empty.
// It first appends `a` to the `b` slice. Then, it iterates through the combined slice and checks if each value is
// empty using the IsEmpty function. If any value is found to be empty, the function returns false immediately.
// If all values are not empty, the function returns true.
func AllNotEmpty(a any, b ...any) bool {
	return !AllEmpty(a, b...)
}

// IsNilOrEmpty checks if a value is nil or empty based on its type.
//
// It calls the IsNil function to check if the value is nil. If it is nil,
// the function returns true immediately.
//
// If the value is not nil, it then uses the IsEmpty function to check if the value is empty.
//
// The IsEmpty function uses reflection to determine the type of the value and checks if it is empty based on its type:
//   - For strings, it trims whitespace from the string and checks if the resulting string has zero length.
//   - For slices and arrays, it checks if the length is zero.
//   - For maps, it checks if the number of keys is zero.
//   - For all other types, it uses the IsZero method of the reflect.Value to check if the value is zero.
//
// If the value is nil or empty, the function returns true.
// Otherwise, it returns false.
func IsNilOrEmpty(a any) bool {
	return IsNil(a) || IsEmpty(a)
}

// IsNotNilOrEmpty is a function that checks if a given value is not nil or empty based on its type.
// It calls the IsNilOrEmpty function to perform the check and returns the negation of its result.
//
// The IsNilOrEmpty function first checks if the value is nil using the IsNil function. If it is nil,
// the function immediately returns true. Otherwise, it uses the IsEmpty function to check if the value is empty.
//
// The IsEmpty function uses reflection to determine the type of the value and checks if it is empty based on its type.
// For strings, it trims whitespace from the string and checks if the resulting string has zero length.
// For slices and arrays, it checks if the length is zero.
// For maps, it checks if the number of keys is zero.
// For all other types, it uses the IsZero method of the reflect.Value to check if the value is zero.
//
// If the value is nil or empty, the IsNilOrEmpty function returns true. Otherwise, it returns false.
//
// Parameters:
//   - a: The value to be checked.
//
// Returns:
//   - true if the value is not nil or empty.
//   - false if the value is nil or empty.
func IsNotNilOrEmpty(a any) bool {
	return !IsNilOrEmpty(a)
}

// AllNilOrEmpty checks if all given values are either nil or empty based on their types.
//
// This function takes a single value as the first argument `a`, followed by variadic values `b`.
// It appends `a` to the `b` slice and iterates over all values in `b`.
// For each value, it calls the IsNotNilOrEmpty function to check if the value is not nil or empty.
// If any value is not nil or empty, the function returns false immediately.
// If all values are either nil or empty, the function returns true.
func AllNilOrEmpty(a any, b ...any) bool {
	b = append([]any{a}, b...)
	for _, v := range b {
		if IsNotNilOrEmpty(v) {
			return false
		}
	}
	return true
}

// AllNotNilOrEmpty checks if all given values are neither nil nor empty based on their types.
//
// This function takes a single value as the first argument `a`, followed by variadic values `b`.
// It calls the AllNilOrEmpty function with `a` and `b` as arguments and returns the negation of its result.
func AllNotNilOrEmpty(a any, b ...any) bool {
	return !AllNilOrEmpty(a, b...)
}
