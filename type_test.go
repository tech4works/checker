package checker

import (
	"errors"
	"testing"
	"time"
)

func TestIsJSON(t *testing.T) {
	tests := []baseCase{
		{
			name: "Valid JSON Object",
			arg:  `{"key": "value"}`,
			want: true,
		},
		{
			name: "Valid JSON Array",
			arg:  `[1, 2, 3]`,
			want: true,
		},
		{
			name: "Empty JSON Object",
			arg:  `{}`,
			want: true,
		},
		{
			name: "Empty JSON Array",
			arg:  `[]`,
			want: true,
		},
		{
			name: "Invalid JSON Object",
			arg:  `{"key": "value`,
			want: false,
		},
		{
			name: "Invalid JSON Array",
			arg:  `[1, 2, 3`,
			want: false,
		},
		{
			name: "Not JSON - Simple String",
			arg:  `Not a JSON string`,
			want: false,
		},
		{
			name: "Not JSON - Number",
			arg:  `123`,
			want: false,
		},
		{
			name: "Not JSON - Boolean",
			arg:  `true`,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsJSON(tt.arg); got != tt.want {
				t.Errorf("IsJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsMap(t *testing.T) {
	testCases := []baseCase{
		{
			name: "JSONObject",
			arg:  []byte(`{"key": "value"}`),
			want: true,
		},
		{
			name: "JSONArray",
			arg:  []byte(`["element1", "element2"]`),
			want: false,
		},
		{
			name: "JSONNumber",
			arg:  []byte(`123`),
			want: false,
		},
		{
			name: "JSONBoolean",
			arg:  []byte(`true`),
			want: false,
		},
		{
			name: "JSONNull",
			arg:  []byte(`null`),
			want: true,
		},
		{
			name: "EmptyJSON",
			arg:  []byte(`{}`),
			want: true,
		},
		{
			name: "NonJSONString",
			arg:  []byte(`hello`),
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsMap(tc.arg); got != tc.want {
				t.Errorf("IsMap(%v) = %v; want %v", tc.arg, got, tc.want)
			}
		})
	}
}

func TestIsSlice(t *testing.T) {
	tests := []baseCase{
		{
			name: "Test With Integer",
			arg:  123,
			want: false,
		},
		{
			name: "Test With String",
			arg:  "123",
			want: false,
		},
		{
			name: "Test With Bool",
			arg:  true,
			want: false,
		},
		{
			name: "Test With Slice",
			arg:  []any{1, 2, 3},
			want: true,
		},
		{
			name: "Test With Empty Slice",
			arg:  []any{},
			want: true,
		},
		{
			name: "Test With Struct",
			arg:  struct{ name string }{"test"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSlice(tt.arg); got != tt.want {
				t.Errorf("IsSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSliceOfMaps(t *testing.T) {
	testCases := []baseCase{
		{
			name: "validSliceOfMaps",
			arg:  []byte(`[{"key":"value"},{"1":2}]`),
			want: true,
		},
		{
			name: "nonSliceInput",
			arg:  []byte("should fail"),
			want: false,
		},
		{
			name: "validSliceNonMapElements",
			arg:  []byte(`[1,2]`),
			want: false,
		},
		{
			name: "validSliceOfInterfaces",
			arg: []interface{}{
				map[string]interface{}{"foo": "bar"},
				map[string]interface{}{"baz": 123},
			},
			want: true,
		},
		{
			name: "validSliceNonMapInterfaces",
			arg:  []interface{}{1, 2},
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if output := IsSliceOfMaps(tc.arg); output != tc.want {
				t.Errorf("IsSliceOfMaps(%v) = %t; want %t", tc.arg, output, tc.want)
			}
		})
	}
}

func TestIsInt(t *testing.T) {
	tt := []baseCase{
		{name: "IntegerAsString", arg: "123", want: true},
		{name: "FloatAsString", arg: "123.45", want: false},
		{name: "IntegerInFloatFormat", arg: "123.00", want: false},
		{name: "NonNumericalString", arg: "hello", want: false},
		{name: "EmptyString", arg: "", want: false},
		{name: "Zero", arg: "0", want: true},
		{name: "NegativeInteger", arg: "-123", want: true},
		{name: "Integer", arg: 123, want: true},
		{name: "Float", arg: 123.45, want: false},
		{name: "NegativeInteger", arg: -123, want: true},
		{name: "ZeroInteger", arg: 0, want: true},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsInt(tc.arg); got != tc.want {
				t.Errorf("expected %v, got %v", tc.want, got)
			}
		})
	}
}

func TestIsBool(t *testing.T) {
	tests := []baseCase{
		{
			name: "boolean true",
			arg:  true,
			want: true,
		},
		{
			name: "boolean false",
			arg:  false,
			want: true,
		},
		{
			name: "true string",
			arg:  "true",
			want: true,
		},
		{
			name: "false string",
			arg:  "false",
			want: true,
		},
		{
			name: "invalid string",
			arg:  "not a boolean",
			want: false,
		},
		{
			name: "int 1",
			arg:  1,
			want: true,
		},
		{
			name: "int 0",
			arg:  0,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBool(tt.arg); got != tt.want {
				t.Errorf("IsBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsFloat(t *testing.T) {
	testCases := []baseCase{
		{name: "FloatAsString", arg: "3.14", want: true},
		{name: "IntAsString", arg: "3", want: true},
		{name: "InvalidStringFloat", arg: "not a float"},
		{name: "NumericFloat", arg: 3.14, want: true},
		{name: "NumericInt", arg: 3, want: true},
		{name: "InvalidType", arg: []int{1, 2, 3}},
		{name: "EmptyString", arg: ""},
		{name: "EdgeCasePositiveInf", arg: "Inf", want: true},
		{name: "EdgeCaseNegativeInf", arg: "-Inf", want: true},
		{name: "EdgeCaseNaN", arg: "NaN", want: true},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFloat(tt.arg); got != tt.want {
				t.Errorf("IsFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsTime(t *testing.T) {
	testCases := []baseCase{
		{
			name: "ValidTime",
			arg:  time.Now(),
			want: true,
		},
		{
			name: "ValidInteger",
			arg:  time.Now().UnixMilli(),
			want: true,
		},
		{
			name: "ValidFloat",
			arg:  float64(time.Now().UnixMilli()),
			want: true,
		},
		{
			name: "InvalidTime",
			arg:  "not a time",
			want: false,
		},
		{
			name: "Integer",
			arg:  5,
			want: true,
		},
		{
			name: "Boolean",
			arg:  true,
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := IsTime(tc.arg)
			if got != tc.want {
				t.Errorf("IsTime(%v) = %v; want %v", tc.arg, got, tc.want)
			}
		})
	}
}

func TestIsDuration(t *testing.T) {
	testCases := []baseCase{
		{
			name: "ValidDuration",
			arg:  "1h30m",
			want: true,
		},
		{
			name: "InvalidDuration",
			arg:  "1h30m30",
		},
		{
			name: "EmptyString",
			arg:  "",
		},
		{
			name: "NonString",
			arg:  123,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := IsDuration(tc.arg)
			if actual != tc.want {
				t.Errorf("Test failed for %s, got: %t, want: %t", tc.name, actual, tc.want)
			}
		})
	}
}

func TestIsByteUnit(t *testing.T) {
	testCases := []baseCase{
		{name: "ByteUnit_Byte", arg: "120B", want: true},
		{name: "ByteUnit_Kilobyte", arg: "120KB", want: true},
		{name: "ByteUnit_Megabyte", arg: "120MB", want: true},
		{name: "ByteUnit_Gigabyte", arg: "120GB", want: true},
		{name: "ByteUnit_Terabyte", arg: "120TB", want: true},
		{name: "ByteUnit_Petabyte", arg: "120PB", want: true},
		{name: "ByteUnit_Invalid", arg: "120XYZ", want: false},
		{name: "ByteUnit_WithoutSuffix", arg: "120", want: false},
		{name: "ByteUnit_EmptyString", arg: "", want: false},
		{name: "ByteUnit_OnlySuffix", arg: "GB", want: false},
		{name: "ByteUnit_ZeroValue", arg: "0B", want: true},
		{name: "ByteUnit_NegativeValue", arg: "-120B", want: false},
		{name: "ByteUnit_WithSpacing", arg: "120 MB", want: false},
		{name: "ByteUnit_LowerCaseSuffix", arg: "120gb", want: false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsByteUnit(tc.arg); got != tc.want {
				t.Errorf("IsByteUnit() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestIsPointerType(t *testing.T) {
	tests := []baseCase{
		{name: "nil", arg: nil, want: false},
		{name: "int", arg: 1, want: false},
		{name: "float", arg: 1.0, want: false},
		{name: "string", arg: "test", want: false},
		{name: "bool", arg: true, want: false},
		{name: "int pointer", arg: new(int), want: true},
		{name: "float pointer", arg: new(float64), want: true},
		{name: "string pointer", arg: new(string), want: true},
		{name: "bool pointer", arg: new(bool), want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPointerType(tt.arg); got != tt.want {
				t.Errorf("IsPointerType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsFuncType(t *testing.T) {
	f := func() {}
	var testCases = []baseCase{
		{name: "Function", arg: func() {}, want: true},
		{name: "AnonFunction", arg: func(i int) string { return "test" }, want: true},
		{name: "PointerToFunction", arg: &f, want: false},
		{name: "String", arg: "foo", want: false},
		{name: "Int", arg: 123, want: false},
		{name: "Float", arg: 1.23, want: false},
		{name: "Bool", arg: true, want: false},
		{name: "Slice", arg: []string{"foo", "bar", "baz"}, want: false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsFuncType(tc.arg); got != tc.want {
				t.Errorf("IsFuncType() = %v; wanted %v", got, tc.want)
			}
		})
	}
}

func TestIsChanType(t *testing.T) {
	tests := []baseCase{
		{
			name: "IntChannel",
			arg:  make(chan int),
			want: true,
		},
		{
			name: "StringChannel",
			arg:  make(chan string),
			want: true,
		},
		{
			name: "CustomStructChannel",
			arg:  make(chan struct{ Name string }),
			want: true,
		},
		{
			name: "Int",
			arg:  5,
			want: false,
		},
		{
			name: "String",
			arg:  "test",
			want: false,
		},
		{
			name: "Nil",
			arg:  nil,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsChanType(tt.arg); got != tt.want {
				t.Errorf("IsChanType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsMapType(t *testing.T) {
	testCases := []baseCase{
		{
			name: "MapType",
			arg:  map[string]int{"test": 1},
			want: true,
		},
		{
			name: "NotMapType",
			arg:  []int{1, 2, 3},
			want: false,
		},
		{
			name: "NilMapType",
			arg:  map[int]int(nil),
			want: true,
		},
		{
			name: "EmptyMapType",
			arg:  make(map[int]int),
			want: true,
		},
		{
			name: "NilValue",
			arg:  nil,
			want: false,
		},
		{
			name: "PointerType",
			arg:  &map[string]int{"test": 1},
			want: false,
		},
		{
			name: "PrimitiveType",
			arg:  1,
			want: false,
		},
		{
			name: "StructureType",
			arg:  struct{ name string }{},
			want: false,
		},
		{
			name: "FunctionType",
			arg:  func() {},
			want: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsMapType(tt.arg); got != tt.want {
				t.Errorf("IsMapType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsStructType(t *testing.T) {
	tests := []baseCase{
		{name: "Struct", arg: struct{}{}, want: true},
		{name: "Int", arg: 1},
		{name: "Float", arg: 1.1},
		{name: "String", arg: ""},
		{name: "Slice", arg: []int{1, 2, 3}},
		{name: "Map", arg: map[string]string{"key": "value"}},
		{name: "Bool", arg: false},
		{name: "Pointer", arg: &struct{}{}},
		{name: "EmptyInterface", arg: new(interface{})},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsStructType(tt.arg); got != tt.want {
				t.Errorf("IsStructType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSliceType(t *testing.T) {
	tests := []baseCase{
		{name: "Test with slice", arg: []int{1, 2, 3}, want: true},
		{name: "Test with empty slice", arg: []string{}, want: true},
		{name: "Test with string", arg: "test", want: false},
		{name: "Test with integer", arg: 123, want: false},
		{name: "Test with float", arg: 123.456, want: false},
		{name: "Test with boolean", arg: true, want: false},
		{name: "Test with map", arg: map[string]string{"a": "b"}, want: false},
		{name: "Test with nil", arg: nil, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSliceType(tt.arg); got != tt.want {
				t.Errorf("IsSliceType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsArrayType(t *testing.T) {
	var intArray [5]int
	var stringArray [5]string
	boolSlice := make([]bool, 5)
	stringSlice := make([]string, 5)

	testCases := []baseCase{
		{
			name: "int array",
			arg:  intArray,
			want: true,
		},
		{
			name: "string array",
			arg:  stringArray,
			want: true,
		},
		{
			name: "bool slice",
			arg:  boolSlice,
			want: false,
		},
		{
			name: "string slice",
			arg:  stringSlice,
			want: false,
		},
		{
			name: "int",
			arg:  5,
			want: false,
		},
		{
			name: "string",
			arg:  "hello",
			want: false,
		},
		{
			name: "bool",
			arg:  true,
			want: false,
		},
		{
			name: "nil",
			arg:  nil,
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			if result := IsArrayType(tc.arg); result != tc.want {
				t.Errorf("Got %v, expected %v", result, tc.want)
			}
		})
	}
}

func TestIsSliceOrArrayType(t *testing.T) {
	cases := []baseCase{
		{name: "Integers Array", arg: []int{1, 2, 3}, want: true},
		{name: "String Array", arg: []string{"Hello", "World"}, want: true},
		{name: "Boolean", arg: true, want: false},
		{name: "Individual String", arg: "Hello", want: false},
		{name: "Empty Slice", arg: []string{}, want: true},
		{name: "Empty Array", arg: [0]int{}, want: true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsSliceOrArrayType(tc.arg); got != tc.want {
				t.Errorf("IsSliceOrArrayType(%v) = %v; want %v", tc.arg, got, tc.want)
			}
		})
	}
}

func TestIsStringType(t *testing.T) {
	tests := []baseCase{
		{
			name: "String",
			arg:  "Hello",
			want: true,
		},
		{
			name: "Integer",
			arg:  1,
			want: false,
		},
		{
			name: "Float",
			arg:  1.1,
			want: false,
		},
		{
			name: "Boolean",
			arg:  true,
			want: false,
		},
		{
			name: "Slice",
			arg:  []int{1, 2, 3},
			want: false,
		},
		{
			name: "Nil",
			arg:  nil,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsStringType(tt.arg); got != tt.want {
				t.Errorf("IsStringType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsIntType(t *testing.T) {
	tests := []baseCase{
		{
			name: "IntValue",
			arg:  5,
			want: true,
		},
		{
			name: "FloatValue",
			arg:  5.5,
			want: false,
		},
		{
			name: "StringValue",
			arg:  "abc",
			want: false,
		},
		{
			name: "BoolValue",
			arg:  false,
			want: false,
		},
		{
			name: "SliceValue",
			arg:  []int{1, 2, 3},
			want: false,
		},
		{
			name: "MapValue",
			arg:  map[string]int{"one": 1},
			want: false,
		},
		{
			name: "IntPointerValue",
			arg:  new(int),
			want: false,
		},
		{
			name: "ZeroValue",
			arg:  0,
			want: true,
		},
		{
			name: "NilValue",
			arg:  nil,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsIntType(tt.arg); got != tt.want {
				t.Errorf("IsIntType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsInt8Type(t *testing.T) {
	testCases := []baseCase{
		{
			name: "Int8",
			arg:  int8(127),
			want: true,
		},
		{
			name: "Int16",
			arg:  int16(32767),
			want: false,
		},
		{
			name: "Int32",
			arg:  int32(2147483647),
			want: false,
		},
		{
			name: "Int64",
			arg:  int64(9223372036854775807),
			want: false,
		},
		{
			name: "Int",
			arg:  9223372036854775807,
			want: false,
		},
		{
			name: "String",
			arg:  "string",
			want: false,
		},
		{
			name: "Bool",
			arg:  true,
			want: false,
		},
		{
			name: "Float32",
			arg:  float32(3.4028235e+38),
			want: false,
		},
		{
			name: "Float64",
			arg:  1.7976931348623157e+308,
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := IsInt8Type(tc.arg)
			if got != tc.want {
				t.Errorf("IsInt8Type() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestIsInt16Type(t *testing.T) {
	tests := []baseCase{
		{
			name: "case where arg is int16",
			arg:  int16(16),
			want: true,
		},
		{
			name: "case where arg is string",
			arg:  "foo",
			want: false,
		},
		{
			name: "case where arg is int32",
			arg:  int32(32),
			want: false,
		},
		{
			name: "case where arg is float64",
			arg:  64.0,
			want: false,
		},
		{
			name: "case where arg is bool",
			arg:  true,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsInt16Type(tt.arg); got != tt.want {
				t.Errorf("IsInt16Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsInt32Type(t *testing.T) {
	tests := []baseCase{
		{
			name: "PositiveInt32Case",
			arg:  int32(32),
			want: true,
		},
		{
			name: "NegativeInt32Case",
			arg:  int32(-32),
			want: true,
		},
		{
			name: "ZeroInt32Case",
			arg:  int32(0),
			want: true,
		},
		{
			name: "StrCase",
			arg:  "32",
			want: false,
		},
		{
			name: "DifferentIntTypesCase",
			arg:  int64(32),
			want: false,
		},
		{
			name: "FloatCase",
			arg:  float32(32.1),
			want: false,
		},
		{
			name: "BoolTypeCase",
			arg:  true,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsInt32Type(tt.arg); got != tt.want {
				t.Errorf("IsInt32Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsInt64Type(t *testing.T) {
	tests := []baseCase{
		{
			name: "Positive Case",
			arg:  int64(1234567890),
			want: true,
		},
		{
			name: "Negative Case Int",
			arg:  123,
			want: false,
		},
		{
			name: "Negative Case Float",
			arg:  123.45,
			want: false,
		},
		{
			name: "Negative Case String",
			arg:  "1234567890",
			want: false,
		},
		{
			name: "Negative Case Boolean",
			arg:  true,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsInt64Type(tt.arg); got != tt.want {
				t.Errorf("IsInt64Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsUintType(t *testing.T) {
	tests := []baseCase{
		{
			name: "checking int value",
			arg:  1,
			want: false,
		},
		{
			name: "checking uint value",
			arg:  uint(1),
			want: true,
		},
		{
			name: "checking string value",
			arg:  "string",
			want: false,
		},
		{
			name: "checking boolean value",
			arg:  true,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUintType(tt.arg); got != tt.want {
				t.Errorf("IsUintType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsUint8Type(t *testing.T) {
	tests := []baseCase{
		{
			name: "Uint8",
			arg:  uint8(1),
			want: true,
		},
		{
			name: "NotUint8_Integer",
			arg:  1,
			want: false,
		},
		{
			name: "NotUint8_String",
			arg:  "1",
			want: false,
		},
		{
			name: "NotUint8_Float",
			arg:  float64(1),
			want: false,
		},
		{
			name: "NotUint8_Bool",
			arg:  true,
			want: false,
		},
		{
			name: "NotUint8_Struct",
			arg:  struct{ a int }{a: 5},
			want: false,
		},
		{
			name: "NotUint8_Slice",
			arg:  []uint8{1, 2, 3},
			want: false,
		},
		{
			name: "NotUint8_Map",
			arg:  map[string]int{"one": 1, "two": 2, "three": 3},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUint8Type(tt.arg); got != tt.want {
				t.Errorf("IsUint8Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsUint16Type(t *testing.T) {
	tests := []baseCase{
		{
			name: "positive - uint16 type",
			arg:  uint16(257),
			want: true,
		},
		{
			name: "negative - uint type",
			arg:  uint(257),
			want: false,
		},
		{
			name: "negative - int type",
			arg:  -257,
			want: false,
		},
		{
			name: "negative - string type",
			arg:  "257",
			want: false,
		},
		{
			name: "negative - float type",
			arg:  257.78,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUint16Type(tt.arg); got != tt.want {
				t.Errorf("IsUint16Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsUint32Type(t *testing.T) {
	cases := []baseCase{
		{
			name: "NilInt",
			arg:  nil,
			want: false,
		},
		{
			name: "PositiveUint32",
			arg:  uint32(10),
			want: true,
		},
		{
			name: "ZeroUint32",
			arg:  uint32(0),
			want: true,
		},
		{
			name: "NegativeInt",
			arg:  -10,
			want: false,
		},
		{
			name: "Float64",
			arg:  10.5,
			want: false,
		},
		{
			name: "String",
			arg:  "hello",
			want: false,
		},
		{
			name: "Bool",
			arg:  true,
			want: false,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUint32Type(tt.arg); got != tt.want {
				t.Errorf("IsUint32Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsUint64Type(t *testing.T) {
	cases := []baseCase{
		{
			name: "Positive Case Uint64",
			arg:  uint64(5),
			want: true,
		},
		{
			name: "Negative Case int",
			arg:  5,
			want: false,
		},
		{
			name: "Negative Case float",
			arg:  1.5,
			want: false,
		},
		{
			name: "Negative Case bool",
			arg:  true,
			want: false,
		},
		{
			name: "Negative Case string",
			arg:  "test",
			want: false,
		},
		{
			name: "Negative Case slice",
			arg:  []int{1, 2, 3},
			want: false,
		},
		{
			name: "Negative Case uint",
			arg:  uint(5),
			want: false,
		},
		{
			name: "Negative Case pointer",
			arg:  new(int),
			want: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsUint64Type(tc.arg); got != tc.want {
				t.Errorf("IsUint64Type(%v) = %v; want %v", tc.arg, got, tc.want)
			}
		})
	}
}

func TestIsFloat32Type(t *testing.T) {
	tests := []baseCase{
		{
			name: "Non-Float32",
			arg:  1,
			want: false,
		},
		{
			name: "Float64",
			arg:  1.0,
			want: false,
		},
		{
			name: "Float32",
			arg:  float32(1.0),
			want: true,
		},
		{
			name: "String",
			arg:  "1.0",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFloat32Type(tt.arg); got != tt.want {
				t.Errorf("IsFloat32Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsFloat64Type(t *testing.T) {
	tests := []baseCase{
		{
			name: "IntInput",
			arg:  3,
			want: false,
		},
		{
			name: "Float64Input",
			arg:  3.3,
			want: true,
		},
		{
			name: "StringInput",
			arg:  "3.3",
			want: false,
		},
		{
			name: "BoolInput",
			arg:  true,
			want: false,
		},
		{
			name: "NilInput",
			arg:  nil,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFloat64Type(tt.arg); got != tt.want {
				t.Errorf("IsFloat64Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsBoolType(t *testing.T) {
	tests := []baseCase{
		{
			name: "Single Bool Variable - True",
			arg:  true,
			want: true,
		},
		{
			name: "Single Bool Variable - False",
			arg:  false,
			want: true,
		},
		{
			name: "Int Variable",
			arg:  42,
			want: false,
		},
		{
			name: "Float Variable",
			arg:  42.42,
			want: false,
		},
		{
			name: "String Variable",
			arg:  "Hello World!",
			want: false,
		},
		{
			name: "Slice Variable",
			arg:  []int{1, 2, 3, 4, 5},
			want: false,
		},
		{
			name: "Map Variable",
			arg:  map[string]int{"one": 1, "two": 2, "three": 3},
			want: false,
		},
		{
			name: "Nil Pointer",
			arg:  (*int)(nil),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBoolType(tt.arg); got != tt.want {
				t.Errorf("IsBoolType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsTimeType(t *testing.T) {
	tests := []baseCase{
		{name: "ValidTimeType", arg: time.Now(), want: true},
		{name: "InvalidTimeTypeString", arg: "hello", want: false},
		{name: "InvalidTimeTypeInt", arg: 123, want: false},
		{name: "InvalidTimeTypeFloat", arg: 123.45, want: false},
		{name: "InvalidTimeTypeSlice", arg: []int{1, 2, 3}, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsTimeType(tt.arg); got != tt.want {
				t.Errorf("IsTimeType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsDurationType(t *testing.T) {
	tests := []baseCase{
		{
			name: "Duration type",
			arg:  time.Duration(0),
			want: true,
		},
		{
			name: "Non-duration type (integer)",
			arg:  5,
			want: false,
		},
		{
			name: "Non-duration type (float)",
			arg:  1.2,
			want: false,
		},
		{
			name: "Non-duration type (string)",
			arg:  "test",
			want: false,
		},
		{
			name: "Non-duration type (nil)",
			arg:  nil,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDurationType(tt.arg); got != tt.want {
				t.Errorf("IsDurationType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsBytesType(t *testing.T) {
	tests := []baseCase{
		{
			name: "NonBytesTypeString",
			arg:  "test",
			want: false,
		},
		{
			name: "NonBytesTypeInteger",
			arg:  123,
			want: false,
		},
		{
			name: "NonBytesTypeFloat",
			arg:  123.45,
			want: false,
		},
		{
			name: "BytesType",
			arg:  []byte("example"),
			want: true,
		},
		{
			name: "NonBytesTypeArray",
			arg:  []int{1, 2, 3},
			want: false,
		},
		{
			name: "NonBytesTypeNil",
			arg:  nil,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBytesType(tt.arg); got != tt.want {
				t.Errorf("IsBytesType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsErrorType(t *testing.T) {
	tests := []baseCase{
		{
			name: "Test With Error Type",
			arg:  errors.New("error"),
			want: true,
		},
		{
			name: "Test With String",
			arg:  "error",
			want: false,
		},
		{
			name: "Test With Int",
			arg:  1,
			want: false,
		},
		{
			name: "Test With Empty Error",
			arg:  errors.New(""),
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsErrorType(tt.arg); got != tt.want {
				t.Errorf("IsErrorType() = %v, want %v", got, tt.want)
			}
		})
	}
}
