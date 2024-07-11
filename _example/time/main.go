package main

import (
	"fmt"
	"github.com/tech4works/checker"
	"time"
)

func main() {
	var a = 1609459200                 // Unix timestamp for January 1, 2021 00:00:00
	var b = "2045-04-23T18:25:43.511Z" // RFC3339 formatted string
	var c = 1614559200                 // Unix timestamp for February 28, 2021 18:00:00
	var d = "2021-02-28T18:00:00Z"     // RFC3339 formatted string
	var t1 = 1577836800                // Unix timestamp for January 1, 2020 00:00:00
	var t2 = "2022-01-01"              // Date in "2006-01-02" format
	var t3 = time.Now()                // Current time

	fmt.Println("IsBeforeNow results:")
	fmt.Println(checker.IsBeforeNow(a)) // Will return true if the current time is after January 1, 2021 00:00:00
	fmt.Println(checker.IsBeforeNow(b)) // Will return true if the current time is before April 23, 2045 18:25:43.511 UTC

	fmt.Println("IsBeforeToday results:")
	fmt.Println(checker.IsBeforeToday(a)) // Will return true if the current date is after January 1, 2021
	fmt.Println(checker.IsBeforeToday(b)) // Will return true if the current date is before April 23, 2045

	fmt.Println("IsBeforeDate results:")
	fmt.Println(checker.IsBeforeDate(t1, t2)) // true if t1 is before t2
	fmt.Println(checker.IsBeforeDate(t2, t1)) // false if t2 is before t1

	fmt.Println("IsBefore results:")
	fmt.Println(checker.IsBefore(t1, t2)) // Will return true if January 1, 2020 is before January 1, 2022
	fmt.Println(checker.IsBefore(t2, t3)) // Will return true if the provided date is before the current time
	fmt.Println(checker.IsBefore(t1, t3)) // Will return true if January 1, 2020 is before the current time

	fmt.Println("IsAfterNow results:")
	fmt.Println(checker.IsAfterNow("2023-01-01"))               // Example output: true
	fmt.Println(checker.IsAfterNow(time.Now().Add(-time.Hour))) // Example output: false

	fmt.Println("IsAfterToday results:")
	fmt.Println(checker.IsAfterToday(c)) // Will return true if the current date is after February 28, 2021
	fmt.Println(checker.IsAfterToday(d)) // Will return false if the current date is before or on February 28, 2021

	fmt.Println("IsAfterDate results:")
	fmt.Println(checker.IsAfterDate(t2, c)) // true if current date is after February 28, 2021
	fmt.Println(checker.IsAfterDate(t1, d)) // true if January 1, 2020 is after or on February 28, 2021

	fmt.Println("IsAfter results:")
	fmt.Println(checker.IsAfter(t1, t2)) // Will return false if January 1, 2020 is after January 1, 2022
	fmt.Println(checker.IsAfter(t2, t1)) // Will return true if January 1, 2022 is after January 1, 2020
	fmt.Println(checker.IsAfter(t3, t2)) // Will return true if current time is after January 1, 2022

	fmt.Println("IsToday results:")
	fmt.Println(checker.IsToday(t1)) // Will return false if the date is not today
	fmt.Println(checker.IsToday(t3)) // Will return true if the current date is today
	fmt.Println(checker.IsToday(t2)) // Will return true if the provided date is today's date
}
