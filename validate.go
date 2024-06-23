package checker

import "github.com/go-playground/validator/v10"

var customValidate *validator.Validate

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

func validateFullName(fl validator.FieldLevel) bool {
	return IsFullName(fl.Field().String())
}

func validateBcrypt(fl validator.FieldLevel) bool {
	return IsBCrypt(fl.Field().Interface())
}

func validateBearer(fl validator.FieldLevel) bool {
	return IsBearer(fl.Field().Interface())
}

func validateBeforeNow(fl validator.FieldLevel) bool {
	return IsBeforeNow(fl.Field().Interface())
}

func validateBeforeToday(fl validator.FieldLevel) bool {
	return IsBeforeToday(fl.Field().Interface())
}

func validateAfterNow(fl validator.FieldLevel) bool {
	return IsAfterNow(fl.Field().Interface())
}

func validateAfterToday(fl validator.FieldLevel) bool {
	return IsAfterToday(fl.Field().Interface())
}

func validateToday(fl validator.FieldLevel) bool {
	return IsToday(fl.Field().Interface())
}

func validateCpf(fl validator.FieldLevel) bool {
	return IsCPF(fl.Field().Interface())
}

func validateCnpj(fl validator.FieldLevel) bool {
	return IsCNPJ(fl.Field().Interface())
}

func validateCpfCnpj(fl validator.FieldLevel) bool {
	a := fl.Field().Interface()
	return IsCPFOrCNPJ(a)
}

func validateDuration(fl validator.FieldLevel) bool {
	return IsDurationType(fl.Field().Interface())
}

func validateByteUnit(fl validator.FieldLevel) bool {
	return IsByteUnit(fl.Field().Interface())
}

func validateEnum(fl validator.FieldLevel) bool {
	value, ok := fl.Field().Interface().(BaseEnum)
	return ok && value.IsEnumValid()
}

func validateUrlPath(fl validator.FieldLevel) bool {
	return IsURLPath(fl.Field().Interface())
}

func validateHttpMethod(fl validator.FieldLevel) bool {
	return IsHTTPMethod(fl.Field().Interface())
}
