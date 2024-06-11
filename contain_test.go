package checker

import (
	"testing"
)

type containsCase struct {
	name  string
	a     any
	b     any
	want  bool
	panic bool
}

func TestContains(t *testing.T) {
	strA := "Test String"
	strB := "Test"

	tests := []containsCase{
		{
			name: "Simple String",
			a:    strA,
			b:    strB,
			want: true,
		},
		{
			name: "Simple String Pointers",
			a:    &strA,
			b:    &strB,
			want: true,
		},
		{
			name: "String with Case Difference",
			a:    "Test String",
			b:    "test",
			want: false,
		},

		{
			name: "String with Space",
			a:    "Test String",
			b:    "Test ",
			want: true,
		},
		{
			name: "Slice of Integers",
			a:    []int{1, 2, 3, 4, 5},
			b:    3,
			want: true,
		},
		{
			name: "Slice of Strings",
			a:    []string{"go", "land"},
			b:    "land",
			want: true,
		},
		{
			name: "Not in Slice",
			a:    []string{"go", "land"},
			b:    "test",
			want: false,
		},
		{
			name: "In Struct",
			a:    struct{ X string }{"go"},
			b:    "go",
			want: true,
		},
		{
			name: "Not in Struct",
			a:    struct{ X string }{"land"},
			b:    "test",
			want: false,
		},
		{
			name: "In Map",
			a:    map[string]string{"X": "go"},
			b:    "go",
			want: true,
		},
		{
			name: "Not In Map",
			a:    map[string]string{"X": "land"},
			b:    "test",
			want: false,
		},
		{
			name: "In Map Struct",
			a:    map[string]any{"X": struct{ X string }{"land"}},
			b:    struct{ X string }{"land"},
			want: true,
		},
		{
			name: "Not In Map Struct",
			a:    map[string]any{"X": struct{ X string }{"land"}},
			b:    struct{ X string }{"test"},
			want: false,
		},
		{
			name:  "Nil Value",
			a:     nil,
			b:     "test",
			panic: true,
		},
		{
			name:  "Unsupported Type",
			a:     1,
			b:     "test",
			panic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}()
			}

			if got := Contains(tt.a, tt.b); got != tt.want {
				t.Errorf("Contains() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestNotContains(t *testing.T) {
	strA := "Test String"
	strB := "Moon"

	tests := []containsCase{
		{
			name: "Simple String",
			a:    strA,
			b:    strB,
			want: true,
		},
		{
			name: "Simple String Pointers",
			a:    &strA,
			b:    &strB,
			want: true,
		},
		{
			name: "NotContains",
			a:    "Hello World",
			b:    "Test",
			want: true,
		},
		{
			name: "Contains",
			a:    "Hello World",
			b:    "Hello",
			want: false,
		},
		{
			name: "EmptyStrings",
			a:    "",
			b:    "",
			want: false,
		},
		{
			name: "OneEmptyString",
			a:    "Hello World",
			b:    "",
			want: false,
		},
		{
			name:  "Nil Value",
			a:     nil,
			b:     "test",
			panic: true,
		},
		{
			name:  "Unsupported Type",
			a:     1,
			b:     "test",
			panic: true,
		},
		{
			name: "Not in Slice",
			a:    []string{"go", "land"},
			b:    "test",
			want: true,
		},
		{
			name: "Not In Map Struct",
			a:    map[string]any{"X": struct{ X string }{"land"}},
			b:    struct{ X string }{"test"},
			want: true,
		},
		{
			name: "Not In Map Struct",
			a:    map[string]any{"X": struct{ X string }{"land"}},
			b:    struct{ X string }{"test"},
			want: true,
		},
		{
			name: "Not in Struct",
			a:    struct{ X string }{"land"},
			b:    "test",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}()
			}

			if got := NotContains(tt.a, tt.b); got != tt.want {
				t.Errorf("NotContains() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestContainsIgnoreCase(t *testing.T) {
	strA := "Test String"
	strB := "test"

	tests := []containsCase{
		{
			name: "SingleCharacterLowerCase",
			a:    "a",
			b:    "a",
			want: true,
		},
		{
			name: "SingleCharacterIgnoreCase",
			a:    "A",
			b:    "a",
			want: true,
		},
		{
			name: "MultipleCharactersIgnoreCase",
			a:    strA,
			b:    strB,
			want: true,
		},
		{
			name: "MultipleCharactersIgnoreCasePointers",
			a:    &strA,
			b:    &strB,
			want: true,
		},
		{
			name: "MultipleCharactersNotExist",
			a:    "HelloWorld",
			b:    "xyz",
			want: false,
		},
		{
			name: "NumericValues",
			a:    "1234567890",
			b:    "1234",
			want: true,
		},
		{
			name: "NumericValuesNotExist",
			a:    "1234567890",
			b:    "9075",
			want: false,
		},
		{
			name: "EmptyString",
			a:    "",
			b:    "",
			want: true,
		},
		{
			name: "EmptyStringAndValue",
			a:    "",
			b:    "abc",
			want: false,
		},
		{
			name: "ValueAndEmptyString",
			a:    "abc",
			b:    "",
			want: true,
		},
		{
			name:  "Nil Value",
			a:     nil,
			b:     "test",
			panic: true,
		},
		{
			name:  "Unsupported Type",
			a:     1,
			b:     "test",
			panic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}()
			}

			if got := ContainsIgnoreCase(tt.a, tt.b); got != tt.want {
				t.Errorf("ContainsIgnoreCase() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestNotContainsIgnoreCase(t *testing.T) {
	strA := "Test String"
	strB := "moon"

	tests := []containsCase{
		{
			name: "Normal Strings",
			a:    strA,
			b:    strB,
			want: true,
		},
		{
			name: "Pointers Strings",
			a:    &strA,
			b:    &strB,
			want: true,
		},
		{
			name: "Both LowerCase",
			a:    "hello",
			b:    "WORLD",
			want: true,
		},
		{
			name: "Both UpperCase",
			a:    "HELLO",
			b:    "HELLO",
			want: false,
		},
		{
			name: "Mixed Case",
			a:    "Hello",
			b:    "WorlD",
			want: true,
		},
		{
			name: "Mixed Case Incorrect",
			a:    "Hello",
			b:    "HELLO",
			want: false,
		},
		{
			name: "Empty String",
			a:    "Hello",
			b:    "",
			want: false,
		},
		{
			name: "Nil String",
			a:    "Hello",
			b:    nil,
			want: true,
		},
		{
			name:  "Nil Value",
			a:     nil,
			b:     "test",
			panic: true,
		},
		{
			name:  "Unsupported Type",
			a:     1,
			b:     "test",
			panic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}()
			}

			if got := NotContainsIgnoreCase(tt.a, tt.b); got != tt.want {
				t.Errorf("NotContainsIgnoreCase() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestContainsKey(t *testing.T) {
	strKey := "key_1"

	tests := []containsCase{
		{
			name: "ContainsKey Map True",
			a:    &map[string]int{"key_1": 1},
			b:    &strKey,
			want: true,
		},
		{
			name: "ContainsKey Map False",
			a:    map[string]int{"key_1": 1},
			b:    "key_2",
			want: false,
		},
		{
			name: "ContainsKey Struct True",
			a:    struct{ Field string }{Field: "value"},
			b:    "Field",
			want: true,
		},
		{
			name: "ContainsKey Struct False",
			a:    struct{ Field string }{Field: "value"},
			b:    "NoField",
			want: false,
		},
		{
			name: "ContainsKey Interface Map True",
			a:    interface{}(map[string]int{"key_1": 1}),
			b:    "key_1",
			want: true,
		},
		{
			name: "ContainsKey Interface Map False",
			a:    interface{}(map[string]int{"key_1": 1}),
			b:    "key_2",
			want: false,
		},
		{
			name: "ContainsKey Map Int True",
			a:    interface{}(map[int]int{1: 1}),
			b:    1,
			want: true,
		},
		{
			name: "ContainsKey Map Int False",
			a:    interface{}(map[int]int{1: 1}),
			b:    2,
			want: false,
		},
		{
			name:  "Nil Value",
			a:     nil,
			b:     "test",
			panic: true,
		},
		{
			name:  "Unsupported Type",
			a:     1,
			b:     "test",
			panic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}()
			}

			if got := ContainsKey(tt.a, tt.b); got != tt.want {
				t.Errorf("ContainsKey() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestNotContainsKey(t *testing.T) {
	strKey := "test"
	tests := []containsCase{
		{name: "STRUCT_EXITS_KEY", a: struct{ test bool }{test: true}, b: &strKey, want: false},
		{name: "STRUCT_EXITS_NO_KEY", a: struct{ test bool }{test: true}, b: "test1", want: true},
		{name: "POINTER_STRUCT_EXITS_NO_KEY", a: &struct{ test bool }{test: true}, b: &strKey, want: false},
		{name: "MAP_INT_EXITS_KEY", a: map[int]bool{1: true, 2: false}, b: 1, want: false},
		{name: "MAP_INT_NO_KEY", a: map[int]bool{1: true, 2: false}, b: 3, want: true},
		{name: "MAP_STRING_EXITS_KEY", a: map[string]bool{"one": true, "two": false}, b: "one", want: false},
		{name: "MAP_STRING_NO_KEY", a: map[string]bool{"one": true, "two": false}, b: "three", want: true},
		{name: "NIL", a: nil, b: "three", panic: true},
		{name: "UNSUPPORTED_TYPE", a: 1, b: "three", panic: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}()
			}

			if got := NotContainsKey(tt.a, tt.b); got != tt.want {
				t.Errorf("NotContainsKey() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestContainsOnSlice(t *testing.T) {
	var tests = []struct {
		input []int
		check func(index, element int) bool
		want  bool
		desc  string
	}{
		{input: []int{1, 2, 3, 4, 5}, check: func(index, element int) bool { return index == 0 }, want: true, desc: "Index Exists"},
		{input: []int{1, 2, 3, 4, 5}, check: func(index, element int) bool { return element == 3 }, want: true, desc: "Element Exists"},
		{input: []int{1, 2, 3, 4, 5}, check: func(index, element int) bool { return index == 7 }, want: false, desc: "Index Not Exists"},
		{input: []int{1, 2, 3, 4, 5}, check: func(index, element int) bool { return element == 6 }, want: false, desc: "Element Not Exists"},
		{input: []int{}, check: func(index, data int) bool { return data == 1 }, want: false, desc: "Empty Slice"},
		{input: nil, check: func(index, data int) bool { return data == 1 }, want: false, desc: "Nil Slice"},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			if got := ContainsOnSlice(tt.input, tt.check); got != tt.want {
				t.Errorf("ContainsOnSlice() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestNotContainsOnSlice(t *testing.T) {
	tests := []struct {
		name      string
		slice     []string
		condition func(index int, data string) bool
		want      bool
	}{
		{
			name:  "Not Contains string in slice",
			slice: []string{"hello", "world"},
			condition: func(index int, ele string) bool {
				return ele == "test"
			},
			want: true,
		},
		{
			name:  "Contains string in slice",
			slice: []string{"hello", "world"},
			condition: func(index int, ele string) bool {
				return ele == "world"
			},
			want: false,
		},
		{
			name:  "Empty slice",
			slice: []string{},
			condition: func(index int, ele string) bool {
				return ele == "test"
			},
			want: true,
		},
		{
			name:  "Nil slice",
			slice: nil,
			condition: func(index int, ele string) bool {
				return ele == "test"
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NotContainsOnSlice(tt.slice, tt.condition); got != tt.want {
				t.Errorf("NotContainsOnSlice() = %v, want = %v", got, tt.want)
			}
		})
	}
}
