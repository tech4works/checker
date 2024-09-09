//	MIT License
//
//	Copyright (c) 2024 Tech4Works
//
//	Permission is hereby granted, free of charge, to any person obtaining a copy
//	of this software and associated documentation files (the "Software"), to deal
//	in the Software without restriction, including without limitation the rights
//	to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//	copies of the Software, and to permit persons to whom the Software is
//	furnished to do so, subject to the following conditions:
//
//	The above copyright notice and this permission notice shall be included in all
//	copies or substantial portions of the Software.
//
//	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//	IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//	AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//	LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//	OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//	SOFTWARE.

package checker

import (
	"encoding/base64"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

// IsURL checks the given value, converts it to string and determines whether it
// forms a valid URL.
//
// Parameters:
//   - a: Any value to be checked if it forms a valid URL.
//
// Returns:
//   - bool: A boolean value indicating whether the given value forms a valid URL.
//
// Panic:
//   - The function will panic if an unsupported value is passed.
//     If the value is not of a string, numeric, bool, array, slice, map, struct, interface, or pointer type.
//
// Example:
//
//	w := "https://example.com"
//	x := "Not a URL"
//	fmt.Println(IsURL(&w)) // true
//	fmt.Println(IsURL(w)) // true
//	fmt.Println(IsURL(x)) // false
func IsURL(a any) bool {
	_, err := url.ParseRequestURI(toString(a))
	return err == nil
}

// IsURLPath checks whether the given value is a valid URL path. It firstly
// converts the input to a string using toString function and then uses a
// regular expression to determine if the string is in URL path format.
//
// Parameters:
//   - a: Any value that will be checked if it's a URL path.
//
// Returns:
//   - bool: A boolean value indicating whether the value is a URL path.
//
// Panic:
//   - The function will panic if an unsupported value is passed.
//     If the value is not of a string, numeric, bool, array, slice, map, struct, interface, or pointer type.
//
// Example:
//
//	fmt.Println(IsURLPath("/test/abc")) // true
//	fmt.Println(IsURLPath("/sas-as/asa_asa/abc")) // true
//	fmt.Println(IsURLPath("/sas-as/asa_asa/abc/")) // true
//	fmt.Println(IsURLPath("not/a/path")) // false
//	fmt.Println(IsURLPath("not/a/path?document=12345")) // false
func IsURLPath(a any) bool {
	s := toString(a)
	regex := regexp.MustCompile(`^(/[^/?\s]*)*$`)
	return IsNotEmpty(s) && regex.MatchString(s)
}

// IsHTTPMethod checks if a given value matches a known HTTP method. It first converts the value to a string, then
// checks it against all predefined HTTP methods in the net/http package. These methods include GET, POST, HEAD,
// PUT, DELETE, CONNECT, OPTIONS, TRACE, PATCH. The comparison is case-sensitive.
//
// Parameters:
//   - a: The value to be checked. This can be any type, and will first be converted to a string.
//
// Returns:
//   - bool: A boolean indicating whether the given value matches a known HTTP method. Returns
//     true if a match is found, otherwise false.
//
// Panic:
//   - This function can potentially panic if its dependency function 'toString' encounters an unexpected type or value.
//     If the value is not of a string, numeric, bool, array, slice, map, struct, interface, or pointer type.
//
// Example:
//
//	var x = "POST"
//	var y = "GET"
//	var z = "random method"
//	fmt.Println(IsHTTPMethod(x)) // true
//	fmt.Println(IsHTTPMethod(&y)) // true
//	fmt.Println(IsHTTPMethod(z)) // false
func IsHTTPMethod(a any) bool {
	switch toString(a) {
	case http.MethodGet, http.MethodPost, http.MethodHead, http.MethodPut, http.MethodDelete, http.MethodConnect,
		http.MethodOptions, http.MethodTrace, http.MethodPatch:
		return true
	}
	return false
}

// IsAlpha checks if the input value, when converted to string, comprises entirely of alphabetic characters.
// It first converts the input to a string using the toString function, and then checks if the resultant string
// matches the Regexp "^\\p{L}+$" which represents a Unicode letter (letter in any language).
// This function does not check for empty strings - an empty string will return `false`.
//
// Parameters:
//   - a: Any value of the types acceptable by toString function. It is the value to be checked for alphabetic
//     character content.
//
// Returns:
//   - bool: A boolean value indicating whether the string representation of the input value is made of alphabetic
//     characters only.
//
// Panics:
//   - If the input value is of an unsupported type for conversion to string, or if the input is a nil pointer, a panic
//     is thrown by the toString function.
//
// Example:
//
//	str := "golang"
//	num := 12345
//	fmt.Println(IsAlpha(str)) // true
//	fmt.Println(IsAlpha(num)) // false
//	fmt.Println(IsAlpha(123)) // panic
//	fmt.Println(IsAlpha([]int{1, 2, 3})) // panic
//	fmt.Println(IsAlpha(nil)) // panic
func IsAlpha(a any) bool {
	s := toString(a)
	regex := regexp.MustCompile("^\\p{L}+$")
	return IsNotEmpty(s) && regex.MatchString(s)
}

// IsAlphaSpace checks the given value, converts it to string and determines whether it
// consists of alphabetic characters and spaces only.
//
// Parameters:
//   - a: Any value to be checked if it consists of alphabetic characters and spaces only.
//
// Returns:
//   - bool: A boolean value indicating whether the given value consists of alphabetic characters and spaces only.
//
// Panic:
//   - The function will panic if an unsupported value is passed.
//     If the value is not of a string, numeric, bool, array, slice, map, struct, interface, or pointer type.
//
// Example:
//
//	s := "Hello World"
//	fmt.Println(IsAlphaSpace(&s)) // true
//	fmt.Println(IsAlphaSpace(s)) // true
//	fmt.Println(IsAlphaSpace("123")) // false
//	fmt.Println(IsAlphaSpace(123)) // panic
//	fmt.Println(IsAlphaSpace([]int{1, 2, 3})) // panic
//	fmt.Println(IsAlphaSpace(nil)) // panic
func IsAlphaSpace(a any) bool {
	s := toString(a)
	regex := regexp.MustCompile("^[\\p{L} ]+$")
	return IsNotEmpty(s) && regex.MatchString(s)
}

// IsNumeric checks the given value, converts it to a string, and determines whether it
// consists of only numeric characters.
//
// Parameters:
//   - a: Any value to be checked if it consists of only numeric characters.
//
// Returns:
//   - bool: A boolean value indicating whether the given value consists of only numeric characters.
//
// Panic:
//   - The function will panic if an unsupported value is passed.
//     If the value is not of a string, numeric, bool, array, slice, map, struct, interface, or pointer type.
//
// Example:
//
//	s := "Hello World"
//	fmt.Println(IsNumeric(&s)) // false
//	fmt.Println(IsNumeric(s)) // false
//	fmt.Println(IsNumeric("123")) // false
//	fmt.Println(IsNumeric("123 1232 3232")) // false
//	fmt.Println(IsNumeric(123)) // panic
//	fmt.Println(IsNumeric([]int{1, 2, 3})) // panic
//	fmt.Println(IsNumeric(nil)) // panic
func IsNumeric(a any) bool {
	s := toString(a)
	regex := regexp.MustCompile("^[-.+0-9]+$")
	return IsNotEmpty(s) && regex.MatchString(s)
}

// IsNotNumeric checks whether a given value consists of non-numeric characters.
// It inverts the result from the IsNumeric function to provide this boolean value.
//
// Parameters:
//   - a: Any value to be checked if it consists of only non-numeric characters.
//
// Returns:
//   - bool: A boolean value indicating whether the given value consists of only non-numeric characters.
//
// Panic:
//   - The function will panic if an unsupported value is passed.
//     If the value is not of a string, numeric, bool, array, slice, map, struct, interface, or pointer type.
//
// Example:
//
//	s := "Hello World"
//	fmt.Println(IsNotNumeric(s)) // true
//	fmt.Println(IsNotNumeric("123")) // false
//	fmt.Println(IsNotNumeric(123)) // panic
//	fmt.Println(IsNotNumeric([]int{1, 2, 3})) // panic
//	fmt.Println(IsNotNumeric(nil)) // panic
func IsNotNumeric(a any) bool {
	return !IsNumeric(a)
}

// IsNumericSpace determines whether a given value is a numeric string or a string containing spaces only.
// It uses the toString function to convert the input to string and a regular expression to check the converted string.
//
// Parameters:
//   - a: Any value to be checked for a numeric or space only string.
//   - regex: A compiled regular expression used for checking the string.
//
// Returns:
//   - bool: A boolean value indicating whether the value is a numeric or space only string.
//
// Example:
//
//	a := "123456"
//	b := "123 456"
//	c := "123abc"
//	fmt.Println(IsNumericSpace(a)) // true
//	fmt.Println(IsNumericSpace(b)) // true
//	fmt.Println(IsNumericSpace(c)) // false
func IsNumericSpace(a any) bool {
	s := toString(a)
	regex := regexp.MustCompile("^[0-9 ]+$")
	return IsNotEmpty(s) && regex.MatchString(s)
}

// IsEmail determines whether a given value is a valid email. It uses the toString function
// to convert the value into a string then uses regex to verify it's a valid email pattern.
//
// Parameters:
//   - a: Any value that is to be checked if it's a valid email.
//
// Returns:
//   - bool: A boolean value indicating whether the value is a valid email.
//
// Panic:
//   - The function will panic if an unsupported value is passed.
//     If the value is not of a string, numeric, bool, array, slice, map, struct, interface, or pointer type.
//
// Example:
//
//	var x string = "test@example.com"
//	y := 12345
//	fmt.Println(IsEmail(x)) // true
//	fmt.Println(IsEmail(y)) // false
//	fmt.Println(IsEmail([]int{1, 2, 3})) // panic
//	fmt.Println(IsEmail(nil)) // panic
func IsEmail(a any) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return regex.MatchString(toString(a))
}

// IsNotEmail verifies whether a given value is not a valid email. It invokes the IsEmail function
// to check the value and flips its returned result.
//
// Parameters:
//   - a: Any interface value to be checked for valid email.
//
// Returns:
//   - bool: A boolean value indicating whether the value is not an email.
//
// Panic:
//   - The function might panic if the passed value is of an unsupported type.
//     If the value is not of a string, numeric, bool, array, slice, map, struct, interface, or pointer type.
//
// Example:
//
//	var x string = "test@example.com"
//	y := 12345
//	fmt.Println(IsNotEmail(x)) // false
//	fmt.Println(IsNotEmail(y)) // true
//	fmt.Println(IsNotEmail([]int{1, 2, 3})) // panic
//	fmt.Println(IsNotEmail(nil)) // panic
func IsNotEmail(a any) bool {
	return !IsEmail(a)
}

// IsDocument determines the type of document (CPF or CNPJ) and checks the value based on the document type.
// It uses the Document custom type to determine the document type, then uses the IsCPF or the IsCNPJ function
// to check if the value is valid for the specified document type.
// Only two document types are allowed: CPF and CNPJ.
//
// Parameters:
//   - documentType: A Document custom type to specify the type of the document.
//     Can be either DocumentCPF for CPF-type documents or DocumentCNPJ for CNPJ-type documents.
//   - a: Any interface value to be checked for validity based on the document type.
//
// Returns:
//   - bool: A boolean value indicating whether the value is valid for the specified document type.
//
// Panic:
//   - The function will panic if an unsupported Document type is passed.
//     Only DocumentCPF and DocumentCNPJ types are supported.
//     The error message will indicate the unsupported type.
//
// Example:
//
//	var docTypeCPF Document = DocumentCPF
//	var docTypeCNPJ Document = DocumentCNPJ
//	w := "12345678909"
//	x := "12345678901234"
//	z := "Not a Document"
//	fmt.Println(IsDocument(docTypeCPF, w)) // true
//	fmt.Println(IsDocument(docTypeCNPJ, x)) // true
//	fmt.Println(IsDocument(docTypeCPF, z)) // false
//	fmt.Println(IsDocument(docTypeCNPJ, z)) // false
//	fmt.Println(Document(0),w) // panic: unknown document type: CNH
func IsDocument(d Document, a any) bool {
	switch d {
	case DocumentCPF:
		return IsCPF(a)
	case DocumentCNPJ:
		return IsCNPJ(a)
	default:
		panic("unknown document type: " + d)
	}
}

// IsCPF checks the given value, converts it to string and determines whether it
// forms a valid CPF (Cadastro de Pessoas Físicas - Brazilian tax ID).
//
// Parameters:
//   - a: Any value to be checked if it forms a valid CPF.
//
// Returns:
//   - bool: A boolean value indicating whether the given value forms a valid CPF.
//
// Panic:
//   - The function will panic if an unsupported value is passed.
//     If the value is not of a string, numeric, bool, array, slice, map, struct,
//     interface, or pointer type.
//
// Example:
//
//	w := "12345678909"
//	x := "Not a CPF"
//	fmt.Println(IsCPF(&w)) // true
//	fmt.Println(IsCPF(w)) // true
//	fmt.Println(IsCPF(x)) // false
//	fmt.Println(IsCPF(nil)) // panic
func IsCPF(a any) bool {
	s := removeNonDigits(toString(a))
	if len(s) != 11 || allDigitsEqual(s) {
		return false
	}

	weights1 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2}
	weights2 := []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}

	firstVerifier, secondVerifier := calculateVerifierDigits(s, weights1, weights2)
	return firstVerifier == int(s[9]-'0') && secondVerifier == int(s[10]-'0')
}

// IsCNPJ checks the given value, converts it to string and determines whether it
// forms a valid CNPJ (Cadastro Nacional da Pessoa Jurídica - Brazilian company ID).
//
// Parameters:
//   - a: Any value to be checked if it forms a valid CNPJ.
//
// Returns:
//   - bool: A boolean value indicating whether the given value forms a valid CNPJ.
//
// Panic:
//   - The function will panic if an unsupported value is passed.
//     If the value is not of a string, numeric, bool, array, slice, map, struct,
//     interface, or pointer type.
//
// Example:
//
//	  w := "12345678901234"
//	  x := "Not a CNPJ"
//	  fmt.Println(IsCNPJ(&w)) // true
//	  fmt.Println(IsCNPJ(w)) // true
//	  fmt.Println(IsCNPJ(x)) // false
//		 fmt.Println(IsCNPJ(nil)) // panic
func IsCNPJ(a any) bool {
	s := removeNonDigits(toString(a))
	if len(s) != 14 || allDigitsEqual(s) {
		return false
	}

	weights1 := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	weights2 := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	firstVerifier, secondVerifier := calculateVerifierDigits(s, weights1, weights2)
	return firstVerifier == int(s[12]-'0') && secondVerifier == int(s[13]-'0')
}

// IsCPFOrCNPJ checks the given value, converts it to string and determines whether it
// forms a valid CPF (Cadastro de Pessoas Físicas - Brazilian tax ID) or a valid CNPJ
// (Cadastro Nacional da Pessoa Jurídica - Brazilian company ID).
// Parameters:
//   - a: Any value to be checked if it forms a valid CPF or CNPJ.
//
// Returns:
//   - bool: A boolean value indicating whether the given value forms a valid CPF or CNPJ.
//
// Panic:
//   - The function will panic if an unsupported value is passed.
//     If the value is not of a string, numeric, bool, array, slice, map, struct,
//     interface, or pointer type.
//
// Example:
//
//	w := "12345678909"
//	x := "Not a CPF"
//	fmt.Println(IsCPFOrCNPJ(&w)) // true
//	fmt.Println(IsCPFOrCNPJ(w)) // true
//	fmt.Println(IsCPFOrCNPJ("12345678901234")) // true
//	fmt.Println(IsCPFOrCNPJ(x)) // false
//	fmt.Println(IsCPFOrCNPJ(nil)) // panic
func IsCPFOrCNPJ(a any) bool {
	return IsCPF(a) || IsCNPJ(a)
}

// IsBase64 checks whether a given value is a Base64 encoded string. It uses string and regex manipulation
// to verify the format of the input.
//
// Parameters:
//   - a: Any value to be checked for Base64 encoding.
//
// Returns:
//   - bool: A boolean value indicating whether the value is a Base64 encoded string.
//
// Panic:
//   - The function will panic if an unsupported value is passed.
//     If the value is not of a string, numeric, bool, array, slice, map, struct,
//     interface, or pointer type.
//
// Example:
//
//	var validBase64 = "SGVsbG8gd29ybGQh" // "Hello world!" in Base64
//	var invalidBase64 = "Hello world!"
//	fmt.Println(IsBase64(validBase64)) // true
//	fmt.Println(IsBase64(invalidBase64)) // false
func IsBase64(a any) bool {
	s := toString(a)
	if IsEmpty(s) {
		return false
	}

	_, err := base64.StdEncoding.DecodeString(s)
	return err == nil
}

// IsBearer checks whether a given value carries a Bearer authentication scheme.
// It uses the toString function to convert the given value to a string.
// It then uses the Split method from the string package to divide the
// string in relation to "Bearer ", and finally checks if the result satisfies
// the Bearer authentication scheme criteria.
//
// Parameters:
//   - a: Any interface value to be processed and checked for a Bearer authentication scheme.
//
// Returns:
//   - bool: A boolean value indicating whether the value adheres to the Bearer authentication scheme.
//
// Panic:
//   - The function will panic if an unsupported value is passed.
//     If the value is not of a string, numeric, bool, array, slice, map, struct,
//     interface, or pointer type.
//
// Example:
//
//	var x string = "Bearer token"
//	y := 12345
//	fmt.Println(IsBearer(x)) // true
//	fmt.Println(IsBearer(y)) // false
func IsBearer(a any) bool {
	const bearer = "Bearer"
	split := strings.Split(toString(a), " ")
	return len(split) > 0 && split[0] == bearer
}

// IsPrivateIP determines whether the provided IP address is a private IP address.
// The function considers both IPv4 and IPv6 ranges for the check.
// It uses net.ParseCIDR to get the IP blocks of reserved private and local IPs.
// The argument is then parsed into an IP using the net.ParseIP function.
// The parsed IP object is examined for being loop back, link-local uni-cast, or link-local multicast.
// If it is not, the function checks whether the IP is within any of the defined private IP blocks.
// If any check is positive, the function returns true, indicating that it is a private IP.
// If every check is negative, the function returns false, indicating that the IP isn't private.
//
// Parameters:
//   - a: The interface value to be checked for being a private IP. Expected to be an IP address in string form.
//
// Returns:
//   - bool: A boolean value indicating whether the IP address is private.
//
// Panic:
//   - The function will panic if an unsupported value is passed.
//     If the value is not of a string, numeric, bool, array, slice, map, struct,
//     interface, or pointer type.
//
// Example:
//
//	x := "192.0.2.1"      // IPv4 public
//	y := "192.168.0.1"    // IPv4 private (RFC1918)
//	z := "2001:db8::1"    // IPv6 public
//	w := "fc00::1"        // IPv6 private (ULA)
//
//	fmt.Println(IsPrivateIP(x)) // false
//	fmt.Println(IsPrivateIP(y)) // true
//	fmt.Println(IsPrivateIP(z)) // false
//	fmt.Println(IsPrivateIP(w)) // true
func IsPrivateIP(a any) bool {
	var privateIPBlocks []*net.IPNet
	for _, cidr := range []string{
		"127.0.0.0/8",    // IPv4
		"10.0.0.0/8",     // RFC1918
		"172.16.0.0/12",  // RFC1918
		"192.168.0.0/16", // RFC1918
		"169.254.0.0/16", // RFC3927 link-local
		"::1/128",        // IPv6
		"fe80::/10",      // IPv6 link-local
		"fc00::/7",       // IPv6 unique local addr
	} {
		_, block, _ := net.ParseCIDR(cidr)
		if block != nil {
			privateIPBlocks = append(privateIPBlocks, block)
		}
	}

	ip := net.ParseIP(toString(a))
	result := ip.IsLoopback() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast()
	for _, block := range privateIPBlocks {
		if block.Contains(ip) {
			result = true
			break
		}
	}
	return result
}

// IsFullName validates if a given value matches a pattern for full names. It treats the given value as a string
// and checks whether it entirely matches the regular expression '^[\p{L}\s'-]+$'. The regular expression indicates
// a full name consisting of unicode letters, spaces, single quotes, or hyphens.
//
// Parameters:
//   - a: any. A value of any type that is to be checked and validated against the pattern for full names.
//
// Returns:
//   - bool: A boolean value. If the given value matches the pattern for full names, true is returned. Otherwise, false is returned.
//
// Panic:
//   - The function will panic if an unsupported value is passed.
//     If the value is not of a string, numeric, bool, array, slice, map, struct,
//     interface, or pointer type.
//
// Example:
//
//	name := "John Doe"
//	invalidName := "John123"
//	fmt.Println(IsFullName(name)) // true
//	fmt.Println(IsFullName(invalidName)) // false
func IsFullName(a any) bool {
	s := toString(a)
	split := strings.Fields(s)
	regex := regexp.MustCompile(`^[\p{L}\s'-]+$`)
	return len(split) > 1 && IsNotEmpty(s) && regex.MatchString(s)
}

// IsNotFullName checks if a given value is not considered a full name. It uses the IsFullName function
// to check if the value is a full name and returns the negation of its result.
//
// Parameters:
//   - a: Any interface value to be checked.
//
// Returns:
//   - bool: A boolean value indicating whether the value is not a full name.
//
// Panic:
//   - The function may panic if an unsupported value is passed to the IsFullName function.
//     Specifically, if the value is not of a string, numeric, bool, array, slice, map, struct,
//     interface, or pointer type, IsFullName will panic, and so will IsNotFullName.
//
// Example:
//
//	name := "John Doe"
//	invalidName := "John"
//	fmt.Println(IsNotFullName(name)) // false
//	fmt.Println(IsNotFullName(invalidName)) // true
func IsNotFullName(a any) bool {
	return !IsFullName(a)
}

// IsIOSDeviceID determines whether a given value adheres to the standard UUID format typically used in iOS device IDs.
// It converts the input to a string and then uses a regular expression to check if it matches the pattern.
//
// Parameters:
//   - a: Any value to be checked if it matches the iOS UUID format.
//
// Returns:
//   - bool: A boolean value indicating whether the given value matches the iOS UUID format.
//
// Panic:
//   - The function will panic if an unsupported value is passed.
//     If the value is not of a string, numeric, bool, array, slice, map, struct,
//     interface, or pointer type.
//
// Example:
//
//	id1 := "A1B2C3D4-E5F6-G7H8-I9J0-K1L2M3N4O5P6"
//	id2 := "incorrect-format"
//	fmt.Println(IsIOSDeviceID(id1)) // true
//	fmt.Println(IsIOSDeviceID(id2)) // false
func IsIOSDeviceID(a any) bool {
	s := toString(a)
	regex := regexp.MustCompile(`^[A-F0-9]{8}-[A-F0-9]{4}-[A-F0-9]{4}-[A-F0-9]{4}-[A-F0-9]{12}$`)
	return IsNotEmpty(s) && regex.MatchString(s)
}

// IsAndroidDeviceID determines whether a given value adheres to the standard format typically used in Android device IDs.
// It converts the input to a string and then uses a regular expression to check if it matches the hexadecimal pattern.
//
// Parameters:
//   - a: Any value to be checked if it matches the Android ID format.
//
// Returns:
//   - bool: A boolean value indicating whether the given value matches the Android ID format.
//
// Panic:
//   - The function will panic if an unsupported value is passed.
//     If the value is not of a string, numeric, bool, array, slice, map, struct,
//     interface, or pointer type.
//
// Example:
//
//	id1 := "abcdef123456"
//	id2 := "incorrect-format"
//	fmt.Println(IsAndroidDeviceId(id1)) // true
//	fmt.Println(IsAndroidDeviceId(id2)) // false
func IsAndroidDeviceID(a any) bool {
	s := toString(a)
	regex := regexp.MustCompile(`^[a-fA-F0-9]{16,}$`)
	return IsNotEmpty(s) && regex.MatchString(s)
}

// IsMobileDeviceID determines whether a given value is a valid Mobile Device ID.
// It uses the IsIOSDeviceID and IsAndroidDeviceId functions
// to check if the value is a valid iOS or Android device ID.
//
// Parameters:
//   - a: Any value to be checked for its validity as either an iOS or Android device ID.
//
// Returns:
//   - bool: A boolean value indicating whether the value is a valid iOS or Android device ID.
//
// Panic:
//   - The function will panic if an unsupported value is passed.
//     If the value is not of a string, numeric, bool, array, slice, map, struct,
//     interface, or pointer type.
//
// Example:
//
//	id1 := "A1B2C3D4-E5F6-G7H8-I9J0-K1L2M3N4O5P6"
//	id2 := "abcdef123456"
//	id3 := "incorrect-format"
//	fmt.Println(IsMobileDeviceID(id1)) // true
//	fmt.Println(IsMobileDeviceID(id2)) // true
//	fmt.Println(IsMobileDeviceID(id3)) // false
func IsMobileDeviceID(a any) bool {
	return IsIOSDeviceID(a) || IsAndroidDeviceID(a)
}

// IsMobilePlatform checks the given value, converts it to lowercase string,
// and determines whether it represents a mobile platform such as Android, iOS, or iPhone OS.
//
// Parameters:
//   - a: Any value to be checked if it represents a mobile platform.
//
// Returns:
//   - bool: A boolean value indicating whether the given value represents a mobile platform.
//
// Panic:
//   - The function will panic if an unsupported value is passed.
//     If the value is not of a string, numeric, bool, array, slice, map, struct,
//     interface, or pointer type.
//
// Example:
//
//	platform := "iOS"
//	fmt.Println(IsMobilePlatform(platform)) // true
func IsMobilePlatform(a any) bool {
	platform := strings.ToLower(toString(a))
	return platform == "android" || platform == "ios" || platform == "iphone os"
}
