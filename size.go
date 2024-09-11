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

// IsGreaterThan compares two values of any type and returns whether the first value is greater than the second value.
// If the supplied values are not of a numeric type, a panic is thrown.
//
// Example usage:
//
//	fmt.Println(IsGreaterThan(5, 4))        // Outputs: true
//	fmt.Println(IsGreaterThan(-2, -3))      // Outputs: true
//	fmt.Println(IsGreaterThan(0, 4))        // Outputs: false
//	fmt.Println(IsGreaterThan(0, -1))       // Outputs: true
//	fmt.Println(IsGreaterThan(5, 5))        // Outputs: false
//	fmt.Println(IsGreaterThan(5.5, 5.4))    // Outputs: true
//
// Returns true if the first value is greater than the second value, false otherwise.
func IsGreaterThan(a, b any) bool {
	return toFloat(a) > toFloat(b)
}

// IsGreaterThanOrEqual compares two values of any type and returns whether the first value is greater than or
// equal to the second value.
// It calls the IsGreaterThan function and the Equals function to perform the comparison.
//
// Example usage:
//
//	fmt.Println(IsGreaterThanOrEqual(5, 4))        // Outputs: true
//	fmt.Println(IsGreaterThanOrEqual(-2, -3))      // Outputs: true
//	fmt.Println(IsGreaterThanOrEqual(0, 4))        // Outputs: false
//	fmt.Println(IsGreaterThanOrEqual(0, -1))       // Outputs: true
//	fmt.Println(IsGreaterThanOrEqual(5, 5))        // Outputs: true
//	fmt.Println(IsGreaterThanOrEqual(5.5, 5.4))    // Outputs: true
//
// Returns true if the first value is greater than or equal to the second value, false otherwise.
func IsGreaterThanOrEqual(a, b any) bool {
	return toFloat(a) >= toFloat(b)
}

// IsLessThan compares two values of any type and returns whether the first value is less than the second value.
// It does this by using the IsGreaterThan function and negating its result.
//
// Example usage:
//
//	fmt.Println(IsLessThan(4, 5))        // Outputs: true
//	fmt.Println(IsLessThan(-2, -3))      // Outputs: false
//	fmt.Println(IsLessThan(0, 4))        // Outputs: true
//	fmt.Println(IsLessThan(0, -1))       // Outputs: false
//	fmt.Println(IsLessThan(5, 5))        // Outputs: false
//	fmt.Println(IsLessThan(5.4, 5.5))    // Outputs: true
//
// Returns true if the first value is less than the second value, false otherwise.
func IsLessThan(a, b any) bool {
	return toFloat(a) < toFloat(b)
}

// IsLessThanOrEqual compares two values of any type and returns whether the first value is less than or equal to the
// second value. It does this by using the IsLessThan function and the Equals function.
//
// Example usage:
//
//	fmt.Println(IsLessThanOrEqual(4, 5))        // Outputs: true
//	fmt.Println(IsLessThanOrEqual(-2, -3))      // Outputs: false
//	fmt.Println(IsLessThanOrEqual(0, 4))        // Outputs: true
//	fmt.Println(IsLessThanOrEqual(0, -1))       // Outputs: false
//	fmt.Println(IsLessThanOrEqual(5, 5))        // Outputs: true
//	fmt.Println(IsLessThanOrEqual(5.4, 5.5))    // Outputs: true
//
// Returns true if the first value is less than or equal to the second value, false otherwise.
// The function is composed of the IsLessThan and the Equals functions.
func IsLessThanOrEqual(a, b any) bool {
	return toFloat(a) <= toFloat(b)
}

// IsLengthEquals checks whether the length or size of two parameters, a and b, are equal.
// It uses the toLen function to get the length of the parameters, which supports various data types.
//
// Example usage:
//
//	fmt.Println(IsLengthEquals("test", "test")) // Outputs: true
//	fmt.Println(IsLengthEquals("test", 4)) // Outputs: true
//	fmt.Println(IsLengthEquals("test", 1)) // Outputs: false
//
// Returns true if the length or size of a and b are equal, false otherwise.
// Panic occurs if a and b are of unsupported types or if the channel, interface, or pointer is nil.
func IsLengthEquals(a, b any) bool {
	return toLength(a) == toLength(b)
}

// IsLengthNotEquals checks if the lengths or sizes of two parameters, a and b, are not equal.
// It uses the toLen function to get the lengths of the parameters, which supports various data types.
// Unlike IsLengthEquals, this function returns true if the lengths or sizes of a and b do not match.
//
// Parameters:
//   - a: The first parameter to compare the length. It can be of any type supported by the toLen function.
//   - b: The second parameter to compare the length. It must be of a type similar to a for the comparison to be relevant.
//
// Returns:
//   - bool: A boolean value indicating whether the lengths/sizes of the two parameters do not match.
//
// Panics:
//
//	May panic if the parameters are of unsupported types or if the channel, interface, or pointer is nil.
//
// Example:
//
//	var a string = "test"
//	var b string = "hello"
//	fmt.Println(IsLengthNotEquals(a,b)) // true
//
//	var x int = 3
//	var y int = 4
//	fmt.Println(IsLengthNotEquals(x,y)) // true
func IsLengthNotEquals(a, b any) bool {
	return !IsLengthEquals(a, b)
}

// IsLengthGreaterThan compares the length or size of two values of any type and returns whether the first value is
// greater than the second value.
// It uses the toLen function to get the length of the values, which supports various data types.
// If the supplied values are of unsupported types or if the channel, interface, or pointer is nil, a panic is thrown.
//
// Example usage:
//
//	fmt.Println(IsLengthGreaterThan("test", "hello"))     // Outputs: false
//	fmt.Println(IsLengthGreaterThan("test", []int{1, 2})) // Outputs: true
//	fmt.Println(IsLengthGreaterThan("test", 2))           // Outputs: true
//	fmt.Println(IsLengthGreaterThan([]int{1, 2}, "test")) // Outputs: false
//	fmt.Println(IsLengthGreaterThan(3, "test")) 		  // Outputs: false
//
// Returns true if the length or size of the first value is greater than the length or size of the second value,
// false otherwise.
func IsLengthGreaterThan(a, b any) bool {
	return toLength(a) > toLength(b)
}

// IsLengthGreaterThanOrEqual compares the length or size of two values of any type and returns whether the first value
// is greater than or equal to the second value.
// It uses the IsLengthGreaterThan and IsLengthEquals functions to determine the result.
//
// Example usage:
//
//	fmt.Println(IsLengthGreaterThanOrEqual("test", "hello"))     // Outputs: true
//	fmt.Println(IsLengthGreaterThanOrEqual("test", []int{1, 2})) // Outputs: true
//	fmt.Println(IsLengthGreaterThanOrEqual("test", 2))           // Outputs: true
//	fmt.Println(IsLengthGreaterThanOrEqual("test", []int{1,2,3,4}))  // Outputs: true
//	fmt.Println(IsLengthGreaterThanOrEqual("test", "test"))  // Outputs: true
//	fmt.Println(IsLengthGreaterThanOrEqual([]int{1, 2}, "test")) // Outputs: false
//	fmt.Println(IsLengthGreaterThanOrEqual(3, "test")) 		   // Outputs: false
//
// Returns true if the length or size of the first value is greater than or equal to the length or size of the second value,
// false otherwise.
func IsLengthGreaterThanOrEqual(a, b any) bool {
	return IsLengthGreaterThan(a, b) || IsLengthEquals(a, b)
}

// IsLengthLessThan compares the length or size of two parameters, a and b, and returns whether the length of an is
// less than the length of b. It uses the toLen function to get the length of the parameters, which supports various
// data types.
//
// Example usage:
//
//	fmt.Println(IsLengthLessThan("test", "testexample")) // Outputs: true
//	fmt.Println(IsLengthLessThan("test", []int{1, 2, 3, 4, 5})) // Outputs: true
//	fmt.Println(IsLengthLessThan("test", 12)) // Outputs: true
//	fmt.Println(IsLengthLessThan("test", 4)) // Outputs: false
//	fmt.Println(IsLengthLessThan("test", 1)) // Outputs: false
//
// Returns true if the length or size of an is less than the length of b, false otherwise.
// Panic occurs if a and b are of unsupported types or if the channel, interface, or pointer is nil.
func IsLengthLessThan(a, b any) bool {
	return toLength(a) < toLength(b)
}

// IsLengthLessThanOrEqual compares the length or size of two parameters, a and b, and returns
// whether the length of an is less than or equal to the length of b. It uses the IsLengthLessThan and IsLengthEquals
// functions to determine the comparison.
//
// Example usage:
//
//	fmt.Println(IsLengthLessThanOrEqual("test", "testexample"))         // Outputs: true
//	fmt.Println(IsLengthLessThanOrEqual("test", []int{1, 2, 3, 4, 5})) // Outputs: true
//	fmt.Println(IsLengthLessThanOrEqual("test", 12))                    // Outputs: true
//	fmt.Println(IsLengthLessThanOrEqual("test", 4))                     // Outputs: false
//	fmt.Println(IsLengthLessThanOrEqual("test", 1))                     // Outputs: false
//
// Returns true if the length or size of an is less than or equal to the length of b, false otherwise.
// Panic occurs if a and b are of unsupported types or if the channel, interface, or pointer is nil.
func IsLengthLessThanOrEqual(a, b any) bool {
	return IsLengthLessThan(a, b) || IsLengthEquals(a, b)
}
