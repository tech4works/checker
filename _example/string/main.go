package main

import (
	"fmt"
	"github.com/tech4works/checker"
)

func main() {
	fmt.Println("IsURL results:")
	fmt.Println(checker.IsURL("https://example.com")) // Should return true.
	fmt.Println(checker.IsURL("Not a URL"))           // Should return false.

	fmt.Println("IsURLPath results:")
	fmt.Println(checker.IsURLPath("/test/abc"))                 // Should return true.
	fmt.Println(checker.IsURLPath("/sas-as/asa_asa/abc"))       // Should return true.
	fmt.Println(checker.IsURLPath("/sas-as/asa_asa/abc/"))      // Should return true.
	fmt.Println(checker.IsURLPath("not/a/path"))                // Should return false.
	fmt.Println(checker.IsURLPath("not/a/path?document=12345")) // Should return false.

	fmt.Println("IsHTTPMethod results:")
	fmt.Println(checker.IsHTTPMethod("POST"))          // Should return true
	fmt.Println(checker.IsHTTPMethod("GET"))           // Should return true
	fmt.Println(checker.IsHTTPMethod("random method")) // Should return false

	fmt.Println("IsAlpha results:")
	str := "golang"
	num := 12345
	fmt.Println(checker.IsAlpha(str)) // Should return true
	fmt.Println(checker.IsAlpha(num)) // Should return false

	fmt.Println("IsAlphaSpace results:")
	s := "Hello World"
	fmt.Println(checker.IsAlphaSpace(&s))    // Should return true
	fmt.Println(checker.IsAlphaSpace(s))     // Should return true
	fmt.Println(checker.IsAlphaSpace("123")) // Should return false

	fmt.Println("IsNumeric results:")
	fmt.Println(checker.IsNumeric(&s))              // Should return false
	fmt.Println(checker.IsNumeric(s))               // Should return false
	fmt.Println(checker.IsNumeric("123"))           // Should return true
	fmt.Println(checker.IsNumeric("123 1232 3232")) // Should return false

	fmt.Println("IsNumericSpace results:")
	fmt.Println(checker.IsNumericSpace("123456"))  // Should return true
	fmt.Println(checker.IsNumericSpace("123 456")) // Should return true
	fmt.Println(checker.IsNumericSpace("123abc"))  // Should return false

	fmt.Println("IsEmail results:")
	fmt.Println(checker.IsEmail("email@example.com")) // Should return true
	fmt.Println(checker.IsEmail("bad email"))         // Should return false

	fmt.Println("IsCPF results:")
	fmt.Println(checker.IsCPF("12101721007"))    // Should return true
	fmt.Println(checker.IsCPF("121.017.210-07")) // Should return true
	fmt.Println(checker.IsCPF("11111111111"))    // Should return false
	fmt.Println(checker.IsCPF("111.111.111-11")) // Should return false
	fmt.Println(checker.IsCPF("invalid"))        // Should return false

	fmt.Println("IsCNPJ results:")
	fmt.Println(checker.IsCNPJ("57309623000168"))     // Should return true
	fmt.Println(checker.IsCNPJ("47.263.759/0001-20")) // Should return true
	fmt.Println(checker.IsCNPJ("00000000000100"))     // Should return false
	fmt.Println(checker.IsCNPJ("00.000.000/0001-00")) // Should return false
	fmt.Println(checker.IsCNPJ("invalid"))            // Should return false

	fmt.Println("IsCPFOrCNPJ results:")
	fmt.Println(checker.IsCPFOrCNPJ("12101721007"))        // Should return true
	fmt.Println(checker.IsCPFOrCNPJ("121.017.210-07"))     // Should return true
	fmt.Println(checker.IsCPFOrCNPJ("57309623000168"))     // Should return true
	fmt.Println(checker.IsCPFOrCNPJ("47.263.759/0001-20")) // Should return true
	fmt.Println(checker.IsCPFOrCNPJ("00000000000100"))     // Should return false
	fmt.Println(checker.IsCPFOrCNPJ("00.000.000/0001-00")) // Should return false
	fmt.Println(checker.IsCPFOrCNPJ("11111111111"))        // Should return false
	fmt.Println(checker.IsCPFOrCNPJ("111.111.111-11"))     // Should return false
	fmt.Println(checker.IsCPFOrCNPJ("invalid"))            // Should return false

	fmt.Println("IsBase64 results:")
	fmt.Println(checker.IsBase64("SGVsbG8gd29ybGQh")) // Should return true
	fmt.Println(checker.IsBase64("Hello world!"))     // Should return false

	fmt.Println("IsBearer results:")
	fmt.Println(checker.IsBearer("Bearer token")) // Should return true
	fmt.Println(checker.IsBearer("token"))        // Should return false
	fmt.Println(checker.IsBearer(12345))          // Should return false

	fmt.Println("IsPrivateIP results:")
	fmt.Println(checker.IsPrivateIP("192.0.2.1"))   // Should return false
	fmt.Println(checker.IsPrivateIP("192.168.0.1")) // Should return true
	fmt.Println(checker.IsPrivateIP("2001:db8::1")) // Should return false
	fmt.Println(checker.IsPrivateIP("fc00::1"))     // Should return true

	fmt.Println("IsFullName results:")
	fmt.Println(checker.IsFullName("John Doe")) // Should return true
	fmt.Println(checker.IsFullName("John123"))  // Should return false

	fmt.Println("IsIOSDeviceId results:")
	fmt.Println(checker.IsIOSDeviceId("A1B2C3D4-E5F6-G7H8-I9J0-K1L2M3N4O5P6")) // Should return true
	fmt.Println(checker.IsIOSDeviceId("asiajsiaaslaks2312oksa"))               // Should return false

	fmt.Println("IsAndroidDeviceId results:")
	fmt.Println(checker.IsAndroidDeviceId("abcdef123456"))     // Should return true
	fmt.Println(checker.IsAndroidDeviceId("incorrect-format")) // Should return false

	fmt.Println("IsMobileDeviceId results:")
	fmt.Println(checker.IsMobileDeviceId("abcdef123456"))                         // Should return true
	fmt.Println(checker.IsMobileDeviceId("A1B2C3D4-E5F6-G7H8-I9J0-K1L2M3N4O5P6")) // Should return true
	fmt.Println(checker.IsMobileDeviceId("incorrect-format"))                     // Should return false

	fmt.Println("IsMobilePlatform results:")
	fmt.Println(checker.IsMobilePlatform("ANDROID"))          // Should return true
	fmt.Println(checker.IsMobilePlatform("android"))          // Should return true
	fmt.Println(checker.IsMobilePlatform("IOS"))              // Should return true
	fmt.Println(checker.IsMobilePlatform("iOS"))              // Should return true
	fmt.Println(checker.IsMobilePlatform("ios"))              // Should return true
	fmt.Println(checker.IsMobilePlatform("iphone os"))        // Should return true
	fmt.Println(checker.IsMobilePlatform("iphone"))           // Should return false
	fmt.Println(checker.IsMobileDeviceId("incorrect-format")) // Should return false
}
