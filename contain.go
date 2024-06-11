package comparator

import (
	"fmt"
	"reflect"
	"strings"
)

// Contains checks if the provided value 'b' is contained within the value 'a'.
// It validates 'a' using the validateContains function.
// If 'a' is a pointer or interface, it calls Contains with the dereferenced value.
// If 'a' is a slice or array, it calls containsValueInSlice to check if 'b' is contained.
// If 'a' is a map, it calls containsValueInMap to check if 'b' is contained.
// If 'a' is a struct, it calls containsValueInStruct to check if 'b' is contained.
// If 'a' is a string and 'b' is a string, it uses strings.Contains to check if 'b' is contained in 'a'.
//
// If none of the above conditions are met, it returns false.
func Contains(a, b any) bool {
	validateContains(a)

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
	} else if reflectValueB.Kind() == reflect.String {
		return strings.Contains(reflectValueA.String(), reflectValueB.String())
	}

	return false
}

// NotContains checks if the provided value 'b' is not contained within the value 'a'.
// It uses the Contains function to determine if 'b' is contained in 'a',
// returning the negation of the result.
// Returns true if 'b' is not found in 'a', false otherwise.
func NotContains(a, b any) bool {
	return !Contains(a, b)
}

// ContainsIgnoreCase checks if the provided value 'b' is contained within the value 'a',
// ignoring case sensitivity. It supports values of type String and recursively checks
// nested values if 'a' is a pointer or interface. It uses lowercase comparison to determine
// if the values are equal. Returns true if 'b' is found in 'a', false otherwise.
func ContainsIgnoreCase(a, b any) bool {
	validateContainsIgnoreCase(a)

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

// NotContainsIgnoreCase checks if the provided value 'b' is not contained within the value 'a',
// ignoring case sensitivity. It supports values of type String and recursively checks
// nested values if 'a' is a pointer or interface. It uses lowercase comparison to determine
// if the values are not equal. Returns true if 'b' is not found in 'a', false otherwise.
func NotContainsIgnoreCase(a, b any) bool {
	return !ContainsIgnoreCase(a, b)
}

// ContainsKey checks if the given key 'key' is present as a field in the value 'a'.
// Supports Map and Struct type values.
// It recursively checks nested values 'if 'a' is a pointer or interface.
// Returns true if 'key' is found in 'a', otherwise returns false.
func ContainsKey(a, key any) bool {
	validateContainsKey(a)

	reflectValue := reflect.ValueOf(a)
	if reflectValue.Kind() == reflect.Ptr || reflectValue.Kind() == reflect.Interface {
		return ContainsKey(reflectValue.Elem().Interface(), key)
	}

	reflectKey := reflect.ValueOf(key)
	if reflectKey.Kind() == reflect.Ptr || reflectKey.Kind() == reflect.Interface {
		return ContainsKey(a, reflectKey.Elem().Interface())
	}

	if reflectValue.Kind() == reflect.Map {
		return reflectValue.MapIndex(reflectKey).IsValid()
	} else if reflectValue.Kind() == reflect.Struct && reflectKey.Kind() == reflect.String {
		_, found := reflectValue.Type().FieldByName(reflectKey.String())
		return found
	}

	return false
}

// NotContainsKey checks if the given key 'key' is not present as a field in the value 'a'.
// It is the negation of the ContainsKey function.
// Returns true if 'key' is not found in 'a', otherwise returns false.
func NotContainsKey(a, key any) bool {
	return !ContainsKey(a, key)
}

// ContainsOnSlice checks if any element in the slice 'a' satisfies the condition specified by the 'check' function.
// It iterates over each element in the slice and applies the 'check' function to determine if the element satisfies the condition.
// If any element satisfies the condition, it returns true; otherwise, it returns false.
func ContainsOnSlice[T any](a []T, check func(index T) bool) bool {
	for _, index := range a {
		if check(index) {
			return true
		}
	}
	return false
}

// NotContainsOnSlice checks if none of the elements in the slice 'a' satisfies the condition specified by
// the 'check' function.
// It calls ContainsOnSlice with the same arguments and negates the result.
func NotContainsOnSlice[T any](a []T, check func(index T) bool) bool {
	return !ContainsOnSlice(a, check)
}

// validateContains validates the value 'a' to ensure it is a supported type for
// the Contains function. If 'a' is nil, it panics with the message "A is nil".
// If 'a' is not one of the supported types (slice, array, map, struct, string),
// it panics with a formatted string message indicating the unsupported type.
func validateContains(a any) {
	reflectValueA := reflect.ValueOf(a)

	if IsNil(a) {
		panic("A is nil")
	} else if reflectValueA.Kind() != reflect.Slice && reflectValueA.Kind() != reflect.Array &&
		reflectValueA.Kind() != reflect.Map && reflectValueA.Kind() != reflect.Struct &&
		reflectValueA.Kind() != reflect.String && reflectValueA.Kind() != reflect.Ptr {
		panic(fmt.Sprintf("Unsupported type: %s", reflectValueA.Kind().String()))
	}
}

// validateContainsIgnoreCase validates the value 'a' to ensure it is not nil and of type string.
// If 'a' is nil, it panics with the message "A is nil".
// If 'a' is not of type string, it panics with a formatted message indicating the unsupported type.
func validateContainsIgnoreCase(a any) {
	reflectValueA := reflect.ValueOf(a)

	if IsNil(a) {
		panic("A is nil")
	} else if reflectValueA.Kind() != reflect.String && reflectValueA.Kind() != reflect.Ptr {
		panic(fmt.Sprintf("Unsupported type: %s", reflectValueA.Kind().String()))
	}
}

// validateContainsKey validates the value 'a' to ensure it is a map or struct.
// If 'a' is nil, it panics with the message "A is nil".
// If 'a' is neither a map nor a struct, it panics with a formatted string
// indicating the unsupported type.
// This function uses reflection to determine the type of 'a'.
// It is used by the ContainsKey function to validate the input value before
// checking if the key is present.
func validateContainsKey(a any) {
	reflectValueA := reflect.ValueOf(a)

	if IsNil(a) {
		panic("A is nil")
	} else if reflectValueA.Kind() != reflect.Map && reflectValueA.Kind() != reflect.Struct {
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
