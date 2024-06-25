package checker

import (
	"encoding/json"
	"fmt"
	"reflect"
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
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(reflectValue.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return float64(reflectValue.Uint())
	case reflect.Float32, reflect.Float64:
		return reflectValue.Float()
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
// If the value is of a channel type, the function returns the length of the channel.
// If the value is of an interface or pointer type, the function recursively calls itself with the dereferenced value.
// If the value is not of a supported type, a panic is thrown.
// Returns: The length or size of the value as an integer.
// Panics: If the value is of unsupported types or if the channel, interface, or pointer is nil.
func toLength(a any) int {
	reflectValue := reflect.ValueOf(a)

	switch reflectValue.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return int(reflectValue.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return int(reflectValue.Uint())
	case reflect.Float32, reflect.Float64:
		return int(reflectValue.Float())
	case reflect.Struct:
		return reflectValue.NumField()
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map:
		return reflectValue.Len()
	case reflect.Complex64, reflect.Complex128:
		return int(real(reflectValue.Complex()))
	case reflect.Chan:
		if reflectValue.IsNil() {
			panic("Error getting the channel size, it is null!")
		} else {
			return reflectValue.Len()
		}
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
		return strconv.FormatInt(reflectValue.Int(), 64)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(reflectValue.Uint(), 64)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(reflectValue.Float(), 'g', -1, 64)
	case reflect.Complex64, reflect.Complex128:
		return strconv.FormatComplex(reflectValue.Complex(), 'g', -1, 64)
	case reflect.Bool:
		return strconv.FormatBool(reflectValue.Bool())
	case reflect.Array, reflect.Slice, reflect.Map, reflect.Struct:
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
// If the value is of a numeric type (int, uint, float), it is converted to a Unix timestamp using time.Unix function.
// If the value is of a string type, multiple time layouts are tried using time.Parse function.
// If the value is not of a numeric or string type, an error is returned.
//
// Returns: The converted time.Time value and a possible error.
func toTimeWithErr(a any) (time.Time, error) {
	reflectValue := reflect.ValueOf(a)
	switch reflectValue.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return time.Unix(reflectValue.Int(), 0), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return time.Unix(int64(reflectValue.Uint()), 0), nil
	case reflect.Float32, reflect.Float64:
		return time.Unix(int64(reflectValue.Float()), 0), nil
	case reflect.String:
		layouts := []string{time.ANSIC, time.UnixDate, time.RubyDate, time.RFC822, time.RFC822Z, time.RFC850,
			time.RFC1123, time.RFC1123Z, time.RFC3339, time.Kitchen, time.Stamp}
		for _, layout := range layouts {
			if t, err := time.Parse(layout, reflectValue.String()); err == nil {
				return t, nil
			}
		}
		return time.Time{}, fmt.Errorf("cannot convert string to time.Time: Unknown format \"%s\"",
			reflectValue.String())
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
