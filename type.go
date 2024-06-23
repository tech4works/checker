package checker

import (
	"encoding/json"
	"reflect"
	"regexp"
	"strconv"
	"time"
)

func IsJSON(a any) bool {
	bs := toBytes(a)

	var js any
	return json.Unmarshal(bs, &js) == nil
}

func IsNotJSON(a any) bool {
	return !IsJSON(a)
}

func IsMap(a any) bool {
	bs := toBytes(a)

	var js map[string]any
	return json.Unmarshal(bs, &js) == nil
}

func IsNotMap(a any) bool {
	return !IsMap(a)
}

func IsSlice(a any) bool {
	bs := toBytes(a)

	var slice []any
	return json.Unmarshal(bs, &slice) == nil
}

func IsNotSlice(a any) bool {
	return !IsSlice(a)
}

func IsSliceOfMaps(a any) bool {
	bs := toBytes(a)

	var slice []map[string]any
	return json.Unmarshal(bs, &slice) == nil
}

func IsNotSliceOfMaps(a any) bool {
	return !IsSliceOfMaps(a)
}

func IsInt(a any) bool {
	s := toString(a)

	_, err := strconv.Atoi(s)
	return err == nil
}

func IsNotInt(a any) bool {
	return !IsInt(a)
}

func IsBool(a any) bool {
	s := toString(a)

	_, err := strconv.ParseBool(s)
	return err == nil
}

func IsNotBool(a any) bool {
	return !IsBool(a)
}

func IsFloat(a any) bool {
	s := toString(a)

	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func IsNotFloat(a any) bool {
	return !IsFloat(a)
}

func IsTime(a any) bool {
	s := toString(a)

	_, err := toTimeWithErr(s)
	return err == nil
}

func IsNotTime(a any) bool {
	return !IsTime(a)
}

func IsDuration(a any) bool {
	s := toString(a)

	_, err := time.ParseDuration(s)
	return err == nil
}

func IsNotDuration(a any) bool {
	return !IsDuration(a)
}

func IsByteUnit(a any) bool {
	s := toString(a)

	re := regexp.MustCompile(`^(\d+)(B|KB|MB|GB|TB|PB)$`)
	return re.MatchString(s)
}

func IsNotByteUnit(a any) bool {
	return !IsByteUnit(a)
}

func IsPointerType(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.Pointer
}

func IsNotPointerType(a any) bool {
	return !IsPointerType(a)
}

func IsFuncType(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.Func
}

func IsNotFuncType(a any) bool {
	return !IsFuncType(a)
}

func IsChanType(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.Chan
}

func IsNotChanType(a any) bool {
	return !IsChanType(a)
}

func IsInterfaceType(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.Interface
}

func IsNotInterfaceType(a any) bool {
	return !IsInterfaceType(a)
}

func IsMapType(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.Map
}

func IsNotMapType(a any) bool {
	return !IsMapType(a)
}

func IsStructType(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.Struct
}

func IsNotStructType(a any) bool {
	return !IsStructType(a)
}

func IsSliceType(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.Slice
}

func IsNotSliceType(a any) bool {
	return !IsSliceType(a)
}

func IsArrayType(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.Array
}

func IsNotArrayType(a any) bool {
	return !IsArrayType(a)
}

func IsSliceOrArrayType(a any) bool {
	return IsSliceType(a) || IsArrayType(a)
}

func IsNotSliceOrArrayType(a any) bool {
	return !IsSliceOrArrayType(a)
}

func IsStringType(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.String
}

func IsNotStringType(a any) bool {
	return !IsStringType(a)
}

func IsIntType(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.Int
}

func IsNotIntType(a any) bool {
	return !IsIntType(a)
}

func IsInt8Type(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.Int8
}

func IsNotInt8Type(a any) bool {
	return !IsInt8Type(a)
}

func IsInt16Type(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.Int16
}

func IsNotInt16Type(a any) bool {
	return !IsInt16Type(a)
}

func IsInt32Type(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.Int32
}

func IsNotInt32Type(a any) bool {
	return !IsInt32Type(a)
}

func IsInt64Type(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.Int64
}

func IsNotInt64Type(a any) bool {
	return !IsInt64Type(a)
}

func IsUintType(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.Uint
}

func IsNotUintType(a any) bool {
	return !IsUintType(a)
}

func IsUint8Type(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.Uint8
}

func IsNotUint8Type(a any) bool {
	return !IsUint8Type(a)
}

func IsUint16Type(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.Uint16
}

func IsNotUint16Type(a any) bool {
	return !IsUint16Type(a)
}

func IsUint32Type(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.Uint32
}

func IsNotUint32Type(a any) bool {
	return !IsUint32Type(a)
}

func IsUint64Type(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.Uint64
}

func IsNotUint64Type(a any) bool {
	return !IsUint64Type(a)
}

func IsFloat32Type(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.Float32
}

func IsNotFloat32Type(a any) bool {
	return !IsFloat32Type(a)
}

func IsFloat64Type(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.Float64
}

func IsNotFloat64Type(a any) bool {
	return !IsFloat64Type(a)
}

func IsBoolType(a any) bool {
	return reflect.ValueOf(a).Kind() == reflect.Bool
}

func IsNotBoolType(a any) bool {
	return !IsBoolType(a)
}

func IsTimeType(a any) bool {
	return a != nil && reflect.TypeOf(a) == reflect.TypeOf(time.Time{})
}

func IsNotTimeType(a any) bool {
	return !IsTimeType(a)
}

func IsDurationType(a any) bool {
	return a != nil && reflect.TypeOf(a) == reflect.TypeOf(time.Duration(0))
}

func IsNotDurationType(a any) bool {
	return !IsDurationType(a)
}

func IsBytesType(a any) bool {
	return a != nil && reflect.TypeOf(a) == reflect.TypeOf([]byte{})
}

func IsNotBytesType(a any) bool {
	return !IsBytesType(a)
}

func IsErrorType(a any) bool {
	if a == nil {
		return false
	}
	_, ok := a.(error)
	return ok
}

func IsNotErrorType(a any) bool {
	return !IsErrorType(a)
}
