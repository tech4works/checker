package checker

import (
	"testing"
)

type emptyCase struct {
	name string
	args []any
	want bool
}

func TestIsNil(t *testing.T) {
	for _, tc := range buildIsNilCases() {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsNil(tc.args[0]); got != tc.want {
				t.Errorf("IsNil() = %v, want = %v", got, tc.want)
			}
		})
	}
}

func TestNonNil(t *testing.T) {
	for _, tc := range buildNonNilCases() {
		t.Run(tc.name, func(t *testing.T) {
			if got := NonNil(tc.args[0]); got != tc.want {
				t.Errorf("NonNil() = %v, want = %v", got, tc.want)
			}
		})
	}
}

func TestAllNil(t *testing.T) {
	for _, tt := range buildAllNilCases() {
		t.Run(tt.name, func(t *testing.T) {
			if got := AllNil(tt.args[0], tt.args[1]); got != tt.want {
				t.Errorf("AllNil() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestAllNonNil(t *testing.T) {
	for _, tt := range buildAllNonNilCases() {
		t.Run(tt.name, func(t *testing.T) {
			if got := AllNonNil(tt.args[0], tt.args[1]); got != tt.want {
				t.Errorf("AllNonNil() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	for _, tc := range buildIsEmptyCases() {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsEmpty(tc.args[0]); got != tc.want {
				t.Errorf("IsEmpty() = %v, want = %v", got, tc.want)
			}
		})
	}
}

func TestIsNotEmpty(t *testing.T) {
	for _, tc := range buildIsNotEmptyCases() {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsNotEmpty(tc.args[0]); got != tc.want {
				t.Errorf("IsNotEmpty() = %v, want = %v", got, tc.want)
			}
		})
	}
}

func TestAllEmpty(t *testing.T) {
	for _, tc := range buildAllEmptyCases() {
		t.Run(tc.name, func(t *testing.T) {
			if got := AllEmpty(tc.args[0], tc.args[1]); got != tc.want {
				t.Errorf("AllEmpty() = %v, want = %v", got, tc.want)
			}
		})
	}
}

func TestAllNotEmpty(t *testing.T) {
	for _, tc := range buildAllNotEmptyCases() {
		t.Run(tc.name, func(t *testing.T) {
			if got := AllNotEmpty(tc.args[0], tc.args[1]); got != tc.want {
				t.Errorf("AllNotEmpty() = %v, want = %v", got, tc.want)
			}
		})
	}
}

func TestIsNilOrEmpty(t *testing.T) {
	for _, tc := range buildIsNilOrEmptyCases() {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsNilOrEmpty(tc.args[0]); got != tc.want {
				t.Errorf("IsNilOrEmpty() = %v, want = %v", got, tc.want)
			}
		})
	}
}

func TestIsNotNilOrEmpty(t *testing.T) {
	for _, tc := range buildIsNotNilOrEmptyCases() {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsNotNilOrEmpty(tc.args[0]); got != tc.want {
				t.Errorf("IsNotNilOrEmpty() = %v, want = %v", got, tc.want)
			}
		})
	}
}

func TestAllNilOrEmpty(t *testing.T) {
	for _, tc := range buildAllNilOrEmptyCases() {
		t.Run(tc.name, func(t *testing.T) {
			if got := AllNilOrEmpty(tc.args[0], tc.args[1]); got != tc.want {
				t.Errorf("AllNilOrEmpty() = %v, want = %v", got, tc.want)
			}
		})
	}
}

func TestAllNotNilOrEmpty(t *testing.T) {
	for _, tc := range buildAllNotNilOrEmptyCases() {
		t.Run(tc.name, func(t *testing.T) {
			if got := AllNotNilOrEmpty(tc.args[0], tc.args[1]); got != tc.want {
				t.Errorf("AllNotNilOrEmpty() = %v, want = %v", got, tc.want)
			}
		})
	}
}

func buildIsNilCases() []emptyCase {
	return []emptyCase{
		{name: "IntNil", args: []any{(*int)(nil)}, want: true},
		{name: "FloatNil", args: []any{(*float64)(nil)}, want: true},
		{name: "StringNil", args: []any{(*string)(nil)}, want: true},
		{name: "SliceNil", args: []any{([]int)(nil)}, want: true},
		{name: "MapNil", args: []any{(map[string]int)(nil)}, want: true},
		{name: "PointerNil", args: []any{(*struct{})(nil)}, want: true},
		{name: "InterfaceNil", args: []any{(interface{})(nil)}, want: true},
		{name: "ChannelNil", args: []any{(chan struct{})(nil)}, want: true},
		{name: "FunctionNil", args: []any{(func())(nil)}, want: true},
		{name: "IntNonNil", args: []any{new(int)}, want: false},
		{name: "FloatNonNil", args: []any{new(float64)}, want: false},
		{name: "StringNonNil", args: []any{new(string)}, want: false},
		{name: "SliceNonNil", args: []any{[]int{}}, want: false},
		{name: "MapNonNil", args: []any{map[string]int{}}, want: false},
		{name: "PointerNonNil", args: []any{new(struct{})}, want: false},
		{name: "InterfaceNonNil", args: []any{new(interface{})}, want: false},
		{name: "ChannelNonNil", args: []any{make(chan struct{})}, want: false},
		{name: "FunctionNonNil", args: []any{func() {}}, want: false},
	}
}

func buildNonNilCases() []emptyCase {
	return []emptyCase{
		{name: "IntNonNil", args: []any{new(int)}, want: true},
		{name: "FloatNonNil", args: []any{new(float64)}, want: true},
		{name: "StringNonNil", args: []any{new(string)}, want: true},
		{name: "SliceNonNil", args: []any{[]int{}}, want: true},
		{name: "MapNonNil", args: []any{map[string]int{}}, want: true},
		{name: "PointerNonNil", args: []any{new(struct{})}, want: true},
		{name: "InterfaceNonNil", args: []any{new(interface{})}, want: true},
		{name: "ChannelNonNil", args: []any{make(chan struct{})}, want: true},
		{name: "FunctionNonNil", args: []any{func() {}}, want: true},
		{name: "IntNil", args: []any{(*int)(nil)}, want: false},
		{name: "FloatNil", args: []any{(*float64)(nil)}, want: false},
		{name: "StringNil", args: []any{(*string)(nil)}, want: false},
		{name: "SliceNil", args: []any{([]int)(nil)}, want: false},
		{name: "MapNil", args: []any{(map[string]int)(nil)}, want: false},
		{name: "PointerNil", args: []any{(*struct{})(nil)}, want: false},
		{name: "InterfaceNil", args: []any{(interface{})(nil)}, want: false},
		{name: "ChannelNil", args: []any{(chan struct{})(nil)}, want: false},
		{name: "FunctionNil", args: []any{(func())(nil)}, want: false},
	}
}

func buildAllNilCases() []emptyCase {
	return []emptyCase{
		{
			name: "AllNil",
			args: []any{(*int)(nil), (*float64)(nil), ([]int)(nil), (*struct{})(nil)},
			want: true,
		},
		{
			name: "SomeNil",
			args: []any{(*int)(nil), make(chan struct{}), new(float64)},
			want: false},
	}
}

func buildAllNonNilCases() []emptyCase {
	return []emptyCase{
		{name: "AllNonNil", args: []any{make(chan struct{}), new(float64)}, want: true},
		{name: "AllNil", args: []any{(*int)(nil), (*float64)(nil), ([]int)(nil), (*struct{})(nil)}, want: false},
	}
}

func buildIsEmptyCases() []emptyCase {
	return []emptyCase{
		{name: "IntNil", args: []any{(*int)(nil)}, want: true},
		{name: "FloatNil", args: []any{(*float64)(nil)}, want: true},
		{name: "StringEmpty", args: []any{""}, want: true},
		{name: "SliceEmpty", args: []any{[]int{}}, want: true},
		{name: "MapEmpty", args: []any{map[string]int{}}, want: true},
		{name: "PointerNil", args: []any{(*struct{})(nil)}, want: true},
		{name: "InterfaceNil", args: []any{(interface{})(nil)}, want: true},
		{name: "ChannelNil", args: []any{(chan struct{})(nil)}, want: true},
		{name: "FunctionNil", args: []any{(func())(nil)}, want: true},
		{name: "IntNonEmpty", args: []any{2}, want: false},
		{name: "FloatNonEmpty", args: []any{2.3}, want: false},
		{name: "StringNonEmpty", args: []any{"Hello World"}, want: false},
		{name: "SliceNonEmpty", args: []any{[]int{1, 2, 3}}, want: false},
		{name: "MapNonEmpty", args: []any{map[string]int{"one": 1}}, want: false},
		{name: "PointerNonNil", args: []any{new(struct{})}, want: true},
		{name: "InterfaceNonNil", args: []any{new(interface{})}, want: true},
		{name: "ChannelNonNil", args: []any{make(chan struct{})}, want: false},
		{name: "FunctionNonNil", args: []any{func() {}}, want: false},
	}
}

func buildIsNotEmptyCases() []emptyCase {
	return []emptyCase{
		{name: "FunctionNonNil", args: []any{func() {}}, want: true},
		{name: "ChannelNonNil", args: []any{make(chan struct{})}, want: true},
		{name: "InterfaceNonNil", args: []any{new(interface{})}, want: false},
		{name: "PointerNonNil", args: []any{new(struct{})}, want: false},
		{name: "MapNonEmpty", args: []any{map[string]int{"one": 1}}, want: true},
		{name: "SliceNonEmpty", args: []any{[]int{1, 2, 3}}, want: true},
		{name: "StringNonEmpty", args: []any{"Hello World"}, want: true},
		{name: "FloatNonEmpty", args: []any{2.5}, want: true},
		{name: "IntNonEmpty", args: []any{11}, want: true},
		{name: "FunctionNil", args: []any{(func())(nil)}, want: false},
		{name: "ChannelNil", args: []any{(chan struct{})(nil)}, want: false},
		{name: "InterfaceNil", args: []any{(interface{})(nil)}, want: false},
		{name: "PointerNil", args: []any{(*struct{})(nil)}, want: false},
		{name: "MapEmpty", args: []any{map[string]int{}}, want: false},
		{name: "SliceEmpty", args: []any{[]int{}}, want: false},
		{name: "StringEmpty", args: []any{""}, want: false},
		{name: "FloatNil", args: []any{(*float64)(nil)}, want: false},
		{name: "IntNil", args: []any{(*int)(nil)}, want: false},
	}
}

func buildAllEmptyCases() []emptyCase {
	return []emptyCase{
		{name: "EmptyString", args: []any{"", "", ""}, want: true},
		{name: "NonEmptyString", args: []any{"test", "dsd", "123"}, want: false},
		{name: "EmptyIntSlice", args: []any{[]int{}, []int{}, []int{}}, want: true},
		{name: "NonEmptyIntSlice", args: []any{[]int{1, 2}, []int{1, 2}, []int{1, 2}}, want: false},
		{name: "EmptyStringSlice", args: []any{[]string{}, []string{}, []string{}}, want: true},
		{name: "NonEmptyStringSlice", args: []any{[]string{"test"}, []string{"test"}, []string{"test"}}, want: false},
		{name: "EmptyStringMap", args: []any{map[string]string{}, map[string]string{}, map[string]string{}}, want: true},
		{name: "NonEmptyStringMap", args: []any{map[string]string{"key": "value"}, map[string]string{"key": "value"}, map[string]string{"key": "value"}}, want: false},
		{name: "EmptyStruct", args: []any{struct{}{}, struct{}{}, struct{}{}}, want: true},
		{name: "NonEmptyStruct", args: []any{struct{ Name string }{Name: "test"}, struct{ Name string }{Name: "test"}, struct{ Name string }{Name: "test"}}, want: false},
	}
}

func buildAllNotEmptyCases() []emptyCase {
	return []emptyCase{
		{name: "NonEmptyString", args: []any{"test", "test"}, want: true},
		{name: "EmptyString", args: []any{"", ""}, want: false},
		{name: "NonEmptyIntSlice", args: []any{[]int{1}, []int{1}}, want: true},
		{name: "EmptyIntSlice", args: []any{[]int{}, []int{}}, want: false},
		{name: "NonEmptyStringSlice", args: []any{[]string{"test"}, []string{"test"}}, want: true},
		{name: "EmptyStringSlice", args: []any{[]string{}, []string{}}, want: false},
		{name: "NonEmptyStringMap", args: []any{map[string]string{"key": "value"}, map[string]string{"key": "value"}}, want: true},
		{name: "EmptyStringMap", args: []any{map[string]string{}, map[string]string{}}, want: false},
		{name: "NonEmptyStruct", args: []any{struct{ Name string }{Name: "test"}, struct{ Name string }{Name: "test"}}, want: true},
		{name: "EmptyStruct", args: []any{struct{}{}, struct{}{}}, want: false},
	}
}

func buildIsNilOrEmptyCases() []emptyCase {
	return []emptyCase{
		{name: "IntNil", args: []any{(*int)(nil)}, want: true},
		{name: "FloatNil", args: []any{(*float64)(nil)}, want: true},
		{name: "StringEmpty", args: []any{""}, want: true},
		{name: "SliceEmpty", args: []any{[]int{}}, want: true},
		{name: "MapEmpty", args: []any{map[string]int{}}, want: true},
		{name: "PointerNil", args: []any{(*struct{})(nil)}, want: true},
		{name: "InterfaceNil", args: []any{(interface{})(nil)}, want: true},
		{name: "ChannelNil", args: []any{(chan struct{})(nil)}, want: true},
		{name: "FunctionNil", args: []any{(func())(nil)}, want: true},
		{name: "StringNil", args: []any{(*string)(nil)}, want: true},
		{name: "SliceNil", args: []any{([]int)(nil)}, want: true},
		{name: "MapNil", args: []any{(map[string]int)(nil)}, want: true},
		{name: "IntNonEmpty", args: []any{3}, want: false},
		{name: "FloatNonEmpty", args: []any{102.23}, want: false},
		{name: "StringNonEmpty", args: []any{"Hello World"}, want: false},
		{name: "SliceNonEmpty", args: []any{[]int{1, 2, 3}}, want: false},
		{name: "MapNonEmpty", args: []any{map[string]int{"one": 1}}, want: false},
		{name: "PointerNonNil", args: []any{new(struct{})}, want: true},
		{name: "InterfaceNonNil", args: []any{new(interface{})}, want: true},
		{name: "ChannelNonNil", args: []any{make(chan struct{})}, want: false},
		{name: "FunctionNonNil", args: []any{func() {}}, want: false},
	}
}

func buildIsNotNilOrEmptyCases() []emptyCase {
	return []emptyCase{
		{name: "IntNil", args: []any{(*int)(nil)}, want: false},
		{name: "FloatNil", args: []any{(*float64)(nil)}, want: false},
		{name: "StringEmpty", args: []any{""}, want: false},
		{name: "SliceEmpty", args: []any{[]int{}}, want: false},
		{name: "MapEmpty", args: []any{map[string]int{}}, want: false},
		{name: "PointerNil", args: []any{(*struct{})(nil)}, want: false},
		{name: "InterfaceNil", args: []any{(interface{})(nil)}, want: false},
		{name: "ChannelNil", args: []any{(chan struct{})(nil)}, want: false},
		{name: "FunctionNil", args: []any{(func())(nil)}, want: false},
		{name: "StringNil", args: []any{(*string)(nil)}, want: false},
		{name: "SliceNil", args: []any{([]int)(nil)}, want: false},
		{name: "MapNil", args: []any{(map[string]int)(nil)}, want: false},
		{name: "IntNonEmpty", args: []any{12}, want: true},
		{name: "FloatNonEmpty", args: []any{2.3}, want: true},
		{name: "StringNonEmpty", args: []any{"Hello World"}, want: true},
		{name: "SliceNonEmpty", args: []any{[]int{1, 2, 3}}, want: true},
		{name: "MapNonEmpty", args: []any{map[string]int{"one": 1}}, want: true},
		{name: "PointerNonNil", args: []any{new(struct{})}, want: false},
		{name: "InterfaceNonNil", args: []any{new(interface{})}, want: false},
		{name: "ChannelNonNil", args: []any{make(chan struct{})}, want: true},
		{name: "FunctionNonNil", args: []any{func() {}}, want: true},
	}
}

func buildAllNilOrEmptyCases() []emptyCase {
	return []emptyCase{
		{
			name: "AllNilOrEmpty",
			args: []any{
				(*int)(nil),
				"",
				[]int{},
				map[string]int{},
				(*struct{})(nil),
				(interface{})(nil),
				(chan struct{})(nil),
				(func())(nil),
			},
			want: true,
		},
		{
			name: "SomeNilOrEmpty",
			args: []any{
				(*int)(nil),
				"NonEmpty",
				[]int{1, 2, 3},
				map[string]int{"one": 1},
				new(struct{}),
				new(interface{}),
				make(chan struct{}),
				func() {},
			},
			want: false,
		},
	}
}

func buildAllNotNilOrEmptyCases() []emptyCase {
	return []emptyCase{
		{
			name: "AllNotNullOrNotEmpty",
			args: []any{
				10,
				"NotEmpty",
				[]int{1, 2, 3},
				map[string]int{"one": 1},
				new(struct{}),
				new(interface{}),
				make(chan struct{}),
				func() {},
			},
			want: true,
		},
		{
			name: "SomeNotNullOrNotEmpty",
			args: []any{
				(*int)(nil),
				"",
				[]int{},
				map[string]int{},
				(*struct{})(nil),
				(interface{})(nil),
				(chan struct{})(nil),
				(func())(nil),
			},
			want: false,
		},
	}
}
