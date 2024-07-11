package main

import (
	"fmt"
	"github.com/tech4works/checker"
	"time"
)

func main() {
	js := `{"key": "value", "key2": "value2", "key3": "value3"}`
	jsArray := `[1, 2, 3]`
	jsArrayOfMap := `[{"key": "value", "key2": "value2", "key3": "value3"}, {"key": "value", "key2": "value2", "key3": "value3"}]`
	text := "Not a JSON string"
	intStr := "10"
	boolStr := "true"
	boolInt := 1
	boolAbbreviation := 'T'
	floatStr := "10.5"
	timeStr := "2020-07-14T"
	timeStr2 := "2020-07-14"
	timeStr3 := "04:12:02Z"
	timeInt := 1609459200
	durationString := "2h45m"
	randomString := "abc123"
	bu1 := "10KB"
	bu2 := "20KB"
	bu3 := "300MB"
	ptr := &intStr
	f := func() {}
	c := make(chan int)
	m := map[string]int{"Alice": 23, "Bob": 24}
	st := struct {
		Name string
		Age  int
	}{}
	sl := []any{2, 3, "test", map[string]any{}}
	arr := [3]int{}
	intV := 10
	int8V := int8(10)
	int16V := int16(10)
	int32V := int32(10)
	int64V := int64(10)
	uintV := uint(10)
	uint8V := uint8(10)
	uint16V := uint16(10)
	uint32V := uint32(10)
	uint64V := uint64(10)
	f32V := float32(10.0)
	f64V := 20.0
	t := time.Now()
	d := 5 * time.Second
	ba := []byte{'a', 'b', 'c'}
	err := fmt.Errorf("this is an error")

	fmt.Println("IsJSON results:")
	fmt.Println(checker.IsJSON(js))      // Should return true.
	fmt.Println(checker.IsJSON(jsArray)) // Should return true.
	fmt.Println(checker.IsJSON(text))    // Should return false.

	fmt.Println("IsMap results:")
	fmt.Println(checker.IsMap(js))      // Should return true.
	fmt.Println(checker.IsMap(jsArray)) // Should return false.

	fmt.Println("IsSlice results:")
	fmt.Println(checker.IsSlice(jsArray)) // Should return true.
	fmt.Println(checker.IsSlice(js))      // Should return false.

	fmt.Println("IsSliceOfMaps results:")
	fmt.Println(checker.IsSliceOfMaps(jsArrayOfMap)) // Should return true
	fmt.Println(checker.IsSliceOfMaps(jsArray))      // Should return false

	fmt.Println("IsInt results:")
	fmt.Println(checker.IsInt(intStr)) // Should return true.
	fmt.Println(checker.IsInt(text))   // Should return false.

	fmt.Println("IsBool results:")
	fmt.Println(checker.IsBool(boolStr))          // Should return true.
	fmt.Println(checker.IsBool(boolInt))          // Should return true.
	fmt.Println(checker.IsBool(boolAbbreviation)) // Should return true.
	fmt.Println(checker.IsBool(text))             // Should return false.

	fmt.Println("IsFloat results:")
	fmt.Println(checker.IsFloat(floatStr)) // Should return true.
	fmt.Println(checker.IsFloat(text))     // Should return false.

	fmt.Println("IsTime results:")
	fmt.Println(checker.IsTime(timeStr))  // Should return true.
	fmt.Println(checker.IsTime(timeStr2)) // Should return true.
	fmt.Println(checker.IsTime(timeStr3)) // Should return true.
	fmt.Println(checker.IsTime(timeInt))  // Should return true.
	fmt.Println(checker.IsTime(text))     // Should return false.

	fmt.Println("IsDuration results:")
	fmt.Println(checker.IsDuration(durationString)) // Should return true.
	fmt.Println(checker.IsDuration(randomString))   // Should return false.

	fmt.Println("IsByteUnit results:")
	fmt.Println(checker.IsByteUnit(bu1))  // Should return true.
	fmt.Println(checker.IsByteUnit(bu2))  // Should return true.
	fmt.Println(checker.IsByteUnit(bu3))  // Should return true.
	fmt.Println(checker.IsByteUnit(text)) // Should return false.

	fmt.Println("IsPointerType results:")
	fmt.Println(checker.IsPointerType(ptr))  // Should return true.
	fmt.Println(checker.IsPointerType(text)) // Should return false.

	fmt.Println("IsFuncType results:")
	fmt.Println(checker.IsFuncType(f))    // Should return true.
	fmt.Println(checker.IsFuncType(text)) // Should return false.

	fmt.Println("IsChanType results:")
	fmt.Println(checker.IsChanType(c))    // Should return true.
	fmt.Println(checker.IsChanType(text)) // Should return false.

	fmt.Println("IsMapType results:")
	fmt.Println(checker.IsMapType(m))    // Should return true.
	fmt.Println(checker.IsMapType(text)) // Should return false.

	fmt.Println("IsStructType results:")
	fmt.Println(checker.IsStructType(st))   // Should return true.
	fmt.Println(checker.IsStructType(text)) // Should return false.

	fmt.Println("IsSliceType results:")
	fmt.Println(checker.IsSliceType(sl))   // Should return true.
	fmt.Println(checker.IsSliceType(text)) // Should return false.

	fmt.Println("IsArrayType results:")
	fmt.Println(checker.IsArrayType(arr))  // Should return true.
	fmt.Println(checker.IsArrayType(text)) // Should return false.

	fmt.Println("IsSliceOrArrayType results:")
	fmt.Println(checker.IsSliceOrArrayType(sl))   // Should return true.
	fmt.Println(checker.IsSliceOrArrayType(arr))  // Should return true.
	fmt.Println(checker.IsSliceOrArrayType(text)) // Should return false.

	fmt.Println("IsStringType results:")
	fmt.Println(checker.IsStringType(text)) // Should return false.
	fmt.Println(checker.IsStringType(intV)) // Should return true.

	fmt.Println("IsIntType results:")
	fmt.Println(checker.IsIntType(intV)) // Should return true.
	fmt.Println(checker.IsIntType(text)) // Should return false.

	fmt.Println("IsInt8Type results:")
	fmt.Println(checker.IsInt8Type(int8V)) // Should return true.
	fmt.Println(checker.IsInt8Type(intV))  // Should return false.

	fmt.Println("IsInt16Type results:")
	fmt.Println(checker.IsInt16Type(int16V)) // Should return true.
	fmt.Println(checker.IsInt16Type(intV))   // Should return false.

	fmt.Println("IsInt32Type results:")
	fmt.Println(checker.IsInt32Type(int32V)) // Should return true.
	fmt.Println(checker.IsInt32Type(intV))   // Should return false.

	fmt.Println("IsInt64Type results:")
	fmt.Println(checker.IsInt64Type(int64V)) // Should return true.
	fmt.Println(checker.IsInt64Type(intV))   // Should return false.

	fmt.Println("IsUintType results:")
	fmt.Println(checker.IsUintType(uintV))  // Should return true.
	fmt.Println(checker.IsUintType(uint8V)) // Should return false.

	fmt.Println("IsUint8Type results:")
	fmt.Println(checker.IsUint8Type(uintV))  // Should return false.
	fmt.Println(checker.IsUint8Type(uint8V)) // Should return true.

	fmt.Println("IsUint16Type results:")
	fmt.Println(checker.IsUint16Type(uintV))   // Should return false.
	fmt.Println(checker.IsUint16Type(uint16V)) // Should return true.

	fmt.Println("IsUint32Type results:")
	fmt.Println(checker.IsUint32Type(uintV))   // Should return false.
	fmt.Println(checker.IsUint32Type(uint32V)) // Should return true.

	fmt.Println("IsUint64Type results:")
	fmt.Println(checker.IsUint64Type(uintV))   // Should return false.
	fmt.Println(checker.IsUint64Type(uint64V)) // Should return true.

	fmt.Println("IsFloat32Type results:")
	fmt.Println(checker.IsFloat32Type(uintV)) // Should return false.
	fmt.Println(checker.IsFloat32Type(f32V))  // Should return true.

	fmt.Println("IsFloat64Type results:")
	fmt.Println(checker.IsFloat64Type(uintV)) // Should return false.
	fmt.Println(checker.IsFloat64Type(f64V))  // Should return true.

	fmt.Println("IsBoolType results:")
	fmt.Println(checker.IsBoolType(true))  // Should return true.
	fmt.Println(checker.IsBoolType(uintV)) // Should return false.

	fmt.Println("IsTimeType results:")
	fmt.Println(checker.IsTimeType(t))    // Should return true.
	fmt.Println(checker.IsTimeType(text)) // Should return false.

	fmt.Println("IsDurationType results:")
	fmt.Println(checker.IsDurationType(d))    // Should return true.
	fmt.Println(checker.IsDurationType(text)) // Should return false.

	fmt.Println("IsBytesType results:")
	fmt.Println(checker.IsBytesType(ba))   // Should return true.
	fmt.Println(checker.IsBytesType(text)) // Should return false.

	fmt.Println("IsErrorType results:")
	fmt.Println(checker.IsErrorType(err))  // Should return true.
	fmt.Println(checker.IsErrorType(text)) // Should return false.
}
