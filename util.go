package checker

import (
	"fmt"
	"reflect"
)

// isNumericType checks whether the given reflect.Value is a numeric type.
// Returns true if it is a numeric type, otherwise false.
func isNumericType(reflectValue reflect.Value) bool {
	switch reflectValue.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8,
		reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64:
		return true
	case reflect.Interface, reflect.Pointer:
		if reflectValue.IsNil() {
			return false
		} else {
			return isNumericType(reflectValue.Elem())
		}
	default:
		return false
	}
}

// isNotNumericType checks whether the given reflect.Value is not a numeric type.
// Returns true if it is not a numeric type, otherwise false.
func isNotNumericType(reflectValue reflect.Value) bool {
	return !isNumericType(reflectValue)
}

// toFloat64 converts the given reflect.Value to a float64 value.
// If the reflect.Value is of a numeric type, it returns the conversion.
// If the reflect.Value is of an interface or pointer type and not nil, it recursively calls itself with the
// reflect.Value's Elem() value.
// If the reflect.Value is nil, it panics with the error message "Error convert interface/pointer to float, it is null!".
// If the reflect.Value is of any other type, it panics with the error message "Error convert <kind> to float, type not supported!",
// where <kind> is the string representation of the reflect.Value's Kind.
// The function returns a float64 value.
func toFloat64(reflectValue reflect.Value) float64 {
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
			return toFloat64(reflectValue.Elem())
		}
	default:
		panic(fmt.Sprintf("Error convert %s to float, type not supported!", reflectValue.Kind().String()))
	}
}

// toLen returns the length or size of the given reflect.Value, which can be of various data types.
// It supports numeric types (int, uint, float), struct, string, array, slice, map, complex, channel,
// interface, and pointer.
// If the reflect.Value is a numeric type, it returns the integer representation of the value.
// If the reflect.Value is a struct, it returns the number of fields in the struct.
// If the reflect.Value is a string, array, slice, or map, it returns the length of the value.
// If the reflect.Value is a complex type, it returns the real part of the complex number as an integer.
// If the reflect.Value is a channel, it returns the number of elements in the channel.
// If the reflect.Value is an interface or pointer, it recursively calls toLen on the dereferenced value.
// If the reflect.Value is of an unsupported type, it panics with an error message.
// Panic occurs if a channel, interface, or pointer is nil.
func toLen(reflectValue reflect.Value) int {
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
			return toLen(reflectValue.Elem())
		}
	default:
		panic(fmt.Sprintf("Error getting %s size, type not supported!", reflectValue.Kind().String()))
	}
}
