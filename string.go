package checker

import (
	"golang.org/x/crypto/bcrypt"
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
//   - a: Any value which will be checked if it's a URL path.
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
//	fmt.Println(IsURLPath("not/a/path")) // false
//	fmt.Println(IsURLPath("not/a/path/")) // false
func IsURLPath(a any) bool {
	regex := regexp.MustCompile(`^/([^/\s]*)+(/[^/\s]+)*$`)
	return regex.MatchString(toString(a))
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
	regex := regexp.MustCompile("^\\p{L}+$")
	return regex.MatchString(toString(a))
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
	regex := regexp.MustCompile("^[\\p{L} ]+$")
	return regex.MatchString(toString(a))
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
	regex := regexp.MustCompile("^[0-9]+$")
	return regex.MatchString(toString(a))
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
	regex := regexp.MustCompile("^[0-9 ]+$")
	return regex.MatchString(toString(a))
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
	regex := regexp.MustCompile(`^([A-Za-z0-9+/]{4})*([A-Za-z0-9+/]{4}|[A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{2}==)?$`)
	return regex.MatchString(toString(a))
}

// IsBCrypt determines whether a given value represents a valid bcrypt cost.
// It uses the bcrypt.Cost function after converting the input to a byte slice with the toBytes function.
//
// Parameters:
//   - a: Any interface value to be checked for a valid bcrypt cost.
//
// Returns:
//   - bool: A boolean value indicating whether the value represents a valid bcrypt cost.
//
// Panic:
//   - The function will panic if an unsupported value is passed.
//     If the value is not of a string, numeric, bool, array, slice, map, struct,
//     interface, or pointer type.
//
// Example:
//
//	password := "mypassword"
//	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
//	fmt.Println(IsBCrypt(hash)) // true
//	fmt.Println(IsBCrypt(password)) // false
func IsBCrypt(a any) bool {
	cost, err := bcrypt.Cost(toBytes(a))
	return err == nil && (cost == bcrypt.MinCost || cost == bcrypt.DefaultCost || cost == bcrypt.MaxCost)
}

// IsBearer checks whether a given value carries a Bearer authentication scheme.
// It utilizes the toString function to convert the given value to a string.
// It then uses the Split method from the strings package to divide the
// string in relation to "Bearer ", and finally checks if the result satisfies
// the Bearer authentication scheme criteria.
//
// Parameters:
//   - a: Any interface value to be processed and checked for Bearer authentication scheme.
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
	const bearer = "Bearer "
	split := strings.Split(toString(a), bearer)
	return len(split) > 0 && split[0] == bearer
}

// IsPrivateIP determines whether the provided IP address is a private IP address.
// The function considers both IPv4 and IPv6 ranges for the check.
// It uses net.ParseCIDR to obtain the IP blocks of reserved private and local IPs.
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
	regex := regexp.MustCompile(`^[\p{L}\s'-]+$`)
	return regex.MatchString(toString(a))
}

// IsIOSDeviceId determines whether a given value adheres to the standard UUID format typically used in iOS device IDs.
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
//	fmt.Println(IsIOSDeviceId(id1)) // true
//	fmt.Println(IsIOSDeviceId(id2)) // false
func IsIOSDeviceId(a any) bool {
	regex := regexp.MustCompile(`[A-F0-9]{8}-[A-F0-9]{4}-[A-F0-9]{4}-[A-F0-9]{4}-[A-F0-9]{12}`)
	return regex.MatchString(toString(a))
}

// IsAndroidDeviceId determines whether a given value adheres to the standard format typically used in Android device IDs.
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
func IsAndroidDeviceId(a any) bool {
	regex := regexp.MustCompile(`[0-9a-fA-F]`)
	return regex.MatchString(toString(a))
}

// IsMobileDeviceId determines whether a given value is a valid Mobile Device ID.
// It uses the IsIOSDeviceId and IsAndroidDeviceId functions
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
//	fmt.Println(IsMobileDeviceId(id1)) // true
//	fmt.Println(IsMobileDeviceId(id2)) // true
//	fmt.Println(IsMobileDeviceId(id3)) // false
func IsMobileDeviceId(a any) bool {
	return IsIOSDeviceId(a) || IsAndroidDeviceId(a)
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
