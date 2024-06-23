package checker

// BaseEnum is an interface that defines a method IsEnumValid.
//
// It is used in conjunction with the BaseEnum interface to validate enum values.
// The IsEnumValid method should return a boolean value indicating if the enum is valid.
type BaseEnum interface {
	// IsEnumValid is a method that returns a boolean value indicating if the enum is valid.
	//
	// It is used in conjunction with the BaseEnum interface to validate enum values.
	IsEnumValid() bool
}
