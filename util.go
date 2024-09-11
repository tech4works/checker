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
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"time"
)

// toFloat converts a value of any type to a float64.
// If the value is of a numeric type, it is directly converted to float64.
// If the value is of an interface or pointer type, the function recursively calls itself with the dereferenced value.
// If the value is not of a numeric, interface, or pointer type, a panic is thrown.
//
// Returns: The converted float64 value.
func toFloat(a any) float64 {
	reflectValue := reflect.ValueOf(a)

	switch reflectValue.Kind() {
	case reflect.String:
		f, err := strconv.ParseFloat(reflectValue.String(), 64)
		if err != nil {
			panic(fmt.Sprintf("Error getting float by string: %s", reflectValue.String()))
		}
		return f
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(reflectValue.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return float64(reflectValue.Uint())
	case reflect.Float32, reflect.Float64:
		return reflectValue.Float()
	case reflect.Complex64, reflect.Complex128:
		c := reflectValue.Complex()
		return real(c) + imag(c)
	case reflect.Interface, reflect.Pointer:
		if reflectValue.IsNil() {
			panic("Error convert interface/pointer to float, it is null!")
		} else {
			return toFloat(reflectValue.Elem().Interface())
		}
	default:
		panic(fmt.Sprintf("Error getting float, type %s not supported!", reflectValue.Kind().String()))
	}
}

// toLength converts a value of any type to its length or size as an integer.
// If the value is of a numeric type, the function returns the integer value of the numeric type.
// If the value is of a struct type, the function returns the number of fields in the struct.
// If the value is of a string, array, slice, or map type, the function returns the length of the string, array, slice, or map.
// If the value is of a complex type, the function returns the integer value of the real part of the complex number.
// If the value is of an interface or pointer type, the function recursively calls itself with the dereferenced value.
// If the value is not of a supported type, a panic is thrown.
// Returns: The length or size of the value as an integer.
// Panics: If the value is of unsupported types or if the channel, interface, or pointer is nil.
func toLength(a any) int {
	reflectValue := reflect.ValueOf(a)

	switch reflectValue.Kind() {
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map:
		return reflectValue.Len()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return int(reflectValue.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return int(reflectValue.Uint())
	case reflect.Float32, reflect.Float64:
		return int(reflectValue.Float())
	case reflect.Struct:
		return reflectValue.NumField()
	case reflect.Complex64, reflect.Complex128:
		return int(real(reflectValue.Complex()))
	case reflect.Interface, reflect.Pointer:
		if reflectValue.IsNil() {
			panic("Error getting the interface/pointer size, it is null!")
		} else {
			return toLength(reflectValue.Elem().Interface())
		}
	default:
		panic(fmt.Sprintf("Error getting %s size, type not supported!", reflectValue.Kind().String()))
	}
}

// toString converts a value of any type to a string.
// If the value is of a string type, it is directly returned as a string.
// If the value is of a numeric type (int, uint, float, complex), it is converted to a string using
// strconv package functions: strconv.FormatInt, strconv.FormatUint, strconv.FormatFloat, strconv.FormatComplex.
// If the value is of a bool type, it is converted to a string using strconv.FormatBool.
// If the value is of an array, slice, map, or struct type, it is marshaled to JSON using json.Marshal
// and then converted to a string.
// If the value is of an interface or pointer type, the function recursively calls itself with the dereferenced value.
// If the value is not of a string, numeric, bool, array, slice, map, struct, interface, or pointer type,
// a panic is thrown.
//
// Returns: The converted string value.
func toString(a any) string {
	reflectValue := reflect.ValueOf(a)
	switch reflectValue.Kind() {
	case reflect.String:
		return reflectValue.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(reflectValue.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(reflectValue.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(reflectValue.Float(), 'g', -1, 64)
	case reflect.Complex64, reflect.Complex128:
		return strconv.FormatComplex(reflectValue.Complex(), 'g', -1, 64)
	case reflect.Bool:
		return strconv.FormatBool(reflectValue.Bool())
	case reflect.Array, reflect.Slice:
		if reflectValue.Type().Elem().Kind() == reflect.Uint8 {
			return string(reflectValue.Bytes())
		} else {
			marshal, _ := json.Marshal(reflectValue.Interface())
			return string(marshal)
		}
	case reflect.Map, reflect.Struct:
		marshal, _ := json.Marshal(reflectValue.Interface())
		return string(marshal)
	case reflect.Ptr, reflect.Interface:
		if reflectValue.IsNil() {
			panic("Error getting a string, it is null!")
		}
		return toString(reflectValue.Elem().Interface())
	default:
		panic(fmt.Sprintf("Error getting a string, unsupported type %s!", reflectValue.Kind().String()))
	}
}

// toBytes converts a value of any type to a byte slice.
// It first converts the value to a string using the toString function,
// and then converts the string to a byte slice using the []byte type conversion.
// If the value is not convertible to a string, a panic is thrown.
//
// Returns: The converted byte slice value.
func toBytes(a any) []byte {
	return []byte(toString(a))
}

// toTimeWithErr converts a value of any type to a time.Time value and returns it along with an error.
// If the value is of a numeric type (int, uint, float), it is converted to a UnixMilli timestamp using
// time.UnixMilli function.
// If the value is of a string type, multiple time layouts are tried using time.Parse function.
// If the value is not of a numeric or string type, an error is returned.
//
// Returns: The converted time.Time value and a possible error.
func toTimeWithErr(a any) (time.Time, error) {
	reflectValue := reflect.ValueOf(a)
	switch reflectValue.Kind() {
	case reflect.String:
		layouts := []string{time.Layout, time.ANSIC, time.UnixDate, time.RubyDate, time.RFC822, time.RFC822Z,
			time.RFC850, time.RFC1123, time.RFC1123Z, time.RFC3339, time.RFC3339Nano, time.Kitchen, time.Stamp,
			time.DateTime, time.DateOnly, time.TimeOnly}
		for _, layout := range layouts {
			if t, err := time.Parse(layout, reflectValue.String()); err == nil {
				return t, nil
			}
		}
		return time.Time{}, fmt.Errorf("cannot convert string to time.Time: Unknown format \"%s\"",
			reflectValue.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return time.UnixMilli(reflectValue.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return time.UnixMilli(int64(reflectValue.Uint())), nil
	case reflect.Float32, reflect.Float64:
		return time.UnixMilli(int64(reflectValue.Float())), nil
	default:
		if reflectValue.Type() == reflect.TypeOf(time.Time{}) {
			return reflectValue.Interface().(time.Time), nil
		}
		return time.Time{}, fmt.Errorf("cannot convert to time.Time from type: %s", reflectValue.Kind().String())
	}
}

// toTime converts a value of any type to a time.Time value.
// It calls toTimeWithErr with the given value and handles the error.
//
// Returns: The converted time.Time value.
func toTime(a any) time.Time {
	t, err := toTimeWithErr(a)
	if err != nil {
		panic(err)
	}
	return t
}

// toDate converts a value of any type to a time.Time value by calling toTime and adjusting it to midnight.
//
// Returns: The converted time.Time value.
func toDate(a any) time.Time {
	t := toTime(a)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// timeNow returns the current local time as a time.Time value.
func timeNow() time.Time {
	return time.Now()
}

// dateNow returns the current date as a time.Time value with the time components set to 0.
// The function uses the time.Now() function to get the current time and then constructs
// a new time.Time value with the same year, month, and day as the current time but with
// the time components (hour, minute, second, nanosecond) set to 0. The location of the
// new time value is set to the same location as the current time.
//
// Returns: The current date as a time.Time value with the time components set to 0.
func dateNow() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

// removeNonDigits removes all non-digit characters from the given string.
// It uses regular expressions to find and replace non-digit characters with an empty string.
// Returns the modified string with only digit characters remaining.
func removeNonDigits(input string) string {
	regex, _ := regexp.Compile(`[^0-9]`)
	return regex.ReplaceAllString(input, "")
}

// allDigitsEqual checks if all characters in the input string are equal.
// It iterates over the string and compares each character to the first character.
// If any character is different, the function returns false.
// Returns: Boolean value indicating if all characters in the input string are equal.
func allDigitsEqual(input string) bool {
	for i := 1; i < len(input); i++ {
		if input[i] != input[0] {
			return false
		}
	}
	return true
}

// calculateVerifierDigits calculates the verifier digits for a given document using the provided weights.
// It iterates over the document string and multiplies each digit by its corresponding weight from weights1 and weights2.
// The sums of the products are then used to calculate the verifier digits.
// The first verifier digit is calculated as the modulo of sum1 by 11.
// If the result is less than 2, the first verifier digit is set to 0, otherwise it is set to 11 minus the result.
// The second verifier digit is calculated in the same way using sum2.
//
// Parameters:
//   - document: The document string for which the verifier digits calculated.
//   - weights1: The weights for the first verifier digit calculation.
//   - weights2: The weights for the second verifier digit calculation.
//
// Returns:
//   - int: The calculated first verifier digit.
//   - int: The calculated second verifier digit.
func calculateVerifierDigits(document string, weights1, weights2 []int) (int, int) {
	sum1, sum2 := 0, 0
	for i := 0; i < len(weights1); i++ {
		num, _ := strconv.Atoi(string(document[i]))
		sum1 += num * weights1[i]
		sum2 += num * weights2[i]
	}
	num, _ := strconv.Atoi(string(document[len(weights1)]))
	sum2 += num * weights2[len(weights1)]

	firstVerifier := sum1 % 11
	if firstVerifier < 2 {
		firstVerifier = 0
	} else {
		firstVerifier = 11 - firstVerifier
	}

	secondVerifier := sum2 % 11
	if secondVerifier < 2 {
		secondVerifier = 0
	} else {
		secondVerifier = 11 - secondVerifier
	}

	return firstVerifier, secondVerifier
}
