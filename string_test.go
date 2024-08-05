package checker

import (
	"net/http"
	"testing"
)

func TestIsURL(t *testing.T) {
	url := "https://example.com"
	testCases := []baseCase{
		{
			name: "ValidHTTPURL",
			arg:  "http://example.com",
			want: true,
		},
		{
			name: "ValidHTTPSURL",
			arg:  "https://example.com",
			want: true,
		},
		{
			name: "ValidURLWithPort",
			arg:  "http://localhost:8080",
			want: true,
		},
		{
			name: "ValidURLPointer",
			arg:  &url,
			want: true,
		},
		{
			name: "InvalidURLWithoutScheme",
			arg:  "example.com",
			want: false,
		},
		{
			name: "InvalidURLWithInvalidCharacters",
			arg:  "http://exa%mple.com",
			want: false,
		},
		{
			name: "EmptyString",
			arg:  "",
			want: false,
		},
		{
			name: "NonStringArgument",
			arg:  1234,
			want: false,
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

			if got := IsURL(tc.arg); got != tc.want {
				t.Errorf("IsURL(%v) = %v, want %v", tc.arg, got, tc.want)
			}
		})
	}
}

func TestIsURLPath(t *testing.T) {
	testCases := []baseCase{
		{
			name: "path with one directory",
			arg:  "/home",
			want: true,
		},
		{
			name: "path with multiple directories",
			arg:  "/home/user/documents",
			want: true,
		},
		{
			name: "root path",
			arg:  "/",
			want: true,
		},
		{
			name: "empty string",
			arg:  "",
			want: false,
		},
		{
			name: "path with space",
			arg:  "/home/ user/documents",
			want: false,
		},
		{
			name: "path with special characters",
			arg:  "/home/user@/documents",
			want: true,
		},
		{
			name: "path with query string",
			arg:  "/home/user/documents?file=1",
			want: false,
		},
		{
			name: "non string",
			arg:  23,
			want: false,
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

			if got := IsURLPath(tc.arg); got != tc.want {
				t.Errorf("IsURLPath(%v) = %v, want %v", tc.arg, got, tc.want)
			}
		})
	}
}

func TestIsHTTPMethod(t *testing.T) {
	testCases := []baseCase{
		{
			name: "Get method",
			arg:  http.MethodGet,
			want: true,
		},
		{
			name: "Post method",
			arg:  http.MethodPost,
			want: true,
		},
		{
			name: "Unknown method",
			arg:  "UNKNOWN",
			want: false,
		},
		{
			name: "Empty string",
			arg:  "",
			want: false,
		},
		{
			name: "Non-string input",
			arg:  12345,
			want: false,
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

			if got := IsHTTPMethod(tc.arg); got != tc.want {
				t.Errorf("IsHTTPMethod(%v) = %v, want %v", tc.arg, got, tc.want)
			}
		})
	}
}

func TestIsAlpha(t *testing.T) {
	tests := []baseCase{
		{
			name: "All Alphabetic Characters",
			arg:  "ABCDE",
			want: true,
		},
		{
			name: "Alphabetic with Numeric Characters",
			arg:  "ALPHA100",
			want: false,
		},
		{
			name: "Empty String",
			arg:  "",
			want: false,
		},
		{
			name: "Alphabetic with Space",
			arg:  "GO LAND",
			want: false,
		},
		{
			name: "Alphabetic with Special Characters",
			arg:  "ABC@#",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAlpha(tt.arg); got != tt.want {
				t.Errorf("IsAlpha() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAlphaSpace(t *testing.T) {
	testCases := []baseCase{
		{
			name: "AllAlphabets",
			arg:  "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
			want: true,
		},
		{
			name: "AlphabetsWithSpace",
			arg:  "This is a test",
			want: true,
		},
		{
			name: "AlphabetsWithNumbers",
			arg:  "Test123",
			want: false,
		},
		{
			name: "AlphabetsWithSpecialChars",
			arg:  "This is @ test",
			want: false,
		},
		{
			name: "EmptyString",
			arg:  "",
			want: false,
		},
		{
			name: "OnlySpaces",
			arg:  "    ",
			want: false,
		},
		{
			name: "NumbersOnly",
			arg:  "1234567890",
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsAlphaSpace(tc.arg); got != tc.want {
				t.Errorf("IsAlphaSpace() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestIsNumeric(t *testing.T) {
	var p *any = nil
	tests := []baseCase{
		{name: "empty string", arg: ""},
		{name: "letters", arg: "abc"},
		{name: "numeric string", arg: "123", want: true},
		{name: "numeric string with spaces", arg: "1 23"},
		{name: "special characters", arg: "!#$%"},
		{name: "alphanumeric", arg: "1a2b3c"},
		{name: "negative integer", arg: -123, want: true},
		{name: "positive float", arg: 123.456, want: true},
		{name: "zero", arg: 0, want: true},
		{name: "uint", arg: uint(23), want: true},
		{name: "complex", arg: complex(23.2, 231), want: false},
		{name: "pointer nil", arg: p, panic: true},
		{name: "nil", panic: true},
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

			if got := IsNumeric(tt.arg); got != tt.want {
				t.Errorf("IsNumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNumericSpace(t *testing.T) {
	tests := []baseCase{
		{
			name: "onlyDigits",
			arg:  "12345678",
			want: true,
		},
		{
			name: "onlySpaces",
			arg:  "        ",
			want: false,
		},
		{
			name: "mixedDigitsAndSpaces",
			arg:  " 345 67 89 ",
			want: true,
		},
		{
			name: "emptyString",
			arg:  "",
			want: false,
		},
		{
			name: "alphabeticalCharacters",
			arg:  "abcd",
			want: false,
		},
		{
			name: "specialCharacters",
			arg:  "()(*&#$%^&",
			want: false,
		},
		{
			name: "digitsWithAlphabets",
			arg:  "123xyz",
			want: false,
		},
		{
			name: "digitsAndSpacesWithAlphabets",
			arg:  "1 2 3xyz",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNumericSpace(tt.arg); got != tt.want {
				t.Errorf("IsNumericSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsEmail(t *testing.T) {
	tests := []baseCase{
		{name: "ValidEmail", arg: "example@test.com", want: true},
		{name: "EmailWithSubdomain", arg: "example.test@test.com", want: true},
		{name: "NoTopLevelDomain", arg: "example@test"},
		{name: "MissingLocalPart", arg: "@test.com"},
		{name: "MissingDomain", arg: "test@.com"},
		{name: "NotAnEmail", arg: "test.test"},
		{name: "MissingAtSymbol", arg: "test"},
		{name: "EmptyString", arg: ""},
		{name: "MissingLocalPartAndTopLevelDomain", arg: "@test."},
		{name: "EndsWithDot", arg: "test.test@com."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmail(tt.arg); got != tt.want {
				t.Errorf("IsEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsDocument(t *testing.T) {
	tests := []struct {
		name         string
		documentType Document
		a            any
		want         bool
		panic        bool
	}{
		{
			name:         "Test_for_Document_Type_CPF",
			documentType: DocumentCPF,
			a:            "37769361001",
			want:         true,
		},
		{
			name:         "Test_for_Document_Type_CNPJ",
			documentType: DocumentCNPJ,
			a:            "53.618.253/0001-90",
			want:         true,
		},
		{
			name:         "Test_for_Document_Type_Not_Valid",
			documentType: Document("Not Valid"),
			a:            "123456",
			panic:        true,
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

			if got := IsDocument(tt.documentType, tt.a); got != tt.want {
				t.Errorf("IsDocument() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsCPF(t *testing.T) {
	tests := []baseCase{
		{
			name: "ValidCPF",
			arg:  "12101721007",
			want: true,
		},
		{
			name: "InvalidCPF",
			arg:  "11111111111",
			want: false,
		},
		{
			name: "ValidCPFWithSpecialChars",
			arg:  "891.595.290-16",
			want: true,
		},
		{
			name: "EmptyCPF",
			arg:  "",
			want: false,
		},
		{
			name: "NonDigitCharsinCPF",
			arg:  "abcdefgijklm",
			want: false,
		},
		{
			name: "IncorrectLength",
			arg:  "1234567",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsCPF(tt.arg); got != tt.want {
				t.Errorf("IsCPF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsCNPJ(t *testing.T) {
	tests := []baseCase{
		{
			name: "ValidCNPJ",
			arg:  "57309623000168",
			want: true,
		},
		{
			name: "ValidCNPJ",
			arg:  "47.263.759/0001-20",
			want: true,
		},
		{
			name: "InvalidCNPJWrongDigits",
			arg:  "00.000.000/0002-90",
			want: false,
		},
		{
			name: "InvalidCNPJAllEqualDigits",
			arg:  "11.111.111/1111-11",
			want: false,
		},
		{
			name: "InvalidCNPJNonNumericCharacters",
			arg:  "33.041.260/065X-90",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsCNPJ(tt.arg); got != tt.want {
				t.Errorf("IsCNPJ() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsCPFOrCNPJ(t *testing.T) {
	testCases := []baseCase{
		{
			name: "Expect success when input is a valid CPF",
			arg:  "58033473029",
			want: true,
		},
		{
			name: "Expect success when input is a valid CPF",
			arg:  "793.786.510-54",
			want: true,
		},
		{
			name: "Expect success when input is a valid CNPJ",
			arg:  "22218636000147",
			want: true,
		},
		{
			name: "Expect success when input is a valid CNPJ",
			arg:  "13.295.729/0001-84",
			want: true,
		},
		{
			name: "Expect failure when input is neither CPF nor CNPJ",
			arg:  "111.222.333-23",
			want: false,
		},
		{
			name: "Expect failure when input is empty string",
			arg:  "",
			want: false,
		},
		{
			name:  "Expect panic nil input",
			arg:   nil,
			panic: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); (r != nil) != tc.panic {
					t.Errorf("IsCPFOrCNPJ() panic = %v, wantPanic = %v", r, tc.panic)
				}
			}()
			got := IsCPFOrCNPJ(tc.arg)
			if got != tc.want {
				t.Errorf("IsCPFOrCNPJ() = %v, want = %v", got, tc.want)
			}
		})
	}
}

func TestIsBase64(t *testing.T) {
	tests := []baseCase{
		{
			name: "ValidBase64",
			arg:  "SGVsbG8gd29ybGQ=",
			want: true,
		},
		{
			name: "InvalidBase64",
			arg:  "Hello world",
			want: false,
		},
		{
			name: "EmptyString",
			arg:  "",
			want: false,
		},
		{
			name:  "NilInput",
			arg:   nil,
			panic: true,
		},
		{
			name: "IntegerInput",
			arg:  12345,
			want: false,
		},
		{
			name: "SpecialCharacters",
			arg:  "!@#$%^&*()_+",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panic {
				defer func() {
					if r := recover(); (r != nil) != tt.panic {
						t.Errorf("IsBase64() panic = %v, wantPanic = %v", r, tt.panic)
					}
				}()
			}

			if got := IsBase64(tt.arg); got != tt.want {
				t.Errorf("IsBase64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsBearer(t *testing.T) {
	tests := []baseCase{
		{
			name: "EmptyStr",
			arg:  "",
			want: false,
		},
		{
			name: "NonString",
			arg:  123,
			want: false,
		},
		{
			name: "NotBearer",
			arg:  "Pizza",
			want: false,
		},
		{
			name: "BearerWithoutSpace",
			arg:  "BearerPizza",
			want: false,
		},
		{
			name: "BearerWithExtraSpaces",
			arg:  "Bearer  Pizza",
			want: true,
		},
		{
			name: "Bearer",
			arg:  "Bearer Pizza",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBearer(tt.arg); got != tt.want {
				t.Errorf("IsBearer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPrivateIP(t *testing.T) {
	tests := []baseCase{
		{"Empty", "", false, false},
		{"Public IPv4", "8.8.8.8", false, false},
		{"Reserved IPv4", "192.168.0.1", true, false},
		{"Loopback IPv4", "127.0.0.1", true, false},
		{"Public IPv6", "2001:4860:4860::8888", false, false},
		{"Loopback IPv6", "::1", true, false},
		{"Invalid", "invalid", false, false},
		{"Empty", "", false, false},
		{"Nil", nil, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got bool
			var gotPanic bool
			func() {
				defer func() {
					if r := recover(); r != nil {
						gotPanic = true
					}
				}()
				got = IsPrivateIP(tt.arg)
			}()
			if got != tt.want {
				t.Errorf("IsPrivateIP() = %v, want %v", got, tt.want)
			}
			if gotPanic != tt.panic {
				t.Errorf("IsPrivateIP() panicked = %v, want panic = %v", gotPanic, tt.panic)
			}
		})
	}
}

func TestIsFullName(t *testing.T) {
	tests := []baseCase{
		{name: "Valid Full Name", arg: "John Doe", want: true},
		{name: "Valid Full Name with special char", arg: "John O'Conner", want: true},
		{name: "Full Name with numeric char", arg: "John2 Doe", want: false},
		{name: "Full Name with special char", arg: "John@ Doe", want: false},
		{name: "Empty Full Name", arg: "", want: false},
		{name: "Whitespace Full Name", arg: "   ", want: false},
		{name: "Full Name with underscore", arg: "John_Doe", want: false},
		{name: "Single word Full Name", arg: "John", want: false},
		{name: "Integer Full Name", arg: 123456, want: false},
		{name: "Non-string Full Name", arg: []int{1, 2, 3}, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFullName(tt.arg); got != tt.want {
				t.Errorf("IsFullName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsIOSDeviceID(t *testing.T) {
	tests := []baseCase{
		{
			name: "Valid IOS device Id",
			arg:  "E241F78F-9477-42B5-A452-2F31E7F20E62",
			want: true,
		},
		{
			name: "IOS device Id with lowercase",
			arg:  "e241f78f-9477-42b5-a452-2f31e7f20e62",
			want: false,
		},
		{
			name: "Invalid length",
			arg:  "E241F78F-9477-42B5-A452-2F31E7F20E62-extra",
			want: false,
		},
		{
			name: "Missing sections",
			arg:  "E241F78F-9477-A452-2F31E7F20E62",
			want: false,
		},
		{
			name: "Empty string",
			arg:  "",
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsIOSDeviceID(tc.arg)
			if got != tc.want {
				t.Fatalf("IsIOSDeviceID() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestIsAndroidDeviceID(t *testing.T) {
	tests := []baseCase{
		{
			name:  "nil input",
			arg:   nil,
			panic: true,
		},
		{
			name: "empty string",
			arg:  "",
			want: false,
		},
		{
			name: "non-string input",
			arg:  123,
			want: false,
		},
		{
			name: "non-hex string",
			arg:  "this-is-not-hex",
			want: false,
		},
		{
			name: "lower case hex string is android device id",
			arg:  "abcdef0123456789",
			want: true,
		},
		{
			name: "upper case hex string is android device id",
			arg:  "ABCDEF0123456789",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); (r != nil) != tt.panic {
					t.Errorf("IsAndroidDeviceID() panic = %v, wantPanic = %v", r, tt.panic)
				}
			}()

			got := IsAndroidDeviceID(tt.arg)
			if got != tt.want {
				t.Errorf("IsAndroidDeviceId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsMobileDeviceID(t *testing.T) {
	tests := []baseCase{
		{
			name: "EmptyString",
			arg:  "",
			want: false,
		},
		{
			name: "AndroidDeviceId",
			arg:  "abcdef0123456789",
			want: true,
		},
		{
			name: "IOSDeviceId",
			arg:  "BA718E20-55BB-4462-B04A-5B372F352124",
			want: true,
		},
		{
			name: "InvalidDeviceId",
			arg:  "invalid-id",
			want: false,
		},
		{
			name: "OtherDataType",
			arg:  12345,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsMobileDeviceID(tt.arg); got != tt.want {
				t.Errorf("IsMobileDeviceID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsMobilePlatform(t *testing.T) {
	testCases := []baseCase{
		{name: "AndroidPlatform", arg: "Android", want: true},
		{name: "iOSPlatform", arg: "iOS", want: true},
		{name: "iPhoneOS", arg: "iPhone Os", want: true},
		{name: "UnrelatedPlatform", arg: "Unrelated", want: false},
		{name: "EmptyInput", arg: "", want: false},
		{name: "NonStringInput", arg: 123, want: false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if result := IsMobilePlatform(tc.arg); result != tc.want {
				t.Errorf("IsMobilePlatform() = %v, want %v", result, tc.want)
			}
		})
	}
}
