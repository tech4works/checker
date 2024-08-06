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

import "reflect"

// Document represents a custom type for different types of documents.
type Document string

const (
	// DocumentCPF represents a constant of type Document that indicates a CPF document.
	DocumentCPF Document = "CPF"
	// DocumentCNPJ represents a constant of type Document that indicates a CNPJ document.
	DocumentCNPJ Document = "CNPJ"
)

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
