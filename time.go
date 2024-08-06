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

// IsBeforeNow determines whether a given time is before the current time. It uses
// the toTime function to convert the provided value to a time.Time object, and
// compares the result with the current time (obtained via timeNow). If the
// converted time is before the current time, IsBeforeNow returns true.
//
// Parameters:
//   - a: Any value to be converted into a time.Time object for comparison.
//
// Returns:
//   - bool: A boolean value indicating whether the provided time is before
//     the current time.
//
// Panic:
//
//	This function will panic if the provided value cannot be converted to a time.Time
//	object through the toTime() function.
//
// Example:
//
//	var a int64 = 1609459200 // Unix timestamp for January 1, 2021 00:00:00
//	var b string = "2045-04-23T18:25:43.511Z" // RFC3339 formatted string
//	fmt.Println(IsBeforeNow(a)) // Will return true if the current time is after January 1, 2021 00:00:00
//	fmt.Println(IsBeforeNow(b)) // Will return true if the current time is before April 23, 2045 18:25:43.511 UTC
func IsBeforeNow(a any) bool {
	return toTime(a).Before(timeNow())
}

// IsBeforeToday determines whether a given time is before the current date without considering time components. It uses
// the toDate function to convert the provided value to a time.Time object and adjusts it to midnight,
// then compares the result with the current date (obtained via dateNow). If the converted date is before the current date,
// IsBeforeToday returns true.
//
// Parameters:
//   - a: Any value to be converted into a time.Time object and compared with the current date.
//
// Returns:
//   - bool: A boolean value indicating whether the provided time is before the current date.
//
// Panics:
//
//	This function will panic if the provided value cannot be converted to a time.Time
//	object through the toDate() function.
//
// Example:
//
//	var a int64 = 1609459200 // Unix timestamp for January 1, 2021
//	var b string = "2045-04-23T18:25:43.511Z" // RFC3339 formatted string
//	fmt.Println(IsBeforeToday(a)) // Will return true if the current date is after January 1, 2021
//	fmt.Println(IsBeforeToday(b)) // Will return true if the current date is before April 23, 2045
func IsBeforeToday(a any) bool {
	return toDate(a).Before(dateNow())
}

// IsBeforeDate determines whether the date represented by the first argument is before the date
// represented by the second argument. It uses the toDate function to convert the arguments to
// time.Time values and then compares them using the Before method of the time.Time type.
//
// Parameters:
//   - a: The first value to be converted to a time.Time value and compared. It can be of any type.
//   - b: The second value to be converted to a time.Time value and compared. It can be of any type.
//
// Returns:
//   - bool: A boolean value indicating whether the date represented by the first argument is before the date
//     represented by the second argument.
//
// Panics:
//   - If the conversion to a time.Time value fails for either of the arguments, a panic is caused.
//
// Example:
//
//	t1 := "2006-01-02"
//	t2 := time.Now()
//	fmt.Println(IsBeforeDate(t1, t2)) // true if t1 is before the current date
//	fmt.Println(IsBeforeDate(t2, t1)) // false if the current date is before t1
func IsBeforeDate(a, b any) bool {
	return toDate(a).Before(toDate(b))
}

// IsBefore determines whether a given time, a, is before another given time, b.
// It uses the toTime function to convert both values into time.Time objects,
// and then compares them using the Before method. If a is before b, IsBefore returns true.
//
// Parameters:
//   - a: Any value to be converted into a time.Time object for comparison.
//   - b: Any value to be converted into a time.Time object for comparison.
//
// Returns:
//   - bool: A boolean value indicating whether the time a is before the time b.
//
// Panics:
//
//	This function will panic if either of the provided values cannot be converted
//	to time.Time objects through the toTime() function.
//
// Example:
//
//	var t1 int64 = 1577836800       // Unix timestamp for January 1, 2020 00:00:00
//	var t2 string = "2022-01-01"    // Date in "2006-01-02" format
//	var t3 time.Time = time.Now()    // Current time
//
//	fmt.Println(IsBefore(t1, t2))   // Will return true if January 1, 2020 is before January 1, 2022
//	fmt.Println(IsBefore(t2, t3))   // Will return true if the provided date is before the current time
//	fmt.Println(IsBefore(t1, t3))   // Will return true if January 1, 2020 is before the current time
func IsBefore(a, b any) bool {
	return toTime(a).Before(toTime(b))
}

// IsAfterNow determines whether a given time is after the current time. It employs the
// toTime method to convert the specified value into a time.Time instance and
// compares this with the current time (retrieved via timeNow). If the
// transformed time is later than the current time, IsAfterNow returns true.
//
// Parameters:
//   - a: Any value to be converted into a time.Time instance for comparison.
//
// Returns:
//   - bool: A boolean value indicating whether the provided time is later than
//     the present time.
//
// Panics:
//
//	This function will panic if the specified value cannot be converted into a time.Time
//	instance using the toTime() function.
//
// Example:
//
//	t := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
//	fmt.Println(IsAfterNow(t)) // Example output: true
//	fmt.Println(IsAfterNow(time.Now().Add(-time.Hour))) // Example output: false
func IsAfterNow(a any) bool {
	return toTime(a).After(timeNow())
}

// IsAfterToday determines whether a given time is after today. It uses the toDate function
// to convert the provided value to a time.Time object at midnight, and compares the result
// with the current date (obtained via dateNow function). If the converted time is after
// today, IsAfterToday returns true.
// Parameters:
//   - a: Any value to be converted into a time.Time object for comparison.
//
// Returns:
//   - bool: A boolean value indicating whether the provided time is after today.
//
// Panics:
//
//	This function will panic if the provided value cannot be converted to a time.Time
//	object through the toDate() function.
//
// Example:
//
//	var a int64 = 1614559200 // Unix timestamp for February 28, 2021 18:00:00
//	var b string = "2021-02-28T18:00:00Z" // RFC3339 formatted string
//	fmt.Println(IsAfterToday(a)) // Will return true if the current date is after February 28, 2021
//	fmt.Println(IsAfterToday(b)) // Will return false if the current date is before or on February 28, 2021
func IsAfterToday(a any) bool {
	return toDate(a).After(dateNow())
}

// IsAfterDate determines whether the first provided date is after the second provided date.
// It uses the toDate function to convert both values into time.Time objects and compares them.
// If the first date is after the second date, IsAfterDate returns true.
// Parameters:
//   - a: Any value to be converted into a time.Time object for comparison.
//   - b: Any value to be converted into a time.Time object for comparison.
//
// Returns:
//   - bool: A boolean value indicating whether the first date is after the second date.
//
// Panics:
//   - This function will panic if either of the provided values cannot be converted to a time.Time
//     object through the toDate() function.
//
// Example:
//
//	var a int64 = 1609459200 // Unix timestamp for January 1, 2021 00:00:00
//	var b string = "2045-04-23T18:25:43.511Z" // RFC3339 formatted string
//	fmt.Println(IsAfterDate(a, b)) // Will return true if January 1, 2021 00:00:00 is after April 23, 2045 18:25:43.511 UTC
func IsAfterDate(a, b any) bool {
	return toDate(a).After(toDate(b))
}

// IsAfter determines whether the first provided time is after the second provided time.
// It converts the two values into time.Time objects using the toTime function and
// compares them using the After method of time.Time. If the first time is after
// the second time, IsAfter returns true.
//
// Parameters:
//   - a: First value to be converted into a time.Time object for comparison.
//   - b: Second value to be converted into a time.Time object for comparison.
//
// Returns:
//   - bool: A boolean value indicating whether the first time is after the second time.
//
// Panics:
//   - This function will panic if either of the provided values cannot be converted
//     to time.Time objects through the toTime function.
//
// Example:
//
//	var a int64 = 1609459200 // Unix timestamp for January 1, 2021 00:00:00
//	var b string = "2045-04-23T18:25:43.511Z" // RFC3339 formatted string
//	fmt.Println(IsAfter(a, b)) // Will return true if the first time is after April 23, 2045 18:25:43.511 UTC
func IsAfter(a, b any) bool {
	return toTime(a).After(toTime(b))
}

// IsToday checks whether the provided value represents the current date. This function
// converts the passed value to a time.Time format by calling toDate function and then
// compares that date with the current date using Equal function from time package.
//
// Parameters:
//   - a: This parameter can be any value that can be converted into a time.Time format.
//
// Returns:
//   - bool: This function returns true if the date represented by input parameter is the
//     current date. Otherwise, it returns false.
//
// Panic:
//
//	This function will panic if it's unable to convert the provided value into a time.Time format.
//
// Example:
//
//	t1 := time.Now()
//	fmt.Println(IsToday(t1)) // true
//
//	t2 := time.Now().Add(-24 * time.Hour)
//	fmt.Println(IsToday(t2)) // false
func IsToday(a any) bool {
	return toDate(a).Equal(dateNow())
}
