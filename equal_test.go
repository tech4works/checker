package checker

import (
	"testing"
)

type equalsCase struct {
	name string
	a, b any
	c    []any
	want bool
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
			b:    map[string]int{"one": 1, "two": 2},
			want: true,
		},
		{
			name: "Unequal Maps",
			a:    map[string]int{"one": 1, "two": 2},
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

func TestIsNotEqualTo(t *testing.T) {
	testCases := []equalsCase{
		{
			name: "integersNotEqual",
			a:    1,
			b:    2,
			want: true,
		},
		{
			name: "integersEqual",
			a:    1,
			b:    1,
			want: false,
		},
		{
			name: "stringsNotEqual",
			a:    "Hello",
			b:    "World",
			want: true,
		},
		{
			name: "stringsEqual",
			a:    "Hello",
			b:    "Hello",
			want: false,
		},
		{
			name: "structuresNotEqual",
			a:    struct{ name string }{"Hello"},
			b:    struct{ name string }{"World"},
			want: true,
		},
		{
			name: "structuresEqual",
			a:    struct{ name string }{"Hello"},
			b:    struct{ name string }{"Hello"},
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := IsNotEqualTo(tc.a, tc.b)
			if got != tc.want {
				t.Errorf("IsNotEqualTo() = %v, want = %v", got, tc.want)
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
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := EqualsIgnoreCase(tc.a, tc.b); got != tc.want {
				t.Errorf("EqualsIgnoreCase() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestIsNotEqualIgnoreCaseTo(t *testing.T) {
	tests := []equalsCase{
		{
			name: "Different case strings",
			a:    "Hello",
			b:    "HELLO",
			want: false,
		},
		{
			name: "Same case strings",
			a:    "Hello",
			b:    "Hello",
			want: false,
		},
		{
			name: "Different case different strings",
			a:    "Hello",
			b:    "World",
			want: true,
		},
		{
			name: "Different case strings numbers",
			a:    "123",
			b:    "123",
			want: false,
		},
		{
			name: "Special chars same strings",
			a:    "@#$%^&*()",
			b:    "@#$%^&*()",
			want: false,
		},
		{
			name: "Empty strings",
			a:    "",
			b:    "",
			want: false,
		},
		{
			name: "Different strings including empty string",
			a:    "Hello",
			b:    "",
			want: true,
		},
		{
			name: "Different strings including whitespace",
			a:    "Hello",
			b:    " ",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNotEqualIgnoreCaseTo(tt.a, tt.b); got != tt.want {
				t.Errorf("IsNotEqualIgnoreCaseTo() = %v, want = %v", got, tt.want)
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

func TestAllNotEquals(t *testing.T) {
	tests := []equalsCase{
		{
			name: "AllDifferent",
			a:    1,
			b:    2,
			c:    []any{3, 4, 5},
			want: true,
		},
		{
			name: "SomeEqual",
			a:    1,
			b:    1,
			c:    []any{2, 3, 4},
			want: true,
		},
		{
			name: "DifferentTypes",
			a:    1,
			b:    "1",
			c:    []any{2.0, "3", struct{}{}},
			want: true,
		},
		{
			name: "AllEqual",
			a:    "test",
			b:    "test",
			c:    []any{"test", "test", "test"},
			want: false,
		},
		{
			name: "NilArguments",
			a:    nil,
			b:    nil,
			c:    []any{nil, nil, nil},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AllNotEquals(tt.a, tt.b, tt.c...); got != tt.want {
				t.Errorf("AllNotEquals() = %v, want = %v", got, tt.want)
			}
		})
	}
}
