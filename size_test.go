package checker

import (
	"testing"
)

type sizeCase struct {
	name string
	a    any
	b    any
	want bool
}

func TestIsGreaterThan(t *testing.T) {
	tests := []sizeCase{
		{
			name: "PositiveNumbers",
			a:    5,
			b:    4,
			want: true,
		},
		{
			name: "NegativeNumbers",
			a:    -2,
			b:    -3,
			want: true,
		},
		{
			name: "ZeroAndPositive",
			a:    0,
			b:    4,
			want: false,
		},
		{
			name: "ZeroAndNegative",
			a:    0,
			b:    -1,
			want: true,
		},
		{
			name: "EqualValues",
			a:    5,
			b:    5,
			want: false,
		},
		{
			name: "FloatingNumbers",
			a:    5.5,
			b:    5.4,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsGreaterThan(tt.a, tt.b); got != tt.want {
				t.Errorf("IsGreaterThan() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestIsGreaterThanOrEqual(t *testing.T) {
	tests := []sizeCase{
		{
			name: "greater_int",
			a:    6,
			b:    5,
			want: true,
		},
		{
			name: "greater_float",
			a:    4.000001245,
			b:    4.00000123,
			want: true,
		},
		{
			name: "less_int",
			a:    4,
			b:    5,
			want: false,
		},
		{
			name: "less_float",
			a:    4.00012,
			b:    4.0023,
			want: false,
		},
		{
			name: "equal_int",
			a:    3312,
			b:    3312,
			want: true,
		},
		{
			name: "equal_float",
			a:    321.23,
			b:    321.23,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsGreaterThanOrEqual(tt.a, tt.b); got != tt.want {
				t.Errorf("IsGreaterThanOrEqual() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestIsLessThan(t *testing.T) {
	tests := []sizeCase{
		{
			name: "greater_int",
			a:    3,
			b:    -5,
			want: false,
		},
		{
			name: "less_int",
			a:    -1,
			b:    2,
			want: true,
		},
		{
			name: "greater_float",
			a:    2.00001,
			b:    2.000001,
			want: false,
		},
		{
			name: "less_float",
			a:    2.0000011,
			b:    2.0000211,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLessThan(tt.a, tt.b); got != tt.want {
				t.Errorf("IsLessThan() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestIsLessThanOrEqual(t *testing.T) {
	var tests = []sizeCase{
		{
			name: "equal_int",
			a:    5,
			b:    5,
			want: true,
		},
		{
			name: "greater_int",
			a:    10,
			b:    5,
			want: false,
		},
		{
			name: "less_int",
			a:    5,
			b:    10,
			want: true,
		},
		{
			name: "equal_float",
			a:    5.0,
			b:    5.0,
			want: true,
		},
		{
			name: "greater_float",
			a:    10.0,
			b:    5.0,
			want: false,
		},
		{
			name: "less_float",
			a:    5.0,
			b:    10.0,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLessThanOrEqual(tt.a, tt.b); got != tt.want {
				t.Errorf("IsLessThanOrEqual() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestIsLengthGreaterThan(t *testing.T) {
	tests := []sizeCase{
		{
			name: "EmptyVsNonEmptyString",
			a:    "",
			b:    "Hello",
			want: false,
		},
		{
			name: "NonEmptyVsEmptyString",
			a:    "Hello, World!",
			b:    "",
			want: true,
		},
		{
			name: "SameSizeStrings",
			a:    "Hello",
			b:    "world",
			want: false,
		},
		{
			name: "EmptyVsNonEmptySlice",
			a:    []int{},
			b:    []int{1, 2, 3},
			want: false,
		},
		{
			name: "NonEmptyVsEmptySlice",
			a:    []int{1, 2, 3, 4},
			b:    []int{},
			want: true,
		},
		{
			name: "SameSizeSlices",
			a:    []int{1, 2, 3},
			b:    []int{4, 5, 6},
			want: false,
		},
		{
			name: "StringVsSlice",
			a:    "Hello",
			b:    []int{1, 2, 3},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLengthGreaterThan(tt.a, tt.b); got != tt.want {
				t.Errorf("IsLengthGreaterThan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsLengthGreaterThanOrEqual(t *testing.T) {
	tests := []sizeCase{
		{
			name: "Same Length",
			a:    "abcdef",
			b:    "uvwxyz",
			want: true,
		},
		{
			name: "Length equal int",
			a:    []string{"abcdef", "uvwxyz"},
			b:    2,
			want: true,
		},
		{
			name: "Length greater int",
			a:    []string{"abcdef", "uvwxyz"},
			b:    1.1,
			want: true,
		},
		{
			name: "Length less int",
			a:    []string{"abcdef", "uvwxyz"},
			b:    10,
			want: false,
		},
		{
			name: "Length A Greater Than B",
			a:    "abcdefg",
			b:    "uvwxy",
			want: true,
		},
		{
			name: "Length A Less Than B",
			a:    "abc",
			b:    "uvwxyz",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLengthGreaterThanOrEqual(tt.a, tt.b); got != tt.want {
				t.Errorf("IsLengthGreaterThanOrEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsLengthLessThan(t *testing.T) {
	tests := []sizeCase{
		{
			name: "LessLengthString",
			a:    "Hello",
			b:    "Hello World",
			want: true,
		},
		{
			name: "EqualLengthString",
			a:    "Hello",
			b:    "World",
			want: false,
		},
		{
			name: "LessLengthSlice",
			a:    []int{1, 2, 3},
			b:    []int{1, 2, 3, 4},
			want: true,
		},
		{
			name: "EqualLengthSlice",
			a:    []int{1, 2, 3, 4},
			b:    []int{1, 2, 3, 4},
			want: false,
		},
		{
			name: "LessLengthMap",
			a:    map[string]int{"one": 1, "two": 2},
			b:    map[string]int{"one": 1, "two": 2, "three": 3},
			want: true,
		},
		{
			name: "EqualLengthMap",
			a:    map[string]int{"one": 1, "two": 2, "three": 3},
			b:    map[string]int{"one": 1, "two": 2, "three": 3},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLengthLessThan(tt.a, tt.b); got != tt.want {
				t.Errorf("IsLengthLessThan() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestIsLengthLessThanOrEqual(t *testing.T) {
	slices := []sizeCase{
		{
			name: "Empty inputs",
			a:    []int{},
			b:    []int{},
			want: true,
		},
		{
			name: "Compared to empty slice",
			a:    []int{1},
			b:    []int{},
			want: false,
		},
		{
			name: "Equal length slices",
			a:    []int{1, 2, 3},
			b:    []int{4, 5, 6},
			want: true,
		},
		{
			name: "a length less than b",
			a:    []int{1, 2},
			b:    []int{1, 2, 3},
			want: true,
		},
		{
			name: "a length greater than b",
			a:    []int{1, 2, 3, 4},
			b:    []int{1, 2, 3},
			want: false,
		},
	}

	for _, tt := range slices {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLengthLessThanOrEqual(tt.a, tt.b); got != tt.want {
				t.Errorf("IsLengthLessThanOrEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
