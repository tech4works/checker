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
