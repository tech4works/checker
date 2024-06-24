package checker

import "github.com/go-playground/validator/v10"

// customValidate is a variable of type *validator.Validate that can be used for putting strict checks and
// validation on the inputs and requests to maintain the data integrity and correctness throughout the application.
//
// It is initialized and registered with multiple custom validation methods upon calling the Validate function
// if the customValidate instance is not already initialized.
var customValidate *validator.Validate

// Validate prepares and returns an instance of *validator.Validate. It registers multiple validation methods
// upon calling if the customValidate instance is not already initialized.
//
// This function implements various custom validations for the provided struct. Each of these validation methods
// has been registered using RegisterValidation method. They include: http_method, url_path, enum, full_name,
// bcrypt, bearer, before_now, before_today, after_now, after_today, today, cpf, cnpj,
// cpfcnpj, duration and byte_unit.
//
// This Validate function can be used for putting strict checks and validation on the inputs and requests
// to maintain the data integrity and correctness throughout the application.
//
// It makes use of the 'validator' package to perform validation.
//
// Returns:
//   - *validator.Validate: An instance of *validator.Validate with all predefined validation methods registered.
//
// Example:
//
//	err := Validate().Struct(initValidatorStruct())
//	if err != nil {
//	  log.Println("Validation error:", err)
//	} else {
//	  log.Println("Validation successful")
//	}
func Validate() *validator.Validate {
	if customValidate != nil {
		return customValidate
	}

	customValidate = validator.New()
	_ = customValidate.RegisterValidation("http_method", validateHttpMethod)
	_ = customValidate.RegisterValidation("url_path", validateUrlPath)
	_ = customValidate.RegisterValidation("enum", validateEnum)
	_ = customValidate.RegisterValidation("full_name", validateFullName)
	_ = customValidate.RegisterValidation("bcrypt", validateBcrypt)
	_ = customValidate.RegisterValidation("bearer", validateBearer)
	_ = customValidate.RegisterValidation("before_now", validateBeforeNow)
	_ = customValidate.RegisterValidation("before_today", validateBeforeToday)
	_ = customValidate.RegisterValidation("after_now", validateAfterNow)
	_ = customValidate.RegisterValidation("after_today", validateAfterToday)
	_ = customValidate.RegisterValidation("today", validateToday)
	_ = customValidate.RegisterValidation("cpf", validateCpf)
	_ = customValidate.RegisterValidation("cnpj", validateCnpj)
	_ = customValidate.RegisterValidation("cpfcnpj", validateCpfCnpj)
	_ = customValidate.RegisterValidation("duration", validateDuration)
	_ = customValidate.RegisterValidation("byte_unit", validateByteUnit)
	return customValidate
}

// validateFullName checks if the provided value is a valid full name.
// It calls the IsFullName function, passing the value as a parameter, and returns the result.
// The IsFullName function converts the value to a string using the toString function and
// checks if the string matches the regular expression '^[\p{L}\s'-]+$' which represents a full name format.
//
// Parameters:
//   - fl: A validator.FieldLevel struct that contains information about the field being validated.
//     The value of the field can be accessed using fl.Field().
//
// Returns:
//   - bool: A boolean value indicating whether the value is a valid full name.
//     Returns true if the value is a valid full name, otherwise false.
func validateFullName(fl validator.FieldLevel) bool {
	return IsFullName(fl.Field().String())
}

// validateBcrypt checks whether the given value represents a valid bcrypt cost.
// It calls the IsBCrypt function, passing the value as a parameter, and returns the result.
// The IsBCrypt function converts the value to a byte slice using the toBytes function and
// checks if the byte slice is a valid bcrypt cost using the bcrypt.Cost function.
//
// Parameters:
//   - fl: A validator.FieldLevel struct that contains information about the field being validated.
//     The value of the field can be accessed using fl.Field().
//
// Returns:
//   - bool: A boolean value indicating whether the value represents a valid bcrypt cost.
//     Returns true if the value represents a valid bcrypt cost, otherwise false.
func validateBcrypt(fl validator.FieldLevel) bool {
	return IsBCrypt(fl.Field().Interface())
}

// validateBearer checks whether the given value carries a Bearer authentication scheme.
// It calls the IsBearer function, passing the value as a parameter, and returns the result.
// The IsBearer function converts the value to a string using the toString function and
// checks if the string satisfies the Bearer authentication scheme criteria.
//
// Parameters:
//   - fl: A validator.FieldLevel struct that contains information about the field being validated.
//     The value of the field can be accessed using fl.Field().
//
// Returns:
//   - bool: A boolean value indicating whether the value carries a Bearer authentication scheme.
//     Returns true if the value carries a Bearer authentication scheme, otherwise false.
func validateBearer(fl validator.FieldLevel) bool {
	return IsBearer(fl.Field().Interface())
}

// validateBeforeNow determines whether the value of the field is before the current time.
// It calls the IsBeforeNow function, passing the value as a parameter, and returns the result.
//
// Parameters:
//   - fl: A validator.FieldLevel struct that contains information about the field being validated.
//     The value of the field can be accessed using fl.Field().
//
// Returns:
//   - bool: A boolean value indicating whether the value is before the current time.
//     Returns true if the value is before the current time, otherwise false.
func validateBeforeNow(fl validator.FieldLevel) bool {
	return IsBeforeNow(fl.Field().Interface())
}

// validateBeforeToday checks if the provided value represents a time that is before the current date.
// It calls the IsBeforeToday function, passing the value as a parameter, and returns the result.
// The IsBeforeToday function converts the value to a time.Time object using the toDate function and
// compares it with the current date obtained from the dateNow function. If the converted date is before
// the current date, IsBeforeToday returns true.
//
// Parameters:
//   - fl: A validator.FieldLevel struct that contains information about the field being validated.
//     The value of the field can be accessed using fl.Field().
//
// Returns:
//   - bool: A boolean value indicating whether the value represents a time before the current date.
//     Returns true if the value represents a time before the current date, otherwise false.
func validateBeforeToday(fl validator.FieldLevel) bool {
	return IsBeforeToday(fl.Field().Interface())
}

// validateAfterNow checks if the provided value is a time value that is after the current time.
// It calls the IsAfterNow function, passing the value as a parameter, and returns the result.
//
// Parameters:
//   - fl: A validator.FieldLevel struct that contains information about the field being validated.
//     The value of the field can be accessed using fl.Field().
//
// Returns:
//   - bool: A boolean value indicating whether the value is after the current time.
//     Returns true if the value is after the current time, otherwise false.
func validateAfterNow(fl validator.FieldLevel) bool {
	return IsAfterNow(fl.Field().Interface())
}

// validateAfterToday checks whether the given value is after today's date.
// It calls the IsAfterToday function, passing the value as a parameter, and returns the result.
// The IsAfterToday function converts the value to a time.Time object using the toDate function.
// Then it compares the converted time with the current date (obtained via the dateNow
// function) to determine if the converted time is after today.
//
// Parameters:
//   - fl: A validator.FieldLevel struct that contains information about the field being validated.
//     The value of the field can be accessed using fl.Field().
//
// Returns:
//   - bool: A boolean value indicating whether the value is after today's date.
//     Returns true if the value is after today's date, otherwise false.
func validateAfterToday(fl validator.FieldLevel) bool {
	return IsAfterToday(fl.Field().Interface())
}

// validateCpfCnpj checks whether the given value forms a valid CPF or CNPJ (Cadastro Nacional da Pessoa Jurídica - Brazilian tax ID).
// It calls the IsCPF or IsCNPJ function based on the length of the value, passing the value as a parameter, and returns the result.
// The IsCPF and IsCNPJ functions respectively convert the value to a string using the toString function and
// checks if the string forms a valid CPF or CNPJ using the cpfcnpj.ValidateCPF and cpfcnpj.ValidateCNPJ functions.
//
// Parameters:
//   - fl: A validator.FieldLevel struct that contains information about the field being validated.
//     The value of the field can be accessed using fl.Field().
//
// Returns:
//   - bool: A boolean value indicating whether the value forms a valid CPF or CNPJ.
//     Returns true if the value forms a valid CPF or CNPJ, otherwise false.
func validateToday(fl validator.FieldLevel) bool {
	return IsToday(fl.Field().Interface())
}

// validateCpf checks whether the given value forms a valid CPF (Cadastro de Pessoas Físicas - Brazilian tax ID).
// It calls the IsCPF function, passing the value as a parameter, and returns the result.
// The IsCPF function converts the value to a string using the toString function and checks if the string
// forms a valid CPF using the cpfcnpj.ValidateCPF function.
//
// Parameters:
//   - fl: A `validator.FieldLevel` struct that contains information about the field being validated.
//     The value of the field can be accessed using `fl.Field()`.
//
// Returns:
//   - bool: A boolean value indicating whether the value forms a valid CPF.
//     Returns `true` if the value forms a valid CPF, otherwise `false`.
func validateCpf(fl validator.FieldLevel) bool {
	return IsCPF(fl.Field().Interface())
}

// validateCnpj validates if the provided value is a valid CNPJ (Cadastro Nacional da Pessoa Jurídica - Brazilian company ID).
//
// Parameters:
//   - fl: A validator.FieldLevel struct that contains information about the field being validated.
//     The value of the field can be accessed using fl.Field().
//
// Returns:
//   - bool: A boolean value indicating whether the value is a valid CNPJ.
//     Returns true if the value is a valid CNPJ, otherwise false.
func validateCnpj(fl validator.FieldLevel) bool {
	return IsCNPJ(fl.Field().Interface())
}

// validateAfterToday checks whether the given value is a date that is after today's date.
// It calls the IsAfterToday function, passing the value as a parameter, and returns the result.
// The IsAfterToday function converts the value to a time.Time type using the toTime function and
// checks if the date is after today's date.
//
// Parameters:
//   - fl: A validator.FieldLevel struct that contains information about the field being validated.
//     The value of the field can be accessed using fl.Field().
//
// Returns:
//   - bool: A boolean value indicating whether the value is a date that is after today's date.
//     Returns true if the value is after today's date, otherwise false.
func validateCpfCnpj(fl validator.FieldLevel) bool {
	a := fl.Field().Interface()
	return IsCPFOrCNPJ(a)
}

// validateDuration checks whether the given value is of the time.Duration type.
// It calls the IsDurationType function, passing the value as a parameter, and returns the result.
// The IsDurationType function performs a comparison between the type of the provided value
// and the type of the time.Duration using the reflect.TypeOf function.
//
// Parameters:
//   - fl: A validator.FieldLevel struct that contains information about the field being validated.
//     The value of the field can be accessed using fl.Field().
//
// Returns:
//   - bool: A boolean value indicating whether the value is of the type time.Duration.
//     Returns true if the value is of type time.Duration, otherwise false.
func validateDuration(fl validator.FieldLevel) bool {
	return IsDurationType(fl.Field().Interface())
}

// validateAfterToday checks if the provided value is after the current date.
// It calls the IsAfterToday function, passing the value as a parameter, and returns the result.
// The IsAfterToday function converts the value to a time.Time type using the parseTime function
// and then checks if the parsed time is after the current date.
//
// Parameters:
//   - fl: A validator.FieldLevel struct that contains information about the field being validated.
//     The value of the field can be accessed using fl.Field().
//
// Returns:
//   - bool: A boolean value indicating whether the value is after the current date.
//     Returns true if the value is after the current date, otherwise false.
func validateByteUnit(fl validator.FieldLevel) bool {
	return IsByteUnit(fl.Field().Interface())
}

// validateEnum checks if the provided value implements the interface BaseEnum
// and if the IsEnumValid method of the value returns true.
//
// Parameters:
//   - fl: A validator.FieldLevel struct that contains information about the field being validated.
//     The value of the field can be accessed using fl.Field().
//
// Returns:
//   - bool: A boolean value indicating whether the value is a valid enum.
//     Returns true if the value implements the BaseEnum interface and its IsEnumValid method returns true,
//     otherwise false.
func validateEnum(fl validator.FieldLevel) bool {
	value, ok := fl.Field().Interface().(BaseEnum)
	return ok && value.IsEnumValid()
}

// validateUrlPath checks whether the given value is a valid URL path.
// It calls the IsURLPath function, passing the value as a parameter, and returns the result.
// The IsURLPath function converts the value to a string using the toString function and
// checks if the string is in URL path format using a regular expression.
//
// Parameters:
//   - fl: A validator.FieldLevel struct that contains information about the field being validated.
//     The value of the field can be accessed using fl.Field().
//
// Returns:
//   - bool: A boolean value indicating whether the value is a valid URL path.
//     Returns true if the value is a valid URL path, otherwise false.
func validateUrlPath(fl validator.FieldLevel) bool {
	return IsURLPath(fl.Field().Interface())
}

// validateHttpMethod checks if the provided value matches a known HTTP method.
// It calls the IsHTTPMethod function, passing the value as a parameter, and returns the result.
// The IsHTTPMethod function converts the value to a string and compares it against predefined
// HTTP methods such as GET, POST, HEAD, PUT, DELETE, CONNECT, OPTIONS, TRACE, and PATCH. The comparison
// is case-sensitive.
//
// Parameters:
//   - fl: A validator.FieldLevel struct that contains information about the field being validated.
//     The value of the field can be accessed using fl.Field().
//
// Returns:
//   - bool: A boolean value indicating whether the value matches a known HTTP method.
//     Returns true if a match is found, otherwise false.
func validateHttpMethod(fl validator.FieldLevel) bool {
	return IsHTTPMethod(fl.Field().Interface())
}
