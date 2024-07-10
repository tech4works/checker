package checker

import (
	"testing"
)

type mockBaseEnum struct {
	isValid bool
}

func (m mockBaseEnum) IsEnumValid() bool {
	return m.isValid
}

func TestIsEnumValid(t *testing.T) {
	tests := []baseCase{
		{
			name: "NonPointerNonInterfaceNonBaseEnum",
			arg:  123,
		},
		{
			name: "PointerToNonBaseEnum",
			arg:  &mockBaseEnum{false},
		},
		{
			name: "InterfaceOfNonBaseEnum",
			arg:  nilInterfaceOfNonBaseEnum(),
		},
		{
			name: "NilPointer",
			arg:  nilPointer(),
		},
		{
			name: "NonNilPointerToBaseEnumInvalid",
			arg:  &mockBaseEnum{true},
			want: true,
		},
		{
			name: "NonNilPointerToBaseEnumValid",
			arg:  &mockBaseEnum{true},
			want: true,
		},
		{
			name: "NilInterface",
			arg:  nilInterface(),
		},
		{
			name: "NonNilInterfaceToBaseEnumInvalid",
			arg:  mockBaseEnum{false},
		},
		{
			name: "NonNilInterfaceToBaseEnumValid",
			arg:  mockBaseEnum{true},
			want: true,
		},
		{
			name: "NonBaseEnum",
			arg:  nonBaseEnum(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEnumValid(tt.arg); got != tt.want {
				t.Errorf("IsEnumValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func nilPointer() any {
	var x *mockBaseEnum
	return x
}

func nilInterface() any {
	var x mockBaseEnum
	return &x
}

func nonBaseEnum() any {
	x := 123
	return x
}

func nilInterfaceOfNonBaseEnum() any {
	var x *int
	return x
}
