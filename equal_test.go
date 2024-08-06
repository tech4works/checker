package checker

import (
	"testing"
)

type equalsCase struct {
	name  string
	a, b  any
	c     []any
	want  bool
	panic bool
}

func TestEquals(t *testing.T) {
	cases := []equalsCase{
		{
			name: "Equal Strings",
			a:    "hello",
			b:    "hello",
			want: true,
		},
		{
			name: "Unequal Strings",
			a:    "hello",
			b:    "world",
			want: false,
		},
		{
			name: "Equal Ints",
			a:    42,
			b:    42,
			want: true,
		},
		{
			name: "Unequal Ints",
			a:    42,
			b:    24,
			want: false,
		},
		{
			name: "Equal Structs",
			a:    struct{ name string }{name: "John"},
			b:    struct{ name string }{name: "John"},
			want: true,
		},
		{
			name: "Unequal Structs",
			a:    struct{ name string }{name: "John"},
			b:    struct{ name string }{name: "Doe"},
			want: false,
		},
		{
			name: "Equal Slices",
			a:    []int{1, 2, 3},
			b:    []int{1, 2, 3},
			want: true,
		},
		{
			name: "Unequal Slices",
			a:    []int{1, 2, 3},
			b:    []int{3, 2, 1},
			want: false,
		},
		{
			name: "Equal Maps",
			a:    map[string]int{"one": 1, "two": 2},
			b:    &map[string]int{"one": 1, "two": 2},
			want: true,
		},
		{
			name: "Unequal Maps",
			a:    &map[string]int{"one": 1, "two": 2},
			b:    map[string]int{"one": 1, "two": 3},
			want: false,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equals(tt.a, tt.b); got != tt.want {
				t.Errorf("Equals() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestNotEquals(t *testing.T) {
	testCases := []equalsCase{
		{
			name: "BasicEqual",
			a:    1,
			b:    1,
			want: false,
		},
		{
			name: "BasicNotEqual",
			a:    1,
			b:    2,
			want: true,
		},
		{
			name: "EmptyString",
			a:    "",
			b:    "",
			want: false,
		},
		{
			name: "DifferentStrings",
			a:    "hello",
			b:    "world",
			want: true,
		},
		{
			name: "Nil",
			a:    nil,
			b:    nil,
			want: false,
		},
		{
			name: "NilAndValue",
			a:    nil,
			b:    2,
			want: true,
		},
		{
			name: "ArrayDifferentLength",
			a:    []int{1, 2, 3},
			b:    []int{1, 2, 3, 4},
			want: true,
		},
		{
			name: "ArraySameElement",
			a:    []int{1, 2, 3},
			b:    []int{1, 2, 3},
			want: false,
		},
		{
			name: "ArrayDifferentElement",
			a:    []int{1, 2, 3},
			b:    []int{4, 5, 6},
			want: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := NotEquals(tc.a, tc.b); got != tc.want {
				t.Errorf("\nCase : %v\nGot  : %v\nWant : %v", tc.name, got, tc.want)
			}
		})
	}
}

func TestEqualsIgnoreCase(t *testing.T) {
	strPointer := "JAVA"

	testCases := []equalsCase{
		{
			name: "ASCII lowercase",
			a:    "golang",
			b:    "GOLANG",
			want: true,
		},
		{
			name: "ASCII uppercase",
			a:    "GOLANG",
			b:    "golang",
			want: true,
		},
		{
			name: "Null strings",
			a:    "",
			b:    "",
			want: true,
		},
		{
			name: "Number string",
			a:    "123456",
			b:    "123456",
			want: true,
		},
		{
			name: "Different length strings",
			a:    "golang",
			b:    "GOLANG123",
			want: false,
		},
		{
			name: "Different strings",
			a:    "golang",
			b:    "JAVA",
			want: false,
		},
		{
			name: "Using pointer to string",
			a:    &strPointer,
			b:    "java",
			want: true,
		},
		{
			name: "Using pointer to string",
			a:    "java",
			b:    &strPointer,
			want: true,
		},
		{
			name:  "Nil",
			a:     nil,
			b:     "test",
			panic: true,
		},
		{
			name:  "Number",
			a:     2,
			b:     3,
			panic: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.panic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}()
			}

			if got := EqualsIgnoreCase(tc.a, tc.b); got != tc.want {
				t.Errorf("EqualsIgnoreCase() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestNotEqualsIgnoreCase(t *testing.T) {
	tests := []equalsCase{
		{name: "same case", a: "Test", b: "Test", want: false},
		{name: "different case", a: "Test", b: "test", want: false},
		{name: "different strings", a: "Test", b: "Another", want: true},
		{name: "empty strings", a: "", b: "", want: false},
		{name: "empty and non-empty", a: "", b: "test", want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NotEqualsIgnoreCase(tt.a, tt.b); got != tt.want {
				t.Errorf("NotEqualsIgnoreCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllEquals(t *testing.T) {
	tests := []equalsCase{
		{
			name: "EmptySlice",
			a:    5,
			b:    5,
			want: true,
		},
		{
			name: "SameElements",
			a:    5,
			b:    5,
			c:    []any{5, 5, 5},
			want: true,
		},
		{
			name: "DiffElements",
			a:    5,
			b:    5,
			c:    []any{5, 4, 5},
			want: false,
		},
		{
			name: "WithStringType",
			a:    "test",
			b:    "test",
			c:    []any{"test", "test"},
			want: true,
		},
		{
			name: "WithDiffStringType",
			a:    "test",
			b:    "testy",
			c:    []any{"tet", "test"},
			want: false,
		},
		{
			name: "WithBoolType",
			a:    false,
			b:    false,
			c:    []any{false, false},
			want: true,
		},
		{
			name: "WithDiffBoolType",
			a:    false,
			b:    false,
			c:    []any{false, true},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AllEquals(tt.a, tt.b, tt.c...); got != tt.want {
				t.Errorf("AllEquals() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestNoneEquals(t *testing.T) {
	testCases := []equalsCase{
		{
			name: "Different Types",
			a:    10,
			b:    "40",
			c:    []any{20, 30.5},
			want: true,
		},
		{
			name: "Same Values",
			a:    10,
			b:    10,
			c:    []any{10, 10},
			want: false,
		},
		{
			name: "Mixed Values",
			a:    10,
			b:    20,
			c:    []any{30, "40"},
			want: true,
		},
		{
			name: "Repeated Values",
			a:    "hello",
			b:    "hello",
			c:    []any{"hello", 20},
			want: false,
		},
		{
			name: "Empty Case",
			a:    "",
			b:    "",
			c:    []any{"", ""},
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := NoneEquals(tc.a, tc.b, tc.c...); got != tc.want {
				t.Errorf("NoneEquals(%v, %v, %v) = %v; want %v", tc.a, tc.b, tc.c, got, tc.want)
			}
		})
	}
}
