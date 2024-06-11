package checker

import (
	"testing"
)

type containsCase struct {
	name string
	a    any
	b    any
	want bool
}

func TestContains(t *testing.T) {
	tests := []containsCase{
		{
			name: "Simple String",
			a:    "Test String",
			b:    "Test",
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
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Contains(tc.a, tc.b); got != tc.want {
				t.Errorf("Contains() = %v, want = %v", got, tc.want)
			}
		})
	}
}

func TestNotContains(t *testing.T) {
	tests := []containsCase{
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
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := NotContains(tc.a, tc.b); got != tc.want {
				t.Errorf("NotContains() = %v, want = %v", got, tc.want)
			}
		})
	}
}

func TestContainsIgnoreCase(t *testing.T) {
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
			a:    "HelloWorld",
			b:    "helloworld",
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
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := ContainsIgnoreCase(tc.a, tc.b); got != tc.want {
				t.Errorf("ContainsIgnoreCase() = %v, want = %v", got, tc.want)
			}
		})
	}
}

func TestNotContainsIgnoreCase(t *testing.T) {
	tests := []containsCase{
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
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := NotContainsIgnoreCase(tc.a, tc.b); got != tc.want {
				t.Errorf("NotContainsIgnoreCase() = %v, want = %v", got, tc.want)
			}
		})
	}
}

func TestContainsKey(t *testing.T) {
	tests := []containsCase{
		{
			name: "ContainsKey Map True",
			a:    map[string]int{"key_1": 1},
			b:    "key_1",
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
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := ContainsKey(tc.a, tc.b); got != tc.want {
				t.Errorf("ContainsKey() = %v, want = %v", got, tc.want)
			}
		})
	}
}

func TestNotContainsKey(t *testing.T) {
	tests := []containsCase{
		{name: "MAP_INT_EXITS_KEY", a: map[int]bool{1: true, 2: false}, b: 1, want: false},
		{name: "MAP_INT_NO_KEY", a: map[int]bool{1: true, 2: false}, b: 3, want: true},
		{name: "MAP_STRING_EXITS_KEY", a: map[string]bool{"one": true, "two": false}, b: "one", want: false},
		{name: "MAP_STRING_NO_KEY", a: map[string]bool{"one": true, "two": false}, b: "three", want: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := NotContainsKey(tc.a, tc.b); got != tc.want {
				t.Errorf("NotContainsKey() = %v, want = %v", got, tc.want)
			}
		})
	}
}

func TestContainsOnSlice(t *testing.T) {
	var tests = []struct {
		input []int
		check func(int) bool
		want  bool
		desc  string
	}{
		{input: []int{1, 2, 3, 4, 5}, check: func(n int) bool { return n == 3 }, want: true, desc: "Element Exists"},
		{input: []int{1, 2, 3, 4, 5}, check: func(n int) bool { return n == 6 }, want: false, desc: "Element Not Exists"},
		{input: []int{}, check: func(n int) bool { return n == 1 }, want: false, desc: "Empty Slice"},
		{input: nil, check: func(n int) bool { return n == 1 }, want: false, desc: "Nil Slice"},
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
		condition func(string) bool
		want      bool
	}{
		{
			name:  "Not Contains string in slice",
			slice: []string{"hello", "world"},
			condition: func(ele string) bool {
				return ele == "test"
			},
			want: true,
		},
		{
			name:  "Contains string in slice",
			slice: []string{"hello", "world"},
			condition: func(ele string) bool {
				return ele == "world"
			},
			want: false,
		},
		{
			name:  "Empty slice",
			slice: []string{},
			condition: func(ele string) bool {
				return ele == "test"
			},
			want: true,
		},
		{
			name:  "Nil slice",
			slice: nil,
			condition: func(ele string) bool {
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
