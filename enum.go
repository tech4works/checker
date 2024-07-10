package checker

import "reflect"

// BaseEnum is an interface that defines a method IsEnumValid.
//
// It is used in conjunction with the BaseEnum interface to validate enum values.
// The IsEnumValid method should return a boolean value indicating if the enum is valid.
type BaseEnum interface {
	// IsEnumValid is a method that returns a boolean value indicating if the enum is valid.
	//
	// It is used in conjunction with the BaseEnum interface to validate enum values.
	IsEnumValid() bool
}

// IsEnumValid checks the validity of an enumerated variable (enum). It uses reflection to unwrap any pointers to enums,
// and then calls the IsEnumValid method of the actual enum (assuming it implements BaseEnum interface).
//
// Parameters:
//   - a: The variable to be checked whether it is a valid enum or not. It can either be a direct enum or pointer to an enum.
//
// Returns:
//   - bool: A boolean value indicating whether the provided variable is a valid enum or not.
//
// Example usage:
//
//	type MyEnum int
//
//	func (e MyEnum) IsEnumValid() bool {
//	  switch e {
//	  case 0, 1, 2:
//	  	return true
//	  default:
//	  	return false
//	  }
//	}
//
//	func main() {
//	   var a MyEnum = 2
//	   var b MyEnum = 3
//	   fmt.Println(IsEnumValid(a)) // true
//	   fmt.Println(IsEnumValid(b)) // false
//	}
func IsEnumValid(a any) bool {
	reflectValueA := reflect.ValueOf(a)
	if (reflectValueA.Kind() == reflect.Ptr || reflectValueA.Kind() == reflect.Interface) && !reflectValueA.IsNil() {
		return IsEnumValid(reflectValueA.Elem().Interface())
	}

	baseEnum, ok := reflectValueA.Interface().(BaseEnum)
	return ok && NonNil(baseEnum) && baseEnum.IsEnumValid()
}
