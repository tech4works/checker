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
	"encoding/json"
	"reflect"
	"regexp"
	"strconv"
	"time"
)

// IsJSON checks if a given value can be a map or a slice in JSON format. It uses the IsMap and
// IsSlice functions to check the given value and returns true if either function returns true.
//
// Parameters:
//   - a: The value of any type to be checked if it can be presented in JSON as map or slice
//
// Returns:
//   - bool: A boolean value indicating whether the value can be represented in JSON format as a map or a slice.
//
// Example:
//
//	jsonMap := `{"key": "value"}`
//	jsonArray := `[1, 2, 3]`
//	notJson := `Not a JSON string`
//
//	fmt.Println(IsJSON(jsonMap)) // Outputs: true
//	fmt.Println(IsJSON(jsonArray)) // Outputs: true
//	fmt.Println(IsJSON(notJson)) // Outputs: false
func IsJSON(a any) bool {
	return IsMap(a) || IsSlice(a)
}

// IsMap determines whether a given value is a map type.
// It does this by attempting to unmarshal JSON from the given value's byte representation.
//
// Parameters:
//   - `a`: The value of any type to be checked if it's a map.
//
// Panics:
//   - If `a` is not convertible to a byte slice, a panic occurs as the underlying 'toBytes' function throws a panic.
//
// Returns:
//   - `bool`: A boolean value indicating whether the value is a map or not.
//
// Example:
//
//	str := `{"key":"value"}`
//	num := 1234
//	fmt.Println(IsMap(str)) // true
//	fmt.Println(IsMap(num)) // false
func IsMap(a any) bool {
	var js map[string]any
	return json.Unmarshal(toBytes(a), &js) == nil
}

// IsSlice checks if a given value is a slice. It uses toBytes function to convert the given
// value into a byte slice. It then uses json.Unmarshal function to unmarshal the byte
// slice into a slice and, if the unmarshal operation is successful, returns true.
//
// Parameters:
//   - a: Any value which should be checked end evaluated if it is a slice.
//
// Returns:
//   - bool: A boolean value indicating whether the input parameter can be unmarshalled into a slice.
//
// Panic:
//   - The function can panic if it encounters any unsupported types during the process or
//     when the provided value is not convertible to string.
//
// Example:
//
//	var x []int = []int{1,2,3}
//	y := "Not a slice"
//	fmt.Println(IsSlice(x)) // Outputs: true
//	fmt.Println(IsSlice(y)) // Outputs: false
func IsSlice(a any) bool {
	var slice []any
	return json.Unmarshal(toBytes(a), &slice) == nil
}

// IsSliceOfMaps checks if the given value is a slice of maps.
// It first converts the value to a byte slice using toBytes function and then attempts to unmarshal
// the byte slice into a slice of maps with the key being string and value being any.
// It returns true if the value can be successfully unmarshalled into a slice of maps and false otherwise.
//
// Parameters:
//   - a: Any value to be checked if it is a slice of maps
//
// Returns:
//   - bool: A boolean value indicating whether the given value is a slice of maps.
//
// Panic:
//   - The function can panic if it encounters any unsupported types during the process or
//     when the provided value is not convertible to string.
//
// Example:
//
//	var x []map[string]int
//	y := "not a slice of maps"
//
//	x = append(x, map[string]int{"foo": 42})
//	fmt.Println(IsSliceOfMaps(x)) // true
//	fmt.Println(IsSliceOfMaps(y)) // false
func IsSliceOfMaps(a any) bool {
	var slice []map[string]any
	return json.Unmarshal(toBytes(a), &slice) == nil
}

// IsInt determines whether a given value can be converted to an integer. It uses the toString function
// to convert the value to a string and strconv.Atoi function to try converting the converted string to an integer.
// The function returns true if the conversion is successful (error from strconv.Atoi is nil), and false otherwise.
//
// Parameters:
//   - a: Any interface value to be checked for its convertibility to an integer.
//
// Returns:
//   - bool: A boolean value indicating whether the value can be converted to an integer.
//
// Panic:
//   - The function can panic if it encounters any unsupported types during the process or
//     when the provided value is not convertible to string.
//
// Example:
//
//	a := "10"
//	b := "hello"
//	fmt.Println(IsInt(a)) // true
//	fmt.Println(IsInt(b)) // false
func IsInt(a any) bool {
	_, err := strconv.Atoi(toString(a))
	return err == nil
}

// IsBool determines whether a given value can be converted to a boolean.
// It takes a value of any type as its parameter, converts it to a string using the toString function,
// and tries to parse the string as a boolean using strconv.ParseBool.
// If the parsing process is successful (i.e., no error occurred), we determine that the input value can be
// converted to a boolean and return true. Otherwise, we return false.
//
// Parameters:
//   - a: Any value to be evaluated for boolean conversion.
//
// Returns:
//   - bool: A boolean value, `true` if the given value can be converted to a boolean and `false` otherwise.
//
// Panics:
//   - If `a` is not convertible to a string, a panic occurs as the underlying 'toString' function throws a panic.
//
// Example:
//
//	var x string = "true"
//	var y string = "false"
//	var z string = "not boolean"
//	var a int = 0
//	var b int = 1
//	fmt.Println(IsBool(x)) // true
//	fmt.Println(IsBool(y)) // true
//	fmt.Println(IsBool(z)) // false
//	fmt.Println(IsBool(a)) // true
//	fmt.Println(IsBool(b)) // true
func IsBool(a any) bool {
	_, err := strconv.ParseBool(toString(a))
	return err == nil
}

// IsFloat determines whether a given value can be parsed into a float64.
// It uses the strconv.ParseFloat function to attempt parsing the value received as a string.
// If parsing succeeds without throwing an error, it returns true. If an error occurs during parsing, it returns false.
//
// Parameters:
//   - a: Any interface value that needs to be tested for float64 parseability.
//
// Returns:
//   - bool: A boolean value indicating whether the given value can be parsed into a float64.
//
// Panics:
//   - If `a` is not convertible to a string, a panic occurs as the underlying 'toString' function throws a panic.
//
// Example:
//
//	var x string = "10.5"
//	var y string = "Hello"
//	var z int = 10
//	fmt.Println(IsFloat(x)) // true
//	fmt.Println(IsFloat(y)) // false
//	fmt.Println(IsFloat(z)) // true
func IsFloat(a any) bool {
	_, err := strconv.ParseFloat(toString(a), 64)
	return err == nil
}

// IsTime checks if a given value can be converted to a time.Time type.
//
// Parameters:
//   - a: The value of any type to be checked for possible conversion to time.Time.
//
// Returns:
//   - bool: A boolean value indicating whether the value can be converted to time.Time.
//
// It internally uses the toTimeWithErr function, which attempts to convert the value to a time.Time.
// If the conversion is possible (i.e., if toTimeWithErr doesn't return an error), true is returned.
// Otherwise, if an error occurs during the conversion, false is returned.
//
// Example:
//
//	var x string = "2020-07-14T04:12:02Z"
//	y := 1594701122
//	fmt.Println(IsTime(x)) // true
//	fmt.Println(IsTime(y)) // true
//	z := "This is a string, not a timestamp"
//	fmt.Println(IsTime(z)) // false
func IsTime(a any) bool {
	_, err := toTimeWithErr(a)
	return err == nil
}

// IsDuration checks if a given value can be parsed as a time.Duration type using the time.ParseDuration
// and returns a boolean value based on the parse result.
//
// Parameters:
//   - a: An interface value to be parsed into a time.Duration type.
//
// Returns:
//   - bool: A boolean value indicating whether the value can be parsed into the time.Duration type.
//     It returns true if the parsing succeeds, false otherwise.
//
// Panics:
//   - Files a panic if the value passed into toString function in this context is invalid in any sense.
//
// Example:
//
//	var durationString string = "2h45m"
//	var randomString string = "abc123"
//	fmt.Println(IsDuration(durationString)) // true
//	fmt.Println(IsDuration(randomString)) // false
func IsDuration(a any) bool {
	_, err := time.ParseDuration(toString(a))
	return err == nil
}

// IsByteUnit validates whether the given value follows the byte unit pattern. Achieves this
// by converting the input to a string and matching it against a regular expression that allows
// any digit followed by a byte unit(B, KB, MB, GB, TB, PB).
//
// Parameters:
//   - a: Input of any type to be checked against the byte unit pattern.
//
// Returns:
//   - bool: A boolean value indicating whether the input matches the byte unit pattern.
//
// Panic:
//   - Throws a panic when the input type is not supported by the underlying toString function.
//
// Example:
//
//	bu1 := 10
//	bu2 := "20KB"
//	bu3 := "300MB"
//	bu4 := "Hello World!"
//
//	fmt.Println(IsByteUnit(bu1)) // false
//	fmt.Println(IsByteUnit(bu2)) // true
//	fmt.Println(IsByteUnit(bu3)) // true
//	fmt.Println(IsByteUnit(bu4)) // false
func IsByteUnit(a any) bool {
	regex := regexp.MustCompile(`^(\d+)(B|KB|MB|GB|TB|PB)$`)
	return regex.MatchString(toString(a))
}

// IsPointerType checks whether the given value's type is a pointer. The
// function uses the reflection package's ValueOf and Kind functions to achieve
// this.
//
// Parameters:
//   - a: An interface value, the type of which is to be checked.
//
// Returns:
//   - bool: A boolean indicating whether the given value's type is a pointer.
//
// Example:
//
//	var x *int
//	y := 10
//	fmt.Println(IsPointerType(x)) // true
//	fmt.Println(IsPointerType(y)) // false
//
// Please note that the function does not handle nil pointers and can panic if a
// nil pointer is passed in. Always check for nil before passing pointers.
func IsPointerType(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.Pointer
}

// IsFuncType checks whether the given value is a function.
// It uses the reflection package to inspect the value and verifies if its kind is Func.
//
// Parameters:
//   - a: An interface{} value for which the functional type check has to be performed.
//
// Returns:
//   - bool: A boolean indicating whether the passed value is of a Func type.
//
// Example:
//
//	func SampleFunction() {}
//	x := 10
//	fmt.Println(IsFuncType(SampleFunction)) // true
//	fmt.Println(IsFuncType(x)) // false
//
// Please note that the function does not handle nil pointers and can panic if a
// nil pointer is passed in. Always check for nil before passing pointers.
func IsFuncType(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.Func
}

// IsChanType checks whether the provided value is of a channel type.
// It uses the reflection package to determine the type of the input value and returns true if it's a channel.
//
// Parameters:
//   - a: Any value (of any type) to be checked for being a channel.
//
// Returns:
//   - bool: A boolean value indicating whether the value is a channel.
//
// Example:
//
//	var x *int
//	c := make(chan int)
//	fmt.Println(IsChanType(x)) // false
//	fmt.Println(IsChanType(c)) // true
//
// Please note that the function does not handle nil pointers and can panic if a
// nil pointer is passed in. Always check for nil before passing pointers.
func IsChanType(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.Chan
}

// IsMapType checks whether the provided value is a map.
// It uses the reflection package to observe the actual type of the given value and compares it with reflection.Map.
//
// Parameters:
//   - a: Any interface value that needs to be checked if it is a map.
//
// Returns:
//   - bool: A boolean value indicating whether the input is a map.
//
// Example:
//
//	m := map[string]int{"Alice": 23, "Bob": 24}
//	i := 10
//	fmt.Println(IsMapType(m)) // true
//	fmt.Println(IsMapType(i)) // false
//
// Please note that the function does not handle nil pointers and can panic if a
// nil pointer is passed in. Always check for nil before passing pointers.
func IsMapType(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.Map
}

// IsStructType determines whether a given value is of a Struct type.
// It uses the reflection package's Kind method
// to check if the value's type is a Struct and returns as a boolean result.
//
// Parameters:
//   - a: An interface value of any type to be checked for a Struct type.
//
// Returns:
//   - bool: A boolean value indicating whether the value's type is Struct.
//
// Example:
//
//	type Person struct {
//		Name string
//		Age  int
//	}
//
//	var p Person
//	x := 10
//	fmt.Println(IsStructType(p)) // true
//	fmt.Println(IsStructType(x)) // false
//
// Please note that the function does not handle nil pointers and can panic if a
// nil pointer is passed in.
// Always check for nil before passing pointers.
func IsStructType(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.Struct
}

// IsSliceType checks whether a given value is of a slice type.
// It uses the Go reflect package's ValueOf method to infer the type of the value and subsequently checks if it is a
// slice using the Kind method.
//
// Parameters:
//   - a: Any interface value to be checked for a slice type.
//
// Returns:
//   - bool: A boolean indicative of whether the value is of slice type or not.
//
// Example:
//
//	var x []int
//	y := 10
//	fmt.Println(IsSliceType(x)) // true
//	fmt.Println(IsSliceType(y)) // false
//
// Please note that the function does not handle nil pointers and can panic if a
// nil pointer is passed in. Always check for nil before passing pointers.
func IsSliceType(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.Slice
}

// IsArrayType checks whether a given value is of an array type.
// It uses the Go reflect package's ValueOf function to infer the type of the variable and subsequently checks if it's
// an array using the Kind function.
//
// Parameters:
//   - a: This is any interface value that is to be checked for an array type.
//
// Returns:
//   - bool: This is a boolean value indicating whether the value is of an array type.
//
// Example:
//
//	var x [3]int
//	y := 10
//	fmt.Println(IsArrayType(x)) // true
//	fmt.Println(IsArrayType(y)) // false
//
// Please note that this function does not handle nil pointers and can panic if a nil pointer is passed in.
// It is always necessary to check for nil before passing pointers.
func IsArrayType(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.Array
}

// IsSliceOrArrayType determines whether a given value is either a slice or an array type. It uses
// IsSliceType and IsArrayType functions to check if the value is a slice or an array type respectively.
//
// Parameters:
//   - a: Any interface value to be checked.
//
// Returns:
//   - bool: A boolean value indicating whether the value is either a slice or an array type.
//
// Example:
//
//	var x []int
//	var y [3]int
//	var z int
//
//	fmt.Println(IsSliceOrArrayType(x)) // true
//	fmt.Println(IsSliceOrArrayType(y)) // true
//	fmt.Println(IsSliceOrArrayType(z)) // false
//
// Note: This function does not handle nil pointers and can panic if a
// nil pointer is passed in. Always check for nil before passing pointers.
func IsSliceOrArrayType(a any) bool {
	return IsSliceType(a) || IsArrayType(a)
}

// IsStringType determines whether a given value is of type string.
// It uses the reflection package to check the kind of the value and compares it with reflect.String.
//
// Parameters:
//   - a: Any interface value to be checked for type string.
//
// Returns:
//   - bool: A boolean value indicating whether the value is of type string.
//
// Example:
//
//	var x int = 10
//	y := "Hello, world!"
//	fmt.Println(IsStringType(x)) // false
//	fmt.Println(IsStringType(y)) // true
//
// Note: This function does not handle nil pointers and can panic if a
// nil pointer is passed in. Always check for nil before passing pointers.
func IsStringType(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.String
}

// IsIntType checks whether the provided interface value is an integer type.
// It uses the reflection package to inspect the type of the value at runtime
// and compares this with the reflect.Int type.
//
// Parameters:
//   - a: Any interface value that should be checked if it is an integer.
//
// Returns:
//   - bool: A boolean value indicating whether the value has an integer type.
//
// Example:
//
//	var a int = 10
//	var b string = "test"
//	fmt.Println(IsIntType(a)) // true
//	fmt.Println(IsIntType(b)) // false
//
// Note: This function does not handle nil pointers and can panic if a
// nil pointer is passed in. Always check for nil before passing pointers.
func IsIntType(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.Int
}

// IsInt8Type checks whether a given value is of an Int8 type.
// It uses the reflect package's ValueOf and Kind functions
// to determine the type of the value.
//
// Parameters:
//   - a: Any interface value whose type is to be checked.
//
// Returns:
//   - bool: A boolean value indicating whether the value is of an Int8 type.
//
// Example:
//
//	var a int8 = 10
//	var b int = 20
//	fmt.Println(IsInt8Type(a)) // true
//	fmt.Println(IsInt8Type(b)) // false
//
// Note: This function does not handle nil pointers and can panic if a
// nil pointer is passed in. Always check for nil before passing pointers.
func IsInt8Type(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.Int8
}

// IsInt16Type checks if the given value is of type int16. It uses the reflect package's ValueOf method
// to create a reflect.Value of the "a" argument and the Kind method to get its exact kind.
// It then checks if the kind is identical to reflect.Int16.
//
// Parameters:
//   - a: Any interface value to be checked for being of type int16.
//
// Returns:
//   - bool: A boolean value indicating whether the kind of the value is int16.
//
// Example:
//
//	var a int16 = 10
//	var b = "string"
//	fmt.Println(IsInt16Type(a)) // true
//	fmt.Println(IsInt16Type(b)) // false
//
// Note: This function does not handle nil pointers and can panic if a
// nil pointer is passed in. Always check for nil before passing pointers.
func IsInt16Type(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.Int16
}

// IsInt32Type checks if the provided input is of type int32. It uses the
// reflect package's ValueOf and Kind methods to determine the kind of the
// value and compares it to reflect.Int32.
//
// Parameters:
//   - a: The value to be checked. This can be of any type.
//
// Returns:
//   - bool: A boolean value indicating whether the value is of type int32.
//
// Example:
//
//	var x int32 = 10
//	var y int = 20
//	fmt.Println(IsInt32Type(x)) // true
//	fmt.Println(IsInt32Type(y)) // false
//
// Note: This function does not handle nil pointers and can panic if a
// nil pointer is passed in. Always check for nil before passing pointers.
func IsInt32Type(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.Int32
}

// IsInt64Type checks if the provided interface value is of type Int64. It uses the reflect package's ValueOf
// function to get the value's Kind and compares it to reflect.Int64.
//
// Parameters:
//   - a: Any interface value to be checked for its type.
//
// Returns:
//   - bool: A boolean value indicating whether the provided value is of type Int64.
//
// Example:
//
//	var x int64
//	y := "Not int64"
//	fmt.Println(IsInt64Type(x)) // true
//	fmt.Println(IsInt64Type(y)) // false
//
// Note: This function does not handle nil pointers and can panic if a
// nil pointer is passed in. Always check for nil before passing pointers.
func IsInt64Type(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.Int64
}

// IsUintType determines whether a given value is of the unsigned integer type.
// It uses the Kind method from reflect package to check the type of the value and compares it with reflect.Uint.
//
// Parameters:
//   - a: Any interface value to be checked for the unsigned integer type.
//
// Returns:
//   - bool: A boolean value indicating whether the type of value is unsigned integer.
//
// Example:
//
//	var x uint = 10
//	var y uint8 = 10
//	fmt.Println(IsUintType(x)) // true
//	fmt.Println(IsUintType(y)) // false
//
// Note: This function does not handle nil pointers and can panic if a
// nil pointer is passed in. Always check for nil before passing pointers.
func IsUintType(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.Uint
}

// IsUint8Type inspects if the provided interface value is of type uint8.
// It uses the reflect package to inspect the value's kind and compares it to reflect.Uint8.
//
// Parameters:
//   - a: This is any interface value to be checked for type uint8.
//
// Returns:
//   - bool: This returns a boolean value indicating whether the value is of type uint8.
//
// Example:
//
//	var x uint8 = 8
//	y := "hello"
//	fmt.Println(IsUint8Type(x)) // true
//	fmt.Println(IsUint8Type(y)) // false
//
// Note: This function does not handle nil pointers and can panic if a
// nil pointer is passed in. Always check for nil before passing pointers.
func IsUint8Type(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.Uint8
}

// IsUint16Type determines whether a given value is of type uint16.
// It uses the reflect.ValueOf function to derive the Kind of the provided interface value,
// and compares it to reflect.Uint16.
//
// Parameters:
//   - a: an interface value of any type.
//
// Returns:
//   - bool: A boolean value indicating whether the given value is of type uint16.
//
// Example:
//
//	var x uint16 = 10
//	y := "test"
//	fmt.Println(IsUint16Type(x)) // true
//	fmt.Println(IsUint16Type(y)) // false
//
// Note: This function does not handle nil pointers and can panic if a
// nil pointer is passed in. Always check for nil before passing pointers.
func IsUint16Type(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.Uint16
}

// IsUint32Type checks if a given value is of an uint32 type.
// It uses the reflect package
// to check for the kind of the given value and returns a boolean indicating if it is uint32 or not.
//
// Parameters:
//   - a: The interface value to be checked for the uint32 type.
//
// Returns:
//   - bool: A boolean value indicating whether the provided value is of an uint32 type.
//
// Example:
//
//	var a uint32 = 10
//	var b string = "hello"
//	fmt.Println(IsUint32Type(a)) // true
//	fmt.Println(IsUint32Type(b)) // false
//
// Note: This function does not handle nil pointers and can panic if a
// nil pointer is passed in. Always check for nil before passing pointers.
func IsUint32Type(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.Uint32
}

// IsUint64Type verifies if a given value is of the uint64 type. It uses reflection to
// fetch the type of 'a' and then check if it equals reflect.Uint64.
//
// Parameters:
//   - a: Any interface value to be checked for its type.
//
// Returns:
//   - bool: A boolean value indicating whether the value type is uint64.
//
// Example:
//
//	var x uint64 = 10
//	y := "test"
//	fmt.Println(IsUint64Type(x)) // Output: true
//	fmt.Println(IsUint64Type(y)) // Output: false
//
// Note: This function does not handle nil pointers and can panic if a
// nil pointer is passed in. Always check for nil before passing pointers.
func IsUint64Type(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.Uint64
}

// IsFloat32Type checks whether the given value is of type float32. This function
// makes use of Go's reflection package to inspect the type of the given value.
//
// Parameters:
//   - a: The value to be checked for type float32.
//
// Returns:
//   - bool: A boolean value indicating whether the type of value is float32.
//
// Example:
//
//	var x int = 5
//	var y float32 = 10.5
//	fmt.Println(IsFloat32Type(x)) // false
//	fmt.Println(IsFloat32Type(y)) // true
//
// Note: This function does not handle nil pointers and can panic if a
// nil pointer is passed in. Always check for nil before passing pointers.
func IsFloat32Type(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.Float32
}

// IsFloat64Type checks if the given value type is Float64. It uses the reflect package's
// ValueOf and Kind methods to inspect the type of the value and verify if it corresponds to Float64.
//
// Parameters:
//   - a: Any interface value whose type has to be checked.
//
// Returns:
//   - bool: A boolean value indicating whether the type of the value is Float64.
//
// Example:
//
//	var a float64 = 10.0
//	var b int = 20
//	fmt.Println(IsFloat64Type(a)) // true
//	fmt.Println(IsFloat64Type(b)) // false
//
// Note: This function does not handle nil pointers and can panic if a
// nil pointer is passed in. Always check for nil before passing pointers.
func IsFloat64Type(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.Float64
}

// IsBoolType determines whether a given value is of a bool type.
// It uses the Kind method of reflect Value to check the kind of the given value and compares it
// against reflect.Bool.
//
// Parameters:
//   - a: Any interface value to be checked for a bool type.
//
// Returns:
//   - bool: A Boolean value indicating whether the given value is of a bool type.
//
// Example:
//
//	var x *int
//	y := true
//	fmt.Println(IsBoolType(x)) // false
//	fmt.Println(IsBoolType(y)) // true
//
// Note: This function does not handle nil pointers and can panic if a
// nil pointer is passed in. Always check for nil before passing pointers.
func IsBoolType(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.Bool
}

// IsTimeType determines whether a given value is of time.Time type. It uses the reflect.TypeOf function
// to check the type of the value and compares it with a type of time.Time.
//
// Parameters:
//   - a: Any interface value to be checked for time.Time type.
//
// Returns:
//   - bool: A boolean value indicating whether the value is of time.Time type.
//
// Example:
//
//	x := time.Now()
//	y := "I am not a time"
//	fmt.Println(IsTimeType(x)) // true
//	fmt.Println(IsTimeType(y)) // false
//
// Note: This function does not handle nil pointers and can panic if a
// nil pointer is passed in. Always check for nil before passing pointers.
func IsTimeType(a any) bool {
	return reflect.TypeOf(a) == reflect.TypeOf(time.Time{})
}

// IsDurationType checks whether the given value is of the time.Duration type.
// It performs a comparison between the type of the provided value and the type
// of the time.Duration using the reflect.TypeOf function.
//
// Parameters:
//   - a: Any interface value to be checked for its type.
//
// Returns:
//   - bool: A boolean value indicating whether the value is of the type time.Duration.
//
// Example:
//
//	var dur = time.Duration(5 * time.Second)
//	var intVal = 10
//	fmt.Println(IsDurationType(dur)) // true
//	fmt.Println(IsDurationType(intVal)) // false
//
// Note: This function does not handle nil pointers and can panic if a
// nil pointer is passed in. Always check for nil before passing pointers.
func IsDurationType(a any) bool {
	return reflect.TypeOf(a) == reflect.TypeOf(time.Duration(0))
}

// IsBytesType checks if the given value is of type []byte. It uses the reflect.TypeOf function
// to check the specific type of the value and compares it with reflect.TypeOf([]byte{}).
//
// Parameters:
//   - a: Any interface value to be checked for type []byte.
//
// Returns:
//   - bool: A boolean value indicating whether the value is of type []byte.
//
// Example:
//
//	x := []byte{'a', 'b', 'c'}
//	y := "abc"
//	fmt.Println(IsBytesType(x)) // true
//	fmt.Println(IsBytesType(y)) // false
//
// Note: This function does not handle nil pointers and can panic if a
// nil pointer is passed in. Always check for nil before passing pointers.
func IsBytesType(a any) bool {
	return reflect.TypeOf(a) == reflect.TypeOf([]byte{})
}

// IsErrorType checks whether the given value is of an error type.
// It does this by performing a type assertion of the value as error and returns the result
// of this operation. If the value is of error type, the function will return true,
// if not - false.
//
// Parameters:
//   - a: Any interface value to be checked for being of an error type.
//
// Returns:
//   - bool: A boolean value indicating whether the value is of an error type.
//
// Example:
//
//	e := errors.New("this is an error")
//	fmt.Println(IsErrorType(e)) // true
//
//	x := "this is not an error"
//	fmt.Println(IsErrorType(x)) // false
//
// Note: This function does not handle nil pointers and can panic if a
// nil pointer is passed in. Always check for nil before passing pointers.
func IsErrorType(a any) bool {
	_, ok := a.(error)
	return ok
}
